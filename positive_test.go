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

func TestPositiveSuccess(t *testing.T) {
	value := 5

	result, err := v2.IsPositive(value)
	if result != value {
		t.Fatalf("result (%d) is not the original value (%d)", result, value)
	}

	if err != nil {
		t.Fatal(err)
	}
}

func TestPositiveZeroIsError(t *testing.T) {
	_, err := v2.IsPositive(0)
	if err == nil {
		t.Fatal("expected error for zero value")
	}
}

func TestPositiveError(t *testing.T) {
	_, err := v2.IsPositive(-5)
	if err == nil {
		t.Fatal("expected error")
	}

	message := "Value must be positive."

	if err.Error() != message {
		t.Fatalf("expected %s actual %s", message, err.Error())
	}
}

func TestReflectPositiveIntSuccess(t *testing.T) {
	type Item struct {
		Quantity int `checkers:"positive"`
	}

	item := &Item{Quantity: 5}

	if _, ok := v2.CheckStruct(item); !ok {
		t.Fatal("expected valid")
	}
}

func TestReflectPositiveIntError(t *testing.T) {
	type Item struct {
		Quantity int `checkers:"positive"`
	}

	item := &Item{Quantity: -5}

	errs, ok := v2.CheckStruct(item)
	if ok {
		t.Fatal("expected errors")
	}

	if !errors.Is(errs["Quantity"], v2.ErrPositive) {
		t.Fatal("expected ErrPositive")
	}
}

func TestReflectPositiveUintSuccess(t *testing.T) {
	type Item struct {
		Quantity uint `checkers:"positive"`
	}

	item := &Item{Quantity: 5}

	if _, ok := v2.CheckStruct(item); !ok {
		t.Fatal("expected valid")
	}
}

func TestReflectPositiveUintError(t *testing.T) {
	type Item struct {
		Quantity uint `checkers:"positive"`
	}

	item := &Item{Quantity: 0}

	errs, ok := v2.CheckStruct(item)
	if ok {
		t.Fatal("expected errors")
	}

	if !errors.Is(errs["Quantity"], v2.ErrPositive) {
		t.Fatal("expected ErrPositive")
	}
}

func TestReflectPositiveFloatSuccess(t *testing.T) {
	type Item struct {
		Price float64 `checkers:"positive"`
	}

	item := &Item{Price: 5.5}

	if _, ok := v2.CheckStruct(item); !ok {
		t.Fatal("expected valid")
	}
}

func TestReflectPositiveFloatError(t *testing.T) {
	type Item struct {
		Price float64 `checkers:"positive"`
	}

	item := &Item{Price: -5.5}

	errs, ok := v2.CheckStruct(item)
	if ok {
		t.Fatal("expected errors")
	}

	if !errors.Is(errs["Price"], v2.ErrPositive) {
		t.Fatal("expected ErrPositive")
	}
}

func TestReflectPositiveInvalidType(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Item struct {
		Price string `checkers:"positive"`
	}

	item := &Item{Price: "5.5"}

	v2.CheckStruct(item)
}
