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

func TestNormalizeLowerNonString(t *testing.T) {
	defer checker.FailIfNoPanic(t)

	type User struct {
		Username int `checkers:"lower"`
	}

	user := &User{}

	checker.Check(user)
}

func TestNormalizeLowerResultValid(t *testing.T) {
	type User struct {
		Username string `checkers:"lower"`
	}

	user := &User{
		Username: "chECker",
	}

	_, valid := checker.Check(user)
	if !valid {
		t.Fail()
	}
}

func TestNormalizeLower(t *testing.T) {
	type User struct {
		Username string `checkers:"lower"`
	}

	user := &User{
		Username: "chECker",
	}

	checker.Check(user)

	if user.Username != "checker" {
		t.Fail()
	}
}
