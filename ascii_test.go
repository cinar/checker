package checker

import "testing"

func TestIsAsciiInvalid(t *testing.T) {
	if IsAscii("𝄞 Music!") == ResultValid {
		t.Fail()
	}
}

func TestIsAsciiValid(t *testing.T) {
	if IsAscii("Checker") != ResultValid {
		t.Fail()
	}
}

func TestCheckAsciiNonString(t *testing.T) {
	defer FailIfNoPanic(t)

	type User struct {
		Age int `checkers:"ascii"`
	}

	user := &User{}

	Check(user)
}

func TestCheckAsciiInvalid(t *testing.T) {
	type User struct {
		Username string `checkers:"ascii"`
	}

	user := &User{
		Username: "𝄞 Music!",
	}

	_, valid := Check(user)
	if valid {
		t.Fail()
	}
}

func TestCheckAsciiValid(t *testing.T) {
	type User struct {
		Username string `checkers:"ascii"`
	}

	user := &User{
		Username: "checker",
	}

	_, valid := Check(user)
	if !valid {
		t.Fail()
	}
}
