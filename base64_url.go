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
	// nameBase64URL is the name of the base64-url check.
	nameBase64URL = "base64-url"
)

var (
	// ErrNotBase64URL indicates that the given string is not a valid base64url-encoded string.
	ErrNotBase64URL = NewCheckError("NOT_BASE64_URL")
)

// IsBase64URL checks if the given string is a valid base64url (RFC 4648 §5) encoded string.
func IsBase64URL(value string) (string, error) {
	if len(value) == 0 {
		return value, ErrNotBase64URL
	}

	if _, err := base64.URLEncoding.DecodeString(value); err != nil {
		return value, ErrNotBase64URL
	}

	return value, nil
}

// checkBase64URL checks if the given string is a valid base64url-encoded string.
func checkBase64URL(value reflect.Value) (reflect.Value, error) {
	_, err := IsBase64URL(reflectString(value))
	return value, err
}

// makeBase64URL makes a checker function for the base64-url checker.
func makeBase64URL(_ string) CheckFunc[reflect.Value] {
	return checkBase64URL
}
