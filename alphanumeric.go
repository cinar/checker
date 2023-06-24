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
	"unicode"
)

// CheckerAlphanumeric is the name of the checker.
const CheckerAlphanumeric = "alphanumeric"

// ResultNotAlphanumeric indicates that the given string contains non-alphanumeric characters.
const ResultNotAlphanumeric = "NOT_ALPHANUMERIC"

// IsAlphanumeric checks if the given string consists of only alphanumeric characters.
func IsAlphanumeric(value string) Result {
	for _, c := range value {
		if !unicode.IsDigit(c) && !unicode.IsLetter(c) {
			return ResultNotAlphanumeric
		}
	}

	return ResultValid
}

// makeAlphanumeric makes a checker function for the alphanumeric checker.
func makeAlphanumeric(_ string) CheckFunc {
	return checkAlphanumeric
}

// checkAlphanumeric checks if the given string consists of only alphanumeric characters.
func checkAlphanumeric(value, _ reflect.Value) Result {
	if value.Kind() != reflect.String {
		panic("string expected")
	}

	return IsAlphanumeric(value.String())
}
