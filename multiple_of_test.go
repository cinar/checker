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

func TestMultipleOfSuccess(t *testing.T) {
	value := 9.0

	result, err := v2.IsMultipleOf(value, 3)
	if result != value {
		t.Fatalf("result (%f) is not the original value (%f)", result, value)
	}

	if err != nil {
		t.Fatal(err)
	}
}

func TestMultipleOfFloatToleranceSuccess(t *testing.T) {
	// 0.3 / 0.1 isn't an exact float64 division, but should still count as
	// a multiple within IsMultipleOf's tolerance.
	if _, err := v2.IsMultipleOf(0.3, 0.1); err != nil {
		t.Fatal(err)
	}
}

func TestMultipleOfError(t *testing.T) {
	_, err := v2.IsMultipleOf(10, 3)
	if err == nil {
		t.Fatal("expected error")
	}

	message := "Value must be a multiple of 3."

	if err.Error() != message {
		t.Fatalf("expected %s actual %s", message, err.Error())
	}
}

func TestReflectMultipleOfIntSuccess(t *testing.T) {
	type Item struct {
		Quantity int `checkers:"multiple-of:5"`
	}

	item := &Item{Quantity: 15}

	if _, ok := v2.CheckStruct(item); !ok {
		t.Fatal("expected valid")
	}
}

func TestReflectMultipleOfIntError(t *testing.T) {
	type Item struct {
		Quantity int `checkers:"multiple-of:5"`
	}

	item := &Item{Quantity: 12}

	errs, ok := v2.CheckStruct(item)
	if ok {
		t.Fatal("expected errors")
	}

	if !errors.Is(errs["Quantity"], v2.ErrMultipleOf) {
		t.Fatal("expected ErrMultipleOf")
	}
}

func TestReflectMultipleOfUintSuccess(t *testing.T) {
	type Item struct {
		Quantity uint `checkers:"multiple-of:5"`
	}

	item := &Item{Quantity: 15}

	if _, ok := v2.CheckStruct(item); !ok {
		t.Fatal("expected valid")
	}
}

func TestReflectMultipleOfFloatSuccess(t *testing.T) {
	type Item struct {
		Price float64 `checkers:"multiple-of:0.1"`
	}

	item := &Item{Price: 0.3}

	if _, ok := v2.CheckStruct(item); !ok {
		t.Fatal("expected valid")
	}
}

func TestReflectMultipleOfInvalidParams(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Item struct {
		Quantity int `checkers:"multiple-of:abcd"`
	}

	item := &Item{Quantity: 15}

	v2.CheckStruct(item)
}

func TestReflectMultipleOfZeroDivisor(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Item struct {
		Quantity int `checkers:"multiple-of:0"`
	}

	item := &Item{Quantity: 15}

	v2.CheckStruct(item)
}

func TestReflectMultipleOfInvalidType(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Item struct {
		Quantity string `checkers:"multiple-of:5"`
	}

	item := &Item{Quantity: "15"}

	v2.CheckStruct(item)
}
