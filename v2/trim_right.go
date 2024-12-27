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
	// nameTrimRight is the name of the trim right normalizer.
	nameTrimRight = "trim-right"
)

// TrimRight returns the value of the string with whitespace removed from the end.
func TrimRight(value string) (string, error) {
	return strings.TrimRight(value, " \t"), nil
}

// reflectTrimRight returns the value of the string with whitespace removed from the end.
func reflectTrimRight(value reflect.Value) (reflect.Value, error) {
	newValue, err := TrimRight(value.Interface().(string))
	return reflect.ValueOf(newValue), err
}

// makeTrimRight returns the trim right normalizer function.
func makeTrimRight(_ string) CheckFunc[reflect.Value] {
	return reflectTrimRight
}
