// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

import (
	"reflect"
	"strings"
)

const (
	// nameTrimLeft is the name of the trim left normalizer.
	nameTrimLeft = "trim-left"
)

// TrimLeft returns the value of the string with whitespace removed from the beginning.
func TrimLeft(value string) (string, error) {
	return strings.TrimLeft(value, " \t"), nil
}

// reflectTrimLeft returns the value of the string with whitespace removed from the beginning.
func reflectTrimLeft(value reflect.Value) (reflect.Value, error) {
	newValue, err := TrimLeft(value.Interface().(string))
	return reflect.ValueOf(newValue), err
}

// makeTrimLeft returns the trim left normalizer function.
func makeTrimLeft(_ string) CheckFunc[reflect.Value] {
	return reflectTrimLeft
}
