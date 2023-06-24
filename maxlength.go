// Package checker is a Go library for validating user input through struct tags.
//
// Copyright (c) 2023 Onur Cinar. All Rights Reserved.
// The source code is provided under MIT License.
//
// https://github.com/cinar/checker
package checker

import (
	"reflect"
	"strconv"
)

// CheckerMaxLength is the name of the checker.
const CheckerMaxLength = "max-length"

// ResultNotMaxLength indicates that the length of the given value is above the defined number.
const ResultNotMaxLength = "NOT_MAX_LENGTH"

// IsMaxLength checks if the length of the given value is less than the given maximum length.
func IsMaxLength(value interface{}, maxLength int) Result {
	return checkMaxLength(reflect.Indirect(reflect.ValueOf(value)), reflect.ValueOf(nil), maxLength)
}

// makeMaxLength makes a checker function for the max length checker.
func makeMaxLength(config string) CheckFunc {
	maxLength, err := strconv.Atoi(config)
	if err != nil {
		panic("unable to parse max length value")
	}

	return func(value, parent reflect.Value) Result {
		return checkMaxLength(value, parent, maxLength)
	}
}

// checkMaxLength checks if the length of the given value is less than the given maximum length.
// The function uses the reflect.Value.Len() function to determaxe the length of the value.
func checkMaxLength(value, _ reflect.Value, maxLength int) Result {
	if value.Len() > maxLength {
		return ResultNotMaxLength
	}

	return ResultValid
}
