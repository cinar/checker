// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

import "reflect"

const (
	// nameEqField is the name of the equal to field check.
	nameEqField = "eq-field"
)

var (
	// ErrEqField indicates that the value is not equal to the given field's value.
	ErrEqField = NewCheckError("NOT_EQ_FIELD")
)

// IsEqField checks if the value is equal to the value of the named field on the
// parent struct. It returns an error if the two values are not equal.
func IsEqField(parent, value reflect.Value, name string) (reflect.Value, error) {
	field := lookupParentField(nameEqField, parent, name)

	if !reflect.DeepEqual(value.Interface(), field.Interface()) {
		return value, newEqFieldError(name)
	}

	return value, nil
}

// makeEqField creates an equal to field check function for the named field.
func makeEqField(params string) CheckFieldFunc {
	return func(parent, value reflect.Value) (reflect.Value, error) {
		return IsEqField(parent, value, params)
	}
}

// newEqFieldError creates a new equal to field error with the given field name.
func newEqFieldError(name string) error {
	return NewCheckErrorWithData(
		ErrEqField.Code,
		map[string]interface{}{
			"field": name,
		},
	)
}
