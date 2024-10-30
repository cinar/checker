// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package checker

import (
	"reflect"
	"strings"
)

// tagLower is the tag of the normalizer.
const tagLower = "lower"

// makeLower makes a normalizer function for the lower normalizer.
func makeLower(_ string) CheckFunc {
	return normalizeLower
}

// normalizeLower maps all Unicode letters in the given value to their lower case.
func normalizeLower(value, _ reflect.Value) error {
	if value.Kind() != reflect.String {
		panic("string expected")
	}

	value.SetString(strings.ToLower(value.String()))

	return nil
}
