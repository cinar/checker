// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2_test

import (
	"fmt"
	"testing"

	v2 "github.com/cinar/checker/v2"
)

func ExampleIsAlphanumeric() {
	_, err := v2.IsAlphanumeric("ABcd1234")
	if err != nil {
		fmt.Println(err)
	}
}

func TestIsAlphanumericInvalid(t *testing.T) {
	_, err := v2.IsAlphanumeric("-/")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestIsAlphanumericValid(t *testing.T) {
	_, err := v2.IsAlphanumeric("ABcd1234")
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

	v2.CheckStruct(person)
}

func TestCheckAlphanumericInvalid(t *testing.T) {
	type Person struct {
		Name string `checkers:"alphanumeric"`
	}

	person := &Person{
		Name: "name-/",
	}

	_, ok := v2.CheckStruct(person)
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

	errs, ok := v2.CheckStruct(person)
	if !ok {
		t.Fatal(errs)
	}
}
