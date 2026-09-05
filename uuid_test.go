// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2_test

import (
	"fmt"
	"testing"

	v2 "github.com/cinar/checker/v2"
)

func ExampleIsUUID() {
	_, err := v2.IsUUID("f47ac10b-58cc-4372-a567-0e02b2c3d479")
	if err != nil {
		fmt.Println(err)
	}
}

func TestIsUUIDInvalid(t *testing.T) {
	_, err := v2.IsUUID("not-a-uuid")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestIsUUIDValid(t *testing.T) {
	_, err := v2.IsUUID("f47ac10b-58cc-4372-a567-0e02b2c3d479")
	if err != nil {
		t.Fatal(err)
	}
}

func TestIsUUIDValidUppercase(t *testing.T) {
	_, err := v2.IsUUID("F47AC10B-58CC-4372-A567-0E02B2C3D479")
	if err != nil {
		t.Fatal(err)
	}
}

func TestIsUUIDInvalidWrongGrouping(t *testing.T) {
	_, err := v2.IsUUID("f47ac10b58cc4372a5670e02b2c3d479")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestCheckUUIDNonString(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Resource struct {
		ID int `checkers:"uuid"`
	}

	resource := &Resource{}

	v2.CheckStruct(resource)
}

func TestCheckUUIDInvalid(t *testing.T) {
	type Resource struct {
		ID string `checkers:"uuid"`
	}

	resource := &Resource{
		ID: "not-a-uuid",
	}

	_, ok := v2.CheckStruct(resource)
	if ok {
		t.Fatal("expected error")
	}
}

func TestCheckUUIDValid(t *testing.T) {
	type Resource struct {
		ID string `checkers:"uuid"`
	}

	resource := &Resource{
		ID: "f47ac10b-58cc-4372-a567-0e02b2c3d479",
	}

	errs, ok := v2.CheckStruct(resource)
	if !ok {
		t.Fatal(errs)
	}
}
