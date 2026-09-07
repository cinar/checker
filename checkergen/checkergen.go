// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

// Package checkergen generates reflection-free Go validation code from
// Checker's "checkers" (and "validate" fallback) struct tags. For a struct
// eligible for generation, it emits a Check<Type>(v *Type)
// (checker.CheckErrors, bool) function that calls the same checker/
// normalizer functions CheckStruct would run through reflection, but
// resolved and called directly at compile time via checker.Check.
//
// It's a separate module, not part of the core checker package: emitting
// Go source is a different kind of concern than validating or reflecting
// into a JSON Schema, and this tool needs golang.org/x/tools (go/packages)
// for reliable type information, which the core module never depends on.
//
// v1 scope: a struct field is only eligible for generation if its Go type
// is exactly one of the predeclared scalar types (string, bool, the int/
// uint/float kinds) -- not a named/defined type (type Email string), a
// pointer, a nested struct, or a slice/map. checkergen and CheckStruct are
// meant to coexist: a struct with a field outside that scope is skipped
// (with a diagnostic explaining why) rather than generated incorrectly,
// and stays on the reflect-based CheckStruct path.
package checkergen

import (
	"fmt"
	"os"
	"path/filepath"
)

// GeneratedFileName is the name of the single output file Generate writes
// into the target package's directory, carrying every struct it
// successfully generated code for.
const GeneratedFileName = "checkergen_generated.go"

// Result summarizes one Generate call: the structs it successfully
// generated code for, and the structs it skipped, each with the reason.
type Result struct {
	// Generated lists the names of every struct a Check<Type> function was
	// emitted for, in the order they were declared.
	Generated []string

	// Skipped maps a struct name to the reason it was skipped: a tagged
	// field with an unmapped checker name, or a type outside v1's scope
	// (see the package doc). A struct with no checkers/validate tag on any
	// field at all isn't a candidate in the first place, so it's not
	// recorded here.
	Skipped map[string]string
}

// Generate scans the package at dir (a directory, loaded as its own
// go/packages pattern) for every named struct type with at least one
// checkers/validate-tagged field, optionally restricted to typeFilter
// (struct names; empty means every struct), and writes GeneratedFileName
// into dir containing a Check<Type> function for each one it could
// generate. A struct it can't generate (see Result.Skipped) is left out of
// the file, not written as broken code, and doesn't prevent the rest of
// the package's structs from being generated.
func Generate(dir string, typeFilter []string) (*Result, error) {
	filter := make(map[string]bool, len(typeFilter))
	for _, name := range typeFilter {
		filter[name] = true
	}

	structs, pkgName, err := findStructs(dir, filter)
	if err != nil {
		return nil, err
	}

	result := &Result{}

	var names []string

	plansByStruct := make(map[string][]fieldPlan)

	for _, st := range structs {
		plans, err := generateStructRecovered(st)
		if err != nil {
			skipf(result, st.name, "%v", err)
			continue
		}

		names = append(names, st.name)
		plansByStruct[st.name] = plans
		result.Generated = append(result.Generated, st.name)
	}

	if len(names) == 0 {
		return result, nil
	}

	data, err := emitFile(pkgName, names, plansByStruct)
	if err != nil {
		return nil, fmt.Errorf("formatting generated source: %w", err)
	}

	if err := os.WriteFile(filepath.Join(dir, GeneratedFileName), data, 0o600); err != nil {
		return nil, fmt.Errorf("writing %s: %w", GeneratedFileName, err)
	}

	return result, nil
}

// skipf records that a struct is being skipped, for a reason string built
// from format/args.
func skipf(result *Result, structName, format string, args ...any) {
	if result.Skipped == nil {
		result.Skipped = make(map[string]string)
	}

	result.Skipped[structName] = fmt.Sprintf(format, args...)
}
