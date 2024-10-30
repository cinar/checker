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

// tagAlphanumeric is the tag of the checker.
const tagAlphanumeric = "alphanumeric"

// ErrNotAlphanumeric indicates that the given string contains non-alphanumeric characters.
var ErrNotAlphanumeric = errors.New("please use only letters and numbers")

// IsAlphanumeric checks if the given string consists of only alphanumeric characters.
func IsAlphanumeric(value string) error {
	for _, c := range value {
		if !unicode.IsDigit(c) && !unicode.IsLetter(c) {
			return ErrNotAlphanumeric
		}
	}

	return nil
}

// makeAlphanumeric makes a checker function for the alphanumeric checker.
func makeAlphanumeric(_ string) CheckFunc {
	return checkAlphanumeric
}

// checkAlphanumeric checks if the given string consists of only alphanumeric characters.
func checkAlphanumeric(value, _ reflect.Value) error {
	if value.Kind() != reflect.String {
		panic("string expected")
	}

	return IsAlphanumeric(value.String())
}
