// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package checker

import (
	"errors"
	"reflect"
	"unicode"
)

// tagDigits is the tag of the checker.
const tagDigits = "digits"

// ErrNotDigits indicates that the given string contains non-digit characters.
var ErrNotDigits = errors.New("please enter a valid number")

// IsDigits checks if the given string consists of only digit characters.
func IsDigits(value string) error {
	for _, c := range value {
		if !unicode.IsDigit(c) {
			return ErrNotDigits
		}
	}

	return nil
}

// makeDigits makes a checker function for the digits checker.
func makeDigits(_ string) CheckFunc {
	return checkDigits
}

// checkDigits checks if the given string consists of only digit characters.
func checkDigits(value, _ reflect.Value) error {
	if value.Kind() != reflect.String {
		panic("string expected")
	}

	return IsDigits(value.String())
}
