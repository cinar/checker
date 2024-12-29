// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

import (
	"reflect"
	"unicode"
)

const (
	// nameASCII is the name of the ASCII check.
	nameASCII = "ascii"
)

var (
	// ErrNotASCII indicates that the given string contains non-ASCII characters.
	ErrNotASCII = NewCheckError("NOT_ASCII")
)

// IsASCII checks if the given string consists of only ASCII characters.
func IsASCII(value string) (string, error) {
	for _, c := range value {
		if c > unicode.MaxASCII {
			return value, ErrNotASCII
		}
	}

	return value, nil
}

// checkASCII checks if the given string consists of only ASCII characters.
func isASCII(value reflect.Value) (reflect.Value, error) {
	_, err := IsASCII(value.Interface().(string))
	return value, err
}

// makeASCII makes a checker function for the ASCII checker.
func makeASCII(_ string) CheckFunc[reflect.Value] {
	return isASCII
}
