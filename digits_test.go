// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package checker_test

import (
	"testing"

	"github.com/cinar/checker"
)

func ExampleIsDigits() {
	result := checker.IsDigits("1234")

	if result != checker.ResultValid {
		// Send the mistakes back to the user
	}
}

func TestIsDigitsInvalid(t *testing.T) {
	if checker.IsDigits("checker") == checker.ResultValid {
		t.Fail()
	}
}

func TestIsDigitsValid(t *testing.T) {
	if checker.IsDigits("1234") != checker.ResultValid {
		t.Fail()
	}
}

func TestCheckDigitsNonString(t *testing.T) {
	defer checker.FailIfNoPanic(t)

	type User struct {
		ID int `checkers:"digits"`
	}

	user := &User{}

	checker.Check(user)
}

func TestCheckDigitsInvalid(t *testing.T) {
	type User struct {
		ID string `checkers:"digits"`
	}

	user := &User{
		ID: "checker",
	}

	_, valid := checker.Check(user)
	if valid {
		t.Fail()
	}
}

func TestCheckDigitsValid(t *testing.T) {
	type User struct {
		ID string `checkers:"digits"`
	}

	user := &User{
		ID: "1234",
	}

	_, valid := checker.Check(user)
	if !valid {
		t.Fail()
	}
}
