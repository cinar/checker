// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

import (
	"reflect"
	"regexp"
)

const (
	// nameNumeric is the name of the numeric check.
	nameNumeric = "numeric"
)

var (
	// ErrNotNumeric indicates that the given string is not a valid numeric string.
	ErrNotNumeric = NewCheckError("NOT_NUMERIC")

	// numericExpression matches an optional leading sign and a decimal point,
	// e.g. "-3.14", "+7", "42", ".5". It deliberately excludes "NaN", "Inf",
	// hex floats, underscore digit separators, and exponents.
	numericExpression = regexp.MustCompile(`^[+-]?(\d+(\.\d+)?|\.\d+)$`)
)

// IsNumeric checks if the given string is a valid numeric string, unlike
// IsDigits, this accepts an optional leading sign and a decimal point, e.g.
// "-3.14", "+7", "42".
func IsNumeric(value string) (string, error) {
	if !numericExpression.MatchString(value) {
		return value, ErrNotNumeric
	}

	return value, nil
}

// checkNumeric checks if the given string is a valid numeric string.
func checkNumeric(value reflect.Value) (reflect.Value, error) {
	_, err := IsNumeric(reflectString(value))
	return value, err
}

// makeNumeric makes a checker function for the numeric checker.
func makeNumeric(_ string) CheckFunc[reflect.Value] {
	return checkNumeric
}
