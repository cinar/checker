// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

import (
	"reflect"
)

const (
	// nameHex is the name of the hex check.
	nameHex = "hex"
)

var (
	// ErrNotHex indicates that the given string contains hex characters.
	ErrNotHex = NewCheckError("NOT_HEX")
)

// IsHex checks if the given string consists of only hex characters.
func IsHex(value string) (string, error) {
	return IsRegexp("^[0-9a-fA-F]+$", value)
}

// isHex checks if the given string consists of only hex characters.
func isHex(value reflect.Value) (reflect.Value, error) {
	_, err := IsHex(value.Interface().(string))
	return value, err
}

// makeAlphanumeric makes a checker function for the alphanumeric checker.
func makeHex(_ string) CheckFunc[reflect.Value] {
	return isHex
}
