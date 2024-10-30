// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
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
