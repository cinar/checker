// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

import (
	"reflect"
	"strconv"
)

const (
	// nameNumeric is the name of the numeric check.
	nameNumeric = "numeric"
)

var (
	// ErrNotNumeric indicates that the given string is not a valid numeric string.
	ErrNotNumeric = NewCheckError("NOT_NUMERIC")
)

// IsNumeric checks if the given string is a valid numeric string, unlike
// IsDigits, this accepts an optional leading sign and a decimal point, e.g.
// "-3.14", "+7", "42".
func IsNumeric(value string) (string, error) {
	if _, err := strconv.ParseFloat(value, 64); err != nil {
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
