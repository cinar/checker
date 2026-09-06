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
	// nameRequiredUnless is the name of the required unless check.
	nameRequiredUnless = "required-unless"
)

// IsRequiredUnless checks if value is required, unless conditionValue
// (typically another field's already-resolved value) stringifies to
// expected. It returns an error if value is missing while the condition is
// not met.
func IsRequiredUnless[T any](value T, conditionValue any, expected string) (T, error) {
	if fmt.Sprintf("%v", conditionValue) != expected {
		return Required(value)
	}

	return value, nil
}

// checkRequiredUnless is the reflect.Value-based variant used internally
// by required-unless's struct-tag checking, where the sibling field's type
// isn't known until runtime.
func checkRequiredUnless(parent, value reflect.Value, name, expected string) (reflect.Value, error) {
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
		return checkRequiredUnless(parent, value, name, expected)
	}
}
