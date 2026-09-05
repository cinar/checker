// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2_test

import (
	"errors"
	"reflect"
	"testing"

	v2 "github.com/cinar/checker/v2"
)

func TestCheckStructAfterFieldSuccess(t *testing.T) {
	type Person struct {
		BornAt      string `checkers:"required"`
		GraduatedAt string `checkers:"after-field:DateOnly:BornAt"`
	}

	person := &Person{
		BornAt:      "2000-01-01",
		GraduatedAt: "2022-06-01",
	}

	errs, ok := v2.CheckStruct(person)
	if !ok {
		t.Fatalf("got unexpected errors %v", errs)
	}
}

func TestCheckStructAfterFieldNotAfter(t *testing.T) {
	type Person struct {
		BornAt      string `checkers:"required"`
		GraduatedAt string `checkers:"after-field:DateOnly:BornAt"`
	}

	person := &Person{
		BornAt:      "2022-06-01",
		GraduatedAt: "2000-01-01",
	}

	errs, ok := v2.CheckStruct(person)
	if ok {
		t.Fatal("expected errors")
	}

	if !errors.Is(errs["GraduatedAt"], v2.ErrNotAfter) {
		t.Fatalf("expected graduated at not after %v", errs)
	}
}

func TestCheckStructAfterFieldEqual(t *testing.T) {
	type Person struct {
		BornAt      string `checkers:"required"`
		GraduatedAt string `checkers:"after-field:DateOnly:BornAt"`
	}

	person := &Person{
		BornAt:      "2000-01-01",
		GraduatedAt: "2000-01-01",
	}

	errs, ok := v2.CheckStruct(person)
	if ok {
		t.Fatal("expected errors")
	}

	if !errors.Is(errs["GraduatedAt"], v2.ErrNotAfter) {
		t.Fatalf("expected graduated at not after %v", errs)
	}
}

func TestCheckStructAfterFieldInvalidValue(t *testing.T) {
	type Person struct {
		BornAt      string `checkers:"required"`
		GraduatedAt string `checkers:"after-field:DateOnly:BornAt"`
	}

	person := &Person{
		BornAt:      "2000-01-01",
		GraduatedAt: "not-a-date",
	}

	errs, ok := v2.CheckStruct(person)
	if ok {
		t.Fatal("expected errors")
	}

	if !errors.Is(errs["GraduatedAt"], v2.ErrTime) {
		t.Fatalf("expected graduated at not a time %v", errs)
	}
}

func TestCheckStructAfterFieldInvalidReference(t *testing.T) {
	type Person struct {
		BornAt      string
		GraduatedAt string `checkers:"after-field:DateOnly:BornAt"`
	}

	person := &Person{
		BornAt:      "not-a-date",
		GraduatedAt: "2022-06-01",
	}

	errs, ok := v2.CheckStruct(person)
	if ok {
		t.Fatal("expected errors")
	}

	if !errors.Is(errs["GraduatedAt"], v2.ErrTime) {
		t.Fatalf("expected graduated at not a time %v", errs)
	}
}

func TestAfterFieldMissingField(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Person struct {
		GraduatedAt string
	}

	person := Person{GraduatedAt: "2022-06-01"}

	v2.IsAfterField(reflect.ValueOf(person), reflect.ValueOf("2022-06-01"), "DateOnly", "Unknown")
}

func TestMakeAfterFieldMissingFieldName(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Person struct {
		BornAt      string
		GraduatedAt string `checkers:"after-field:DateOnly"`
	}

	person := &Person{
		BornAt:      "2000-01-01",
		GraduatedAt: "2022-06-01",
	}

	v2.CheckStruct(person)
}
