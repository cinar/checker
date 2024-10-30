// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package checker

import (
	"reflect"
	"strings"
)

// NormalizerLower is the name of the normalizer.
const NormalizerLower = "lower"

// makeLower makes a normalizer function for the lower normalizer.
func makeLower(_ string) CheckFunc {
	return normalizeLower
}

// normalizeLower maps all Unicode letters in the given value to their lower case.
func normalizeLower(value, _ reflect.Value) Result {
	if value.Kind() != reflect.String {
		panic("string expected")
	}

	value.SetString(strings.ToLower(value.String()))

	return ResultValid
}
