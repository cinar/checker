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

func TestSameValid(t *testing.T) {
	type User struct {
		Password string
		Confirm  string `checkers:"same:Password"`
	}

	user := &User{
		Password: "1234",
		Confirm:  "1234",
	}

	_, valid := checker.Check(user)
	if !valid {
		t.Fail()
	}
}

func TestSameInvalid(t *testing.T) {
	type User struct {
		Password string
		Confirm  string `checkers:"same:Password"`
	}

	user := &User{
		Password: "1234",
		Confirm:  "12",
	}

	_, valid := checker.Check(user)
	if valid {
		t.Fail()
	}
}

func TestSameInvalidName(t *testing.T) {
	defer checker.FailIfNoPanic(t)

	type User struct {
		Password string
		Confirm  string `checkers:"same:Unknown"`
	}

	user := &User{
		Password: "1234",
		Confirm:  "1234",
	}

	checker.Check(user)
}
