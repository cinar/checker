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

func ExampleIsASCII() {
	_, err := checker.IsASCII("Checker")
	if err != nil {
		fmt.Println(err)
	}
}

func TestIsASCIIInvalid(t *testing.T) {
	_, err := checker.IsASCII("𝄞 Music!")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestIsASCIIValid(t *testing.T) {
	_, err := checker.IsASCII("Checker")
	if err != nil {
		t.Fatal(err)
	}
}

func TestCheckASCIINonString(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Person struct {
		Name int `checkers:"ascii"`
	}

	person := &Person{}

	checker.CheckStruct(person)
}

func TestCheckASCIIInvalid(t *testing.T) {
	type Person struct {
		Name string `checkers:"ascii"`
	}

	person := &Person{
		Name: "𝄞 Music!",
	}

	_, valid := checker.CheckStruct(person)
	if valid {
		t.Fail()
	}
}

func TestCheckASCIIValid(t *testing.T) {
	type Person struct {
		Name string `checkers:"ascii"`
	}

	person := &Person{
		Name: "checker",
	}

	_, valid := checker.CheckStruct(person)
	if !valid {
		t.Fail()
	}
}
