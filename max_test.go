package checker

import "testing"

func TestIsMaxValid(t *testing.T) {
	n := 5

	if IsMax(n, 10) != ResultValid {
		t.Fail()
	}
}

func TestCheckMaxInvalidConfig(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fail()
		}
	}()

	type Order struct {
		Quantity int `checkers:"max:AB"`
	}

	order := &Order{}

	Check(order)
}

func TestCheckMaxValid(t *testing.T) {
	type Order struct {
		Quantity int `checkers:"max:10"`
	}

	order := &Order{
		Quantity: 5,
	}

	_, valid := Check(order)
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

	_, valid := Check(order)
	if valid {
		t.Fail()
	}
}
