package checker

import "testing"

func TestIsMinLengthValid(t *testing.T) {
	s := "1234"

	if IsMinLength(s, 4) != ResultValid {
		t.Fail()
	}
}

func TestCheckMinLengthInvalidConfig(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fail()
		}
	}()

	type User struct {
		Password string `checkers:"min-length:AB"`
	}

	user := &User{}

	Check(user)
}

func TestCheckMinLengthValid(t *testing.T) {
	type User struct {
		Password string `checkers:"min-length:4"`
	}

	user := &User{
		Password: "1234",
	}

	_, valid := Check(user)
	if !valid {
		t.Fail()
	}
}

func TestCheckMinLengthInvalid(t *testing.T) {
	type User struct {
		Password string `checkers:"min-length:4"`
	}

	user := &User{
		Password: "12",
	}

	_, valid := Check(user)
	if valid {
		t.Fail()
	}
}
