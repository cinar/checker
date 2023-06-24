// Package checker is a Go library for validating user input through struct tags.
//
// Copyright (c) 2023 Onur Cinar. All Rights Reserved.
// The source code is provided under MIT License.
//
// https://github.com/cinar/checker
package checker_test

import (
	"testing"

	"github.com/cinar/checker"
)

func TestNormalizeTitleNonString(t *testing.T) {
	defer checker.FailIfNoPanic(t)

	type Book struct {
		Chapter int `checkers:"title"`
	}

	book := &Book{}

	checker.Check(book)
}

func TestNormalizeTitleResultValid(t *testing.T) {
	type Book struct {
		Chapter string `checkers:"title"`
	}

	book := &Book{
		Chapter: "THE checker",
	}

	_, valid := checker.Check(book)
	if !valid {
		t.Fail()
	}
}

func TestNormalizeTitle(t *testing.T) {
	type Book struct {
		Chapter string `checkers:"title"`
	}

	book := &Book{
		Chapter: "THE checker",
	}

	checker.Check(book)

	if book.Chapter != "The Checker" {
		t.Fail()
	}
}
