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
	// nameMaxLen is the name of the maximum length check.
	nameMaxLen = "max-len"
)

var (
	// ErrMaxLen indicates that the value's length is greater than the specified maximum.
	ErrMaxLen = NewCheckError("MAX_LEN")
)

// MaxLen checks if the length of the given value (string, slice, or map) is at most n.
// Returns an error if the length is greater than n.
func MaxLen[T any](n int) CheckFunc[T] {
	return func(value T) (T, error) {
		v, ok := any(value).(reflect.Value)
		if !ok {
			v = reflect.ValueOf(value)
		}

		v = reflect.Indirect(v)

		if v.Len() > n {
			return value, ErrMaxLen
		}

		return value, nil
	}
}

// makeMaxLen creates a maximum length check function from a string parameter.
// Panics if the parameter cannot be parsed as an integer.
func makeMaxLen(params string) CheckFunc[reflect.Value] {
	n, err := strconv.Atoi(params)
	if err != nil {
		panic("unable to parse max length")
	}

	return MaxLen[reflect.Value](n)
}
