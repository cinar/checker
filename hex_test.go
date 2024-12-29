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

func ExampleIsHex() {
	_, err := v2.IsHex("0123456789abcdefABCDEF")
	if err != nil {
		fmt.Println(err)
	}
}

func TestIsHexInvalid(t *testing.T) {
	_, err := v2.IsHex("ONUR")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestIsHexValid(t *testing.T) {
	_, err := v2.IsHex("0123456789abcdefABCDEF")
	if err != nil {
		t.Fatal(err)
	}
}

func TestCheckHexNonString(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Car struct {
		Color int `checkers:"hex"`
	}

	car := &Car{}

	v2.CheckStruct(car)
}

func TestCheckHexInvalid(t *testing.T) {
	type Car struct {
		Color string `checkers:"hex"`
	}

	car := &Car{
		Color: "red",
	}

	_, ok := v2.CheckStruct(car)
	if ok {
		t.Fatal("expected error")
	}
}

func TestCheckHexValid(t *testing.T) {
	type Car struct {
		Color string `checkers:"hex"`
	}

	car := &Car{
		Color: "ABcd1234",
	}

	errs, ok := v2.CheckStruct(car)
	if !ok {
		t.Fatal(errs)
	}
}
