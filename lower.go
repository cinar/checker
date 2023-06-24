// Package checker is a Go library for validating user input through struct tags.
//
// Copyright (c) 2023 Onur Cinar. All Rights Reserved.
// The source code is provided under MIT License.
//
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
