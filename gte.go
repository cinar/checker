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
	// nameGte is the name of the greater than or equal to check.
	nameGte = "gte"
)

var (
	// ErrGte indicates that the value is not greater than or equal to the given value.
	ErrGte = NewCheckError("NOT_GTE")
)

// IsGte checks if the value is greater than or equal to the given value.
func IsGte[T cmp.Ordered](value, n T) (T, error) {
	if cmp.Compare(value, n) < 0 {
		return value, newGteError(n)
	}

	return value, nil
}

// makeGte creates a greater than or equal to check function from a string parameter.
// Panics if the parameter cannot be parsed as a number.
func makeGte(params string) CheckFunc[reflect.Value] {
	n, err := strconv.ParseFloat(params, 64)
	if err != nil {
		panic("unable to parse params as float")
	}

	return func(value reflect.Value) (reflect.Value, error) {
		v := reflect.Indirect(value)

		switch {
		case v.CanInt():
			_, err := IsGte(float64(v.Int()), n)
			return v, err

		case v.CanFloat():
			_, err := IsGte(v.Float(), n)
			return v, err

		default:
			panic("value is not numeric")
		}
	}
}

// newGteError creates a new greater than or equal to error with the given value.
func newGteError[T cmp.Ordered](n T) error {
	return NewCheckErrorWithData(
		ErrGte.Code,
		map[string]interface{}{
			"n": n,
		},
	)
}
