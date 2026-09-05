// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

import (
	"fmt"
	"reflect"
	"strings"
)

const (
	// nameOneOf is the name of the one-of check.
	nameOneOf = "oneof"
)

var (
	// ErrNotOneOf indicates that the value does not match any of the allowed values.
	ErrNotOneOf = NewCheckError("NOT_ONE_OF")
)

// IsOneOf checks if the value is equal to one of the allowed values.
func IsOneOf[T comparable](value T, allowed ...T) (T, error) {
	for _, a := range allowed {
		if value == a {
			return value, nil
		}
	}

	return value, newOneOfError(allowed)
}

// checkOneOf makes a reflect-based check function for the given allowed values.
func checkOneOf(allowed []string) CheckFunc[reflect.Value] {
	return func(value reflect.Value) (reflect.Value, error) {
		_, err := IsOneOf(reflectString(value), allowed...)
		return value, err
	}
}

// makeOneOf makes a checker function from a comma-separated list of allowed
// values, e.g. "oneof:admin,user,guest". A single value is comma-separated,
// not space-separated, because checker tag fields are already split on
// whitespace before an individual checker's own params are parsed (the same
// convention "credit-card:visa,mastercard" uses). Panics if no allowed
// values are given.
func makeOneOf(config string) CheckFunc[reflect.Value] {
	if config == "" {
		panic("oneof requires at least one allowed value, e.g. oneof:admin,user,guest")
	}

	return checkOneOf(strings.Split(config, ","))
}

// newOneOfError creates a new one-of error listing the allowed values.
func newOneOfError[T comparable](allowed []T) error {
	names := make([]string, len(allowed))
	for i, a := range allowed {
		names[i] = fmt.Sprint(a)
	}

	return NewCheckErrorWithData(
		ErrNotOneOf.Code,
		map[string]interface{}{
			"allowed": strings.Join(names, ", "),
		},
	)
}
