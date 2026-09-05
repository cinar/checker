// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

// Package checkerlint provides a static analyzer that validates
// cinar/checker "checkers" (and "validate" fallback) struct tags at
// compile time: unknown checker/normalizer names, checkers applied to a
// field of an incompatible type, and field-relative checkers whose target
// field doesn't exist on the same struct.
package checkerlint

import (
	"go/ast"
	"go/types"
	"reflect"
	"strconv"
	"strings"

	checker "github.com/cinar/checker/v2"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const (
	checkerTag        = "checkers"
	validateTag       = "validate"
	sliceConfigPrefix = "@"
	omitEmptyName     = "omitempty"

	checkerPackagePath = "github.com/cinar/checker/v2"
	registerMakerName  = "RegisterMaker"
	registerFieldMaker = "RegisterFieldMaker"
)

const doc = `check that cinar/checker struct tags are syntactically and semantically sound

The checkerlint analyzer parses "checkers" (and "validate" fallback) struct
tags and reports:

  - unknown checker/normalizer names
  - checkers applied to a field of a type their implementation can't
    handle (e.g. a string-only checker on a numeric field)
  - field-relative checkers (eq-field, after-field, before-field,
    required-if, required-unless) whose target field name doesn't exist
    on the same struct

It knows the built-in checker vocabulary from the version of
github.com/cinar/checker/v2 this analyzer is built against, plus any
custom checker registered in the analyzed packages via a
checker.RegisterMaker/RegisterFieldMaker call with a string-literal name.
It cannot see a custom checker registered only at runtime in a different
package (e.g. dynamically, or via a name built from a non-literal
expression).`

// Analyzer reports invalid cinar/checker struct tags at compile time.
var Analyzer = &analysis.Analyzer{
	Name:     "checkerlint",
	Doc:      doc,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
	Run:      run,
}

// stringOnly is the set of built-in checker/normalizer names whose
// implementation panics if applied to a field whose reflect.Kind isn't
// String (see reflectString in the core module). Kept in sync by hand
// against the core module's checker files that call reflectString(value).
var stringOnly = map[string]bool{
	"after-field": true, "alpha": true, "alphanumeric": true, "ascii": true,
	"before-field": true, "cidr": true, "contains": true, "credit-card": true,
	"digits": true, "email": true, "ends-with": true, "eoa": true, "eq": true,
	"fqdn": true, "hash": true, "hex": true, "html-escape": true,
	"html-unescape": true, "ip": true, "ipv4": true, "ipv6": true, "isbn": true,
	"iso3166-1-alpha-2": true, "iso3166-1-alpha-3": true, "iso639-1": true,
	"lower": true, "luhn": true, "mac": true, "ne": true, "numeric": true,
	"oneof": true, "regexp": true, "starts-with": true, "title": true,
	"trim-left": true, "trim-right": true, "trim": true, "upper": true,
	"url": true, "url-escape": true, "url-unescape": true, "uuid": true,
}

// numericOnly is the set of built-in checker names whose implementation
// panics if applied to a field that isn't an integer, unsigned integer, or
// float kind (see gt.go/gte.go/lt.go/lte.go's "value is not numeric" panic).
var numericOnly = map[string]bool{
	"gt": true, "gte": true, "lt": true, "lte": true,
}

// fieldTargetParam extracts the target sibling field name from a
// field-relative checker's params, given the checker's name. ok is false
// if name isn't a field-relative checker with a field-name parameter, or
// the parameter doesn't contain one.
func fieldTargetParam(name, params string) (field string, ok bool) {
	switch name {
	case "eq-field", "after-field", "before-field":
		return params, params != ""
	case "required-if", "required-unless":
		field, _, found := strings.Cut(params, ":")
		return field, found && field != ""
	default:
		return "", false
	}
}

// knownNames returns the checker vocabulary this analyzer run considers
// valid: every name currently registered in the checker module this
// analyzer was built against, plus "omitempty".
func knownNames() map[string]bool {
	known := map[string]bool{omitEmptyName: true}

	for _, name := range checker.RegisteredMakerNames() {
		known[name] = true
	}

	for _, name := range checker.RegisteredFieldMakerNames() {
		known[name] = true
	}

	return known
}

func run(pass *analysis.Pass) (interface{}, error) {
	insp := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	known := knownNames()

	// First pass: record custom checkers registered in the analyzed
	// packages via a checker.RegisterMaker/RegisterFieldMaker call with a
	// string-literal name, so real projects using their own checkers
	// don't get flagged as unknown.
	insp.Preorder([]ast.Node{(*ast.CallExpr)(nil)}, func(n ast.Node) {
		recordCustomRegistration(pass, n.(*ast.CallExpr), known)
	})

	// Second pass: check every struct's field tags against the now-complete
	// vocabulary.
	insp.Preorder([]ast.Node{(*ast.StructType)(nil)}, func(n ast.Node) {
		checkStruct(pass, n.(*ast.StructType), known)
	})

	return nil, nil
}

// recordCustomRegistration adds name to known if call is a
// checker.RegisterMaker/RegisterFieldMaker(name, ...) call with a
// string-literal first argument.
func recordCustomRegistration(pass *analysis.Pass, call *ast.CallExpr, known map[string]bool) {
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok || len(call.Args) == 0 {
		return
	}

	fn, ok := pass.TypesInfo.Uses[sel.Sel].(*types.Func)
	if !ok || fn.Pkg() == nil || fn.Pkg().Path() != checkerPackagePath {
		return
	}

	if fn.Name() != registerMakerName && fn.Name() != registerFieldMaker {
		return
	}

	// The real RegisterMaker/RegisterFieldMaker's first parameter is a
	// string, so a BasicLit reaching here can only be a string literal;
	// anything else assignable to that parameter (a variable, a constant
	// reference, a concatenation, ...) isn't a BasicLit at all and is
	// skipped by the type assertion below, deliberately: this analyzer
	// only recognizes literal, statically-visible checker names.
	lit, ok := call.Args[0].(*ast.BasicLit)
	if !ok {
		return
	}

	name, _ := strconv.Unquote(lit.Value)
	known[name] = true
}

// checkStruct validates every tagged field of a struct type literal.
func checkStruct(pass *analysis.Pass, st *ast.StructType, known map[string]bool) {
	siblings := siblingFieldNames(st)

	for _, field := range st.Fields.List {
		if field.Tag == nil {
			continue
		}

		config, ok := tagConfig(field.Tag.Value)
		if !ok || config == "" {
			continue
		}

		fieldType := pass.TypesInfo.TypeOf(field.Type)

		checkFieldConfig(pass, field, fieldType, config, known, siblings)
	}
}

// tagConfig returns the effective checkers/validate tag config from a raw
// (still-quoted) struct tag literal, and whether one was present at all.
// rawTag is field.Tag.Value from a successfully parsed Go source file, so
// it's always a syntactically valid string literal.
func tagConfig(rawTag string) (string, bool) {
	unquoted, _ := strconv.Unquote(rawTag)

	tag := reflect.StructTag(unquoted)

	if config, ok := tag.Lookup(checkerTag); ok {
		return config, true
	}

	return tag.Lookup(validateTag)
}

// siblingFieldNames returns the set of explicitly named fields on a struct
// type literal (embedded/anonymous fields are approximated by their type
// name).
func siblingFieldNames(st *ast.StructType) map[string]bool {
	names := make(map[string]bool)

	for _, field := range st.Fields.List {
		if len(field.Names) == 0 {
			names[embeddedFieldName(field.Type)] = true
			continue
		}

		for _, name := range field.Names {
			names[name.Name] = true
		}
	}

	return names
}

// embeddedFieldName approximates the promoted field name of an anonymous
// struct field from its type expression (e.g. "*pkg.Address" -> "Address").
func embeddedFieldName(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.SelectorExpr:
		return t.Sel.Name
	case *ast.StarExpr:
		return embeddedFieldName(t.X)
	default:
		return ""
	}
}

// checkFieldConfig validates every token of a single field's checkers
// config against the known vocabulary, field-relative targets, and (for a
// well-scoped set of built-in checkers) type compatibility.
func checkFieldConfig(pass *analysis.Pass, field *ast.Field, fieldType types.Type, config string, known, siblings map[string]bool) {
	container, elemType := containerElem(fieldType)

	var sliceTokens, itemTokens []string

	itemType := fieldType
	if container {
		sliceTokens, itemTokens = splitSliceConfig(config)
		itemType = elemType
	} else {
		itemTokens = strings.Fields(config)
	}

	for _, token := range sliceTokens {
		checkToken(pass, field, nil, token, known, siblings)
	}

	for _, token := range itemTokens {
		checkToken(pass, field, itemType, token, known, siblings)
	}
}

// checkToken validates a single checkers-tag token: vocabulary, then (if
// valueType is non-nil) type compatibility, then (if field-relative)
// cross-field target existence.
func checkToken(pass *analysis.Pass, field *ast.Field, valueType types.Type, token string, known, siblings map[string]bool) {
	name, params, _ := strings.Cut(token, ":")

	if name == omitEmptyName {
		return
	}

	if !known[name] {
		pass.Reportf(field.Tag.Pos(), "checkerlint: unknown checker %q", name)
		return
	}

	if valueType != nil {
		checkTypeCompatibility(pass, field, valueType, name)
	}

	if target, ok := fieldTargetParam(name, params); ok && !siblings[target] {
		pass.Reportf(field.Tag.Pos(), "checkerlint: %s references field %q, which doesn't exist on this struct", name, target)
	}
}

// checkTypeCompatibility reports a diagnostic if a type-constrained
// built-in checker is applied to a value of an incompatible kind.
func checkTypeCompatibility(pass *analysis.Pass, field *ast.Field, valueType types.Type, name string) {
	underlying := indirectOnce(valueType).Underlying()

	basic, ok := underlying.(*types.Basic)
	if !ok {
		return
	}

	switch {
	case stringOnly[name] && basic.Info()&types.IsString == 0:
		pass.Reportf(field.Tag.Pos(), "checkerlint: %s requires a string, but the field's type is %s", name, valueType)

	case numericOnly[name] && basic.Info()&(types.IsInteger|types.IsFloat) == 0:
		pass.Reportf(field.Tag.Pos(), "checkerlint: %s requires a numeric type, but the field's type is %s", name, valueType)
	}
}

// containerElem reports whether t (after at most one level of pointer
// indirection) is a slice, array, or map, and if so, its element type --
// mirroring how CheckStruct itself descends into these kinds.
func containerElem(t types.Type) (bool, types.Type) {
	switch u := indirectOnce(t).Underlying().(type) {
	case *types.Slice:
		return true, u.Elem()
	case *types.Array:
		return true, u.Elem()
	case *types.Map:
		return true, u.Elem()
	default:
		return false, nil
	}
}

// indirectOnce strips a single level of pointer indirection from t, like
// the core module's indirectOrNilPointer/reflect.Indirect do at runtime.
func indirectOnce(t types.Type) types.Type {
	if ptr, ok := t.Underlying().(*types.Pointer); ok {
		return ptr.Elem()
	}

	return t
}

// splitSliceConfig splits config into slice/map-level ("@"-prefixed) and
// item-level tokens, mirroring the core module's splitSliceConfig.
func splitSliceConfig(config string) (sliceTokens, itemTokens []string) {
	for _, token := range strings.Fields(config) {
		if strings.HasPrefix(token, sliceConfigPrefix) {
			sliceTokens = append(sliceTokens, strings.TrimPrefix(token, sliceConfigPrefix))
		} else {
			itemTokens = append(itemTokens, token)
		}
	}

	return sliceTokens, itemTokens
}
