package checker

import "testing"

func TestNormalizeLowerNonString(t *testing.T) {
	defer FailIfNoPanic(t)

	type User struct {
		Username int `checkers:"lower"`
	}

	user := &User{}

	Check(user)
}

func TestNormalizeLowerResultValid(t *testing.T) {
	type User struct {
		Username string `checkers:"lower"`
	}

	user := &User{
		Username: "chECker",
	}

	_, valid := Check(user)
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

	Check(user)

	if user.Username != "checker" {
		t.Fail()
	}
}
