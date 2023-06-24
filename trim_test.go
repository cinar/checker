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

func TestNormalizeTrimNonString(t *testing.T) {
	defer checker.FailIfNoPanic(t)

	type User struct {
		Username int `checkers:"trim"`
	}

	user := &User{}

	checker.Check(user)
}

func TestNormalizeTrimResultValid(t *testing.T) {
	type User struct {
		Username string `checkers:"trim"`
	}

	user := &User{
		Username: "      normalizer      ",
	}

	_, valid := checker.Check(user)
	if !valid {
		t.Fail()
	}
}

func TestNormalizeTrim(t *testing.T) {
	type User struct {
		Username string `checkers:"trim"`
	}

	user := &User{
		Username: "      normalizer      ",
	}

	checker.Check(user)

	if user.Username != "normalizer" {
		t.Fail()
	}
}
