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

// IsAfterField checks if the value, parsed using the given layout, is after
// the value of the named field on the parent struct, also parsed using the
// same layout. Unlike IsAfter, an unparsable reference does not panic: the
// reference here is another field's data, not checker configuration, so
// ErrTime is returned instead, same as an unparsable value.
func IsAfterField(parent, value reflect.Value, layout, name string) (reflect.Value, error) {
	field := lookupParentField(nameAfterField, parent, name)

	resolvedLayout := resolveTimeLayout(layout)
	reference := reflectString(field)

	referenceTime, err := time.Parse(resolvedLayout, reference)
	if err != nil {
		return value, ErrTime
	}

	valueTime, err := time.Parse(resolvedLayout, reflectString(value))
	if err != nil {
		return value, ErrTime
	}

	if !valueTime.After(referenceTime) {
		return value, newAfterError(reference)
	}

	return value, nil
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
		return IsAfterField(parent, value, layout, name)
	}
}
