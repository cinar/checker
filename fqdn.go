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

// CheckerFqdn is the name of the checker.
const CheckerFqdn = "fqdn"

// ErrNotFqdn indicates that the given string is not a valid FQDN.
var ErrNotFqdn = errors.New("please enter a valid domain name")

// Valid characters excluding full-width characters.
var fqdnValidChars = regexp.MustCompile("^[a-z0-9\u00a1-\uff00\uff06-\uffff\\-]+$")

// IsFqdn checks if the given string is a fully qualified domain name.
func IsFqdn(domain string) error {
	parts := strings.Split(domain, ".")

	// Require TLD
	if len(parts) < 2 {
		return ErrNotFqdn
	}

	tld := parts[len(parts)-1]

	// Should be all numeric TLD
	if IsDigits(tld) == nil {
		return ErrNotFqdn
	}

	// Short TLD
	if len(tld) < 2 {
		return ErrNotFqdn
	}

	for _, part := range parts {
		// Cannot be more than 63 characters
		if len(part) > 63 {
			return ErrNotFqdn
		}

		// Check for valid characters
		if !fqdnValidChars.MatchString(part) {
			return ErrNotFqdn
		}

		// Should not start or end with a hyphen (-) character.
		if part[0] == '-' || part[len(part)-1] == '-' {
			return ErrNotFqdn
		}
	}

	return nil
}

// makeFqdn makes a checker function for the fqdn checker.
func makeFqdn(_ string) CheckFunc {
	return checkFqdn
}

// checkFqdn checks if the given string is a fully qualified domain name.
func checkFqdn(value, _ reflect.Value) error {
	if value.Kind() != reflect.String {
		panic("string expected")
	}

	return IsFqdn(value.String())
}
