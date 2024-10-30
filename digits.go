// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package checker

import (
	"reflect"
	"unicode"
)

// CheckerDigits is the name of the checker.
const CheckerDigits = "digits"

// ResultNotDigits indicates that the given string contains non-digit characters.
const ResultNotDigits = "NOT_DIGITS"

// IsDigits checks if the given string consists of only digit characters.
func IsDigits(value string) Result {
	for _, c := range value {
		if !unicode.IsDigit(c) {
			return ResultNotDigits
		}
	}

	return ResultValid
}

// makeDigits makes a checker function for the digits checker.
func makeDigits(_ string) CheckFunc {
	return checkDigits
}

// checkDigits checks if the given string consists of only digit characters.
func checkDigits(value, _ reflect.Value) Result {
	if value.Kind() != reflect.String {
		panic("string expected")
	}

	return IsDigits(value.String())
}
