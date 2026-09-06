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

func ExampleIsISBN() {
	_, err := v2.IsISBN("1430248270")
	if err != nil {
		fmt.Println(err)
	}
}

func TestIsISBNInvalid(t *testing.T) {
	_, err := v2.IsISBN("invalid-isbn")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestIsISBNValid(t *testing.T) {
	_, err := v2.IsISBN("1430248270")
	if err != nil {
		t.Fatal(err)
	}
}

func TestIsISBN10BadChecksumInvalid(t *testing.T) {
	_, err := v2.IsISBN("0306406153")
	if err == nil {
		t.Fatal("expected error for bad ISBN-10 checksum")
	}
}

func TestIsISBN13BadChecksumInvalid(t *testing.T) {
	_, err := v2.IsISBN("9780306406158")
	if err == nil {
		t.Fatal("expected error for bad ISBN-13 checksum")
	}
}

func TestIsISBN10XCheckDigitValid(t *testing.T) {
	_, err := v2.IsISBN("080442957X")
	if err != nil {
		t.Fatal(err)
	}
}

func TestIsISBN10HyphenatedValid(t *testing.T) {
	_, err := v2.IsISBN("0-306-40615-2")
	if err != nil {
		t.Fatal(err)
	}
}

func TestIsISBN13HyphenatedValid(t *testing.T) {
	_, err := v2.IsISBN("978-0-306-40615-7")
	if err != nil {
		t.Fatal(err)
	}
}

func TestCheckISBNNonString(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Book struct {
		ISBN int `checkers:"isbn"`
	}

	book := &Book{}

	v2.CheckStruct(book)
}

func TestCheckISBNInvalid(t *testing.T) {
	type Book struct {
		ISBN string `checkers:"isbn"`
	}

	book := &Book{
		ISBN: "invalid-isbn",
	}

	_, ok := v2.CheckStruct(book)
	if ok {
		t.Fatal("expected error")
	}
}

func TestCheckISBNValid(t *testing.T) {
	type Book struct {
		ISBN string `checkers:"isbn"`
	}

	book := &Book{
		ISBN: "9783161484100",
	}

	_, ok := v2.CheckStruct(book)
	if !ok {
		t.Fatal("expected valid")
	}
}
