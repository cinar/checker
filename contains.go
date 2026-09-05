// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

import (
	"reflect"
	"strings"
)

const (
	// nameContains is the name of the contains check.
	nameContains = "contains"
)

var (
	// ErrNotContains indicates that the given string does not contain the given substring.
	ErrNotContains = NewCheckError("NOT_CONTAINS")
)

// IsContains checks if the given string contains the given substring.
func IsContains(substr, value string) (string, error) {
	if !strings.Contains(value, substr) {
		return value, newContainsError(substr)
	}

	return value, nil
}

// checkContains makes a reflect-based check function for the given substring.
func checkContains(substr string) CheckFunc[reflect.Value] {
	return func(value reflect.Value) (reflect.Value, error) {
		_, err := IsContains(substr, reflectString(value))
		return value, err
	}
}

// makeContains makes a checker function for the contains checker.
func makeContains(config string) CheckFunc[reflect.Value] {
	return checkContains(config)
}

// newContainsError creates a new contains error with the expected substring.
func newContainsError(substr string) error {
	return NewCheckErrorWithData(
		ErrNotContains.Code,
		map[string]interface{}{
			"substr": substr,
		},
	)
}
