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

func TestNonnegativeSuccess(t *testing.T) {
	value := 0

	result, err := v2.IsNonnegative(value)
	if result != value {
		t.Fatalf("result (%d) is not the original value (%d)", result, value)
	}

	if err != nil {
		t.Fatal(err)
	}
}

func TestNonnegativeError(t *testing.T) {
	_, err := v2.IsNonnegative(-1)
	if err == nil {
		t.Fatal("expected error")
	}

	message := "Value cannot be negative."

	if err.Error() != message {
		t.Fatalf("expected %s actual %s", message, err.Error())
	}
}

func TestReflectNonnegativeIntSuccess(t *testing.T) {
	type Item struct {
		Balance int `checkers:"nonnegative"`
	}

	item := &Item{Balance: 0}

	if _, ok := v2.CheckStruct(item); !ok {
		t.Fatal("expected valid")
	}
}

func TestReflectNonnegativeIntError(t *testing.T) {
	type Item struct {
		Balance int `checkers:"nonnegative"`
	}

	item := &Item{Balance: -1}

	errs, ok := v2.CheckStruct(item)
	if ok {
		t.Fatal("expected errors")
	}

	if !errors.Is(errs["Balance"], v2.ErrNonnegative) {
		t.Fatal("expected ErrNonnegative")
	}
}

func TestReflectNonnegativeUintSuccess(t *testing.T) {
	type Item struct {
		Balance uint `checkers:"nonnegative"`
	}

	item := &Item{Balance: 0}

	if _, ok := v2.CheckStruct(item); !ok {
		t.Fatal("expected valid")
	}
}

func TestReflectNonnegativeFloatSuccess(t *testing.T) {
	type Item struct {
		Balance float64 `checkers:"nonnegative"`
	}

	item := &Item{Balance: 0}

	if _, ok := v2.CheckStruct(item); !ok {
		t.Fatal("expected valid")
	}
}

func TestReflectNonnegativeFloatError(t *testing.T) {
	type Item struct {
		Balance float64 `checkers:"nonnegative"`
	}

	item := &Item{Balance: -0.5}

	errs, ok := v2.CheckStruct(item)
	if ok {
		t.Fatal("expected errors")
	}

	if !errors.Is(errs["Balance"], v2.ErrNonnegative) {
		t.Fatal("expected ErrNonnegative")
	}
}

func TestReflectNonnegativeInvalidType(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Item struct {
		Balance string `checkers:"nonnegative"`
	}

	item := &Item{Balance: "-0.5"}

	v2.CheckStruct(item)
}
