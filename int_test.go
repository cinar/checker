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

func TestIntSuccess(t *testing.T) {
	value := 4.0

	result, err := v2.IsInt(value)
	if result != value {
		t.Fatalf("result (%f) is not the original value (%f)", result, value)
	}

	if err != nil {
		t.Fatal(err)
	}
}

func TestIntError(t *testing.T) {
	value := 4.5

	_, err := v2.IsInt(value)
	if err == nil {
		t.Fatal("expected error")
	}

	message := "Value must be a whole number."

	if err.Error() != message {
		t.Fatalf("expected %s actual %s", message, err.Error())
	}
}

func TestReflectIntFloatSuccess(t *testing.T) {
	type Item struct {
		Quantity float64 `checkers:"int"`
	}

	item := &Item{Quantity: 4}

	if _, ok := v2.CheckStruct(item); !ok {
		t.Fatal("expected valid")
	}
}

func TestReflectIntFloatError(t *testing.T) {
	type Item struct {
		Quantity float64 `checkers:"int"`
	}

	item := &Item{Quantity: 4.5}

	errs, ok := v2.CheckStruct(item)
	if ok {
		t.Fatal("expected errors")
	}

	if !errors.Is(errs["Quantity"], v2.ErrInt) {
		t.Fatal("expected ErrInt")
	}
}

func TestReflectIntIntSuccess(t *testing.T) {
	type Item struct {
		Quantity int `checkers:"int"`
	}

	item := &Item{Quantity: 4}

	if _, ok := v2.CheckStruct(item); !ok {
		t.Fatal("expected valid")
	}
}

func TestReflectIntUintSuccess(t *testing.T) {
	type Item struct {
		Quantity uint `checkers:"int"`
	}

	item := &Item{Quantity: 4}

	if _, ok := v2.CheckStruct(item); !ok {
		t.Fatal("expected valid")
	}
}

func TestReflectIntInvalidType(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Item struct {
		Quantity string `checkers:"int"`
	}

	item := &Item{Quantity: "4"}

	v2.CheckStruct(item)
}
