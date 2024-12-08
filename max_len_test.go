// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package checker_test

import (
	"errors"
	"log"
	"testing"

	"github.com/cinar/checker"
)

func TestMaxLenSuccess(t *testing.T) {
	value := "test"

	check := checker.MaxLen[string](4)

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

	check := checker.MaxLen[string](5)

	result, err := check(value)
	if result != value {
		t.Fatalf("result (%s) is not the original value (%s)", result, value)
	}

	if !errors.Is(err, checker.ErrMaxLen) {
		t.Fatalf("got unexpected error %v", err)
	}

	log.Println(err)
}

func TestReflectMaxLenError(t *testing.T) {
	type Person struct {
		Name string `checkers:"max-len:2"`
	}

	person := &Person{
		Name: "Onur",
	}

	errs, ok := checker.CheckStruct(person)
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

	checker.CheckStruct(person)
}

func TestReflectMaxLenInvalidType(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Person struct {
		Name int `checkers:"max-len:8"`
	}

	person := &Person{
		Name: 1,
	}

	checker.CheckStruct(person)
}
