// Copyright (c) 2023-2024 Onur Cinar.
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
		DepartAt string
	}

	trip := Trip{DepartAt: "2024-01-01"}

	v2.IsBeforeField(reflect.ValueOf(trip), reflect.ValueOf("2024-01-01"), "DateOnly", "Unknown")
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
