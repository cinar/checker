// Package checker is a Go library for validating user input through struct tags.
//
// https://github.com/cinar/checker
//
// Copyright 2023 Onur Cinar. All rights reserved. 
// Use of this source code is governed by a MIT-style 
// license that can be found in the LICENSE file. 
//
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
