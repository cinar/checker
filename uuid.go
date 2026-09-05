// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

import (
	"reflect"
)

const (
	// nameUUID is the name of the UUID check.
	nameUUID = "uuid"

	// uuidExpression matches an RFC 4122 UUID: 32 hex digits grouped
	// 8-4-4-4-12 and separated by hyphens, case-insensitive. This accepts
	// any UUID version/variant, not just v4.
	uuidExpression = "^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$"
)

var (
	// ErrNotUUID indicates that the given string is not a valid UUID.
	ErrNotUUID = NewCheckError("NOT_UUID")
)

// IsUUID checks if the given string is a valid RFC 4122 UUID.
func IsUUID(value string) (string, error) {
	if _, err := IsRegexp(uuidExpression, value); err != nil {
		return value, ErrNotUUID
	}

	return value, nil
}

// checkUUID checks if the given string is a valid UUID.
func checkUUID(value reflect.Value) (reflect.Value, error) {
	_, err := IsUUID(reflectString(value))
	return value, err
}

// makeUUID makes a checker function for the UUID checker.
func makeUUID(_ string) CheckFunc[reflect.Value] {
	return checkUUID
}
