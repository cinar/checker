package checker

import "testing"

func TestIsMinValid(t *testing.T) {
	n := 45

	if IsMin(n, 21) != ResultValid {
		t.Fail()
	}
}

func TestCheckMinInvalidConfig(t *testing.T) {
	defer FailIfNoPanic(t)

	type User struct {
		Age int `checkers:"min:AB"`
	}

	user := &User{}

	Check(user)
}

func TestCheckMinValid(t *testing.T) {
	type User struct {
		Age int `checkers:"min:21"`
	}

	user := &User{
		Age: 45,
	}

	_, valid := Check(user)
	if !valid {
		t.Fail()
	}
}

func TestCheckMinInvalid(t *testing.T) {
	type User struct {
		Age int `checkers:"min:21"`
	}

	user := &User{
		Age: 18,
	}

	_, valid := Check(user)
	if valid {
		t.Fail()
	}
}
