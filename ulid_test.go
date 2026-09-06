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

func ExampleIsULID() {
	_, err := v2.IsULID("01ARZ3NDEKTSV4RRFFQ69G5FAV")
	if err != nil {
		fmt.Println(err)
	}
}

func TestIsULIDInvalid(t *testing.T) {
	_, err := v2.IsULID("not-a-ulid")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestIsULIDAmbiguousLetterInvalid(t *testing.T) {
	_, err := v2.IsULID("01ARZ3NDEKTSVORRFFQ69G5FAV")
	if err == nil {
		t.Fatal("expected error for ambiguous 'O' letter")
	}
}

func TestIsULIDValid(t *testing.T) {
	_, err := v2.IsULID("01ARZ3NDEKTSV4RRFFQ69G5FAV")
	if err != nil {
		t.Fatal(err)
	}
}

func TestCheckULIDNonString(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Record struct {
		ID int `checkers:"ulid"`
	}

	record := &Record{}

	v2.CheckStruct(record)
}

func TestCheckULIDInvalid(t *testing.T) {
	type Record struct {
		ID string `checkers:"ulid"`
	}

	record := &Record{
		ID: "not-a-ulid",
	}

	_, ok := v2.CheckStruct(record)
	if ok {
		t.Fatal("expected error")
	}
}

func TestCheckULIDValid(t *testing.T) {
	type Record struct {
		ID string `checkers:"ulid"`
	}

	record := &Record{
		ID: "01ARZ3NDEKTSV4RRFFQ69G5FAV",
	}

	_, ok := v2.CheckStruct(record)
	if !ok {
		t.Fatal("expected valid")
	}
}
