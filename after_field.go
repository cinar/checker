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
	// nameAfterField is the name of the after field check.
	nameAfterField = "after-field"
)

// IsAfterField checks if value, parsed using the given layout, is after
// reference (typically another field's already-resolved value), also
// parsed using the same layout. Unlike IsAfter, an unparsable reference
// does not panic: the reference here is another field's data, not checker
// configuration, so ErrTime is returned instead, same as an unparsable
// value.
func IsAfterField(layout, reference, value string) (string, error) {
	resolvedLayout := resolveTimeLayout(layout)

	referenceTime, err := time.Parse(resolvedLayout, reference)
	if err != nil {
		return value, ErrTime
	}

	valueTime, err := time.Parse(resolvedLayout, value)
	if err != nil {
		return value, ErrTime
	}

	if !valueTime.After(referenceTime) {
		return value, newAfterError(reference)
	}

	return value, nil
}

// checkAfterField is the reflect.Value-based variant used internally by
// after-field's struct-tag checking: it resolves the named sibling field
// via parent, then delegates the actual comparison to IsAfterField.
func checkAfterField(parent, value reflect.Value, layout, name string) (reflect.Value, error) {
	field := lookupParentField(nameAfterField, parent, name)
	reference := reflectString(field)

	newValue, err := IsAfterField(layout, reference, reflectString(value))

	return reflect.ValueOf(newValue).Convert(value.Type()), err
}

// makeAfterField makes an after field check function from the
// "layout:field" parameter. Panics if the parameter does not contain both
// the layout and the field name.
func makeAfterField(params string) CheckFieldFunc {
	layout, name, found := strings.Cut(params, ":")
	if !found {
		panic("after-field requires a layout and a field name, e.g. after-field:DateOnly:BornAt")
	}

	return func(parent, value reflect.Value) (reflect.Value, error) {
		return checkAfterField(parent, value, layout, name)
	}
}
