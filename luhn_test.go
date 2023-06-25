// Package checker is a Go library for validating user input through struct tags.
//
// https://github.com/cinar/checker
//
// Copyright 2023 Onur Cinar. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package checker_test

import (
	"testing"

	"github.com/cinar/checker"
)

func ExampleIsLuhn() {
	result := checker.IsLuhn("4012888888881881")

	if result != checker.ResultValid {
		// Send the mistakes back to the user
	}
}

func TestIsLuhnValid(t *testing.T) {
	numbers := []string{
		"4012888888881881",
		"4222222222222",
		"5555555555554444",
		"5105105105105100",
	}

	for _, number := range numbers {
		if checker.IsLuhn(number) != checker.ResultValid {
			t.Fail()
		}
	}
}

func TestCheckLuhnNonString(t *testing.T) {
	defer checker.FailIfNoPanic(t)

	type Order struct {
		CreditCard int `checkers:"luhn"`
	}

	order := &Order{}

	checker.Check(order)
}

func TestCheckLuhnInvalid(t *testing.T) {
	type Order struct {
		CreditCard string `checkers:"luhn"`
	}

	order := &Order{
		CreditCard: "4012888888881884",
	}

	_, valid := checker.Check(order)
	if valid {
		t.Fail()
	}
}

func TestCheckLuhnValid(t *testing.T) {
	type Order struct {
		CreditCard string `checkers:"luhn"`
	}

	order := &Order{
		CreditCard: "4012888888881881",
	}

	_, valid := checker.Check(order)
	if !valid {
		t.Fail()
	}
}
