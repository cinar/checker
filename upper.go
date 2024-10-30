// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package checker

import (
	"reflect"
	"strings"
)

// tagUpper is the tag of the normalizer.
const tagUpper = "upper"

// makeUpper makes a normalizer function for the upper normalizer.
func makeUpper(_ string) CheckFunc {
	return normalizeUpper
}

// normalizeUpper maps all Unicode letters in the given value to their upper case.
func normalizeUpper(value, _ reflect.Value) error {
	if value.Kind() != reflect.String {
		panic("string expected")
	}

	value.SetString(strings.ToUpper(value.String()))

	return nil
}
