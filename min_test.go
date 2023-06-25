package checker_test

import (
	"testing"

	"github.com/cinar/checker"
)

func ExampleIsMin() {
	age := 45

	result := checker.IsMin(age, 21)

	if result != checker.ResultValid {
		// Send the mistakes back to the user
	}
}

func TestIsMinValid(t *testing.T) {
	n := 45

	if checker.IsMin(n, 21) != checker.ResultValid {
		t.Fail()
	}
}

func TestCheckMinInvalidConfig(t *testing.T) {
	defer checker.FailIfNoPanic(t)

	type User struct {
		Age int `checkers:"min:AB"`
	}

	user := &User{}

	checker.Check(user)
}

func TestCheckMinValid(t *testing.T) {
	type User struct {
		Age int `checkers:"min:21"`
	}

	user := &User{
		Age: 45,
	}

	_, valid := checker.Check(user)
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

	_, valid := checker.Check(user)
	if valid {
		t.Fail()
	}
}
