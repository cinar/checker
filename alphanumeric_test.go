package checker

import "testing"

func TestIsAlphanumericInvalid(t *testing.T) {
	if IsAlphanumeric("-/") == ResultValid {
		t.Fail()
	}
}

func TestIsAlphanumericValid(t *testing.T) {
	if IsAlphanumeric("ABcd1234") != ResultValid {
		t.Fail()
	}
}

func TestCheckAlphanumericNonString(t *testing.T) {
	defer FailIfNoPanic(t)

	type User struct {
		Username int `checkers:"alphanumeric"`
	}

	user := &User{}

	Check(user)
}

func TestCheckAlphanumericInvalid(t *testing.T) {
	type User struct {
		Username string `checkers:"alphanumeric"`
	}

	user := &User{
		Username: "user-/",
	}

	_, valid := Check(user)
	if valid {
		t.Fail()
	}
}

func TestCheckAlphanumericValid(t *testing.T) {
	type User struct {
		Username string `checkers:"alphanumeric"`
	}

	user := &User{
		Username: "ABcd1234",
	}

	_, valid := Check(user)
	if !valid {
		t.Fail()
	}
}
