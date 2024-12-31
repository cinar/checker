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

func TestLteIntSuccess(t *testing.T) {
	value := 4

	result, err := v2.IsLte(value, 4)
	if result != value {
		t.Fatalf("result (%d) is not the original value (%d)", result, value)
	}

	if err != nil {
		t.Fatal(err)
	}
}

func TestLteIntError(t *testing.T) {
	value := 6

	result, err := v2.IsLte(value, 5)
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

func TestReflectLteIntError(t *testing.T) {
	type Person struct {
		Age int `checkers:"lte:18"`
	}

	person := &Person{
		Age: 21,
	}

	errs, ok := v2.CheckStruct(person)
	if ok {
		t.Fatalf("expected errors")
	}

	if !errors.Is(errs["Age"], v2.ErrLte) {
		t.Fatalf("expected ErrLte")
	}
}

func TestReflectLteIntInvalidLte(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Person struct {
		Age int `checkers:"lte:abcd"`
	}

	person := &Person{
		Age: 16,
	}

	v2.CheckStruct(person)
}

func TestReflectLteIntInvalidType(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Person struct {
		Age string `checkers:"lte:18"`
	}

	person := &Person{
		Age: "18",
	}

	v2.CheckStruct(person)
}

func TestReflectLteFloatError(t *testing.T) {
	type Person struct {
		Weight float64 `checkers:"lte:165.0"`
	}

	person := &Person{
		Weight: 170,
	}

	errs, ok := v2.CheckStruct(person)
	if ok {
		t.Fatalf("expected errors")
	}

	if !errors.Is(errs["Weight"], v2.ErrLte) {
		t.Fatalf("expected ErrLte")
	}
}

func TestReflectLteFloatInvalidLte(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Person struct {
		Weight float64 `checkers:"lte:abcd"`
	}

	person := &Person{
		Weight: 170,
	}

	v2.CheckStruct(person)
}

func TestReflectLteFloatInvalidType(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Person struct {
		Weight string `checkers:"lte:165.0"`
	}

	person := &Person{
		Weight: "170",
	}

	v2.CheckStruct(person)
}
