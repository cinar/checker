// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2_test

import (
	"fmt"
	"testing"

	v2 "github.com/cinar/checker/v2"
)

func ExampleIsHexColor() {
	_, err := v2.IsHexColor("#ff0000")
	if err != nil {
		fmt.Println(err)
	}
}

func TestIsHexColorInvalid(t *testing.T) {
	_, err := v2.IsHexColor("ff0000")
	if err == nil {
		t.Fatal("expected error for missing '#'")
	}
}

func TestIsHexColorShortValid(t *testing.T) {
	_, err := v2.IsHexColor("#f00")
	if err != nil {
		t.Fatal(err)
	}
}

func TestIsHexColorAlphaValid(t *testing.T) {
	_, err := v2.IsHexColor("#ff0000ff")
	if err != nil {
		t.Fatal(err)
	}
}

func TestIsHexColorValid(t *testing.T) {
	_, err := v2.IsHexColor("#ff0000")
	if err != nil {
		t.Fatal(err)
	}
}

func TestCheckHexColorNonString(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Theme struct {
		Color int `checkers:"hex-color"`
	}

	theme := &Theme{}

	v2.CheckStruct(theme)
}

func TestCheckHexColorInvalid(t *testing.T) {
	type Theme struct {
		Color string `checkers:"hex-color"`
	}

	theme := &Theme{
		Color: "ff0000",
	}

	_, ok := v2.CheckStruct(theme)
	if ok {
		t.Fatal("expected error")
	}
}

func TestCheckHexColorValid(t *testing.T) {
	type Theme struct {
		Color string `checkers:"hex-color"`
	}

	theme := &Theme{
		Color: "#ff0000",
	}

	_, ok := v2.CheckStruct(theme)
	if !ok {
		t.Fatal("expected valid")
	}
}
