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
	// nameRequiredUnless is the name of the required unless check.
	nameRequiredUnless = "required-unless"
)

// IsRequiredUnless checks if the value is required, unless the named field on the
// parent struct is equal to the expected value. It returns an error if the value
// is missing while the condition is not met.
func IsRequiredUnless(parent, value reflect.Value, name, expected string) (reflect.Value, error) {
	field := lookupParentField(nameRequiredUnless, parent, name)

	if fmt.Sprintf("%v", field.Interface()) != expected {
		return reflectRequired(value)
	}

	return value, nil
}

// makeRequiredUnless creates a required unless check function from the "field:value" parameter.
// Panics if the parameter does not contain both the field name and the expected value.
func makeRequiredUnless(params string) CheckFieldFunc {
	name, expected, found := strings.Cut(params, ":")
	if !found {
		panic("required-unless requires a field name and value, e.g. required-unless:Field:Value")
	}

	return func(parent, value reflect.Value) (reflect.Value, error) {
		return IsRequiredUnless(parent, value, name, expected)
	}
}
