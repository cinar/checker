package checker

import "testing"

func TestNormalizeTrimNonString(t *testing.T) {
	defer FailIfNoPanic(t)

	type User struct {
		Username int `checkers:"trim"`
	}

	user := &User{}

	Check(user)
}

func TestNormalizeTrimResultValid(t *testing.T) {
	type User struct {
		Username string `checkers:"trim"`
	}

	user := &User{
		Username: "      normalizer      ",
	}

	_, valid := Check(user)
	if !valid {
		t.Fail()
	}
}

func TestNormalizeTrim(t *testing.T) {
	type User struct {
		Username string `checkers:"trim"`
	}

	user := &User{
		Username: "      normalizer      ",
	}

	Check(user)

	if user.Username != "normalizer" {
		t.Fail()
	}
}
