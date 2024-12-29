// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2_test

import (
	"testing"

	v2 "github.com/cinar/checker/v2"
)

func TestTrimLeft(t *testing.T) {
	input := "    test    "
	expected := "test    "

	actual, err := v2.TrimLeft(input)
	if err != nil {
		t.Fatal(err)
	}

	if actual != expected {
		t.Fatalf("actual %s expected %s", actual, expected)
	}
}

func TestReflectTrimLeft(t *testing.T) {
	type Person struct {
		Name string `checkers:"trim-left"`
	}

	person := &Person{
		Name: "    test    ",
	}

	expected := "test    "

	errs, ok := v2.CheckStruct(person)
	if !ok {
		t.Fatalf("got unexpected errors %v", errs)
	}

	if person.Name != expected {
		t.Fatalf("actual %s expected %s", person.Name, expected)
	}
}
