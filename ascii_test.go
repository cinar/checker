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

func TestIsASCIIInvalid(t *testing.T) {
	if checker.IsASCII("𝄞 Music!") == checker.ResultValid {
		t.Fail()
	}
}

func TestIsASCIIValid(t *testing.T) {
	if checker.IsASCII("Checker") != checker.ResultValid {
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
