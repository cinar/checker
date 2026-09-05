// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

import (
	"cmp"
	"reflect"
	"strconv"
)

const (
	// nameGt is the name of the greater than check.
	nameGt = "gt"
)

var (
	// ErrGt indicates that the value is not greater than the given value.
	ErrGt = NewCheckError("NOT_GT")
)

// IsGt checks if the value is strictly greater than the given value.
func IsGt[T cmp.Ordered](value, n T) (T, error) {
	if cmp.Compare(value, n) <= 0 {
		return value, newGtError(n)
	}

	return value, nil
}

// makeGt creates a greater than check function from a string parameter.
// Panics if the parameter cannot be parsed as a number.
func makeGt(params string) CheckFunc[reflect.Value] {
	n, err := strconv.ParseFloat(params, 64)
	if err != nil {
		panic("unable to parse params as float")
	}

	return func(value reflect.Value) (reflect.Value, error) {
		v := reflect.Indirect(value)

		switch {
		case v.CanInt():
			_, err := IsGt(float64(v.Int()), n)
			return v, err

		case v.CanUint():
			_, err := IsGt(float64(v.Uint()), n)
			return v, err

		case v.CanFloat():
			_, err := IsGt(v.Float(), n)
			return v, err

		default:
			panic("value is not numeric")
		}
	}
}

// newGtError creates a new greater than error with the given value.
func newGtError[T cmp.Ordered](n T) error {
	return NewCheckErrorWithData(
		ErrGt.Code,
		map[string]interface{}{
			"n": n,
		},
	)
}
