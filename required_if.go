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
	// nameRequiredIf is the name of the required if check.
	nameRequiredIf = "required-if"
)

// IsRequiredIf checks if the value is required, given that the named field on the
// parent struct is equal to the expected value. It returns an error if the value
// is missing while the condition is met.
func IsRequiredIf(parent, value reflect.Value, name, expected string) (reflect.Value, error) {
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
		return IsRequiredIf(parent, value, name, expected)
	}
}
