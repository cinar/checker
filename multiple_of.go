// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

import (
	"math"
	"reflect"
	"strconv"
)

const (
	// nameMultipleOf is the name of the multiple of check.
	nameMultipleOf = "multiple-of"

	// multipleOfTolerance is the floating-point tolerance IsMultipleOf allows
	// between value/n and its nearest whole number, to absorb the rounding
	// error inherent in float64 division (e.g. 0.3 is a multiple of 0.1).
	// Not exact for values or divisors of extreme magnitude.
	multipleOfTolerance = 1e-9
)

var (
	// ErrMultipleOf indicates that the value is not a multiple of the given value.
	ErrMultipleOf = NewCheckError("NOT_MULTIPLE_OF")
)

// IsMultipleOf checks if the given value is a multiple of n, within a small
// floating-point tolerance (see multipleOfTolerance).
func IsMultipleOf(value, n float64) (float64, error) {
	q := value / n
	if math.Abs(q-math.Round(q)) >= multipleOfTolerance {
		return value, newMultipleOfError(n)
	}

	return value, nil
}

// makeMultipleOf creates a multiple of check function from a string parameter.
// Panics if the parameter cannot be parsed as a non-zero number.
func makeMultipleOf(params string) CheckFunc[reflect.Value] {
	n, err := strconv.ParseFloat(params, 64)
	if err != nil {
		panic("unable to parse params as float")
	}

	if n == 0 {
		panic("multiple-of divisor cannot be zero")
	}

	return func(value reflect.Value) (reflect.Value, error) {
		v := reflect.Indirect(value)

		switch {
		case v.CanInt():
			_, err := IsMultipleOf(float64(v.Int()), n)
			return v, err

		case v.CanUint():
			_, err := IsMultipleOf(float64(v.Uint()), n)
			return v, err

		case v.CanFloat():
			_, err := IsMultipleOf(v.Float(), n)
			return v, err

		default:
			panic("value is not numeric")
		}
	}
}

// newMultipleOfError creates a new multiple of error with the given divisor.
func newMultipleOfError(n float64) error {
	return NewCheckErrorWithData(
		ErrMultipleOf.Code,
		map[string]interface{}{
			"n": n,
		},
	)
}
