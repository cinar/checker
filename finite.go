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
	// nameFinite is the name of the finite number check.
	nameFinite = "finite"
)

var (
	// ErrFinite indicates that the value is not a finite number.
	ErrFinite = NewCheckError("NOT_FINITE")
)

// IsFinite checks if the given value is neither NaN nor an infinity.
func IsFinite(value float64) (float64, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return value, ErrFinite
	}

	return value, nil
}

// makeFinite creates a finite number check function.
func makeFinite(_ string) CheckFunc[reflect.Value] {
	return func(value reflect.Value) (reflect.Value, error) {
		v := reflect.Indirect(value)

		switch {
		case v.CanInt(), v.CanUint():
			// An int/uint kind can never hold NaN or an infinity.
			return v, nil

		case v.CanFloat():
			_, err := IsFinite(v.Float())
			return v, err

		default:
			panic("value is not numeric")
		}
	}
}
