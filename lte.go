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
	// nameLte is the name of the less than or equal to check.
	nameLte = "lte"
)

var (
	// ErrLte indicates that the value is not less than or equal to the given value.
	ErrLte = NewCheckError("NOT_LTE")
)

// IsLte checks if the value is less than or equal to the given value.
func IsLte[T cmp.Ordered](value, n T) (T, error) {
	if cmp.Compare(value, n) > 0 {
		return value, newLteError(n)
	}

	return value, nil
}

// makeLte creates a less than or equal to check function from a string parameter.
// Panics if the parameter cannot be parsed as a number.
func makeLte(params string) CheckFunc[reflect.Value] {
	n, err := strconv.ParseFloat(params, 64)
	if err != nil {
		panic("unable to parse params as float")
	}

	return func(value reflect.Value) (reflect.Value, error) {
		v := reflect.Indirect(value)

		switch {
		case v.CanInt():
			_, err := IsLte(float64(v.Int()), n)
			return v, err

		case v.CanFloat():
			_, err := IsLte(v.Float(), n)
			return v, err

		default:
			panic("value is not numeric")
		}
	}
}

// newLteError creates a new less than or equal to error with the given value.
func newLteError[T cmp.Ordered](n T) error {
	return NewCheckErrorWithData(
		ErrLte.Code,
		map[string]interface{}{
			"n": n,
		},
	)
}
