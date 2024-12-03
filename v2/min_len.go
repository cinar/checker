// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

import (
	"reflect"
	"strconv"
)

const (
	// nameMinLen is the name of the minimum length check.
	nameMinLen = "min-len"
)

var (
	// ErrMinLen indicates that the value's length is less than the specified minimum.
	ErrMinLen = NewCheckError("MIN_LEN")
)

// MinLen checks if the length of the given value (string, slice, or map) is at least n.
// Returns an error if the length is less than n.
func MinLen[T any](n int) CheckFunc[T] {
	return func(value T) (T, error) {
		v, ok := any(value).(reflect.Value)
		if !ok {
			v = reflect.ValueOf(value)
		}

		v = reflect.Indirect(v)

		if v.Len() < n {
			return value, ErrMinLen
		}

		return value, nil
	}
}

// makeMinLen creates a minimum length check function from a string parameter.
// Panics if the parameter cannot be parsed as an integer.
func makeMinLen(params string) CheckFunc[reflect.Value] {
	n, err := strconv.Atoi(params)
	if err != nil {
		panic("unable to parse min length")
	}

	return MinLen[reflect.Value](n)
}
