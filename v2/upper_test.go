// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2_test

import (
	"testing"

	v2 "github.com/cinar/checker/v2"
)

func TestUpper(t *testing.T) {
	input := "checker"
	expected := "CHECKER"

	actual, err := v2.Upper(input)
	if err != nil {
		t.Fatal(err)
	}

	if actual != expected {
		t.Fatalf("actual %s expected %s", actual, expected)
	}
}

func TestReflectUpper(t *testing.T) {
	type Person struct {
		Name string `checkers:"upper"`
	}

	person := &Person{
		Name: "checker",
	}

	expected := "CHECKER"

	errs, ok := v2.CheckStruct(person)
	if !ok {
		t.Fatalf("got unexpected errors %v", errs)
	}

	if person.Name != expected {
		t.Fatalf("actual %s expected %s", person.Name, expected)
	}
}