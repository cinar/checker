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
	// nameStartsWith is the name of the starts-with check.
	nameStartsWith = "starts-with"
)

var (
	// ErrNotStartsWith indicates that the given string does not start with the given prefix.
	ErrNotStartsWith = NewCheckError("NOT_STARTS_WITH")
)

// IsStartsWith checks if the given string starts with the given prefix.
func IsStartsWith(prefix, value string) (string, error) {
	if !strings.HasPrefix(value, prefix) {
		return value, newStartsWithError(prefix)
	}

	return value, nil
}

// checkStartsWith makes a reflect-based check function for the given prefix.
func checkStartsWith(prefix string) CheckFunc[reflect.Value] {
	return func(value reflect.Value) (reflect.Value, error) {
		_, err := IsStartsWith(prefix, reflectString(value))
		return value, err
	}
}

// makeStartsWith makes a checker function for the starts-with checker.
func makeStartsWith(config string) CheckFunc[reflect.Value] {
	return checkStartsWith(config)
}

// newStartsWithError creates a new starts-with error with the expected prefix.
func newStartsWithError(prefix string) error {
	return NewCheckErrorWithData(
		ErrNotStartsWith.Code,
		map[string]interface{}{
			"prefix": prefix,
		},
	)
}
