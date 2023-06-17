package checker

import "testing"

func TestNormalizeTrimLeftNonString(t *testing.T) {
	defer FailIfNoPanic(t)

	type User struct {
		Username int `checkers:"trim-left"`
	}

	user := &User{}

	Check(user)
}

func TestNormalizeTrimLeftResultValid(t *testing.T) {
	type User struct {
		Username string `checkers:"trim-left"`
	}

	user := &User{
		Username: "      normalizer      ",
	}

	_, valid := Check(user)
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

	Check(user)

	if user.Username != "normalizer      " {
		t.Fail()
	}
}
