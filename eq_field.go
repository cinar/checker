// Copyright (c) 2023-2026 Onur Cinar.
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

// IsEqField checks if value is equal to other, typically another field's
// already-resolved value on the same struct. It returns an error, naming
// name, if the two values are not equal. Unlike checkEqField's tag-driven
// path (reflect.DeepEqual, so it can compare slice/map/struct field values
// too), this compares with ==, so T must be comparable; that matches IsEq
// and IsNe, and covers every real-world eq-field use (confirmation fields,
// scalar cross-checks).
func IsEqField[T comparable](value, other T, name string) (T, error) {
	if value != other {
		return value, newEqFieldError(name)
	}

	return value, nil
}

// checkEqField is the reflect.Value-based variant used internally by
// eq-field's struct-tag checking, where the sibling field's type isn't
// known until runtime. It keeps reflect.DeepEqual, unlike IsEqField's ==,
// so a slice/map/struct field tagged eq-field keeps working exactly as
// before.
func checkEqField(parent, value reflect.Value, name string) (reflect.Value, error) {
	field := lookupParentField(nameEqField, parent, name)

	if !reflect.DeepEqual(value.Interface(), field.Interface()) {
		return value, newEqFieldError(name)
	}

	return value, nil
}

// makeEqField creates an equal to field check function for the named field.
func makeEqField(params string) CheckFieldFunc {
	return func(parent, value reflect.Value) (reflect.Value, error) {
		return checkEqField(parent, value, params)
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
