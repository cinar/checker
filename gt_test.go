// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2_test

import (
	"errors"
	"testing"

	v2 "github.com/cinar/checker/v2"
)

func TestGtIntSuccess(t *testing.T) {
	value := 5

	result, err := v2.IsGt(value, 4)
	if result != value {
		t.Fatalf("result (%d) is not the original value (%d)", result, value)
	}

	if err != nil {
		t.Fatal(err)
	}
}

func TestGtIntEqualIsError(t *testing.T) {
	// gt is strict: a value equal to the bound must fail, unlike gte.
	value := 4

	_, err := v2.IsGt(value, 4)
	if err == nil {
		t.Fatal("expected error for equal value")
	}
}

func TestGtIntError(t *testing.T) {
	value := 4

	result, err := v2.IsGt(value, 5)
	if result != value {
		t.Fatalf("result (%d) is not the original value (%d)", result, value)
	}

	if err == nil {
		t.Fatal("expected error")
	}

	message := "Value must be greater than 5."

	if err.Error() != message {
		t.Fatalf("expected %s actual %s", message, err.Error())
	}
}

func TestReflectGtIntError(t *testing.T) {
	type Person struct {
		Age int `checkers:"gt:18"`
	}

	person := &Person{
		Age: 18,
	}

	errs, ok := v2.CheckStruct(person)
	if ok {
		t.Fatalf("expected errors")
	}

	if !errors.Is(errs["Age"], v2.ErrGt) {
		t.Fatalf("expected ErrGt")
	}
}

func TestReflectGtIntSuccess(t *testing.T) {
	type Person struct {
		Age int `checkers:"gt:18"`
	}

	person := &Person{
		Age: 19,
	}

	if _, ok := v2.CheckStruct(person); !ok {
		t.Fatal("expected valid")
	}
}

func TestReflectGtIntInvalidGt(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Person struct {
		Age int `checkers:"gt:abcd"`
	}

	person := &Person{
		Age: 16,
	}

	v2.CheckStruct(person)
}

func TestReflectGtIntInvalidType(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Person struct {
		Age string `checkers:"gt:18"`
	}

	person := &Person{
		Age: "18",
	}

	v2.CheckStruct(person)
}

func TestReflectGtUintError(t *testing.T) {
	type Order struct {
		Quantity uint64 `checkers:"gt:1"`
	}

	order := &Order{
		Quantity: 1,
	}

	errs, ok := v2.CheckStruct(order)
	if ok {
		t.Fatalf("expected errors")
	}

	if !errors.Is(errs["Quantity"], v2.ErrGt) {
		t.Fatalf("expected ErrGt")
	}
}

func TestReflectGtUintSuccess(t *testing.T) {
	type Order struct {
		Quantity uint64 `checkers:"gt:1"`
	}

	order := &Order{
		Quantity: 5,
	}

	if _, ok := v2.CheckStruct(order); !ok {
		t.Fatal("expected valid")
	}
}

func TestReflectGtFloatError(t *testing.T) {
	type Person struct {
		Weight float64 `checkers:"gt:165.0"`
	}

	person := &Person{
		Weight: 150,
	}

	errs, ok := v2.CheckStruct(person)
	if ok {
		t.Fatalf("expected errors")
	}

	if !errors.Is(errs["Weight"], v2.ErrGt) {
		t.Fatalf("expected ErrGt")
	}
}

func TestReflectGtFloatInvalidGt(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Person struct {
		Weight float64 `checkers:"gt:abcd"`
	}

	person := &Person{
		Weight: 170,
	}

	v2.CheckStruct(person)
}

func TestReflectGtFloatInvalidType(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Person struct {
		Weight string `checkers:"gt:165.0"`
	}

	person := &Person{
		Weight: "170",
	}

	v2.CheckStruct(person)
}
