// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2_test

import (
	"errors"
	"fmt"
	"testing"

	v2 "github.com/cinar/checker/v2"
)

func ExampleIsEndsWith() {
	_, err := v2.IsEndsWith(".com", "example.com")
	if err != nil {
		fmt.Println(err)
	}
}

func TestIsEndsWithValid(t *testing.T) {
	_, err := v2.IsEndsWith(".com", "example.com")
	if err != nil {
		t.Fatal(err)
	}
}

func TestIsEndsWithInvalid(t *testing.T) {
	_, err := v2.IsEndsWith(".com", "example.org")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestCheckEndsWithNonString(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Link struct {
		URL int `checkers:"ends-with:.com"`
	}

	link := &Link{}

	v2.CheckStruct(link)
}

func TestCheckEndsWithInvalid(t *testing.T) {
	type Link struct {
		URL string `checkers:"ends-with:.com"`
	}

	link := &Link{
		URL: "example.org",
	}

	errs, ok := v2.CheckStruct(link)
	if ok {
		t.Fatal("expected error")
	}

	if !errors.Is(errs["URL"], v2.ErrNotEndsWith) {
		t.Fatalf("expected ErrNotEndsWith, got %v", errs)
	}
}

func TestCheckEndsWithValid(t *testing.T) {
	type Link struct {
		URL string `checkers:"ends-with:.com"`
	}

	link := &Link{
		URL: "example.com",
	}

	errs, ok := v2.CheckStruct(link)
	if !ok {
		t.Fatal(errs)
	}
}
