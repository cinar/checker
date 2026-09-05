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

func TestCheckStructRequiredIfConditionMetSuccess(t *testing.T) {
	type Address struct {
		Country string `checkers:"required"`
		State   string `checkers:"required-if:Country:US"`
	}

	address := &Address{
		Country: "US",
		State:   "CA",
	}

	errs, ok := v2.CheckStruct(address)
	if !ok {
		t.Fatalf("got unexpected errors %v", errs)
	}
}

func TestCheckStructRequiredIfConditionMetMissing(t *testing.T) {
	type Address struct {
		Country string `checkers:"required"`
		State   string `checkers:"required-if:Country:US"`
	}

	address := &Address{
		Country: "US",
	}

	errs, ok := v2.CheckStruct(address)
	if ok {
		t.Fatal("expected errors")
	}

	if !errors.Is(errs["State"], v2.ErrRequired) {
		t.Fatalf("expected state required %v", errs)
	}
}

func TestCheckStructRequiredIfConditionNotMet(t *testing.T) {
	type Address struct {
		Country string `checkers:"required"`
		State   string `checkers:"required-if:Country:US"`
	}

	address := &Address{
		Country: "CA",
	}

	errs, ok := v2.CheckStruct(address)
	if !ok {
		t.Fatalf("got unexpected errors %v", errs)
	}
}

func TestMakeRequiredIfMissingValue(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Address struct {
		Country string
		State   string `checkers:"required-if:Country"`
	}

	address := &Address{
		Country: "US",
	}

	v2.CheckStruct(address)
}

func TestRequiredIfMissingField(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Address struct {
		Country string
	}

	address := Address{Country: "US"}

	v2.IsRequiredIf(reflect.ValueOf(address), reflect.ValueOf(""), "Unknown", "US")
}
