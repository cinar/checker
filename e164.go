// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

import (
	"reflect"
)

const (
	// nameE164 is the name of the E.164 check.
	nameE164 = "e164"

	// e164Expression matches an E.164 phone number: a leading "+", a
	// non-zero first digit, and up to 15 digits in total.
	e164Expression = `^\+[1-9]\d{1,14}$`
)

var (
	// ErrNotE164 indicates that the given string is not a valid E.164 phone number.
	ErrNotE164 = NewCheckError("NOT_E164")
)

// IsE164 checks if the given string is a valid E.164 phone number, e.g. "+14155552671".
func IsE164(value string) (string, error) {
	if _, err := IsRegexp(e164Expression, value); err != nil {
		return value, ErrNotE164
	}

	return value, nil
}

// checkE164 checks if the given string is a valid E.164 phone number.
func checkE164(value reflect.Value) (reflect.Value, error) {
	_, err := IsE164(reflectString(value))
	return value, err
}

// makeE164 makes a checker function for the E.164 checker.
func makeE164(_ string) CheckFunc[reflect.Value] {
	return checkE164
}
