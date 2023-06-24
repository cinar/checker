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

func TestIsMinValid(t *testing.T) {
	n := 45

	if checker.IsMin(n, 21) != checker.ResultValid {
		t.Fail()
	}
}

func TestCheckMinInvalidConfig(t *testing.T) {
	defer checker.FailIfNoPanic(t)

	type User struct {
		Age int `checkers:"min:AB"`
	}

	user := &User{}

	checker.Check(user)
}

func TestCheckMinValid(t *testing.T) {
	type User struct {
		Age int `checkers:"min:21"`
	}

	user := &User{
		Age: 45,
	}

	_, valid := checker.Check(user)
	if !valid {
		t.Fail()
	}
}

func TestCheckMinInvalid(t *testing.T) {
	type User struct {
		Age int `checkers:"min:21"`
	}

	user := &User{
		Age: 18,
	}

	_, valid := checker.Check(user)
	if valid {
		t.Fail()
	}
}
