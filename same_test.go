package checker

import (
	"reflect"
	"testing"
)

func TestSameValid(t *testing.T) {
	type User struct {
		Password string
		Confirm  string `checkers:"same:Password"`
	}

	user := &User{
		Password: "1234",
		Confirm:  "1234",
	}

	_, valid := Check(user)
	if !valid {
		t.Fail()
	}
}

func TestSameInvalid(t *testing.T) {
	type User struct {
		Password string
		Confirm  string `checkers:"same:Password"`
	}

	user := &User{
		Password: "1234",
		Confirm:  "12",
	}

	_, valid := Check(user)
	if valid {
		t.Fail()
	}
}

func TestSameWithoutParent(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fail()
		}
	}()

	type User struct {
		Password string
		Confirm  string `checkers:"same:Password"`
	}

	user := &User{
		Password: "1234",
		Confirm:  "12",
	}

	checkSame(reflect.ValueOf(user.Confirm), reflect.ValueOf(nil), "Password")
}

func TestSameInvalidName(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fail()
		}
	}()

	type User struct {
		Password string
		Confirm  string `checkers:"same:Unknown"`
	}

	user := &User{
		Password: "1234",
		Confirm:  "1234",
	}

	Check(user)
}
