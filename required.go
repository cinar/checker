// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package checker

import (
	"errors"
	"reflect"
)

// tagRequired is the tag of the checker.
const tagRequired = "required"

// ErrRequired indicates that the required value is missing.
var ErrRequired = errors.New("is required")

// IsRequired checks if the given required value is present.
func IsRequired(v interface{}) error {
	return checkRequired(reflect.ValueOf(v), reflect.ValueOf(nil))
}

// makeRequired makes a checker function for required.
func makeRequired(_ string) CheckFunc {
	return checkRequired
}

// checkRequired checks if the required value is present.
func checkRequired(value, _ reflect.Value) error {
	if value.IsZero() {
		return ErrRequired
	}

	k := value.Kind()

	if (k == reflect.Array || k == reflect.Map || k == reflect.Slice) && value.Len() == 0 {
		return ErrRequired
	}

	return nil
}
