// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

import (
	"reflect"
)

const (
	// nameNe is the name of the not equal to check.
	nameNe = "ne"
)

var (
	// ErrEq indicates that the value unexpectedly equals the forbidden value.
	ErrEq = NewCheckError("EQ")
)

// IsNe checks if the value is not equal to the forbidden value.
func IsNe[T comparable](value, forbidden T) (T, error) {
	if value == forbidden {
		return value, newNeError(forbidden)
	}

	return value, nil
}

// checkNe makes a reflect-based check function for the given forbidden value.
func checkNe(forbidden string) CheckFunc[reflect.Value] {
	return func(value reflect.Value) (reflect.Value, error) {
		_, err := IsNe(reflectString(value), forbidden)
		return value, err
	}
}

// makeNe makes a checker function for the ne checker.
func makeNe(config string) CheckFunc[reflect.Value] {
	return checkNe(config)
}

// newNeError creates a new not equal to error with the forbidden value.
func newNeError[T comparable](forbidden T) error {
	return NewCheckErrorWithData(
		ErrEq.Code,
		map[string]interface{}{
			"forbidden": forbidden,
		},
	)
}
