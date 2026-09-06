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

func TestIsRequiredUnlessConditionNotMetMissing(t *testing.T) {
	_, err := v2.IsRequiredUnless("", "member", "guest")

	if !errors.Is(err, v2.ErrRequired) {
		t.Fatalf("got unexpected error %v", err)
	}
}

func TestIsRequiredUnlessConditionNotMetPresent(t *testing.T) {
	value, err := v2.IsRequiredUnless("onur@example.com", "member", "guest")
	if err != nil {
		t.Fatal(err)
	}

	if value != "onur@example.com" {
		t.Fatalf("actual %s expected %s", value, "onur@example.com")
	}
}

func TestIsRequiredUnlessConditionMet(t *testing.T) {
	value, err := v2.IsRequiredUnless("", "guest", "guest")
	if err != nil {
		t.Fatal(err)
	}

	if value != "" {
		t.Fatalf("actual %q expected empty", value)
	}
}

func TestCheckStructRequiredUnlessConditionNotMetSuccess(t *testing.T) {
	type Person struct {
		Type  string `checkers:"required"`
		Email string `checkers:"required-unless:Type:guest"`
	}

	person := &Person{
		Type:  "member",
		Email: "onur@example.com",
	}

	errs, ok := v2.CheckStruct(person)
	if !ok {
		t.Fatalf("got unexpected errors %v", errs)
	}
}

func TestCheckStructRequiredUnlessConditionNotMetMissing(t *testing.T) {
	type Person struct {
		Type  string `checkers:"required"`
		Email string `checkers:"required-unless:Type:guest"`
	}

	person := &Person{
		Type: "member",
	}

	errs, ok := v2.CheckStruct(person)
	if ok {
		t.Fatal("expected errors")
	}

	if !errors.Is(errs["Email"], v2.ErrRequired) {
		t.Fatalf("expected email required %v", errs)
	}
}

func TestCheckStructRequiredUnlessConditionMet(t *testing.T) {
	type Person struct {
		Type  string `checkers:"required"`
		Email string `checkers:"required-unless:Type:guest"`
	}

	person := &Person{
		Type: "guest",
	}

	errs, ok := v2.CheckStruct(person)
	if !ok {
		t.Fatalf("got unexpected errors %v", errs)
	}
}

func TestMakeRequiredUnlessMissingValue(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Person struct {
		Type  string
		Email string `checkers:"required-unless:Type"`
	}

	person := &Person{
		Type: "guest",
	}

	v2.CheckStruct(person)
}

func TestRequiredUnlessMissingField(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Person struct {
		Type  string
		Email string `checkers:"required-unless:Unknown:guest"`
	}

	person := &Person{
		Type: "guest",
	}

	v2.CheckStruct(person)
}
