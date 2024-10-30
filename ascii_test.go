// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package checker_test

import (
	"testing"

	"github.com/cinar/checker"
)

func ExampleIsASCII() {
	err := checker.IsASCII("Checker")
	if err != nil {
		// Send the errors back to the user
	}
}

func TestIsASCIIInvalid(t *testing.T) {
	if checker.IsASCII("𝄞 Music!") == nil {
		t.Fail()
	}
}

func TestIsASCIIValid(t *testing.T) {
	if checker.IsASCII("Checker") != nil {
		t.Fail()
	}
}

func TestCheckASCIINonString(t *testing.T) {
	defer checker.FailIfNoPanic(t)

	type User struct {
		Age int `checkers:"ascii"`
	}

	user := &User{}

	checker.Check(user)
}

func TestCheckASCIIInvalid(t *testing.T) {
	type User struct {
		Username string `checkers:"ascii"`
	}

	user := &User{
		Username: "𝄞 Music!",
	}

	_, valid := checker.Check(user)
	if valid {
		t.Fail()
	}
}

func TestCheckASCIIValid(t *testing.T) {
	type User struct {
		Username string `checkers:"ascii"`
	}

	user := &User{
		Username: "checker",
	}

	_, valid := checker.Check(user)
	if !valid {
		t.Fail()
	}
}
