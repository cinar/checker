// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2_test

import (
	"strconv"
	"testing"

	v2 "github.com/cinar/checker/v2"
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

func ExampleIsAnyCreditCard() {
	_, err := v2.IsAnyCreditCard("6011111111111117")

	if err != nil {
		// Send the errors back to the user
	}
}

func TestIsAnyCreditCardValid(t *testing.T) {
	_, err := v2.IsAnyCreditCard(amexCard)
	if err != nil {
		t.Error(err)
	}
}

func TestIsAnyCreditCardInvalidPattern(t *testing.T) {
	if _, err := v2.IsAnyCreditCard(invalidCard); err == nil {
		t.Error("expected error for invalid card pattern")
	}
}

func TestIsAnyCreditCardInvalidLuhn(t *testing.T) {
	if _, err := v2.IsAnyCreditCard(changeToInvalidLuhn(amexCard)); err == nil {
		t.Error("expected error for invalid Luhn")
	}
}

func ExampleIsAmexCreditCard() {
	_, err := v2.IsAmexCreditCard("378282246310005")

	if err != nil {
		// Send the errors back to the user
	}
}

func TestIsAmexCreditCardValid(t *testing.T) {
	if _, err := v2.IsAmexCreditCard(amexCard); err != nil {
		t.Error(err)
	}
}

func TestIsAmexCreditCardInvalidPattern(t *testing.T) {
	if _, err := v2.IsAmexCreditCard(invalidCard); err == nil {
		t.Error("expected error for invalid card pattern")
	}
}

func TestIsAmexCreditCardInvalidLuhn(t *testing.T) {
	if _, err := v2.IsAmexCreditCard(changeToInvalidLuhn(amexCard)); err == nil {
		t.Error("expected error for invalid Luhn")
	}
}

func ExampleIsDinersCreditCard() {
	_, err := v2.IsDinersCreditCard("36227206271667")

	if err != nil {
		// Send the errors back to the user
	}
}
func TestIsDinersCreditCardValid(t *testing.T) {
	if _, err := v2.IsDinersCreditCard(dinersCard); err != nil {
		t.Error(err)
	}
}

func TestIsDinersCreditCardInvalidPattern(t *testing.T) {
	if _, err := v2.IsDinersCreditCard(invalidCard); err == nil {
		t.Error("expected error for invalid card pattern")
	}
}

func TestIsDinersCreditCardInvalidLuhn(t *testing.T) {
	if _, err := v2.IsDinersCreditCard(changeToInvalidLuhn(dinersCard)); err == nil {
		t.Error("expected error for invalid Luhn")
	}
}

func ExampleIsDiscoverCreditCard() {
	_, err := v2.IsDiscoverCreditCard("6011111111111117")

	if err != nil {
		// Send the errors back to the user
	}
}
func TestIsDiscoverCreditCardValid(t *testing.T) {
	if _, err := v2.IsDiscoverCreditCard(discoverCard); err != nil {
		t.Error(err)
	}
}

func TestIsDiscoverCreditCardInvalidPattern(t *testing.T) {
	if _, err := v2.IsDiscoverCreditCard(invalidCard); err == nil {
		t.Error("expected error for invalid card pattern")
	}
}

func TestIsDiscoverCreditCardInvalidLuhn(t *testing.T) {
	if _, err := v2.IsDiscoverCreditCard(changeToInvalidLuhn(discoverCard)); err == nil {
		t.Error("expected error for invalid Luhn")
	}
}

func ExampleIsJcbCreditCard() {
	_, err := v2.IsJcbCreditCard("3530111333300000")

	if err != nil {
		// Send the errors back to the user
	}
}

func TestIsJcbCreditCardValid(t *testing.T) {
	if _, err := v2.IsJcbCreditCard(jcbCard); err != nil {
		t.Error(err)
	}
}

func TestIsJcbCreditCardInvalidPattern(t *testing.T) {
	if _, err := v2.IsJcbCreditCard(invalidCard); err == nil {
		t.Error("expected error for invalid card pattern")
	}
}

func TestIsJcbCreditCardInvalidLuhn(t *testing.T) {
	if _, err := v2.IsJcbCreditCard(changeToInvalidLuhn(jcbCard)); err == nil {
		t.Error("expected error for invalid Luhn")
	}
}

func ExampleIsMasterCardCreditCard() {
	_, err := v2.IsMasterCardCreditCard("5555555555554444")

	if err != nil {
		// Send the errors back to the user
	}
}

func TestIsMasterCardCreditCardValid(t *testing.T) {
	if _, err := v2.IsMasterCardCreditCard(masterCard); err != nil {
		t.Error(err)
	}
}

func TestIsMasterCardCreditCardInvalidPattern(t *testing.T) {
	if _, err := v2.IsMasterCardCreditCard(invalidCard); err == nil {
		t.Error("expected error for invalid card pattern")
	}
}

func TestIsMasterCardCreditCardInvalidLuhn(t *testing.T) {
	if _, err := v2.IsMasterCardCreditCard(changeToInvalidLuhn(masterCard)); err == nil {
		t.Error("expected error for invalid Luhn")
	}
}

func ExampleIsUnionPayCreditCard() {
	_, err := v2.IsUnionPayCreditCard("6200000000000005")

	if err != nil {
		// Send the errors back to the user
	}
}

func TestIsUnionPayCreditCardValid(t *testing.T) {
	if _, err := v2.IsUnionPayCreditCard(unionPayCard); err != nil {
		t.Error(err)
	}
}

func TestIsUnionPayCreditCardInvalidPattern(t *testing.T) {
	if _, err := v2.IsUnionPayCreditCard(invalidCard); err == nil {
		t.Error("expected error for invalid card pattern")
	}
}

func TestIsUnionPayCreditCardInvalidLuhn(t *testing.T) {
	if _, err := v2.IsUnionPayCreditCard(changeToInvalidLuhn(unionPayCard)); err == nil {
		t.Error("expected error for invalid Luhn")
	}
}

func ExampleIsVisaCreditCard() {
	_, err := v2.IsVisaCreditCard("4111111111111111")

	if err != nil {
		// Send the errors back to the user
	}
}
func TestIsVisaCreditCardValid(t *testing.T) {
	if _, err := v2.IsVisaCreditCard(visaCard); err != nil {
		t.Error(err)
	}
}

func TestIsVisaCreditCardInvalidPattern(t *testing.T) {
	if _, err := v2.IsVisaCreditCard(invalidCard); err == nil {
		t.Error("expected error for invalid card pattern")
	}
}

func TestIsVisaCreditCardInvalidLuhn(t *testing.T) {
	if _, err := v2.IsVisaCreditCard(changeToInvalidLuhn(visaCard)); err == nil {
		t.Error("expected error for invalid Luhn")
	}
}

func TestCheckCreditCardNonString(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic for non-string credit card")

	type Order struct {
		CreditCard int `checkers:"credit-card"`
	}

	order := &Order{}

	v2.CheckStruct(order)
}

func TestCheckCreditCardValid(t *testing.T) {
	type Order struct {
		CreditCard string `checkers:"credit-card"`
	}

	order := &Order{
		CreditCard: amexCard,
	}

	_, valid := v2.CheckStruct(order)
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

	_, valid := v2.CheckStruct(order)
	if valid {
		t.Fail()
	}
}

func TestCheckCreditCardMultipleUnknown(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic for unknown credit card")

	type Order struct {
		CreditCard string `checkers:"credit-card:amex,unknown"`
	}

	order := &Order{
		CreditCard: amexCard,
	}

	v2.CheckStruct(order)
}

func TestCheckCreditCardMultipleInvalid(t *testing.T) {
	type Order struct {
		CreditCard string `checkers:"credit-card:amex,visa"`
	}

	order := &Order{
		CreditCard: discoverCard,
	}

	_, valid := v2.CheckStruct(order)
	if valid {
		t.Fail()
	}
}
