// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2_test

import (
	"errors"
	"fmt"
	"testing"

	v2 "github.com/cinar/checker/v2"
)

func ExampleIsEq() {
	_, err := v2.IsEq("active", "active")
	if err != nil {
		fmt.Println(err)
	}
}

func TestIsEqValid(t *testing.T) {
	_, err := v2.IsEq("active", "active")
	if err != nil {
		t.Fatal(err)
	}
}

func TestIsEqInvalid(t *testing.T) {
	_, err := v2.IsEq("inactive", "active")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestIsEqInt(t *testing.T) {
	_, err := v2.IsEq(5, 5)
	if err != nil {
		t.Fatal(err)
	}
}

func TestEqErrorMessage(t *testing.T) {
	_, err := v2.IsEq("inactive", "active")

	expected := "Value must equal active."

	if err.Error() != expected {
		t.Fatalf("actual %q expected %q", err.Error(), expected)
	}
}

func TestCheckEqNonString(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Status struct {
		Value int `checkers:"eq:5"`
	}

	status := &Status{}

	v2.CheckStruct(status)
}

func TestCheckEqInvalid(t *testing.T) {
	type Status struct {
		Value string `checkers:"eq:active"`
	}

	status := &Status{
		Value: "inactive",
	}

	errs, ok := v2.CheckStruct(status)
	if ok {
		t.Fatal("expected error")
	}

	if !errors.Is(errs["Value"], v2.ErrNotEq) {
		t.Fatalf("expected ErrNotEq, got %v", errs)
	}
}

func TestCheckEqValid(t *testing.T) {
	type Status struct {
		Value string `checkers:"eq:active"`
	}

	status := &Status{
		Value: "active",
	}

	errs, ok := v2.CheckStruct(status)
	if !ok {
		t.Fatal(errs)
	}
}
