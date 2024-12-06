// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

import (
	"reflect"
	"unicode"
)

const (
	// nameAlphanumeric is the name of the alphanumeric check.
	nameAlphanumeric = "alphanumeric"
)

var (
	// ErrNotAlphanumeric indicates that the given string contains non-alphanumeric characters.
	ErrNotAlphanumeric = NewCheckError("NOT_ALPHANUMERIC")
)

// IsAlphanumeric checks if the given string consists of only alphanumeric characters.
func IsAlphanumeric(value string) (string, error) {
	for _, c := range value {
		if !unicode.IsDigit(c) && !unicode.IsLetter(c) {
			return value, ErrNotAlphanumeric
		}
	}

	return value, nil
}

// checkAlphanumeric checks if the given string consists of only alphanumeric characters.
func isAlphanumeric(value reflect.Value) (reflect.Value, error) {
	_, err := IsAlphanumeric(value.Interface().(string))
	return value, err
}

// makeAlphanumeric makes a checker function for the alphanumeric checker.
func makeAlphanumeric(_ string) CheckFunc[reflect.Value] {
	return isAlphanumeric
}
