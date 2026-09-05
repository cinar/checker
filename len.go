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
	// nameLen is the name of the exact length check.
	nameLen = "len"
)

var (
	// ErrLen indicates that the value's length is not the specified length.
	ErrLen = NewCheckError("NOT_LEN")
)

// Len checks if the length of the given value (string, slice, or map) is exactly n.
// Returns an error if the length is not n.
func Len[T any](n int) CheckFunc[T] {
	return func(value T) (T, error) {
		v, ok := any(value).(reflect.Value)
		if !ok {
			v = reflect.ValueOf(value)
		}

		v = reflect.Indirect(v)

		if v.Len() != n {
			return value, newLenError(n)
		}

		return value, nil
	}
}

// makeLen creates an exact length check function from a string parameter.
// Panics if the parameter cannot be parsed as an integer.
func makeLen(params string) CheckFunc[reflect.Value] {
	n, err := strconv.Atoi(params)
	if err != nil {
		panic("unable to parse len")
	}

	return Len[reflect.Value](n)
}

// newLenError creates a new exact length error with the expected length.
func newLenError(n int) error {
	return NewCheckErrorWithData(
		ErrLen.Code,
		map[string]interface{}{
			"len": n,
		},
	)
}
