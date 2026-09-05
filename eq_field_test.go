// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2_test

import (
	"errors"
	"reflect"
	"testing"

	v2 "github.com/cinar/checker/v2"
)

func TestCheckStructEqFieldSuccess(t *testing.T) {
	type Person struct {
		Password        string `checkers:"required"`
		ConfirmPassword string `checkers:"eq-field:Password"`
	}

	person := &Person{
		Password:        "secret",
		ConfirmPassword: "secret",
	}

	errs, ok := v2.CheckStruct(person)
	if !ok {
		t.Fatalf("got unexpected errors %v", errs)
	}
}

func TestCheckStructEqFieldMismatch(t *testing.T) {
	type Person struct {
		Password        string `checkers:"required"`
		ConfirmPassword string `checkers:"eq-field:Password"`
	}

	person := &Person{
		Password:        "secret",
		ConfirmPassword: "not-secret",
	}

	errs, ok := v2.CheckStruct(person)
	if ok {
		t.Fatal("expected errors")
	}

	if !errors.Is(errs["ConfirmPassword"], v2.ErrEqField) {
		t.Fatalf("expected eq-field error %v", errs)
	}
}

func TestEqFieldMissingParent(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	v2.IsEqField(reflect.Value{}, reflect.ValueOf("secret"), "Password")
}

func TestEqFieldMissingField(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Person struct {
		Password string
	}

	person := Person{Password: "secret"}

	v2.IsEqField(reflect.ValueOf(person), reflect.ValueOf("secret"), "Unknown")
}
