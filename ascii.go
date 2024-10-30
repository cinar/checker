// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package checker

import (
	"errors"
	"reflect"
	"unicode"
)

// tagASCII is the tag of the checker.
const tagASCII = "ascii"

// ErrNotASCII indicates that the given string contains non-ASCII characters.
var ErrNotASCII = errors.New("please use standard English characters only")

// IsASCII checks if the given string consists of only ASCII characters.
func IsASCII(value string) error {
	for _, c := range value {
		if c > unicode.MaxASCII {
			return ErrNotASCII
		}
	}

	return nil
}

// makeASCII makes a checker function for the ASCII checker.
func makeASCII(_ string) CheckFunc {
	return checkASCII
}

// checkASCII checks if the given string consists of only ASCII characters.
func checkASCII(value, _ reflect.Value) error {
	if value.Kind() != reflect.String {
		panic("string expected")
	}

	return IsASCII(value.String())
}
