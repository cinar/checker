// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2_test

import (
	"errors"
	"testing"

	v2 "github.com/cinar/checker/v2"
)

func TestGteIntSuccess(t *testing.T) {
	value := 4

	result, err := v2.IsGte(value, 4)
	if result != value {
		t.Fatalf("result (%d) is not the original value (%d)", result, value)
	}

	if err != nil {
		t.Fatal(err)
	}
}

func TestGteIntError(t *testing.T) {
	value := 4

	result, err := v2.IsGte(value, 5)
	if result != value {
		t.Fatalf("result (%d) is not the original value (%d)", result, value)
	}

	if err == nil {
		t.Fatal("expected error")
	}

	message := "Value cannot be less than 5."

	if err.Error() != message {
		t.Fatalf("expected %s actual %s", message, err.Error())
	}
}

func TestReflectGteIntError(t *testing.T) {
	type Person struct {
		Age int `checkers:"gte:18"`
	}

	person := &Person{
		Age: 16,
	}

	errs, ok := v2.CheckStruct(person)
	if ok {
		t.Fatalf("expected errors")
	}

	if !errors.Is(errs["Age"], v2.ErrGte) {
		t.Fatalf("expected ErrGte")
	}
}

func TestReflectGteIntInvalidGte(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Person struct {
		Age int `checkers:"gte:abcd"`
	}

	person := &Person{
		Age: 16,
	}

	v2.CheckStruct(person)
}

func TestReflectGteIntInvalidType(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Person struct {
		Age string `checkers:"gte:18"`
	}

	person := &Person{
		Age: "18",
	}

	v2.CheckStruct(person)
}

func TestReflectGteFloatError(t *testing.T) {
	type Person struct {
		Weight float64 `checkers:"gte:165.0"`
	}

	person := &Person{
		Weight: 150,
	}

	errs, ok := v2.CheckStruct(person)
	if ok {
		t.Fatalf("expected errors")
	}

	if !errors.Is(errs["Weight"], v2.ErrGte) {
		t.Fatalf("expected ErrGte")
	}
}

func TestReflectGteFloatInvalidGte(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Person struct {
		Weight float64 `checkers:"gte:abcd"`
	}

	person := &Person{
		Weight: 170,
	}

	v2.CheckStruct(person)
}

func TestReflectGteFloatInvalidType(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Person struct {
		Weight string `checkers:"gte:165.0"`
	}

	person := &Person{
		Weight: "170",
	}

	v2.CheckStruct(person)
}
