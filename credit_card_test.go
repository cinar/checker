package checker

import (
	"strconv"
	"testing"
)

// Test numbers from https://stripe.com/docs/testing
var invalidCard = "1234123412341234"
var amexCard = "378282246310005"
var dinersCard = "36227206271667"
var discoverCard = "6011111111111117"
var jcbCard = "3530111333300000"
var masterCard = "5555555555554444"
var unionPayCard = "6200000000000005"
var visaCard = "4111111111111111"

// changeToInvalidLuhn increments the luhn digit to make the number invalid. It assumes that the given number is valid.
func changeToInvalidLuhn(number string) string {
	luhn, err := strconv.Atoi(number[len(number)-1:])
	if err != nil {
		panic(err)
	}

	luhn = (luhn + 1) % 10

	return number[:len(number)-1] + strconv.Itoa(luhn)
}

func TestIsAnyCreditCardValid(t *testing.T) {
	if IsAnyCreditCard(amexCard) != ResultValid {
		t.Fail()
	}
}

func TestIsAnyCreditCardInvalidPattern(t *testing.T) {
	if IsAnyCreditCard(invalidCard) == ResultValid {
		t.Fail()
	}
}

func TestIsAnyCreditCardInvalidLuhn(t *testing.T) {
	if IsAnyCreditCard(changeToInvalidLuhn(amexCard)) == ResultValid {
		t.Fail()
	}
}

func TestIsAmexCreditCardValid(t *testing.T) {
	if IsAmexCreditCard(amexCard) != ResultValid {
		t.Fail()
	}
}

func TestIsAmexCreditCardInvalidPattern(t *testing.T) {
	if IsAmexCreditCard(invalidCard) == ResultValid {
		t.Fail()
	}
}

func TestIsAmexCreditCardInvalidLuhn(t *testing.T) {
	if IsAmexCreditCard(changeToInvalidLuhn(amexCard)) == ResultValid {
		t.Fail()
	}
}

func TestIsDinersCreditCardValid(t *testing.T) {
	if IsDinersCreditCard(dinersCard) != ResultValid {
		t.Fail()
	}
}

func TestIsDinersCreditCardInvalidPattern(t *testing.T) {
	if IsDinersCreditCard(invalidCard) == ResultValid {
		t.Fail()
	}
}

func TestIsDinersCreditCardInvalidLuhn(t *testing.T) {
	if IsDinersCreditCard(changeToInvalidLuhn(dinersCard)) == ResultValid {
		t.Fail()
	}
}

func TestIsDiscoverCreditCardValid(t *testing.T) {
	if IsDiscoveryCreditCard(discoverCard) != ResultValid {
		t.Fail()
	}
}

func TestIsDiscoverCreditCardInvalidPattern(t *testing.T) {
	if IsDiscoveryCreditCard(invalidCard) == ResultValid {
		t.Fail()
	}
}

func TestIsDiscoverCreditCardInvalidLuhn(t *testing.T) {
	if IsDiscoveryCreditCard(changeToInvalidLuhn(discoverCard)) == ResultValid {
		t.Fail()
	}
}

func TestIsJcbCreditCardValid(t *testing.T) {
	if IsJcbCreditCard(jcbCard) != ResultValid {
		t.Fail()
	}
}

func TestIsJcbCreditCardInvalidPattern(t *testing.T) {
	if IsJcbCreditCard(invalidCard) == ResultValid {
		t.Fail()
	}
}

func TestIsJcbCreditCardInvalidLuhn(t *testing.T) {
	if IsJcbCreditCard(changeToInvalidLuhn(jcbCard)) == ResultValid {
		t.Fail()
	}
}

func TestIsMasterCardCreditCardValid(t *testing.T) {
	if IsMasterCardCreditCard(masterCard) != ResultValid {
		t.Fail()
	}
}

func TestIsMasterCardCreditCardInvalidPattern(t *testing.T) {
	if IsMasterCardCreditCard(invalidCard) == ResultValid {
		t.Fail()
	}
}

func TestIsMasterCardCreditCardInvalidLuhn(t *testing.T) {
	if IsMasterCardCreditCard(changeToInvalidLuhn(masterCard)) == ResultValid {
		t.Fail()
	}
}

func TestIsUnionPayCreditCardValid(t *testing.T) {
	if IsUnionPayCreditCard(unionPayCard) != ResultValid {
		t.Fail()
	}
}

func TestIsUnionPayCreditCardInvalidPattern(t *testing.T) {
	if IsUnionPayCreditCard(invalidCard) == ResultValid {
		t.Fail()
	}
}

func TestIsUnionPayCreditCardInvalidLuhn(t *testing.T) {
	if IsUnionPayCreditCard(changeToInvalidLuhn(unionPayCard)) == ResultValid {
		t.Fail()
	}
}

func TestIsVisaCreditCardValid(t *testing.T) {
	if IsVisaCreditCard(visaCard) != ResultValid {
		t.Fail()
	}
}

func TestIsVisaCreditCardInvalidPattern(t *testing.T) {
	if IsVisaCreditCard(invalidCard) == ResultValid {
		t.Fail()
	}
}

func TestIsVisaCreditCardInvalidLuhn(t *testing.T) {
	if IsVisaCreditCard(changeToInvalidLuhn(visaCard)) == ResultValid {
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
		CreditCard: amexCard,
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
		CreditCard: invalidCard,
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
		CreditCard: amexCard,
	}

	Check(order)
}

func TestCheckCreditCardMultipleInvalid(t *testing.T) {
	type Order struct {
		CreditCard string `checkers:"credit-card:amex,visa"`
	}

	order := &Order{
		CreditCard: discoverCard,
	}

	_, valid := Check(order)
	if valid {
		t.Fail()
	}
}
