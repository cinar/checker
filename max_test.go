package checker_test

import (
	"testing"

	"github.com/cinar/checker"
)

func ExampleIsMax() {
	quantity := 5

	result := checker.IsMax(quantity, 10)

	if result != checker.ResultValid {
		// Send the mistakes back to the user
	}
}

func TestIsMaxValid(t *testing.T) {
	n := 5

	if checker.IsMax(n, 10) != checker.ResultValid {
		t.Fail()
	}
}

func TestCheckMaxInvalidConfig(t *testing.T) {
	defer checker.FailIfNoPanic(t)

	type Order struct {
		Quantity int `checkers:"max:AB"`
	}

	order := &Order{}

	checker.Check(order)
}

func TestCheckMaxValid(t *testing.T) {
	type Order struct {
		Quantity int `checkers:"max:10"`
	}

	order := &Order{
		Quantity: 5,
	}

	_, valid := checker.Check(order)
	if !valid {
		t.Fail()
	}
}

func TestCheckMaxInvalid(t *testing.T) {
	type Order struct {
		Quantity int `checkers:"max:10"`
	}

	order := &Order{
		Quantity: 20,
	}

	_, valid := checker.Check(order)
	if valid {
		t.Fail()
	}
}
