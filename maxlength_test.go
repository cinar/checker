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

func ExampleIsMaxLength() {
	s := "1234"

	result := checker.IsMaxLength(s, 4)

	if result != checker.ResultValid {
		// Send the mistakes back to the user
	}
}

func TestIsMaxLengthValid(t *testing.T) {
	s := "1234"

	if checker.IsMaxLength(s, 4) != checker.ResultValid {
		t.Fail()
	}
}

func TestCheckMaxLengthInvalidConfig(t *testing.T) {
	defer checker.FailIfNoPanic(t)

	type User struct {
		Password string `checkers:"max-length:AB"`
	}

	user := &User{}

	checker.Check(user)
}

func TestCheckMaxLengthValid(t *testing.T) {
	type User struct {
		Password string `checkers:"max-length:4"`
	}

	user := &User{
		Password: "1234",
	}

	_, valid := checker.Check(user)
	if !valid {
		t.Fail()
	}
}

func TestCheckMaxLengthInvalid(t *testing.T) {
	type User struct {
		Password string `checkers:"max-length:4"`
	}

	user := &User{
		Password: "123456",
	}

	_, valid := checker.Check(user)
	if valid {
		t.Fail()
	}
}
