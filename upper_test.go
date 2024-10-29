// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package checker_test

import (
	"testing"

	"github.com/cinar/checker"
)

func TestNormalizeUpperNonString(t *testing.T) {
	defer checker.FailIfNoPanic(t)

	type User struct {
		Username int `checkers:"upper"`
	}

	user := &User{}

	checker.Check(user)
}

func TestNormalizeUpperResultValid(t *testing.T) {
	type User struct {
		Username string `checkers:"upper"`
	}

	user := &User{
		Username: "chECker",
	}

	_, valid := checker.Check(user)
	if !valid {
		t.Fail()
	}
}

func TestNormalizeUpper(t *testing.T) {
	type User struct {
		Username string `checkers:"upper"`
	}

	user := &User{
		Username: "chECker",
	}

	checker.Check(user)

	if user.Username != "CHECKER" {
		t.Fail()
	}
}
