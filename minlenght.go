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
	"strconv"
)

// CheckerMinLength is the name of the checker.
const CheckerMinLength = "min-length"

// ResultNotMinLength indicates that the length of the given value is below the defined number.
const ResultNotMinLength = "NOT_MIN_LENGTH"

// IsMinLength checks if the length of the given value is greather than the given minimum length.
func IsMinLength(value interface{}, minLength int) Result {
	return checkMinLength(reflect.Indirect(reflect.ValueOf(value)), reflect.ValueOf(nil), minLength)
}

// makeMinLength makes a checker function for the min length checker.
func makeMinLength(config string) CheckFunc {
	minLength, err := strconv.Atoi(config)
	if err != nil {
		panic("unable to parse min length value")
	}

	return func(value, parent reflect.Value) Result {
		return checkMinLength(value, parent, minLength)
	}
}

// checkMinLength checks if the length of the given value is greather than the given minimum length.
// The function uses the reflect.Value.Len() function to determine the length of the value.
func checkMinLength(value, _ reflect.Value, minLength int) Result {
	if value.Len() < minLength {
		return ResultNotMinLength
	}

	return ResultValid
}
