// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

const (
	// nameBefore is the name of the before check.
	nameBefore = "before"
)

var (
	// ErrNotBefore indicates that the value is not before the given reference time.
	ErrNotBefore = NewCheckError("NOT_BEFORE")
)

// IsBefore checks if the value, parsed using the given layout, is before the
// given reference time, which is parsed using the same layout. Panics if the
// reference cannot be parsed, as that indicates a configuration error. If the
// value itself cannot be parsed, ErrTime is returned.
func IsBefore(layout, reference, value string) (string, error) {
	layout = resolveTimeLayout(layout)

	referenceTime, err := time.Parse(layout, reference)
	if err != nil {
		panic(fmt.Sprintf("unable to parse before reference time: %v", err))
	}

	valueTime, err := time.Parse(layout, value)
	if err != nil {
		return value, ErrTime
	}

	if !valueTime.Before(referenceTime) {
		return value, newBeforeError(reference)
	}

	return value, nil
}

// makeBefore makes a before check function from the "layout:reference" parameter.
// Panics if the parameter does not contain both the layout and the reference time.
func makeBefore(params string) CheckFunc[reflect.Value] {
	layout, reference, found := strings.Cut(params, ":")
	if !found {
		panic("before requires a layout and a reference time, e.g. before:DateOnly:2024-01-01")
	}

	return func(value reflect.Value) (reflect.Value, error) {
		_, err := IsBefore(layout, reference, value.String())
		return value, err
	}
}

// newBeforeError creates a new before error with the given reference time.
func newBeforeError(reference string) error {
	return NewCheckErrorWithData(
		ErrNotBefore.Code,
		map[string]interface{}{
			"reference": reference,
		},
	)
}
