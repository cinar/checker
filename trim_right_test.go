package checker_test

import (
	"testing"

	"github.com/cinar/checker"
)

func TestNormalizeTrimRightNonString(t *testing.T) {
	defer checker.FailIfNoPanic(t)

	type User struct {
		Username int `checkers:"trim-right"`
	}

	user := &User{}

	checker.Check(user)
}

func TestNormalizeTrimRightResultValid(t *testing.T) {
	type User struct {
		Username string `checkers:"trim-right"`
	}

	user := &User{
		Username: "      normalizer      ",
	}

	_, valid := checker.Check(user)
	if !valid {
		t.Fail()
	}
}

func TestNormalizeTrimRight(t *testing.T) {
	type User struct {
		Username string `checkers:"trim-right"`
	}

	user := &User{
		Username: "      normalizer      ",
	}

	checker.Check(user)

	if user.Username != "      normalizer" {
		t.Fail()
	}
}
