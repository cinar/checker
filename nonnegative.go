// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

import (
	"cmp"
	"reflect"
)

const (
	// nameNonnegative is the name of the non-negative number check.
	nameNonnegative = "nonnegative"
)

var (
	// ErrNonnegative indicates that the value is a negative number.
	ErrNonnegative = NewCheckError("NOT_NONNEGATIVE")
)

// IsNonnegative checks if the given value is greater than or equal to zero.
func IsNonnegative[T cmp.Ordered](value T) (T, error) {
	var zero T

	if _, err := IsGte(value, zero); err != nil {
		return value, ErrNonnegative
	}

	return value, nil
}

// makeNonnegative creates a non-negative number check function.
func makeNonnegative(_ string) CheckFunc[reflect.Value] {
	return func(value reflect.Value) (reflect.Value, error) {
		v := reflect.Indirect(value)

		switch {
		case v.CanInt():
			_, err := IsNonnegative(float64(v.Int()))
			return v, err

		case v.CanUint():
			// An unsigned integer is always non-negative.
			return v, nil

		case v.CanFloat():
			_, err := IsNonnegative(v.Float())
			return v, err

		default:
			panic("value is not numeric")
		}
	}
}
