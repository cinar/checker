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

func ExampleIsNumeric() {
	_, err := v2.IsNumeric("-3.14")
	if err != nil {
		fmt.Println(err)
	}
}

func TestIsNumericInvalid(t *testing.T) {
	_, err := v2.IsNumeric("abc")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestIsNumericEmptyInvalid(t *testing.T) {
	_, err := v2.IsNumeric("")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestIsNumericValidInteger(t *testing.T) {
	_, err := v2.IsNumeric("42")
	if err != nil {
		t.Fatal(err)
	}
}

func TestIsNumericValidSigned(t *testing.T) {
	_, err := v2.IsNumeric("-3.14")
	if err != nil {
		t.Fatal(err)
	}
}

func TestIsNumericValidPlusSign(t *testing.T) {
	_, err := v2.IsNumeric("+7")
	if err != nil {
		t.Fatal(err)
	}
}

func TestCheckNumericNonString(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Product struct {
		Price int `checkers:"numeric"`
	}

	product := &Product{}

	v2.CheckStruct(product)
}

func TestCheckNumericInvalid(t *testing.T) {
	type Product struct {
		Price string `checkers:"numeric"`
	}

	product := &Product{
		Price: "free",
	}

	_, ok := v2.CheckStruct(product)
	if ok {
		t.Fatal("expected error")
	}
}

func TestCheckNumericValid(t *testing.T) {
	type Product struct {
		Price string `checkers:"numeric"`
	}

	product := &Product{
		Price: "19.99",
	}

	errs, ok := v2.CheckStruct(product)
	if !ok {
		t.Fatal(errs)
	}
}
