// Copyright (c) 2023-2024 Onur Cinar.
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
	// nameAfter is the name of the after check.
	nameAfter = "after"
)

var (
	// ErrNotAfter indicates that the value is not after the given reference time.
	ErrNotAfter = NewCheckError("NOT_AFTER")
)

// IsAfter checks if the value, parsed using the given layout, is after the
// given reference time, which is parsed using the same layout. Panics if the
// reference cannot be parsed, as that indicates a configuration error. If the
// value itself cannot be parsed, ErrTime is returned.
func IsAfter(layout, reference, value string) (string, error) {
	layout = resolveTimeLayout(layout)

	referenceTime, err := time.Parse(layout, reference)
	if err != nil {
		panic(fmt.Sprintf("unable to parse after reference time: %v", err))
	}

	valueTime, err := time.Parse(layout, value)
	if err != nil {
		return value, ErrTime
	}

	if !valueTime.After(referenceTime) {
		return value, newAfterError(reference)
	}

	return value, nil
}

// makeAfter makes an after check function from the "layout:reference" parameter.
// Panics if the parameter does not contain both the layout and the reference time.
func makeAfter(params string) CheckFunc[reflect.Value] {
	layout, reference, found := strings.Cut(params, ":")
	if !found {
		panic("after requires a layout and a reference time, e.g. after:DateOnly:2024-01-01")
	}

	return func(value reflect.Value) (reflect.Value, error) {
		_, err := IsAfter(layout, reference, value.String())
		return value, err
	}
}

// newAfterError creates a new after error with the given reference time.
func newAfterError(reference string) error {
	return NewCheckErrorWithData(
		ErrNotAfter.Code,
		map[string]interface{}{
			"reference": reference,
		},
	)
}
