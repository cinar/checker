// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2_test

import (
	"testing"

	v2 "github.com/cinar/checker/v2"
)

func TestTitle(t *testing.T) {
	input := "the checker"
	expected := "The Checker"

	actual, err := v2.Title(input)
	if err != nil {
		t.Fatal(err)
	}

	if actual != expected {
		t.Fatalf("actual %s expected %s", actual, expected)
	}
}

func TestReflectTitle(t *testing.T) {
	type Book struct {
		Chapter string `checkers:"title"`
	}

	book := &Book{
		Chapter: "the checker",
	}

	expected := "The Checker"

	errs, ok := v2.CheckStruct(book)
	if !ok {
		t.Fatalf("got unexpected errors %v", errs)
	}

	if book.Chapter != expected {
		t.Fatalf("actual %s expected %s", book.Chapter, expected)
	}
}
