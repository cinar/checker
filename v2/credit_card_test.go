// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2_test

import (
	"testing"

	v2 "github.com/cinar/checker/v2"
)

const (
	visaCard    = "4111111111111111"
	masterCard  = "5555555555554444"
	amexCard    = "378282246310005"
	invalidCard = "1234567890123456"
)

func changeToInvalidLuhn(card string) string {
	return card[:len(card)-1] + "0"
}

func TestIsCreditCardValid(t *testing.T) {
	if _, err := v2.IsCreditCard(visaCard); err != nil {
		t.Fail()
	}
}

func TestIsCreditCardInvalid(t *testing.T) {
	if _, err := v2.IsCreditCard(invalidCard); err == nil {
		t.Fail()
	}
}

func TestIsAnyCreditCardInvalidLuhn(t *testing.T) {
	if _, err := v2.IsCreditCard(changeToInvalidLuhn(amexCard)); err == nil {
		t.Fail()
	}
}

func ExampleIsAmexCreditCard() {
	_, err := v2.IsCreditCard("378282246310005")

	if err != nil {
		// Send the errors back to the user
	}
}

func TestIsAmexCreditCardValid(t *testing.T) {
	if _, err := v2.IsCreditCard(amexCard); err != nil {
		t.Fail()
	}
}

func TestIsAmexCreditCardInvalidPattern(t *testing.T) {
	if _, err := v2.IsCreditCard(invalidCard); err == nil {
		t.Fail()
	}
}
