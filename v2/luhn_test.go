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

func ExampleIsLUHN() {
	_, err := v2.IsLUHN("4012888888881881")
	if err != nil {
		fmt.Println(err)
	}
}

func TestIsLUHNInvalid(t *testing.T) {
	_, err := v2.IsLUHN("123456789")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestIsLUHNInvalidDigits(t *testing.T) {
	_, err := v2.IsLUHN("ABCD")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestIsLUHNValid(t *testing.T) {
	_, err := v2.IsLUHN("4012888888881881")
	if err != nil {
		t.Fatal(err)
	}
}

func TestCheckLUHNNonString(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Card struct {
		Number int `checkers:"luhn"`
	}

	card := &Card{}

	v2.CheckStruct(card)
}

func TestCheckLUHNInvalid(t *testing.T) {
	type Card struct {
		Number string `checkers:"luhn"`
	}

	card := &Card{
		Number: "123456789",
	}

	_, ok := v2.CheckStruct(card)
	if ok {
		t.Fatal("expected error")
	}
}

func TestCheckLUHNValid(t *testing.T) {
	type Card struct {
		Number string `checkers:"luhn"`
	}

	card := &Card{
		Number: "79927398713",
	}

	_, ok := v2.CheckStruct(card)
	if !ok {
		t.Fatal("expected valid")
	}
}
