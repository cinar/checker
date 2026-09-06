// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

import (
	"reflect"
	"strings"
	"time"
)

const (
	// nameBeforeField is the name of the before field check.
	nameBeforeField = "before-field"
)

// IsBeforeField checks if value, parsed using the given layout, is before
// reference (typically another field's already-resolved value), also
// parsed using the same layout. Unlike IsBefore, an unparsable reference
// does not panic: the reference here is another field's data, not checker
// configuration, so ErrTime is returned instead, same as an unparsable
// value.
func IsBeforeField(layout, reference, value string) (string, error) {
	resolvedLayout := resolveTimeLayout(layout)

	referenceTime, err := time.Parse(resolvedLayout, reference)
	if err != nil {
		return value, ErrTime
	}

	valueTime, err := time.Parse(resolvedLayout, value)
	if err != nil {
		return value, ErrTime
	}

	if !valueTime.Before(referenceTime) {
		return value, newBeforeError(reference)
	}

	return value, nil
}

// checkBeforeField is the reflect.Value-based variant used internally by
// before-field's struct-tag checking: it resolves the named sibling field
// via parent, then delegates the actual comparison to IsBeforeField.
func checkBeforeField(parent, value reflect.Value, layout, name string) (reflect.Value, error) {
	field := lookupParentField(nameBeforeField, parent, name)
	reference := reflectString(field)

	newValue, err := IsBeforeField(layout, reference, reflectString(value))

	return reflect.ValueOf(newValue).Convert(value.Type()), err
}

// makeBeforeField makes a before field check function from the
// "layout:field" parameter. Panics if the parameter does not contain both
// the layout and the field name.
func makeBeforeField(params string) CheckFieldFunc {
	layout, name, found := strings.Cut(params, ":")
	if !found {
		panic("before-field requires a layout and a field name, e.g. before-field:DateOnly:ExpiresAt")
	}

	return func(parent, value reflect.Value) (reflect.Value, error) {
		return checkBeforeField(parent, value, layout, name)
	}
}
