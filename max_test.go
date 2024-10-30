// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package checker_test

import (
	"testing"

	"github.com/cinar/checker"
)

func ExampleIsMax() {
	quantity := 5

	err := checker.IsMax(quantity, 10)
	if err != nil {
		// Send the errors back to the user
	}
}

func TestIsMaxValid(t *testing.T) {
	n := 5

	if checker.IsMax(n, 10) != nil {
		t.Fail()
	}
}

func TestCheckMaxInvalidConfig(t *testing.T) {
	defer checker.FailIfNoPanic(t)

	type Order struct {
		Quantity int `checkers:"max:AB"`
	}

	order := &Order{}

	checker.Check(order)
}

func TestCheckMaxValid(t *testing.T) {
	type Order struct {
		Quantity int `checkers:"max:10"`
	}

	order := &Order{
		Quantity: 5,
	}

	_, valid := checker.Check(order)
	if !valid {
		t.Fail()
	}
}

func TestCheckMaxInvalid(t *testing.T) {
	type Order struct {
		Quantity int `checkers:"max:10"`
	}

	order := &Order{
		Quantity: 20,
	}

	_, valid := checker.Check(order)
	if valid {
		t.Fail()
	}
}
