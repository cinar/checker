// Copyright (c) 2023-2024 Onur Cinar.
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

func ExampleIsOneOf() {
	_, err := v2.IsOneOf("admin", "admin", "user", "guest")
	if err != nil {
		fmt.Println(err)
	}
}

func TestIsOneOfValid(t *testing.T) {
	_, err := v2.IsOneOf("admin", "admin", "user", "guest")
	if err != nil {
		t.Fatal(err)
	}
}

func TestIsOneOfInvalid(t *testing.T) {
	_, err := v2.IsOneOf("owner", "admin", "user", "guest")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestIsOneOfInt(t *testing.T) {
	_, err := v2.IsOneOf(2, 1, 2, 3)
	if err != nil {
		t.Fatal(err)
	}
}

func TestOneOfErrorMessage(t *testing.T) {
	_, err := v2.IsOneOf("owner", "admin", "user", "guest")

	expected := "Value must be one of admin, user, guest."

	if err.Error() != expected {
		t.Fatalf("actual %q expected %q", err.Error(), expected)
	}
}

func TestMakeOneOfNoValuesPanics(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Role struct {
		Name string `checkers:"oneof"`
	}

	v2.CheckStruct(&Role{Name: "admin"})
}

func TestCheckOneOfNonString(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Role struct {
		Level int `checkers:"oneof:1,2,3"`
	}

	role := &Role{}

	v2.CheckStruct(role)
}

func TestCheckOneOfInvalid(t *testing.T) {
	type Role struct {
		Name string `checkers:"oneof:admin,user,guest"`
	}

	role := &Role{
		Name: "owner",
	}

	errs, ok := v2.CheckStruct(role)
	if ok {
		t.Fatal("expected error")
	}

	if !errors.Is(errs["Name"], v2.ErrNotOneOf) {
		t.Fatalf("expected ErrNotOneOf, got %v", errs)
	}
}

func TestCheckOneOfValid(t *testing.T) {
	type Role struct {
		Name string `checkers:"oneof:admin,user,guest"`
	}

	role := &Role{
		Name: "user",
	}

	errs, ok := v2.CheckStruct(role)
	if !ok {
		t.Fatal(errs)
	}
}
