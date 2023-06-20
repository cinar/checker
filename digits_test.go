package checker

import "testing"

func TestIsDigitsInvalid(t *testing.T) {
	if IsDigits("checker") == ResultValid {
		t.Fail()
	}
}

func TestIsDigitsValid(t *testing.T) {
	if IsDigits("1234") != ResultValid {
		t.Fail()
	}
}

func TestCheckDigitsNonString(t *testing.T) {
	defer FailIfNoPanic(t)

	type User struct {
		ID int `checkers:"digits"`
	}

	user := &User{}

	Check(user)
}

func TestCheckDigitsInvalid(t *testing.T) {
	type User struct {
		ID string `checkers:"digits"`
	}

	user := &User{
		ID: "checker",
	}

	_, valid := Check(user)
	if valid {
		t.Fail()
	}
}

func TestCheckDigitsValid(t *testing.T) {
	type User struct {
		ID string `checkers:"digits"`
	}

	user := &User{
		ID: "1234",
	}

	_, valid := Check(user)
	if !valid {
		t.Fail()
	}
}
