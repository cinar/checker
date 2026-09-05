// Copyright (c) 2023-2026 Onur Cinar.
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
	// nameLt is the name of the less than check.
	nameLt = "lt"
)

var (
	// ErrLt indicates that the value is not less than the given value.
	ErrLt = NewCheckError("NOT_LT")
)

// IsLt checks if the value is strictly less than the given value.
func IsLt[T cmp.Ordered](value, n T) (T, error) {
	if cmp.Compare(value, n) >= 0 {
		return value, newLtError(n)
	}

	return value, nil
}

// makeLt creates a less than check function from a string parameter.
// Panics if the parameter cannot be parsed as a number.
func makeLt(params string) CheckFunc[reflect.Value] {
	n, err := strconv.ParseFloat(params, 64)
	if err != nil {
		panic("unable to parse params as float")
	}

	return func(value reflect.Value) (reflect.Value, error) {
		v := reflect.Indirect(value)

		switch {
		case v.CanInt():
			_, err := IsLt(float64(v.Int()), n)
			return v, err

		case v.CanUint():
			_, err := IsLt(float64(v.Uint()), n)
			return v, err

		case v.CanFloat():
			_, err := IsLt(v.Float(), n)
			return v, err

		default:
			panic("value is not numeric")
		}
	}
}

// newLtError creates a new less than error with the given value.
func newLtError[T cmp.Ordered](n T) error {
	return NewCheckErrorWithData(
		ErrLt.Code,
		map[string]interface{}{
			"n": n,
		},
	)
}
