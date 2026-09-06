// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2_test

import (
	"errors"
	"math"
	"testing"

	v2 "github.com/cinar/checker/v2"
)

func TestFiniteSuccess(t *testing.T) {
	value := 5.5

	result, err := v2.IsFinite(value)
	if result != value {
		t.Fatalf("result (%f) is not the original value (%f)", result, value)
	}

	if err != nil {
		t.Fatal(err)
	}
}

func TestFiniteNaNIsError(t *testing.T) {
	_, err := v2.IsFinite(math.NaN())
	if err == nil {
		t.Fatal("expected error for NaN")
	}

	message := "Value must be a finite number."

	if err.Error() != message {
		t.Fatalf("expected %s actual %s", message, err.Error())
	}
}

func TestFiniteInfIsError(t *testing.T) {
	_, err := v2.IsFinite(math.Inf(1))
	if err == nil {
		t.Fatal("expected error for +Inf")
	}

	_, err = v2.IsFinite(math.Inf(-1))
	if err == nil {
		t.Fatal("expected error for -Inf")
	}
}

func TestReflectFiniteFloatSuccess(t *testing.T) {
	type Item struct {
		Price float64 `checkers:"finite"`
	}

	item := &Item{Price: 5.5}

	if _, ok := v2.CheckStruct(item); !ok {
		t.Fatal("expected valid")
	}
}

func TestReflectFiniteFloatError(t *testing.T) {
	type Item struct {
		Price float64 `checkers:"finite"`
	}

	item := &Item{Price: math.Inf(1)}

	errs, ok := v2.CheckStruct(item)
	if ok {
		t.Fatal("expected errors")
	}

	if !errors.Is(errs["Price"], v2.ErrFinite) {
		t.Fatal("expected ErrFinite")
	}
}

func TestReflectFiniteIntSuccess(t *testing.T) {
	type Item struct {
		Quantity int `checkers:"finite"`
	}

	item := &Item{Quantity: 5}

	if _, ok := v2.CheckStruct(item); !ok {
		t.Fatal("expected valid")
	}
}

func TestReflectFiniteUintSuccess(t *testing.T) {
	type Item struct {
		Quantity uint `checkers:"finite"`
	}

	item := &Item{Quantity: 5}

	if _, ok := v2.CheckStruct(item); !ok {
		t.Fatal("expected valid")
	}
}

func TestReflectFiniteInvalidType(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Item struct {
		Price string `checkers:"finite"`
	}

	item := &Item{Price: "5.5"}

	v2.CheckStruct(item)
}
