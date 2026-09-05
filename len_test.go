// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2_test

import (
	"testing"

	v2 "github.com/cinar/checker/v2"
)

func TestLenSuccess(t *testing.T) {
	value := "test"

	check := v2.Len[string](4)

	result, err := check(value)
	if result != value {
		t.Fatalf("result (%s) is not the original value (%s)", result, value)
	}

	if err != nil {
		t.Fatal(err)
	}
}

func TestLenErrorTooShort(t *testing.T) {
	value := "test"

	check := v2.Len[string](5)

	result, err := check(value)
	if result != value {
		t.Fatalf("result (%s) is not the original value (%s)", result, value)
	}

	message := "Value must have a length of 5."

	if err.Error() != message {
		t.Fatalf("expected %s actual %s", message, err.Error())
	}
}

func TestLenErrorTooLong(t *testing.T) {
	value := "test"

	check := v2.Len[string](3)

	_, err := check(value)
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestReflectLenSliceSuccess(t *testing.T) {
	type Person struct {
		Codes []string `checkers:"@len:2"`
	}

	person := &Person{
		Codes: []string{"a", "b"},
	}

	if _, ok := v2.CheckStruct(person); !ok {
		t.Fatal("expected valid")
	}
}

func TestReflectLenError(t *testing.T) {
	type Person struct {
		Zip string `checkers:"len:5"`
	}

	person := &Person{
		Zip: "123",
	}

	errs, ok := v2.CheckStruct(person)
	if ok {
		t.Fatalf("expected errors")
	}

	if errs["Zip"] == nil {
		t.Fatalf("expected len error")
	}
}

func TestReflectLenInvalidLen(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Person struct {
		Zip string `checkers:"len:abcd"`
	}

	person := &Person{
		Zip: "12345",
	}

	v2.CheckStruct(person)
}

func TestReflectLenInvalidType(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Person struct {
		Age int `checkers:"len:2"`
	}

	person := &Person{
		Age: 1,
	}

	v2.CheckStruct(person)
}
