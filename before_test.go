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

func TestIsBeforeSuccess(t *testing.T) {
	value, err := v2.IsBefore("DateOnly", "2024-06-01", "2024-01-01")
	if err != nil {
		t.Fatal(err)
	}

	if value != "2024-01-01" {
		t.Fatalf("actual %s expected %s", value, "2024-01-01")
	}
}

func TestIsBeforeNotBefore(t *testing.T) {
	_, err := v2.IsBefore("DateOnly", "2024-01-01", "2024-06-01")

	if !errors.Is(err, v2.ErrNotBefore) {
		t.Fatalf("got unexpected error %v", err)
	}
}

func TestIsBeforeEqual(t *testing.T) {
	_, err := v2.IsBefore("DateOnly", "2024-01-01", "2024-01-01")

	if !errors.Is(err, v2.ErrNotBefore) {
		t.Fatalf("got unexpected error %v", err)
	}
}

func TestIsBeforeInvalidValue(t *testing.T) {
	_, err := v2.IsBefore("DateOnly", "2024-01-01", "not-a-date")

	if !errors.Is(err, v2.ErrTime) {
		t.Fatalf("got unexpected error %v", err)
	}
}

func TestIsBeforeInvalidReference(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	v2.IsBefore("DateOnly", "not-a-date", "2024-01-01")
}

func TestIsBeforeLiteralLayout(t *testing.T) {
	value, err := v2.IsBefore("2006-01-02", "2024-06-01", "2024-01-01")
	if err != nil {
		t.Fatal(err)
	}

	if value != "2024-01-01" {
		t.Fatalf("actual %s expected %s", value, "2024-01-01")
	}
}

func TestMakeBeforeMissingReference(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Coupon struct {
		ExpiresAt string `checkers:"before:DateOnly"`
	}

	coupon := &Coupon{
		ExpiresAt: "2024-01-01",
	}

	v2.CheckStruct(coupon)
}

func TestCheckStructBefore(t *testing.T) {
	type Coupon struct {
		ExpiresAt string `checkers:"before:DateOnly:2024-06-01"`
	}

	coupon := &Coupon{
		ExpiresAt: "2024-12-01",
	}

	errs, ok := v2.CheckStruct(coupon)
	if ok {
		t.Fatal("expected errors")
	}

	if !errors.Is(errs["ExpiresAt"], v2.ErrNotBefore) {
		t.Fatalf("expected expires at not before %v", errs)
	}

	coupon.ExpiresAt = "2024-01-01"

	errs, ok = v2.CheckStruct(coupon)
	if !ok {
		t.Fatalf("got unexpected errors %v", errs)
	}
}
