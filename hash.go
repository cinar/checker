// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

import (
	"fmt"
	"reflect"
)

const (
	// nameHash is the name of the hash check.
	nameHash = "hash"
)

var (
	// ErrNotHash indicates that the given value is not a valid hash for the given algorithm.
	ErrNotHash = NewCheckError("NOT_HASH")

	// hashLengths is a map of the hex-encoded digest length for each supported hash algorithm.
	hashLengths = map[string]int{
		"md5":    32,
		"sha1":   40,
		"sha256": 64,
		"sha384": 96,
		"sha512": 128,
	}
)

// IsHash checks if the value is a valid hex-encoded hash for the given algorithm.
// The supported algorithms are md5, sha1, sha256, sha384, and sha512. Panics if
// the algorithm is not one of these, as that indicates a configuration error.
func IsHash(algorithm, value string) (string, error) {
	length, ok := hashLengths[algorithm]
	if !ok {
		panic(fmt.Sprintf("unknown hash algorithm %s", algorithm))
	}

	if len(value) != length {
		return value, newHashError(algorithm)
	}

	if _, err := IsHex(value); err != nil {
		return value, newHashError(algorithm)
	}

	return value, nil
}

// makeHash makes a hash check function for the given algorithm parameter.
func makeHash(params string) CheckFunc[reflect.Value] {
	return func(value reflect.Value) (reflect.Value, error) {
		_, err := IsHash(params, reflectString(value))
		return value, err
	}
}

// newHashError creates a new hash error with the given algorithm.
func newHashError(algorithm string) error {
	return NewCheckErrorWithData(
		ErrNotHash.Code,
		map[string]interface{}{
			"algorithm": algorithm,
		},
	)
}
