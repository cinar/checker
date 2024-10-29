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

// CheckerMax is the name of the checker.
const CheckerMax = "max"

// IsMax checks if the given value is below than the given maximum.
func IsMax(value interface{}, max float64) error {
	return checkMax(reflect.Indirect(reflect.ValueOf(value)), reflect.ValueOf(nil), max)
}

// makeMax makes a checker function for the max checker.
func makeMax(config string) CheckFunc {
	max, err := strconv.ParseFloat(config, 64)
	if err != nil {
		panic("unable to parse max")
	}

	return func(value, parent reflect.Value) error {
		return checkMax(value, parent, max)
	}
}

// checkMax checks if the given value is less than the given maximum.
func checkMax(value, _ reflect.Value, max float64) error {
	n := numberOf(value)

	if n > max {
		return fmt.Errorf("please enter a number less than %g", max)
	}

	return nil
}
