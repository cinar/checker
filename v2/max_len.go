// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

import (
	"fmt"
	"reflect"
	"strconv"
)

const (
	// nameMaxLen is the name of the maximum length check.
	nameMaxLen = "max-len"
)

// MaxLenError defines the maximum length error.
type MaxLenError struct {
	// MaxLen is the maximum number of characters required.
	MaxLen int
}

// newMaxLenError initializes a new maximum length error based on the given maximum length.
func newMaxLenError(MaxLen int) *MaxLenError {
	return &MaxLenError{
		MaxLen: MaxLen,
	}
}

// Error provides the string representation of the error.
func (m *MaxLenError) Error() string {
	return fmt.Sprintf("must be at most %d characters", m.MaxLen)
}

// MaxLen checks if a string has at least n characters. It returns an
// error if the string is shorter than n.
func MaxLen(n int) CheckFunc[string] {
	return func(value string) (string, error) {
		if len(value) > n {
			return value, newMaxLenError(n)
		}

		return value, nil
	}
}

// makeMaxLen returns the maximum length check function.
func makeMaxLen(params string) ReflectCheckFunc {
	n, err := strconv.Atoi(params)
	if err != nil {
		panic("unable to parse max length")
	}

	check := MaxLen(n)

	return func(value reflect.Value) error {
		value = reflect.Indirect(value)
		if value.Kind() != reflect.String {
			panic("expected string")
		}

		_, err := check(value.String())

		return err
	}
}
