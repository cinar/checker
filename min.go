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

// CheckerMin is the name of the checker.
const CheckerMin = "min"

// ResultNotMin indicates that the given value is below the defined minimum.
const ResultNotMin = "NOT_MIN"

// IsMin checks if the given value is above than the given minimum.
func IsMin(value interface{}, min float64) Result {
	return checkMin(reflect.Indirect(reflect.ValueOf(value)), reflect.ValueOf(nil), min)
}

// makeMin makes a checker function for the min checker.
func makeMin(config string) CheckFunc {
	min, err := strconv.ParseFloat(config, 64)
	if err != nil {
		panic("unable to parse min")
	}

	return func(value, parent reflect.Value) Result {
		return checkMin(value, parent, min)
	}
}

// checkMin checks if the given value is greather than the given minimum.
func checkMin(value, _ reflect.Value, min float64) Result {
	n := numberOf(value)

	if n < min {
		return ResultNotMin
	}

	return ResultValid
}
