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

func ExampleIsAlpha() {
	_, err := v2.IsAlpha("ABcd")
	if err != nil {
		fmt.Println(err)
	}
}

func TestIsAlphaInvalid(t *testing.T) {
	_, err := v2.IsAlpha("abc123")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestIsAlphaValid(t *testing.T) {
	_, err := v2.IsAlpha("ABcd")
	if err != nil {
		t.Fatal(err)
	}
}

func TestCheckAlphaNonString(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Person struct {
		Name int `checkers:"alpha"`
	}

	person := &Person{}

	v2.CheckStruct(person)
}

func TestCheckAlphaInvalid(t *testing.T) {
	type Person struct {
		Name string `checkers:"alpha"`
	}

	person := &Person{
		Name: "name123",
	}

	_, ok := v2.CheckStruct(person)
	if ok {
		t.Fatal("expected error")
	}
}

func TestCheckAlphaValid(t *testing.T) {
	type Person struct {
		Name string `checkers:"alpha"`
	}

	person := &Person{
		Name: "ABcd",
	}

	errs, ok := v2.CheckStruct(person)
	if !ok {
		t.Fatal(errs)
	}
}
