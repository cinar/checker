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

// NormalizerTrimRight is the name of the normalizer.
const NormalizerTrimRight = "trim-right"

// makeTrimRight makes a normalizer function for the trim right normalizer.
func makeTrimRight(_ string) CheckFunc {
	return normalizeTrimRight
}

// normalizeTrimRight removes the whitespaces at the end of the given value.
func normalizeTrimRight(value, _ reflect.Value) Result {
	if value.Kind() != reflect.String {
		panic("string expected")
	}

	value.SetString(strings.TrimRight(value.String(), " \t"))

	return ResultValid
}
