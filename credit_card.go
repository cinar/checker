// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package checker

import (
	"errors"
	"reflect"
	"regexp"
	"strings"
)

// CheckerCreditCard is the name of the checker.
const CheckerCreditCard = "credit-card"

// ErrNotCreditCard indicates that the given value is not a valid credit card number.
var ErrNotCreditCard = errors.New("please enter a valid credit card number")

// amexExpression is the regexp for the AMEX cards. They start with 34 or 37, and has 15 digits.
var amexExpression = "(?:^(?:3[47])[0-9]{13}$)"
var amexPattern = regexp.MustCompile(amexExpression)

// dinersExpression is the regexp for the Diners cards. They start with 305, 36, 38, and has 14 digits.
var dinersExpression = "(?:^3(?:(?:05[0-9]{11})|(?:[68][0-9]{12}))$)"
var dinersPattern = regexp.MustCompile(dinersExpression)

// discoverExpression is the regexp for the Discover cards. They start with 6011 and has 16 digits.
var discoverExpression = "(?:^6011[0-9]{12}$)"
var discoverPattern = regexp.MustCompile(discoverExpression)

// jcbExpression is the regexp for the JCB 15 cards. They start with 2131, 1800, and has 15 digits, or start with 35 and has 16 digits.
var jcbExpression = "(?:^(?:(?:2131)|(?:1800)|(?:35[0-9]{3}))[0-9]{11}$)"
var jcbPattern = regexp.MustCompile(jcbExpression)

// masterCardExpression is the regexp for the MasterCard cards. They start with 51, 52, 53, 54, or 55, and has 15 digits.
var masterCardExpression = "(?:^5[12345][0-9]{14}$)"
var masterCardPattern = regexp.MustCompile(masterCardExpression)

// unionPayExpression is the regexp for the UnionPay cards. They start either with 62 or 67, and has 16 digits, or they start with 81 and has 16 to 19 digits.
var unionPayExpression = "(?:(?:6[27][0-9]{14})|(?:81[0-9]{14,17})^$)"
var unionPayPattern = regexp.MustCompile(unionPayExpression)

// visaExpression is the regexp for the Visa cards. They start with 4 and has 13 or 16 digits.
var visaExpression = "(?:^4[0-9]{12}(?:[0-9]{3})?$)"
var visaPattern = regexp.MustCompile(visaExpression)

// anyCreditCardPattern is the regexp for any credit card.
var anyCreditCardPattern = regexp.MustCompile(strings.Join([]string{
	amexExpression,
	dinersExpression,
	discoverExpression,
	jcbExpression,
	masterCardExpression,
	unionPayExpression,
	visaExpression,
}, "|"))

// creditCardPatterns is the mapping for credit card names to patterns.
var creditCardPatterns = map[string]*regexp.Regexp{
	"amex":       amexPattern,
	"diners":     dinersPattern,
	"discover":   discoverPattern,
	"jcb":        jcbPattern,
	"mastercard": masterCardPattern,
	"unionpay":   unionPayPattern,
	"visa":       visaPattern,
}

// IsAnyCreditCard checks if the given value is a valid credit card number.
func IsAnyCreditCard(number string) error {
	return isCreditCard(number, anyCreditCardPattern)
}

// IsAmexCreditCard checks if the given valie is a valid AMEX credit card.
func IsAmexCreditCard(number string) error {
	return isCreditCard(number, amexPattern)
}

// IsDinersCreditCard checks if the given valie is a valid Diners credit card.
func IsDinersCreditCard(number string) error {
	return isCreditCard(number, dinersPattern)
}

// IsDiscoverCreditCard checks if the given valie is a valid Discover credit card.
func IsDiscoverCreditCard(number string) error {
	return isCreditCard(number, discoverPattern)
}

// IsJcbCreditCard checks if the given valie is a valid JCB 15 credit card.
func IsJcbCreditCard(number string) error {
	return isCreditCard(number, jcbPattern)
}

// IsMasterCardCreditCard checks if the given valie is a valid MasterCard credit card.
func IsMasterCardCreditCard(number string) error {
	return isCreditCard(number, masterCardPattern)
}

// IsUnionPayCreditCard checks if the given valie is a valid UnionPay credit card.
func IsUnionPayCreditCard(number string) error {
	return isCreditCard(number, unionPayPattern)
}

// IsVisaCreditCard checks if the given valie is a valid Visa credit card.
func IsVisaCreditCard(number string) error {
	return isCreditCard(number, visaPattern)
}

// makeCreditCard makes a checker function for the credit card checker.
func makeCreditCard(config string) CheckFunc {
	patterns := []*regexp.Regexp{}

	if config != "" {
		for _, card := range strings.Split(config, ",") {
			pattern, ok := creditCardPatterns[card]
			if !ok {
				panic("unknown credit card name")
			}

			patterns = append(patterns, pattern)
		}
	} else {
		patterns = append(patterns, anyCreditCardPattern)
	}

	return func(value, _ reflect.Value) error {
		if value.Kind() != reflect.String {
			panic("string expected")
		}

		number := value.String()

		for _, pattern := range patterns {
			if isCreditCard(number, pattern) == nil {
				return nil
			}
		}

		return ErrNotCreditCard
	}
}

// isCreditCard checks if the given number based on the given credit card pattern and the Luhn algorithm check digit.
func isCreditCard(number string, pattern *regexp.Regexp) error {
	if !pattern.MatchString(number) {
		return ErrNotCreditCard
	}

	return IsLuhn(number)
}
