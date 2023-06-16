package checker

import "testing"

func TestIsMaxLengthValid(t *testing.T) {
	s := "1234"

	if IsMaxLength(s, 4) != ResultValid {
		t.Fail()
	}
}

func TestCheckMaxLengthInvalidConfig(t *testing.T) {
	defer FailIfNoPanic(t)

	type User struct {
		Password string `checkers:"max-length:AB"`
	}

	user := &User{}

	Check(user)
}

func TestCheckMaxLengthValid(t *testing.T) {
	type User struct {
		Password string `checkers:"max-length:4"`
	}

	user := &User{
		Password: "1234",
	}

	_, valid := Check(user)
	if !valid {
		t.Fail()
	}
}

func TestCheckMaxLengthInvalid(t *testing.T) {
	type User struct {
		Password string `checkers:"max-length:4"`
	}

	user := &User{
		Password: "123456",
	}

	_, valid := Check(user)
	if valid {
		t.Fail()
	}
}
