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

func TestIsAlphanumericInvalid(t *testing.T) {
	if checker.IsAlphanumeric("-/") == checker.ResultValid {
		t.Fail()
	}
}

func TestIsAlphanumericValid(t *testing.T) {
	if checker.IsAlphanumeric("ABcd1234") != checker.ResultValid {
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
