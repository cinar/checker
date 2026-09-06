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

func TestIsAfterFieldSuccess(t *testing.T) {
	value, err := v2.IsAfterField("DateOnly", "2000-01-01", "2022-06-01")
	if err != nil {
		t.Fatal(err)
	}

	if value != "2022-06-01" {
		t.Fatalf("actual %s expected %s", value, "2022-06-01")
	}
}

func TestIsAfterFieldNotAfter(t *testing.T) {
	_, err := v2.IsAfterField("DateOnly", "2022-06-01", "2000-01-01")

	if !errors.Is(err, v2.ErrNotAfter) {
		t.Fatalf("got unexpected error %v", err)
	}
}

func TestIsAfterFieldEqual(t *testing.T) {
	_, err := v2.IsAfterField("DateOnly", "2000-01-01", "2000-01-01")

	if !errors.Is(err, v2.ErrNotAfter) {
		t.Fatalf("got unexpected error %v", err)
	}
}

func TestIsAfterFieldInvalidValue(t *testing.T) {
	_, err := v2.IsAfterField("DateOnly", "2000-01-01", "not-a-date")

	if !errors.Is(err, v2.ErrTime) {
		t.Fatalf("got unexpected error %v", err)
	}
}

func TestIsAfterFieldInvalidReference(t *testing.T) {
	_, err := v2.IsAfterField("DateOnly", "not-a-date", "2022-06-01")

	if !errors.Is(err, v2.ErrTime) {
		t.Fatalf("got unexpected error %v", err)
	}
}

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
		GraduatedAt string `checkers:"after-field:DateOnly:Unknown"`
	}

	person := &Person{
		GraduatedAt: "2022-06-01",
	}

	v2.CheckStruct(person)
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
