// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2_test

import (
	"errors"
	"log"
	"testing"

	v2 "github.com/cinar/checker/v2"
)

func TestMinLenSuccess(t *testing.T) {
	value := "test"

	check := v2.MinLen[string](4)

	result, err := check(value)
	if result != value {
		t.Fatalf("result (%s) is not the original value (%s)", result, value)
	}

	if err != nil {
		t.Fatal(err)
	}
}

func TestMinLenError(t *testing.T) {
	value := "test"

	check := v2.MinLen[string](5)

	result, err := check(value)
	if result != value {
		t.Fatalf("result (%s) is not the original value (%s)", result, value)
	}

	if !errors.Is(err, v2.ErrMinLen) {
		t.Fatalf("got unexpected error %v", err)
	}

	log.Println(err)
}

func TestReflectMinLenError(t *testing.T) {
	type Person struct {
		Name string `checker:"trim min-len:8"`
	}

	person := &Person{
		Name: "    Onur    ",
	}

	errs, ok := v2.CheckStruct(person)
	if ok {
		t.Fatalf("expected errors")
	}

	if errs["Name"] == nil {
		t.Fatalf("expected minimum length error")
	}
}

func TestReflectMinLenInvalidMinLen(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Person struct {
		Name string `checker:"min-len:abcd"`
	}

	person := &Person{
		Name: "Onur",
	}

	v2.CheckStruct(person)
}

func TestReflectMinLenInvalidType(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Person struct {
		Name int `checker:"min-len:8"`
	}

	person := &Person{
		Name: 1,
	}

	v2.CheckStruct(person)
}
