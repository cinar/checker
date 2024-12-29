// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2_test

import (
	"testing"

	v2 "github.com/cinar/checker/v2"
)

func TestMaxLenSuccess(t *testing.T) {
	value := "test"

	check := v2.MaxLen[string](4)

	result, err := check(value)
	if result != value {
		t.Fatalf("result (%s) is not the original value (%s)", result, value)
	}

	if err != nil {
		t.Fatal(err)
	}
}

func TestMaxLenError(t *testing.T) {
	value := "test test"

	check := v2.MaxLen[string](5)

	result, err := check(value)
	if result != value {
		t.Fatalf("result (%s) is not the original value (%s)", result, value)
	}

	message := "Value cannot be greater than 5."

	if err.Error() != message {
		t.Fatalf("expected %s actual %s", message, err.Error())
	}
}

func TestReflectMaxLenError(t *testing.T) {
	type Person struct {
		Name string `checkers:"max-len:2"`
	}

	person := &Person{
		Name: "Onur",
	}

	errs, ok := v2.CheckStruct(person)
	if ok {
		t.Fatalf("expected errors")
	}

	if errs["Name"] == nil {
		t.Fatalf("expected maximum length error")
	}
}

func TestReflectMaxLenInvalidMaxLen(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Person struct {
		Name string `checkers:"max-len:abcd"`
	}

	person := &Person{
		Name: "Onur",
	}

	v2.CheckStruct(person)
}

func TestReflectMaxLenInvalidType(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Person struct {
		Name int `checkers:"max-len:8"`
	}

	person := &Person{
		Name: 1,
	}

	v2.CheckStruct(person)
}
