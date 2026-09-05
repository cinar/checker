// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

import (
	"reflect"
)

const (
	// nameEq is the name of the equal to check.
	nameEq = "eq"
)

var (
	// ErrNotEq indicates that the value does not equal the expected value.
	ErrNotEq = NewCheckError("NOT_EQ")
)

// IsEq checks if the value is equal to the expected value.
func IsEq[T comparable](value, expected T) (T, error) {
	if value != expected {
		return value, newEqError(expected)
	}

	return value, nil
}

// checkEq makes a reflect-based check function for the given expected value.
func checkEq(expected string) CheckFunc[reflect.Value] {
	return func(value reflect.Value) (reflect.Value, error) {
		_, err := IsEq(reflectString(value), expected)
		return value, err
	}
}

// makeEq makes a checker function for the eq checker.
func makeEq(config string) CheckFunc[reflect.Value] {
	return checkEq(config)
}

// newEqError creates a new equal to error with the expected value.
func newEqError[T comparable](expected T) error {
	return NewCheckErrorWithData(
		ErrNotEq.Code,
		map[string]interface{}{
			"expected": expected,
		},
	)
}
