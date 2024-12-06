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

func ExampleIsASCII() {
	_, err := v2.IsASCII("Checker")
	if err != nil {
		fmt.Println(err)
	}
}

func TestIsASCIIInvalid(t *testing.T) {
	_, err := v2.IsASCII("𝄞 Music!")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestIsASCIIValid(t *testing.T) {
	_, err := v2.IsASCII("Checker")
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

	v2.CheckStruct(person)
}

func TestCheckASCIIInvalid(t *testing.T) {
	type Person struct {
		Name string `checkers:"ascii"`
	}

	person := &Person{
		Name: "𝄞 Music!",
	}

	_, valid := v2.CheckStruct(person)
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

	_, valid := v2.CheckStruct(person)
	if !valid {
		t.Fail()
	}
}
