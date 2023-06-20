package checker

import "testing"

func TestIsLuhnValid(t *testing.T) {
	numbers := []string{
		"4012888888881881",
		"4222222222222",
		"5555555555554444",
		"5105105105105100",
	}

	for _, number := range numbers {
		if IsLuhn(number) != ResultValid {
			t.Fail()
		}
	}
}

func TestCheckLuhnNonString(t *testing.T) {
	defer FailIfNoPanic(t)

	type Order struct {
		CreditCard int `checkers:"luhn"`
	}

	order := &Order{}

	Check(order)
}

func TestCheckLuhnInvalid(t *testing.T) {
	type Order struct {
		CreditCard string `checkers:"luhn"`
	}

	order := &Order{
		CreditCard: "4012888888881884",
	}

	_, valid := Check(order)
	if valid {
		t.Fail()
	}
}

func TestCheckLuhnValid(t *testing.T) {
	type Order struct {
		CreditCard string `checkers:"luhn"`
	}

	order := &Order{
		CreditCard: "4012888888881881",
	}

	_, valid := Check(order)
	if !valid {
		t.Fail()
	}
}
