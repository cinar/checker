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

func TestIsBeforeFieldSuccess(t *testing.T) {
	value, err := v2.IsBeforeField("DateOnly", "2024-06-01", "2024-01-01")
	if err != nil {
		t.Fatal(err)
	}

	if value != "2024-01-01" {
		t.Fatalf("actual %s expected %s", value, "2024-01-01")
	}
}

func TestIsBeforeFieldNotBefore(t *testing.T) {
	_, err := v2.IsBeforeField("DateOnly", "2024-01-01", "2024-06-01")

	if !errors.Is(err, v2.ErrNotBefore) {
		t.Fatalf("got unexpected error %v", err)
	}
}

func TestIsBeforeFieldEqual(t *testing.T) {
	_, err := v2.IsBeforeField("DateOnly", "2024-01-01", "2024-01-01")

	if !errors.Is(err, v2.ErrNotBefore) {
		t.Fatalf("got unexpected error %v", err)
	}
}

func TestIsBeforeFieldInvalidValue(t *testing.T) {
	_, err := v2.IsBeforeField("DateOnly", "2024-06-01", "not-a-date")

	if !errors.Is(err, v2.ErrTime) {
		t.Fatalf("got unexpected error %v", err)
	}
}

func TestIsBeforeFieldInvalidReference(t *testing.T) {
	_, err := v2.IsBeforeField("DateOnly", "not-a-date", "2024-01-01")

	if !errors.Is(err, v2.ErrTime) {
		t.Fatalf("got unexpected error %v", err)
	}
}

func TestCheckStructBeforeFieldSuccess(t *testing.T) {
	type Trip struct {
		ReturnAt string `checkers:"required"`
		DepartAt string `checkers:"before-field:DateOnly:ReturnAt"`
	}

	trip := &Trip{
		ReturnAt: "2024-06-01",
		DepartAt: "2024-01-01",
	}

	errs, ok := v2.CheckStruct(trip)
	if !ok {
		t.Fatalf("got unexpected errors %v", errs)
	}
}

func TestCheckStructBeforeFieldNotBefore(t *testing.T) {
	type Trip struct {
		ReturnAt string `checkers:"required"`
		DepartAt string `checkers:"before-field:DateOnly:ReturnAt"`
	}

	trip := &Trip{
		ReturnAt: "2024-01-01",
		DepartAt: "2024-06-01",
	}

	errs, ok := v2.CheckStruct(trip)
	if ok {
		t.Fatal("expected errors")
	}

	if !errors.Is(errs["DepartAt"], v2.ErrNotBefore) {
		t.Fatalf("expected depart at not before %v", errs)
	}
}

func TestCheckStructBeforeFieldEqual(t *testing.T) {
	type Trip struct {
		ReturnAt string `checkers:"required"`
		DepartAt string `checkers:"before-field:DateOnly:ReturnAt"`
	}

	trip := &Trip{
		ReturnAt: "2024-01-01",
		DepartAt: "2024-01-01",
	}

	errs, ok := v2.CheckStruct(trip)
	if ok {
		t.Fatal("expected errors")
	}

	if !errors.Is(errs["DepartAt"], v2.ErrNotBefore) {
		t.Fatalf("expected depart at not before %v", errs)
	}
}

func TestCheckStructBeforeFieldInvalidValue(t *testing.T) {
	type Trip struct {
		ReturnAt string `checkers:"required"`
		DepartAt string `checkers:"before-field:DateOnly:ReturnAt"`
	}

	trip := &Trip{
		ReturnAt: "2024-06-01",
		DepartAt: "not-a-date",
	}

	errs, ok := v2.CheckStruct(trip)
	if ok {
		t.Fatal("expected errors")
	}

	if !errors.Is(errs["DepartAt"], v2.ErrTime) {
		t.Fatalf("expected depart at not a time %v", errs)
	}
}

func TestCheckStructBeforeFieldInvalidReference(t *testing.T) {
	type Trip struct {
		ReturnAt string
		DepartAt string `checkers:"before-field:DateOnly:ReturnAt"`
	}

	trip := &Trip{
		ReturnAt: "not-a-date",
		DepartAt: "2024-01-01",
	}

	errs, ok := v2.CheckStruct(trip)
	if ok {
		t.Fatal("expected errors")
	}

	if !errors.Is(errs["DepartAt"], v2.ErrTime) {
		t.Fatalf("expected depart at not a time %v", errs)
	}
}

func TestBeforeFieldMissingField(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Trip struct {
		DepartAt string `checkers:"before-field:DateOnly:Unknown"`
	}

	trip := &Trip{
		DepartAt: "2024-01-01",
	}

	v2.CheckStruct(trip)
}

func TestMakeBeforeFieldMissingFieldName(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Trip struct {
		ReturnAt string
		DepartAt string `checkers:"before-field:DateOnly"`
	}

	trip := &Trip{
		ReturnAt: "2024-06-01",
		DepartAt: "2024-01-01",
	}

	v2.CheckStruct(trip)
}
