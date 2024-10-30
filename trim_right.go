// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package checker

import (
	"reflect"
	"strings"
)

// tagTrimRight is the tag of the normalizer.
const tagTrimRight = "trim-right"

// makeTrimRight makes a normalizer function for the trim right normalizer.
func makeTrimRight(_ string) CheckFunc {
	return normalizeTrimRight
}

// normalizeTrimRight removes the whitespaces at the end of the given value.
func normalizeTrimRight(value, _ reflect.Value) error {
	if value.Kind() != reflect.String {
		panic("string expected")
	}

	value.SetString(strings.TrimRight(value.String(), " \t"))

	return nil
}
