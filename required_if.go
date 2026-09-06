// Copyright (c) 2023-2026 Onur Cinar.
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
	// nameRequiredIf is the name of the required if check.
	nameRequiredIf = "required-if"
)

// IsRequiredIf checks if value is required, given that conditionValue
// (typically another field's already-resolved value) stringifies to
// expected. It returns an error if value is missing while the condition is
// met.
func IsRequiredIf[T any](value T, conditionValue any, expected string) (T, error) {
	if fmt.Sprintf("%v", conditionValue) == expected {
		return Required(value)
	}

	return value, nil
}

// checkRequiredIf is the reflect.Value-based variant used internally by
// required-if's struct-tag checking, where the sibling field's type isn't
// known until runtime.
func checkRequiredIf(parent, value reflect.Value, name, expected string) (reflect.Value, error) {
	field := lookupParentField(nameRequiredIf, parent, name)

	if fmt.Sprintf("%v", field.Interface()) == expected {
		return reflectRequired(value)
	}

	return value, nil
}

// makeRequiredIf creates a required if check function from the "field:value" parameter.
// Panics if the parameter does not contain both the field name and the expected value.
func makeRequiredIf(params string) CheckFieldFunc {
	name, expected, found := strings.Cut(params, ":")
	if !found {
		panic("required-if requires a field name and value, e.g. required-if:Field:Value")
	}

	return func(parent, value reflect.Value) (reflect.Value, error) {
		return checkRequiredIf(parent, value, name, expected)
	}
}
