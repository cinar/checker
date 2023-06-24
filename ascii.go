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

// CheckerASCII is the name of the checker.
const CheckerASCII = "ascii"

// ResultNotASCII indicates that the given string contains non-ASCII characters.
const ResultNotASCII = "NOT_ASCII"

// IsASCII checks if the given string consists of only ASCII characters.
func IsASCII(value string) Result {
	for _, c := range value {
		if c > unicode.MaxASCII {
			return ResultNotASCII
		}
	}

	return ResultValid
}

// makeASCII makes a checker function for the ASCII checker.
func makeASCII(_ string) CheckFunc {
	return checkASCII
}

// checkASCII checks if the given string consists of only ASCII characters.
func checkASCII(value, _ reflect.Value) Result {
	if value.Kind() != reflect.String {
		panic("string expected")
	}

	return IsASCII(value.String())
}
