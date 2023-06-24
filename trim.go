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

// NormalizerTrim is the name of the normalizer.
const NormalizerTrim = "trim"

// makeTrim makes a normalizer function for the trim normalizer.
func makeTrim(_ string) CheckFunc {
	return normalizeTrim
}

// normalizeTrim removes the whitespaces at the beginning and at the end of the given value.
func normalizeTrim(value, _ reflect.Value) Result {
	if value.Kind() != reflect.String {
		panic("string expected")
	}

	value.SetString(strings.Trim(value.String(), " \t"))

	return ResultValid
}
