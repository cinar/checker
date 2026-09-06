// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

import (
	"math"
	"reflect"
)

const (
	// nameInt is the name of the whole number check.
	nameInt = "int"
)

var (
	// ErrInt indicates that the value is not a whole number.
	ErrInt = NewCheckError("NOT_INT")
)

// IsInt checks if the given value is a whole number, with no fractional part.
func IsInt(value float64) (float64, error) {
	if value != math.Trunc(value) {
		return value, ErrInt
	}

	return value, nil
}

// makeInt creates a whole number check function.
func makeInt(_ string) CheckFunc[reflect.Value] {
	return func(value reflect.Value) (reflect.Value, error) {
		v := reflect.Indirect(value)

		switch {
		case v.CanInt(), v.CanUint():
			// An int/uint kind is always a whole number.
			return v, nil

		case v.CanFloat():
			_, err := IsInt(v.Float())
			return v, err

		default:
			panic("value is not numeric")
		}
	}
}
