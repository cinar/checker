// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2_test

import (
	"errors"
	"testing"

	v2 "github.com/cinar/checker/v2"
)

func TestCheckStructValidateTagFallback(t *testing.T) {
	type Person struct {
		Name string `validate:"trim required"`
	}

	person := &Person{
		Name: "  Onur Cinar  ",
	}

	errs, ok := v2.CheckStruct(person)
	if !ok {
		t.Fatal(errs)
	}

	if person.Name != "Onur Cinar" {
		t.Fatalf("expected trimmed value, got %q", person.Name)
	}
}

func TestCheckStructValidateTagFallbackReportsErrors(t *testing.T) {
	type Person struct {
		Name string `validate:"required"`
	}

	person := &Person{}

	errs, ok := v2.CheckStruct(person)
	if ok {
		t.Fatal("expected errors")
	}

	if !errors.Is(errs["Name"], v2.ErrRequired) {
		t.Fatalf("expected ErrRequired, got %v", errs)
	}
}

func TestCheckStructCheckersTagWinsOverValidateTag(t *testing.T) {
	type Person struct {
		// checkers is present and takes precedence, even though it's a
		// weaker rule than the validate tag it's paired with here -- this
		// intentionally proves precedence, not a realistic dual-tag setup.
		Name string `checkers:"trim" validate:"required"`
	}

	person := &Person{
		Name: "   ",
	}

	if _, ok := v2.CheckStruct(person); !ok {
		t.Fatal("expected valid: checkers tag (trim only) should win over validate tag (required)")
	}
}

func TestCheckStructEmptyCheckersTagDoesNotFallBackToValidate(t *testing.T) {
	type Person struct {
		// An explicit, empty checkers tag means "no checks for this field",
		// not "fall back to validate".
		Name string `checkers:"" validate:"required"`
	}

	person := &Person{}

	if _, ok := v2.CheckStruct(person); !ok {
		t.Fatal("expected valid: explicit empty checkers tag should not fall back to validate")
	}
}

func TestJSONSchemaValidateTagFallback(t *testing.T) {
	type Person struct {
		Name string `validate:"required"`
	}

	schema := v2.JSONSchema(&Person{})

	if len(schema.Required) != 1 || schema.Required[0] != "Name" {
		t.Fatalf("expected Name to be required via validate tag fallback, got %v", schema.Required)
	}
}
