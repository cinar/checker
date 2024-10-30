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

// tagMaxLength is the tag of the checker.
const tagMaxLength = "max-length"

// IsMaxLength checks if the length of the given value is less than the given maximum length.
func IsMaxLength(value interface{}, maxLength int) error {
	return checkMaxLength(reflect.Indirect(reflect.ValueOf(value)), reflect.ValueOf(nil), maxLength)
}

// makeMaxLength makes a checker function for the max length checker.
func makeMaxLength(config string) CheckFunc {
	maxLength, err := strconv.Atoi(config)
	if err != nil {
		panic("unable to parse max length value")
	}

	return func(value, parent reflect.Value) error {
		return checkMaxLength(value, parent, maxLength)
	}
}

// checkMaxLength checks if the length of the given value is less than the given maximum length.
// The function uses the reflect.Value.Len() function to determaxe the length of the value.
func checkMaxLength(value, _ reflect.Value, maxLength int) error {
	if value.Len() > maxLength {
		return fmt.Errorf("please enter %d characters or less", maxLength-1)
	}

	return nil
}
