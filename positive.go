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
	// namePositive is the name of the positive number check.
	namePositive = "positive"
)

var (
	// ErrPositive indicates that the value is not a positive number.
	ErrPositive = NewCheckError("NOT_POSITIVE")
)

// IsPositive checks if the given value is strictly greater than zero.
func IsPositive[T cmp.Ordered](value T) (T, error) {
	var zero T

	if _, err := IsGt(value, zero); err != nil {
		return value, ErrPositive
	}

	return value, nil
}

// makePositive creates a positive number check function.
func makePositive(_ string) CheckFunc[reflect.Value] {
	return func(value reflect.Value) (reflect.Value, error) {
		v := reflect.Indirect(value)

		switch {
		case v.CanInt():
			_, err := IsPositive(float64(v.Int()))
			return v, err

		case v.CanUint():
			_, err := IsPositive(float64(v.Uint()))
			return v, err

		case v.CanFloat():
			_, err := IsPositive(v.Float())
			return v, err

		default:
			panic("value is not numeric")
		}
	}
}
