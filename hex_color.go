// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

import (
	"reflect"
)

const (
	// nameHexColor is the name of the hex-color check.
	nameHexColor = "hex-color"

	// hexColorExpression matches a "#" followed by 3, 4, 6, or 8 hex
	// digits, covering RGB, RGBA, RRGGBB, and RRGGBBAA forms.
	hexColorExpression = `^#(?:[0-9a-fA-F]{3,4}|[0-9a-fA-F]{6}|[0-9a-fA-F]{8})$`
)

var (
	// ErrNotHexColor indicates that the given string is not a valid hex color code.
	ErrNotHexColor = NewCheckError("NOT_HEX_COLOR")
)

// IsHexColor checks if the given string is a valid hex color code, e.g. "#ff0000".
func IsHexColor(value string) (string, error) {
	if _, err := IsRegexp(hexColorExpression, value); err != nil {
		return value, ErrNotHexColor
	}

	return value, nil
}

// checkHexColor checks if the given string is a valid hex color code.
func checkHexColor(value reflect.Value) (reflect.Value, error) {
	_, err := IsHexColor(reflectString(value))
	return value, err
}

// makeHexColor makes a checker function for the hex-color checker.
func makeHexColor(_ string) CheckFunc[reflect.Value] {
	return checkHexColor
}
