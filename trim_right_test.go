package checker

import "testing"

func TestNormalizeTrimRightNonString(t *testing.T) {
	defer FailIfNoPanic(t)

	type User struct {
		Username int `checkers:"trim-right"`
	}

	user := &User{}

	Check(user)
}

func TestNormalizeTrimRightResultValid(t *testing.T) {
	type User struct {
		Username string `checkers:"trim-right"`
	}

	user := &User{
		Username: "      normalizer      ",
	}

	_, valid := Check(user)
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

	Check(user)

	if user.Username != "      normalizer" {
		t.Fail()
	}
}
