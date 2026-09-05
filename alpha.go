// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

import (
	"reflect"
	"unicode"
)

const (
	// nameAlpha is the name of the alpha check.
	nameAlpha = "alpha"
)

var (
	// ErrNotAlpha indicates that the given string contains non-letter characters.
	ErrNotAlpha = NewCheckError("NOT_ALPHA")
)

// IsAlpha checks if the given string consists of only letters.
func IsAlpha(value string) (string, error) {
	for _, c := range value {
		if !unicode.IsLetter(c) {
			return value, ErrNotAlpha
		}
	}

	return value, nil
}

// checkAlpha checks if the given string consists of only letters.
func checkAlpha(value reflect.Value) (reflect.Value, error) {
	_, err := IsAlpha(reflectString(value))
	return value, err
}

// makeAlpha makes a checker function for the alpha checker.
func makeAlpha(_ string) CheckFunc[reflect.Value] {
	return checkAlpha
}
