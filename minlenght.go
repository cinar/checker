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

// tagMinLength is the tag of the checker.
const tagMinLength = "min-length"

// IsMinLength checks if the length of the given value is greather than the given minimum length.
func IsMinLength(value interface{}, minLength int) error {
	return checkMinLength(reflect.Indirect(reflect.ValueOf(value)), reflect.ValueOf(nil), minLength)
}

// makeMinLength makes a checker function for the min length checker.
func makeMinLength(config string) CheckFunc {
	minLength, err := strconv.Atoi(config)
	if err != nil {
		panic("unable to parse min length value")
	}

	return func(value, parent reflect.Value) error {
		return checkMinLength(value, parent, minLength)
	}
}

// checkMinLength checks if the length of the given value is greather than the given minimum length.
// The function uses the reflect.Value.Len() function to determine the length of the value.
func checkMinLength(value, _ reflect.Value, minLength int) error {
	if value.Len() < minLength {
		return fmt.Errorf("please enter at least %d characters", minLength)
	}

	return nil
}
