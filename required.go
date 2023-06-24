// Package checker is a Go library for validating user input through struct tags.
//
// https://github.com/cinar/checker
//
// Copyright 2023 Onur Cinar. All rights reserved. 
// Use of this source code is governed by a MIT-style 
// license that can be found in the LICENSE file. 
//
package checker

import "reflect"

// CheckerRequired is the name of the checker.
const CheckerRequired = "required"

// ResultRequired indicates that the required value is missing.
const ResultRequired Result = "REQUIRED"

// IsRequired checks if the given required value is present.
func IsRequired(v interface{}) Result {
	return checkRequired(reflect.ValueOf(v), reflect.ValueOf(nil))
}

// makeRequired makes a checker function for required.
func makeRequired(_ string) CheckFunc {
	return checkRequired
}

// checkRequired checks if the required value is present.
func checkRequired(value, _ reflect.Value) Result {
	if value.IsZero() {
		return ResultRequired
	}

	k := value.Kind()

	if (k == reflect.Array || k == reflect.Map || k == reflect.Slice) && value.Len() == 0 {
		return ResultRequired
	}

	return ResultValid
}
