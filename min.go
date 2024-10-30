// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package checker

import (
	"fmt"
	"reflect"
	"strconv"
)

// tagMin is the tag of the checker.
const tagMin = "min"

// IsMin checks if the given value is above than the given minimum.
func IsMin(value interface{}, min float64) error {
	return checkMin(reflect.Indirect(reflect.ValueOf(value)), reflect.ValueOf(nil), min)
}

// makeMin makes a checker function for the min checker.
func makeMin(config string) CheckFunc {
	min, err := strconv.ParseFloat(config, 64)
	if err != nil {
		panic("unable to parse min")
	}

	return func(value, parent reflect.Value) error {
		return checkMin(value, parent, min)
	}
}

// checkMin checks if the given value is greather than the given minimum.
func checkMin(value, _ reflect.Value, min float64) error {
	n := numberOf(value)

	if n < min {
		return fmt.Errorf("please enter a number less than %g", min)
	}

	return nil
}
