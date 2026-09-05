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

func TestIsAfterSuccess(t *testing.T) {
	value, err := v2.IsAfter("DateOnly", "2024-01-01", "2024-06-01")
	if err != nil {
		t.Fatal(err)
	}

	if value != "2024-06-01" {
		t.Fatalf("actual %s expected %s", value, "2024-06-01")
	}
}

func TestIsAfterNotAfter(t *testing.T) {
	_, err := v2.IsAfter("DateOnly", "2024-06-01", "2024-01-01")

	if !errors.Is(err, v2.ErrNotAfter) {
		t.Fatalf("got unexpected error %v", err)
	}
}

func TestIsAfterEqual(t *testing.T) {
	_, err := v2.IsAfter("DateOnly", "2024-01-01", "2024-01-01")

	if !errors.Is(err, v2.ErrNotAfter) {
		t.Fatalf("got unexpected error %v", err)
	}
}

func TestIsAfterInvalidValue(t *testing.T) {
	_, err := v2.IsAfter("DateOnly", "2024-01-01", "not-a-date")

	if !errors.Is(err, v2.ErrTime) {
		t.Fatalf("got unexpected error %v", err)
	}
}

func TestIsAfterInvalidReference(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	v2.IsAfter("DateOnly", "not-a-date", "2024-01-01")
}

func TestIsAfterLiteralLayout(t *testing.T) {
	value, err := v2.IsAfter("2006-01-02", "2024-01-01", "2024-06-01")
	if err != nil {
		t.Fatal(err)
	}

	if value != "2024-06-01" {
		t.Fatalf("actual %s expected %s", value, "2024-06-01")
	}
}

func TestMakeAfterMissingReference(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Person struct {
		BornAt string `checkers:"after:DateOnly"`
	}

	person := &Person{
		BornAt: "2024-01-01",
	}

	v2.CheckStruct(person)
}

func TestCheckStructAfter(t *testing.T) {
	type Person struct {
		BornAt string `checkers:"after:DateOnly:2000-01-01"`
	}

	person := &Person{
		BornAt: "1999-01-01",
	}

	errs, ok := v2.CheckStruct(person)
	if ok {
		t.Fatal("expected errors")
	}

	if !errors.Is(errs["BornAt"], v2.ErrNotAfter) {
		t.Fatalf("expected born at not after %v", errs)
	}

	person.BornAt = "2001-01-01"

	errs, ok = v2.CheckStruct(person)
	if !ok {
		t.Fatalf("got unexpected errors %v", errs)
	}
}
