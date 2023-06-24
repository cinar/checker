// Package checker is a Go library for validating user input through struct tags.
//
// https://github.com/cinar/checker
//
// Copyright 2023 Onur Cinar. All rights reserved. 
// Use of this source code is governed by a MIT-style 
// license that can be found in the LICENSE file. 
//
package checker_test

import (
	"testing"

	"github.com/cinar/checker"
)

func TestIsMinLengthValid(t *testing.T) {
	s := "1234"

	if checker.IsMinLength(s, 4) != checker.ResultValid {
		t.Fail()
	}
}

func TestCheckMinLengthInvalidConfig(t *testing.T) {
	defer checker.FailIfNoPanic(t)

	type User struct {
		Password string `checkers:"min-length:AB"`
	}

	user := &User{}

	checker.Check(user)
}

func TestCheckMinLengthValid(t *testing.T) {
	type User struct {
		Password string `checkers:"min-length:4"`
	}

	user := &User{
		Password: "1234",
	}

	_, valid := checker.Check(user)
	if !valid {
		t.Fail()
	}
}

func TestCheckMinLengthInvalid(t *testing.T) {
	type User struct {
		Password string `checkers:"min-length:4"`
	}

	user := &User{
		Password: "12",
	}

	_, valid := checker.Check(user)
	if valid {
		t.Fail()
	}
}
