package checker

import "testing"

func TestIsASCIIInvalid(t *testing.T) {
	if IsASCII("𝄞 Music!") == ResultValid {
		t.Fail()
	}
}

func TestIsASCIIValid(t *testing.T) {
	if IsASCII("Checker") != ResultValid {
		t.Fail()
	}
}

func TestCheckASCIINonString(t *testing.T) {
	defer FailIfNoPanic(t)

	type User struct {
		Age int `checkers:"ascii"`
	}

	user := &User{}

	Check(user)
}

func TestCheckASCIIInvalid(t *testing.T) {
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

func TestCheckASCIIValid(t *testing.T) {
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
