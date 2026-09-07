// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package checkergen

import (
	"fmt"
	"go/format"
	"strings"
)

// fieldPlan is one tagged, eligible field's resolved generation plan.
type fieldPlan struct {
	goName    string // the field's Go identifier, e.g. "Email"
	tagName   string // the error-map key, e.g. "email"
	goType    string // one of scalarGoTypes
	omitEmpty bool
	callExprs []string // Go source, one checker.CheckFunc[goType] expression per tag token, in order
}

// generateStructRecovered calls generateStruct, converting a panic (a
// malformed tag parameter, e.g. after-field missing its ":field" half --
// the same "closer to user input than the compile-time struct tag the core
// module assumes" situation the cli module's CheckWithConfig recovery
// handles, see CLAUDE.md) into a returned error, so one struct's malformed
// tag can't crash generation for the rest of the package.
func generateStructRecovered(st loadedStruct) (plans []fieldPlan, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()

	return generateStruct(st)
}

// generateStruct builds fieldPlans for every eligible field of st, or
// returns an error naming the first field that isn't eligible (an unmapped
// checker name, or a field whose type isn't one of scalarGoTypes) -- the
// caller skips the whole struct on error, per the package doc.
func generateStruct(st loadedStruct) ([]fieldPlan, error) {
	// All field declarations are resolved to a Go type first, in one pass,
	// so a field-relative checker can reference a sibling declared later
	// in the struct.
	var plans []fieldPlan

	// tagged[i] mirrors plans[i]: whether that field actually has a
	// checkers/validate tag to fill in during the second pass below.
	var tagged []bool

	byName := make(map[string]int)

	for _, field := range st.fields {
		if len(field.Names) != 1 {
			// Embedded or multi-name field declaration; out of scope for v1.
			continue
		}

		config, hasTag := tagConfig(field)

		goType, eligible := goTypeName(st.info.TypeOf(field.Type))
		if hasTag && config != "" && !eligible {
			return nil, fmt.Errorf("field %s: type %s is not eligible for generation (must be a predeclared scalar type)", field.Names[0].Name, goType)
		}

		byName[field.Names[0].Name] = len(plans)

		plans = append(plans, fieldPlan{
			goName:  field.Names[0].Name,
			tagName: jsonName(field),
			goType:  goType,
		})

		tagged = append(tagged, hasTag && config != "")
	}

	sib := func(name string) (expr string, err error) {
		i, ok := byName[name]
		if !ok {
			return "", fmt.Errorf("field %s not found", name)
		}

		return "v." + plans[i].goName, nil
	}

	for _, field := range st.fields {
		if len(field.Names) != 1 {
			continue
		}

		i := byName[field.Names[0].Name]
		if !tagged[i] {
			continue
		}

		config, _ := tagConfig(field)

		if err := fillCallExprs(&plans[i], config, sib); err != nil {
			return nil, fmt.Errorf("field %s: %w", plans[i].goName, err)
		}
	}

	return plans, nil
}

// fillCallExprs parses config's space-separated tokens (extracting
// omitempty) and resolves each remaining one to a Go source expression via
// callSpecs/fieldCallSpecs, in order, appending to plan.callExprs. Returns
// an error naming the first unmapped checker.
func fillCallExprs(plan *fieldPlan, config string, sib sibling) error {
	for _, token := range strings.Fields(config) {
		if token == omitEmptyName {
			plan.omitEmpty = true
			continue
		}

		name, params, _ := strings.Cut(token, ":")

		spec, ok := callSpecs[name]
		if !ok {
			spec, ok = fieldCallSpecs[name]
		}

		if !ok {
			return fmt.Errorf("checker %q has no checkergen mapping", name)
		}

		expr, err := spec(params, plan.goType, sib)
		if err != nil {
			return fmt.Errorf("checker %q: %w", name, err)
		}

		plan.callExprs = append(plan.callExprs, expr)
	}

	return nil
}

// emitFunc renders the Check<Name>(v *Name) (checker.CheckErrors, bool)
// function for a struct's resolved field plans.
func emitFunc(structName string, plans []fieldPlan) string {
	var b strings.Builder

	fmt.Fprintf(&b, "func Check%s(v *%s) (checker.CheckErrors, bool) {\n", structName, structName)
	b.WriteString("errs := make(checker.CheckErrors)\n\n")

	for _, plan := range plans {
		if len(plan.callExprs) == 0 {
			continue
		}

		b.WriteString("{\n")

		if plan.omitEmpty {
			fmt.Fprintf(&b, "if v.%s != %s {\n", plan.goName, zeroLiteral(plan.goType))
		}

		fmt.Fprintf(&b, "newValue, err := checker.Check(v.%s, %s)\n", plan.goName, strings.Join(plan.callExprs, ", "))
		fmt.Fprintf(&b, "v.%s = newValue\n", plan.goName)
		fmt.Fprintf(&b, "if err != nil {\n errs[%q] = err\n}\n", plan.tagName)

		if plan.omitEmpty {
			b.WriteString("}\n")
		}

		b.WriteString("}\n\n")
	}

	b.WriteString("return errs, len(errs) == 0\n}\n")

	return b.String()
}

// emitFile renders a complete generated Go source file for pkgName,
// containing one Check<Type> function per struct in plansByStruct (in the
// given order).
func emitFile(pkgName string, structNames []string, plansByStruct map[string][]fieldPlan) ([]byte, error) {
	var b strings.Builder

	b.WriteString("// Code generated by checkergen. DO NOT EDIT.\n\n")
	fmt.Fprintf(&b, "package %s\n\n", pkgName)
	b.WriteString("import (\n\tchecker \"github.com/cinar/checker/v2\"\n)\n\n")

	for _, name := range structNames {
		b.WriteString(emitFunc(name, plansByStruct[name]))
		b.WriteString("\n")
	}

	return format.Source([]byte(b.String()))
}
