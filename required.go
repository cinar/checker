// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

import "reflect"

const (
	// nameRequired is the name of the required check.
	nameRequired = "required"
)

var (
	// ErrRequired indicates that a required value was missing.
	ErrRequired = NewCheckError("REQUIRED")
)

// Required checks if the given value of type T is its zero value. It
// returns an error if the value is zero.
func Required[T any](value T) (T, error) {
	_, err := reflectRequired(reflect.ValueOf(value))
	return value, err
}

// reflectRequired checks if the given value is its zero value. It
// returns an error if the value is zero.
func reflectRequired(value reflect.Value) (reflect.Value, error) {
	var err error

	if value.IsZero() {
		err = ErrRequired
	}

	return value, err
}

// makeRequired returns the required check function.
func makeRequired(_ string) CheckFunc[reflect.Value] {
	return reflectRequired
}
