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

func ExampleIsDigits() {
	_, err := v2.IsDigits("123456")
	if err != nil {
		fmt.Println(err)
	}
}

func TestIsDigitsInvalid(t *testing.T) {
	_, err := v2.IsDigits("123a456")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestIsDigitsValid(t *testing.T) {
	_, err := v2.IsDigits("123456")
	if err != nil {
		t.Fatal(err)
	}
}

func TestCheckDigitsNonString(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Code struct {
		Value int `checkers:"digits"`
	}

	code := &Code{}

	v2.CheckStruct(code)
}

func TestCheckDigitsInvalid(t *testing.T) {
	type Code struct {
		Value string `checkers:"digits"`
	}

	code := &Code{
		Value: "123a456",
	}

	_, ok := v2.CheckStruct(code)
	if ok {
		t.Fatal("expected error")
	}
}

func TestCheckDigitsValid(t *testing.T) {
	type Code struct {
		Value string `checkers:"digits"`
	}

	code := &Code{
		Value: "123456",
	}

	_, ok := v2.CheckStruct(code)
	if !ok {
		t.Fatal("expected valid")
	}
}
