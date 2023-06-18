package checker

import "testing"

func TestNormalizeUpperNonString(t *testing.T) {
	defer FailIfNoPanic(t)

	type User struct {
		Username int `checkers:"upper"`
	}

	user := &User{}

	Check(user)
}

func TestNormalizeUpperResultValid(t *testing.T) {
	type User struct {
		Username string `checkers:"upper"`
	}

	user := &User{
		Username: "chECker",
	}

	_, valid := Check(user)
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

	Check(user)

	if user.Username != "CHECKER" {
		t.Fail()
	}
}
