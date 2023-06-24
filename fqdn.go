// Package checker is a Go library for validating user input through struct tags.
//
// Copyright (c) 2023 Onur Cinar. All Rights Reserved.
// The source code is provided under MIT License.
//
// https://github.com/cinar/checker
package checker

import (
	"reflect"
	"regexp"
	"strings"
)

// CheckerFqdn is the name of the checker.
const CheckerFqdn = "fqdn"

// ResultNotFqdn indicates that the given string is not a valid FQDN.
const ResultNotFqdn = "NOT_FQDN"

// Valid characters excluding full-width characters.
var fqdnValidChars = regexp.MustCompile("^[a-z0-9\u00a1-\uff00\uff06-\uffff\\-]+$")

// IsFqdn checks if the given string is a fully qualified domain name.
func IsFqdn(domain string) Result {
	parts := strings.Split(domain, ".")

	// Require TLD
	if len(parts) < 2 {
		return ResultNotFqdn
	}

	tld := parts[len(parts)-1]

	// Should be all numeric TLD
	if IsDigits(tld) == ResultValid {
		return ResultNotFqdn
	}

	// Short TLD
	if len(tld) < 2 {
		return ResultNotFqdn
	}

	for _, part := range parts {
		// Cannot be more than 63 characters
		if len(part) > 63 {
			return ResultNotFqdn
		}

		// Check for valid characters
		if !fqdnValidChars.MatchString(part) {
			return ResultNotFqdn
		}

		// Should not start or end with a hyphen (-) character.
		if part[0] == '-' || part[len(part)-1] == '-' {
			return ResultNotFqdn
		}
	}

	return ResultValid
}

// makeFqdn makes a checker function for the fqdn checker.
func makeFqdn(_ string) CheckFunc {
	return checkFqdn
}

// checkFqdn checks if the given string is a fully qualified domain name.
func checkFqdn(value, _ reflect.Value) Result {
	if value.Kind() != reflect.String {
		panic("string expected")
	}

	return IsFqdn(value.String())
}
