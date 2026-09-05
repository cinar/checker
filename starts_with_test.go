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

func ExampleIsStartsWith() {
	_, err := v2.IsStartsWith("https://", "https://example.com")
	if err != nil {
		fmt.Println(err)
	}
}

func TestIsStartsWithValid(t *testing.T) {
	_, err := v2.IsStartsWith("https://", "https://example.com")
	if err != nil {
		t.Fatal(err)
	}
}

func TestIsStartsWithInvalid(t *testing.T) {
	_, err := v2.IsStartsWith("https://", "http://example.com")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestCheckStartsWithNonString(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Link struct {
		URL int `checkers:"starts-with:https://"`
	}

	link := &Link{}

	v2.CheckStruct(link)
}

func TestCheckStartsWithInvalid(t *testing.T) {
	type Link struct {
		URL string `checkers:"starts-with:https://"`
	}

	link := &Link{
		URL: "http://example.com",
	}

	errs, ok := v2.CheckStruct(link)
	if ok {
		t.Fatal("expected error")
	}

	if !errors.Is(errs["URL"], v2.ErrNotStartsWith) {
		t.Fatalf("expected ErrNotStartsWith, got %v", errs)
	}
}

func TestCheckStartsWithValid(t *testing.T) {
	type Link struct {
		URL string `checkers:"starts-with:https://"`
	}

	link := &Link{
		URL: "https://example.com",
	}

	errs, ok := v2.CheckStruct(link)
	if !ok {
		t.Fatal(errs)
	}
}
