// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

import (
	"encoding/base64"
	"reflect"
)

const (
	// nameBase64 is the name of the base64 check.
	nameBase64 = "base64"
)

var (
	// ErrNotBase64 indicates that the given string is not a valid base64-encoded string.
	ErrNotBase64 = NewCheckError("NOT_BASE64")
)

// IsBase64 checks if the given string is a valid standard (RFC 4648) base64-encoded string.
func IsBase64(value string) (string, error) {
	if len(value) == 0 {
		return value, ErrNotBase64
	}

	if _, err := base64.StdEncoding.DecodeString(value); err != nil {
		return value, ErrNotBase64
	}

	return value, nil
}

// checkBase64 checks if the given string is a valid base64-encoded string.
func checkBase64(value reflect.Value) (reflect.Value, error) {
	_, err := IsBase64(reflectString(value))
	return value, err
}

// makeBase64 makes a checker function for the base64 checker.
func makeBase64(_ string) CheckFunc[reflect.Value] {
	return checkBase64
}
