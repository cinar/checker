// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package checker_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/cinar/checker"
)

func TestCheckRegexpNonString(t *testing.T) {
	defer checker.FailIfNoPanic(t)

	type User struct {
		Username int `checkers:"regexp:^[A-Za-z]$"`
	}

	user := &User{}

	checker.Check(user)
}

func TestCheckRegexpInvalid(t *testing.T) {
	type User struct {
		Username string `checkers:"regexp:^[A-Za-z]+$"`
	}

	user := &User{
		Username: "abcd1234",
	}

	_, valid := checker.Check(user)
	if valid {
		t.Fail()
	}
}

func TestCheckRegexpValid(t *testing.T) {
	type User struct {
		Username string `checkers:"regexp:^[A-Za-z]+$"`
	}

	user := &User{
		Username: "abcd",
	}

	_, valid := checker.Check(user)
	if !valid {
		t.Fail()
	}
}

func TestMakeRegexpChecker(t *testing.T) {
	checkHex := checker.MakeRegexpChecker("^[A-Fa-f0-9]+$", errors.New("Not Hex"))

	err := checkHex(reflect.ValueOf("f0f0f0"), reflect.ValueOf(nil))
	if err != nil {
		t.Fail()
	}
}

func TestMakeRegexpMaker(t *testing.T) {
	checker.Register("hex", checker.MakeRegexpMaker("^[A-Fa-f0-9]+$", errors.New("Not Hex")))

	type Theme struct {
		Color string `checkers:"hex"`
	}

	theme := &Theme{
		Color: "f0f0f0",
	}

	_, valid := checker.Check(theme)
	if !valid {
		t.Fail()
	}
}
