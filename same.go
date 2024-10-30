// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package checker

import (
	"errors"
	"reflect"
)

// tagSame is the tag of the checker.
const tagSame = "same"

// ErrNotSame indicates that the given two values are not equal to each other.
var ErrNotSame = errors.New("does not match the other")

// makeSame makes a checker function for the same checker.
func makeSame(config string) CheckFunc {
	return func(value, parent reflect.Value) error {
		return checkSame(value, parent, config)
	}
}

// checkSame checks if the given value is equal to the value of the field with the given name.
func checkSame(value, parent reflect.Value, name string) error {
	other := parent.FieldByName(name)

	if !other.IsValid() {
		panic("other field not found")
	}

	other = reflect.Indirect(other)

	if !value.Equal(other) {
		return ErrNotSame
	}

	return nil
}
