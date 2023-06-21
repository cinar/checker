package checker

import "testing"

var amexCard = "371449635398431"
var dinersCard = "30569309025904"
var discoverCard = "6011111111111117"
var jcbCard = "3530111333300000"
var mastercardCard = "5555555555554444"
var visaCard = "4111111111111111"

func TestIsAmexCreditCardValid(t *testing.T) {
	if IsAmexCreditCard("371449635398431") != ResultValid {
		t.Fail()
	}
}

func TestIsAmexCreditCardInvalidPattern(t *testing.T) {
	if IsAmexCreditCard("30569309025904") == ResultValid {
		t.Fail()
	}
}

func TestIsAmexCreditCardInvalidLuhn(t *testing.T) {
	if IsAmexCreditCard("371449635398432") == ResultValid {
		t.Fail()
	}
}

func TestIsDinersCreditCardValid(t *testing.T) {
	if IsDinersCreditCard("30569309025904") != ResultValid {
		t.Fail()
	}
}

func TestIsDinersCreditCardInvalidPattern(t *testing.T) {
	if IsDinersCreditCard("371449635398431") == ResultValid {
		t.Fail()
	}
}

func TestIsDinersCreditCardInvalidLuhn(t *testing.T) {
	if IsDinersCreditCard("30569309025906") == ResultValid {
		t.Fail()
	}
}

func TestIsDiscoverCreditCardValid(t *testing.T) {
	if IsDiscoveryCreditCard("6011111111111117") != ResultValid {
		t.Fail()
	}
}

func TestIsDiscoverCreditCardInvalidPattern(t *testing.T) {
	if IsDiscoveryCreditCard("30569309025904") == ResultValid {
		t.Fail()
	}
}

func TestIsDiscoverCreditCardInvalidLuhn(t *testing.T) {
	if IsDiscoveryCreditCard("6011111111111118") == ResultValid {
		t.Fail()
	}
}

func TestIsJcbCreditCardValid(t *testing.T) {
	if IsJcbCreditCard("3530111333300000") != ResultValid {
		t.Fail()
	}
}

func TestIsJcbCreditCardInvalidPattern(t *testing.T) {
	if IsJcbCreditCard("6011111111111117") == ResultValid {
		t.Fail()
	}
}

func TestIsJcbCreditCardInvalidLuhn(t *testing.T) {
	if IsJcbCreditCard("3530111333300002") == ResultValid {
		t.Fail()
	}
}

func TestIsMastercardCreditCardValid(t *testing.T) {
	if IsMasterCardCreditCard("5555555555554444") != ResultValid {
		t.Fail()
	}
}

func TestIsMastercardCreditCardInvalidPattern(t *testing.T) {
	if IsMasterCardCreditCard("3530111333300000") == ResultValid {
		t.Fail()
	}
}

func TestIsMastercardCreditCardInvalidLuhn(t *testing.T) {
	if IsMasterCardCreditCard("5555555555554446") == ResultValid {
		t.Fail()
	}
}

func TestIsVisaCreditCardValid(t *testing.T) {
	if IsVisaCreditCard("4111111111111111") != ResultValid {
		t.Fail()
	}
}

func TestIsVisaCreditCardInvalidPattern(t *testing.T) {
	if IsVisaCreditCard("5555555555554444") == ResultValid {
		t.Fail()
	}
}

func TestIsVisaCreditCardInvalidLuhn(t *testing.T) {
	if IsVisaCreditCard("4111111111111112") == ResultValid {
		t.Fail()
	}
}

func TestCheckCreditCardNonString(t *testing.T) {
	defer FailIfNoPanic(t)

	type Order struct {
		CreditCard int `checkers:"credit-card"`
	}

	order := &Order{}

	Check(order)
}

func TestCheckCreditCardValid(t *testing.T) {
	type Order struct {
		CreditCard string `checkers:"credit-card"`
	}

	order := &Order{
		CreditCard: "371449635398431",
	}

	_, valid := Check(order)
	if !valid {
		t.Fail()
	}
}

func TestCheckCreditCardInvalid(t *testing.T) {
	type Order struct {
		CreditCard string `checkers:"credit-card"`
	}

	order := &Order{
		CreditCard: "371449635398432",
	}

	_, valid := Check(order)
	if valid {
		t.Fail()
	}
}

func TestCheckCreditCardMultipleUnknown(t *testing.T) {
	defer FailIfNoPanic(t)

	type Order struct {
		CreditCard string `checkers:"credit-card:amex,unknown"`
	}

	order := &Order{
		CreditCard: "371449635398431",
	}

	Check(order)
}

func TestCheckCreditCardMultipleInvalid(t *testing.T) {
	type Order struct {
		CreditCard string `checkers:"credit-card:amex,visa"`
	}

	order := &Order{
		CreditCard: "6011111111111117",
	}

	_, valid := Check(order)
	if valid {
		t.Fail()
	}
}
