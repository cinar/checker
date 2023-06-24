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

// NormalizerTrimLeft is the name of the normalizer.
const NormalizerTrimLeft = "trim-left"

// makeTrimLeft makes a normalizer function for the trim left normalizer.
func makeTrimLeft(_ string) CheckFunc {
	return normalizeTrimLeft
}

// normalizeTrim removes the whitespaces at the beginning of the given value.
func normalizeTrimLeft(value, _ reflect.Value) Result {
	if value.Kind() != reflect.String {
		panic("string expected")
	}

	value.SetString(strings.TrimLeft(value.String(), " \t"))

	return ResultValid
}
