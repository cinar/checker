package checker

import (
	"reflect"
	"testing"
)

func TestCheckRegexpNonString(t *testing.T) {
	defer FailIfNoPanic(t)

	type User struct {
		Username int `checkers:"regexp:^[A-Za-z]$"`
	}

	user := &User{}

	Check(user)
}

func TestCheckRegexpInvalid(t *testing.T) {
	type User struct {
		Username string `checkers:"regexp:^[A-Za-z]+$"`
	}

	user := &User{
		Username: "abcd1234",
	}

	_, valid := Check(user)
	if valid {
		t.Fail()
	}
}

func TestCheckRegexpValid(t *testing.T) {
	type User struct {
		Username string `checkers:"regexp:^[A-Za-z]+$"`
	}

	user := &User{
		Username: "abcd",
	}

	_, valid := Check(user)
	if !valid {
		t.Fail()
	}
}

func TestMakeRegexpChecker(t *testing.T) {
	checkHex := MakeRegexpChecker("^[A-Fa-f0-9]+$", "NOT_HEX")

	result := checkHex(reflect.ValueOf("f0f0f0"), reflect.ValueOf(nil))
	if result != ResultValid {
		t.Fail()
	}
}

func TestMakeRegexpMaker(t *testing.T) {
	Register("hex", MakeRegexpMaker("^[A-Fa-f0-9]+$", "NOT_HEX"))

	type Theme struct {
		Color string `checkers:hex`
	}

	theme := &Theme{
		Color: "f0f0f0",
	}

	_, valid := Check(theme)
	if !valid {
		t.Fail()
	}
}
