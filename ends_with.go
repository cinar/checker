// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

import (
	"reflect"
	"strings"
)

const (
	// nameEndsWith is the name of the ends-with check.
	nameEndsWith = "ends-with"
)

var (
	// ErrNotEndsWith indicates that the given string does not end with the given suffix.
	ErrNotEndsWith = NewCheckError("NOT_ENDS_WITH")
)

// IsEndsWith checks if the given string ends with the given suffix.
func IsEndsWith(suffix, value string) (string, error) {
	if !strings.HasSuffix(value, suffix) {
		return value, newEndsWithError(suffix)
	}

	return value, nil
}

// checkEndsWith makes a reflect-based check function for the given suffix.
func checkEndsWith(suffix string) CheckFunc[reflect.Value] {
	return func(value reflect.Value) (reflect.Value, error) {
		_, err := IsEndsWith(suffix, reflectString(value))
		return value, err
	}
}

// makeEndsWith makes a checker function for the ends-with checker.
func makeEndsWith(config string) CheckFunc[reflect.Value] {
	return checkEndsWith(config)
}

// newEndsWithError creates a new ends-with error with the expected suffix.
func newEndsWithError(suffix string) error {
	return NewCheckErrorWithData(
		ErrNotEndsWith.Code,
		map[string]interface{}{
			"suffix": suffix,
		},
	)
}
