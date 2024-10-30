// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package checker_test

import (
	"testing"

	"github.com/cinar/checker"
)

func TestNormalizeTrimLeftNonString(t *testing.T) {
	defer checker.FailIfNoPanic(t)

	type User struct {
		Username int `checkers:"trim-left"`
	}

	user := &User{}

	checker.Check(user)
}

func TestNormalizeTrimLeftErrValid(t *testing.T) {
	type User struct {
		Username string `checkers:"trim-left"`
	}

	user := &User{
		Username: "      normalizer      ",
	}

	_, valid := checker.Check(user)
	if !valid {
		t.Fail()
	}
}

func TestNormalizeTrimLeft(t *testing.T) {
	type User struct {
		Username string `checkers:"trim-left"`
	}

	user := &User{
		Username: "      normalizer      ",
	}

	checker.Check(user)

	if user.Username != "normalizer      " {
		t.Fail()
	}
}
