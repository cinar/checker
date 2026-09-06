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

func TestNegativeSuccess(t *testing.T) {
	value := -5

	result, err := v2.IsNegative(value)
	if result != value {
		t.Fatalf("result (%d) is not the original value (%d)", result, value)
	}

	if err != nil {
		t.Fatal(err)
	}
}

func TestNegativeZeroIsError(t *testing.T) {
	_, err := v2.IsNegative(0)
	if err == nil {
		t.Fatal("expected error for zero value")
	}
}

func TestNegativeError(t *testing.T) {
	_, err := v2.IsNegative(5)
	if err == nil {
		t.Fatal("expected error")
	}

	message := "Value must be negative."

	if err.Error() != message {
		t.Fatalf("expected %s actual %s", message, err.Error())
	}
}

func TestReflectNegativeIntSuccess(t *testing.T) {
	type Item struct {
		Delta int `checkers:"negative"`
	}

	item := &Item{Delta: -5}

	if _, ok := v2.CheckStruct(item); !ok {
		t.Fatal("expected valid")
	}
}

func TestReflectNegativeIntError(t *testing.T) {
	type Item struct {
		Delta int `checkers:"negative"`
	}

	item := &Item{Delta: 5}

	errs, ok := v2.CheckStruct(item)
	if ok {
		t.Fatal("expected errors")
	}

	if !errors.Is(errs["Delta"], v2.ErrNegative) {
		t.Fatal("expected ErrNegative")
	}
}

func TestReflectNegativeUintAlwaysError(t *testing.T) {
	type Item struct {
		Delta uint `checkers:"negative"`
	}

	item := &Item{Delta: 0}

	errs, ok := v2.CheckStruct(item)
	if ok {
		t.Fatal("expected errors")
	}

	if !errors.Is(errs["Delta"], v2.ErrNegative) {
		t.Fatal("expected ErrNegative")
	}
}

func TestReflectNegativeFloatSuccess(t *testing.T) {
	type Item struct {
		Delta float64 `checkers:"negative"`
	}

	item := &Item{Delta: -5.5}

	if _, ok := v2.CheckStruct(item); !ok {
		t.Fatal("expected valid")
	}
}

func TestReflectNegativeFloatError(t *testing.T) {
	type Item struct {
		Delta float64 `checkers:"negative"`
	}

	item := &Item{Delta: 5.5}

	errs, ok := v2.CheckStruct(item)
	if ok {
		t.Fatal("expected errors")
	}

	if !errors.Is(errs["Delta"], v2.ErrNegative) {
		t.Fatal("expected ErrNegative")
	}
}

func TestReflectNegativeInvalidType(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Item struct {
		Delta string `checkers:"negative"`
	}

	item := &Item{Delta: "-5.5"}

	v2.CheckStruct(item)
}
