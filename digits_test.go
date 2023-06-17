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
		Id int `checkers:"digits"`
	}

	user := &User{}

	Check(user)
}

func TestCheckDigitsInvalid(t *testing.T) {
	type User struct {
		Id string `checkers:"digits"`
	}

	user := &User{
		Id: "checker",
	}

	_, valid := Check(user)
	if valid {
		t.Fail()
	}
}

func TestCheckDigitsValid(t *testing.T) {
	type User struct {
		Id string `checkers:"digits"`
	}

	user := &User{
		Id: "1234",
	}

	_, valid := Check(user)
	if !valid {
		t.Fail()
	}
}
