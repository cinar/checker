// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

import "reflect"

var (
	// ErrRequired indicates that a required value was missing.
	ErrRequired = NewCheckError("REQUIRED")
)

// Required checks if the given value of type T is its zero value. It
// returns an error if the value is zero.
func Required[T any](value T) (T, error) {
	if reflect.ValueOf(value).IsZero() {
		return value, ErrRequired
	}

	return value, nil
}
