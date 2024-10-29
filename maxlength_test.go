// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package checker_test

import (
	"testing"

	"github.com/cinar/checker"
)

func ExampleIsMaxLength() {
	s := "1234"

	err := checker.IsMaxLength(s, 4)
	if err != nil {
		// Send the mistakes back to the user
	}
}

func TestIsMaxLengthValid(t *testing.T) {
	s := "1234"

	if checker.IsMaxLength(s, 4) != nil {
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
