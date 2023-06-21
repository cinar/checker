package checker

import (
	"regexp"
	"strings"
)

// Used information from the following resources:
//
// - https://dnschecker.org/credit-card-validator.php
// - https://www.validcreditcardnumber.com/

// CheckerCreditCard is the name of the checker.
const CheckerCreditCard = "credit-card"

// ResultNotCreditCard indicates that the given value is not a valid credit card number.
const ResultNotCreditCard = "NOT_CREDIT_CARD"

// amexExpression is the regexp for the AMEX cards. They start with 34 or 37, and has 15 digits.
var amexExpression = "(?:^(?:3[47])[0-9]{13}$)"
var amexPattern = regexp.MustCompile(amexExpression)

// dinersExpression is the regexp for the Diners cards. They start with 305, 36, 38, and has 14 digits.
var dinersExpression = "(?:^3(?:(?:05[0-9]{11})|(?:[68][0-9]{12}))$)"
var dinersPattern = regexp.MustCompile(dinersExpression)

// discoverExpression is the regexp for the Discover cards. They start with 6011 and has 16 digits.
var discoverExpression = "(?:^6011[0-9]{12}$)"
var discoverPattern = regexp.MustCompile(discoverExpression)

// enrouteExpression is the regexp for the Enroute cards. They start with 2014 or 2149, and has 16 digits.
var enrouteExpression = "(?:^2(?:(?:014)|(?:149))[0-9]{12}$)"
var enroutePattern = regexp.MustCompile(enrouteExpression)

// jcbExpression is the regexp for the JCB 15 cards. They start with 2131, 1800, and has 15 digits, or start with 35 and has 16 digits.
var jcbExpression = "(?:^(?:(?:2131)|(?:1800)|(?:35[0-9]{3}))[0-9]{11}$)"
var jcbPattern = regexp.MustCompile(jcbExpression)

// masterCardExpression is the regexp for the MasterCard cards. They start with 51, 52, 53, 54, or 55, and has 15 digits.
var masterCardExpression = "(?:^5[12345][0-9]{14}$)"
var masterCardPattern = regexp.MustCompile(masterCardExpression)

// visaExpression is the regexp for the Visa cards. They start with 4 and has 13 or 16 digits.
var visaExpression = "(?:^4[0-9]{12}(?:[0-9]{3})?$)"
var visaPattern = regexp.MustCompile(visaExpression)

// voyagerExpression is the regexp for the Voyager cards. They start with 8699 and has 13 or 16 digits.
var voyagerExpression = "(?:^8699[0-9]{12}(?:[0-9]{3})?$)"
var voyagerPattern = regexp.MustCompile(voyagerExpression)

// anyCreditCardPattern is the regexp for any credit card.
var anyCreditCardPattern = regexp.MustCompile(strings.Join([]string{
	amexExpression,
	dinersExpression,
	discoverExpression,
	enrouteExpression,
	jcbExpression,
	masterCardExpression,
	visaExpression,
	voyagerExpression,
}, "|"))

// IsAnyCreditCard checks if the given value is a valid credit card number.
func IsAnyCreditCard(number string) Result {
	return isCreditCard(number, anyCreditCardPattern)
}

// IsAmexCreditCard checks if the given valie is a valid AMEX credit card.
func IsAmexCreditCard(number string) Result {
	return isCreditCard(number, amexPattern)
}

// IsDinersCreditCard checks if the given valie is a valid Diners credit card.
func IsDinersCreditCard(number string) Result {
	return isCreditCard(number, dinersPattern)
}

// IsDiscoveryCreditCard checks if the given valie is a valid Discovery credit card.
func IsDiscoveryCreditCard(number string) Result {
	return isCreditCard(number, discoverPattern)
}

// IsEnrouteCreditCard checks if the given valie is a valid Enroute credit card.
func IsEnrouteCreditCard(number string) Result {
	return isCreditCard(number, enroutePattern)
}

// IsJcbCreditCard checks if the given valie is a valid JCB 15 credit card.
func IsJcbCreditCard(number string) Result {
	return isCreditCard(number, jcbPattern)
}

// IsMasterCardCreditCard checks if the given valie is a valid MasterCard credit card.
func IsMasterCardCreditCard(number string) Result {
	return isCreditCard(number, masterCardPattern)
}

// IsVisaCreditCard checks if the given valie is a valid Visa credit card.
func IsVisaCreditCard(number string) Result {
	return isCreditCard(number, visaPattern)
}

// IsVoyagerCreditCard checks if the given valie is a valid Voyager credit card.
func IsVoyagerCreditCard(number string) Result {
	return isCreditCard(number, voyagerPattern)
}

// isCreditCard checks if the given number based on the given credit card pattern and the Luhn algorithm check digit.
func isCreditCard(number string, pattern *regexp.Regexp) Result {
	if !pattern.MatchString(number) {
		return ResultNotCreditCard
	}

	return IsLuhn(number)
}
