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

// CheckerMax is the name of the checker.
const CheckerMax = "max"

// ResultNotMax indicates that the given value is above the defined maximum.
const ResultNotMax = "NOT_MIN"

// IsMax checks if the given value is below than the given maximum.
func IsMax(value interface{}, max float64) Result {
	return checkMax(reflect.Indirect(reflect.ValueOf(value)), reflect.ValueOf(nil), max)
}

// makeMax makes a checker function for the max checker.
func makeMax(config string) CheckFunc {
	max, err := strconv.ParseFloat(config, 64)
	if err != nil {
		panic("unable to parse max")
	}

	return func(value, parent reflect.Value) Result {
		return checkMax(value, parent, max)
	}
}

// checkMax checks if the given value is less than the given maximum.
func checkMax(value, _ reflect.Value, max float64) Result {
	n := numberOf(value)

	if n > max {
		return ResultNotMax
	}

	return ResultValid
}
