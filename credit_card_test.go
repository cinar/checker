// Package checker is a Go library for validating user input through struct tags.
//
// https://github.com/cinar/checker
//
// Copyright 2023 Onur Cinar. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package checker_test

import (
	"strconv"
	"testing"

	"github.com/cinar/checker"
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
	result := checker.IsAnyCreditCard("6011111111111117")

	if result != checker.ResultValid {
		// Send the mistakes back to the user
	}
}

func TestIsAnyCreditCardValid(t *testing.T) {
	if checker.IsAnyCreditCard(amexCard) != checker.ResultValid {
		t.Fail()
	}
}

func TestIsAnyCreditCardInvalidPattern(t *testing.T) {
	if checker.IsAnyCreditCard(invalidCard) == checker.ResultValid {
		t.Fail()
	}
}

func TestIsAnyCreditCardInvalidLuhn(t *testing.T) {
	if checker.IsAnyCreditCard(changeToInvalidLuhn(amexCard)) == checker.ResultValid {
		t.Fail()
	}
}

func ExampleIsAmexCreditCard() {
	result := checker.IsAmexCreditCard("378282246310005")

	if result != checker.ResultValid {
		// Send the mistakes back to the user
	}
}

func TestIsAmexCreditCardValid(t *testing.T) {
	if checker.IsAmexCreditCard(amexCard) != checker.ResultValid {
		t.Fail()
	}
}

func TestIsAmexCreditCardInvalidPattern(t *testing.T) {
	if checker.IsAmexCreditCard(invalidCard) == checker.ResultValid {
		t.Fail()
	}
}

func TestIsAmexCreditCardInvalidLuhn(t *testing.T) {
	if checker.IsAmexCreditCard(changeToInvalidLuhn(amexCard)) == checker.ResultValid {
		t.Fail()
	}
}

func ExampleIsDinersCreditCard() {
	result := checker.IsDinersCreditCard("36227206271667")

	if result != checker.ResultValid {
		// Send the mistakes back to the user
	}
}
func TestIsDinersCreditCardValid(t *testing.T) {
	if checker.IsDinersCreditCard(dinersCard) != checker.ResultValid {
		t.Fail()
	}
}

func TestIsDinersCreditCardInvalidPattern(t *testing.T) {
	if checker.IsDinersCreditCard(invalidCard) == checker.ResultValid {
		t.Fail()
	}
}

func TestIsDinersCreditCardInvalidLuhn(t *testing.T) {
	if checker.IsDinersCreditCard(changeToInvalidLuhn(dinersCard)) == checker.ResultValid {
		t.Fail()
	}
}

func ExampleIsDiscoverCreditCard() {
	result := checker.IsDiscoverCreditCard("6011111111111117")

	if result != checker.ResultValid {
		// Send the mistakes back to the user
	}
}
func TestIsDiscoverCreditCardValid(t *testing.T) {
	if checker.IsDiscoverCreditCard(discoverCard) != checker.ResultValid {
		t.Fail()
	}
}

func TestIsDiscoverCreditCardInvalidPattern(t *testing.T) {
	if checker.IsDiscoverCreditCard(invalidCard) == checker.ResultValid {
		t.Fail()
	}
}

func TestIsDiscoverCreditCardInvalidLuhn(t *testing.T) {
	if checker.IsDiscoverCreditCard(changeToInvalidLuhn(discoverCard)) == checker.ResultValid {
		t.Fail()
	}
}

func ExampleIsJcbCreditCard() {
	result := checker.IsJcbCreditCard("3530111333300000")

	if result != checker.ResultValid {
		// Send the mistakes back to the user
	}
}

func TestIsJcbCreditCardValid(t *testing.T) {
	if checker.IsJcbCreditCard(jcbCard) != checker.ResultValid {
		t.Fail()
	}
}

func TestIsJcbCreditCardInvalidPattern(t *testing.T) {
	if checker.IsJcbCreditCard(invalidCard) == checker.ResultValid {
		t.Fail()
	}
}

func TestIsJcbCreditCardInvalidLuhn(t *testing.T) {
	if checker.IsJcbCreditCard(changeToInvalidLuhn(jcbCard)) == checker.ResultValid {
		t.Fail()
	}
}

func ExampleIsMasterCardCreditCard() {
	result := checker.IsMasterCardCreditCard("5555555555554444")

	if result != checker.ResultValid {
		// Send the mistakes back to the user
	}
}

func TestIsMasterCardCreditCardValid(t *testing.T) {
	if checker.IsMasterCardCreditCard(masterCard) != checker.ResultValid {
		t.Fail()
	}
}

func TestIsMasterCardCreditCardInvalidPattern(t *testing.T) {
	if checker.IsMasterCardCreditCard(invalidCard) == checker.ResultValid {
		t.Fail()
	}
}

func TestIsMasterCardCreditCardInvalidLuhn(t *testing.T) {
	if checker.IsMasterCardCreditCard(changeToInvalidLuhn(masterCard)) == checker.ResultValid {
		t.Fail()
	}
}

func ExampleIsUnionPayCreditCard() {
	result := checker.IsUnionPayCreditCard("6200000000000005")

	if result != checker.ResultValid {
		// Send the mistakes back to the user
	}
}

func TestIsUnionPayCreditCardValid(t *testing.T) {
	if checker.IsUnionPayCreditCard(unionPayCard) != checker.ResultValid {
		t.Fail()
	}
}

func TestIsUnionPayCreditCardInvalidPattern(t *testing.T) {
	if checker.IsUnionPayCreditCard(invalidCard) == checker.ResultValid {
		t.Fail()
	}
}

func TestIsUnionPayCreditCardInvalidLuhn(t *testing.T) {
	if checker.IsUnionPayCreditCard(changeToInvalidLuhn(unionPayCard)) == checker.ResultValid {
		t.Fail()
	}
}

func ExampleIsVisaCreditCard() {
	result := checker.IsVisaCreditCard("4111111111111111")

	if result != checker.ResultValid {
		// Send the mistakes back to the user
	}
}
func TestIsVisaCreditCardValid(t *testing.T) {
	if checker.IsVisaCreditCard(visaCard) != checker.ResultValid {
		t.Fail()
	}
}

func TestIsVisaCreditCardInvalidPattern(t *testing.T) {
	if checker.IsVisaCreditCard(invalidCard) == checker.ResultValid {
		t.Fail()
	}
}

func TestIsVisaCreditCardInvalidLuhn(t *testing.T) {
	if checker.IsVisaCreditCard(changeToInvalidLuhn(visaCard)) == checker.ResultValid {
		t.Fail()
	}
}

func TestCheckCreditCardNonString(t *testing.T) {
	defer checker.FailIfNoPanic(t)

	type Order struct {
		CreditCard int `checkers:"credit-card"`
	}

	order := &Order{}

	checker.Check(order)
}

func TestCheckCreditCardValid(t *testing.T) {
	type Order struct {
		CreditCard string `checkers:"credit-card"`
	}

	order := &Order{
		CreditCard: amexCard,
	}

	_, valid := checker.Check(order)
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

	_, valid := checker.Check(order)
	if valid {
		t.Fail()
	}
}

func TestCheckCreditCardMultipleUnknown(t *testing.T) {
	defer checker.FailIfNoPanic(t)

	type Order struct {
		CreditCard string `checkers:"credit-card:amex,unknown"`
	}

	order := &Order{
		CreditCard: amexCard,
	}

	checker.Check(order)
}

func TestCheckCreditCardMultipleInvalid(t *testing.T) {
	type Order struct {
		CreditCard string `checkers:"credit-card:amex,visa"`
	}

	order := &Order{
		CreditCard: discoverCard,
	}

	_, valid := checker.Check(order)
	if valid {
		t.Fail()
	}
}
