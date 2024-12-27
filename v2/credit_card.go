// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

import (
	"reflect"
	"regexp"
	"strings"
)

const (
	// nameCreditCard is the name of the credit card check.
	nameCreditCard = "credit_card"
)

var (
	// ErrNotCreditCard indicates that the given value is not a valid credit card number.
	ErrNotCreditCard = NewCheckError("CreditCard")

	// creditCardPatterns contains the regular expressions for validating different credit card types.
	creditCardPatterns = map[string]*regexp.Regexp{
		"visa":       regexp.MustCompile(`^4[0-9]{12}(?:[0-9]{3})?$`),
		"mastercard": regexp.MustCompile(`^5[1-5][0-9]{14}$`),
		"amex":       regexp.MustCompile(`^3[47][0-9]{13}$`),
		"discover":   regexp.MustCompile(`^6(?:011|5[0-9]{2})[0-9]{12}$`),
		"jcb":        regexp.MustCompile(`^(?:2131|1800|35\d{3})\d{11}$`),
		"diners":     regexp.MustCompile(`^3(?:0[0-5]|[68][0-9])[0-9]{11}$`),
	}

	// anyCreditCardPattern is the regular expression for validating any credit card number.
	anyCreditCardPattern = regexp.MustCompile(strings.Join([]string{
		creditCardPatterns["visa"].String(),
		creditCardPatterns["mastercard"].String(),
		creditCardPatterns["amex"].String(),
		creditCardPatterns["discover"].String(),
		creditCardPatterns["jcb"].String(),
		creditCardPatterns["diners"].String(),
	}, "|"))
)

// IsCreditCard checks if the value is a valid credit card number.
func IsCreditCard(value string) (string, error) {
	if !anyCreditCardPattern.MatchString(value) {
		return value, ErrNotCreditCard
	}
	return value, nil
}

// IsVisaCreditCard checks if the value is a valid Visa credit card number.
func IsVisaCreditCard(value string) (string, error) {
	if !creditCardPatterns["visa"].MatchString(value) {
		return value, ErrNotCreditCard
	}
	return value, nil
}

// IsMasterCardCreditCard checks if the value is a valid MasterCard credit card number.
func IsMasterCardCreditCard(value string) (string, error) {
	if !creditCardPatterns["mastercard"].MatchString(value) {
		return value, ErrNotCreditCard
	}
	return value, nil
}

// IsAmexCreditCard checks if the value is a valid American Express credit card number.
func IsAmexCreditCard(value string) (string, error) {
	if !creditCardPatterns["amex"].MatchString(value) {
		return value, ErrNotCreditCard
	}
	return value, nil
}

// checkCreditCard checks if the value is a valid credit card number.
func checkCreditCard(value reflect.Value) (reflect.Value, error) {
	_, err := IsCreditCard(value.Interface().(string))
	return value, err
}

// makeCreditCard makes a checker function for the credit card checker.
func makeCreditCard(config string) CheckFunc[reflect.Value] {
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

	return func(value reflect.Value) (reflect.Value, error) {
		if value.Kind() != reflect.String {
			panic("string expected")
		}

		number := value.String()

		for _, pattern := range patterns {
			if pattern.MatchString(number) {
				return value, nil
			}
		}

		return value, ErrNotCreditCard
	}
}
