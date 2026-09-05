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

func TestLtIntSuccess(t *testing.T) {
	value := 3

	result, err := v2.IsLt(value, 4)
	if result != value {
		t.Fatalf("result (%d) is not the original value (%d)", result, value)
	}

	if err != nil {
		t.Fatal(err)
	}
}

func TestLtIntEqualIsError(t *testing.T) {
	// lt is strict: a value equal to the bound must fail, unlike lte.
	value := 4

	_, err := v2.IsLt(value, 4)
	if err == nil {
		t.Fatal("expected error for equal value")
	}
}

func TestLtIntError(t *testing.T) {
	value := 6

	result, err := v2.IsLt(value, 5)
	if result != value {
		t.Fatalf("result (%d) is not the original value (%d)", result, value)
	}

	if err == nil {
		t.Fatal("expected error")
	}

	message := "Value must be less than 5."

	if err.Error() != message {
		t.Fatalf("expected %s actual %s", message, err.Error())
	}
}

func TestReflectLtIntError(t *testing.T) {
	type Person struct {
		Age int `checkers:"lt:18"`
	}

	person := &Person{
		Age: 18,
	}

	errs, ok := v2.CheckStruct(person)
	if ok {
		t.Fatalf("expected errors")
	}

	if !errors.Is(errs["Age"], v2.ErrLt) {
		t.Fatalf("expected ErrLt")
	}
}

func TestReflectLtIntSuccess(t *testing.T) {
	type Person struct {
		Age int `checkers:"lt:18"`
	}

	person := &Person{
		Age: 17,
	}

	if _, ok := v2.CheckStruct(person); !ok {
		t.Fatal("expected valid")
	}
}

func TestReflectLtIntInvalidLt(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Person struct {
		Age int `checkers:"lt:abcd"`
	}

	person := &Person{
		Age: 16,
	}

	v2.CheckStruct(person)
}

func TestReflectLtIntInvalidType(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Person struct {
		Age string `checkers:"lt:18"`
	}

	person := &Person{
		Age: "18",
	}

	v2.CheckStruct(person)
}

func TestReflectLtUintError(t *testing.T) {
	type Order struct {
		Quantity uint64 `checkers:"lt:10"`
	}

	order := &Order{
		Quantity: 10,
	}

	errs, ok := v2.CheckStruct(order)
	if ok {
		t.Fatalf("expected errors")
	}

	if !errors.Is(errs["Quantity"], v2.ErrLt) {
		t.Fatalf("expected ErrLt")
	}
}

func TestReflectLtUintSuccess(t *testing.T) {
	type Order struct {
		Quantity uint64 `checkers:"lt:10"`
	}

	order := &Order{
		Quantity: 5,
	}

	if _, ok := v2.CheckStruct(order); !ok {
		t.Fatal("expected valid")
	}
}

func TestReflectLtFloatError(t *testing.T) {
	type Person struct {
		Weight float64 `checkers:"lt:165.0"`
	}

	person := &Person{
		Weight: 170,
	}

	errs, ok := v2.CheckStruct(person)
	if ok {
		t.Fatalf("expected errors")
	}

	if !errors.Is(errs["Weight"], v2.ErrLt) {
		t.Fatalf("expected ErrLt")
	}
}

func TestReflectLtFloatInvalidLt(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Person struct {
		Weight float64 `checkers:"lt:abcd"`
	}

	person := &Person{
		Weight: 150,
	}

	v2.CheckStruct(person)
}

func TestReflectLtFloatInvalidType(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Person struct {
		Weight string `checkers:"lt:165.0"`
	}

	person := &Person{
		Weight: "150",
	}

	v2.CheckStruct(person)
}
