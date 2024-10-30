// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package checker_test

import (
	"testing"

	"github.com/cinar/checker"
)

func ExampleIsAlphanumeric() {
	err := checker.IsAlphanumeric("ABcd1234")
	if err != nil {
		// Send the mistakes back to the user
	}
}

func TestIsAlphanumericInvalid(t *testing.T) {
	if checker.IsAlphanumeric("-/") == nil {
		t.Fail()
	}
}

func TestIsAlphanumericValid(t *testing.T) {
	if checker.IsAlphanumeric("ABcd1234") != nil {
		t.Fail()
	}
}

func TestCheckAlphanumericNonString(t *testing.T) {
	defer checker.FailIfNoPanic(t)

	type User struct {
		Username int `checkers:"alphanumeric"`
	}

	user := &User{}

	checker.Check(user)
}

func TestCheckAlphanumericInvalid(t *testing.T) {
	type User struct {
		Username string `checkers:"alphanumeric"`
	}

	user := &User{
		Username: "user-/",
	}

	_, valid := checker.Check(user)
	if valid {
		t.Fail()
	}
}

func TestCheckAlphanumericValid(t *testing.T) {
	type User struct {
		Username string `checkers:"alphanumeric"`
	}

	user := &User{
		Username: "ABcd1234",
	}

	_, valid := checker.Check(user)
	if !valid {
		t.Fail()
	}
}
