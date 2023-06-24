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

func TestIsMaxValid(t *testing.T) {
	n := 5

	if checker.IsMax(n, 10) != checker.ResultValid {
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
