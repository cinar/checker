// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

import (
	"reflect"
)

const (
	// nameULID is the name of the ULID check.
	nameULID = "ulid"

	// ulidExpression matches a 26-character Crockford base32 ULID
	// (case-insensitive), excluding the ambiguous I, L, O, and U letters.
	ulidExpression = `(?i)^[0-7][0-9A-HJKMNP-TV-Z]{25}$`
)

var (
	// ErrNotULID indicates that the given string is not a valid ULID.
	ErrNotULID = NewCheckError("NOT_ULID")
)

// IsULID checks if the given string is a valid ULID, e.g. "01ARZ3NDEKTSV4RRFFQ69G5FAV".
func IsULID(value string) (string, error) {
	if _, err := IsRegexp(ulidExpression, value); err != nil {
		return value, ErrNotULID
	}

	return value, nil
}

// checkULID checks if the given string is a valid ULID.
func checkULID(value reflect.Value) (reflect.Value, error) {
	_, err := IsULID(reflectString(value))
	return value, err
}

// makeULID makes a checker function for the ULID checker.
func makeULID(_ string) CheckFunc[reflect.Value] {
	return checkULID
}
