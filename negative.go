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
	// nameNegative is the name of the negative number check.
	nameNegative = "negative"
)

var (
	// ErrNegative indicates that the value is not a negative number.
	ErrNegative = NewCheckError("NOT_NEGATIVE")
)

// IsNegative checks if the given value is strictly less than zero.
func IsNegative[T cmp.Ordered](value T) (T, error) {
	var zero T

	if _, err := IsLt(value, zero); err != nil {
		return value, ErrNegative
	}

	return value, nil
}

// makeNegative creates a negative number check function.
func makeNegative(_ string) CheckFunc[reflect.Value] {
	return func(value reflect.Value) (reflect.Value, error) {
		v := reflect.Indirect(value)

		switch {
		case v.CanInt():
			_, err := IsNegative(float64(v.Int()))
			return v, err

		case v.CanUint():
			_, err := IsNegative(float64(v.Uint()))
			return v, err

		case v.CanFloat():
			_, err := IsNegative(v.Float())
			return v, err

		default:
			panic("value is not numeric")
		}
	}
}
