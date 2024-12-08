// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package checker_test

import (
	"fmt"
	"testing"

	"github.com/cinar/checker"
)

func ExampleIsAlphanumeric() {
	_, err := checker.IsAlphanumeric("ABcd1234")
	if err != nil {
		fmt.Println(err)
	}
}

func TestIsAlphanumericInvalid(t *testing.T) {
	_, err := checker.IsAlphanumeric("-/")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestIsAlphanumericValid(t *testing.T) {
	_, err := checker.IsAlphanumeric("ABcd1234")
	if err != nil {
		t.Fatal(err)
	}
}

func TestCheckAlphanumericNonString(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Person struct {
		Name int `checkers:"alphanumeric"`
	}

	person := &Person{}

	checker.CheckStruct(person)
}

func TestCheckAlphanumericInvalid(t *testing.T) {
	type Person struct {
		Name string `checkers:"alphanumeric"`
	}

	person := &Person{
		Name: "name-/",
	}

	_, ok := checker.CheckStruct(person)
	if ok {
		t.Fatal("expected error")
	}
}

func TestCheckAlphanumericValid(t *testing.T) {
	type Person struct {
		Name string `checkers:"alphanumeric"`
	}

	person := &Person{
		Name: "ABcd1234",
	}

	errs, ok := checker.CheckStruct(person)
	if !ok {
		t.Fatal(errs)
	}
}
