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
	// nameTrimSpace is the name of the trim normalizer.
	nameTrimSpace = "trim"
)

// TrimSpace returns the value of the string with whitespace removed from both ends.
func TrimSpace(value string) (string, error) {
	return strings.TrimSpace(value), nil
}

// reflectTrimSpace returns the value of the string with whitespace removed from both ends.
func reflectTrimSpace(value reflect.Value) (reflect.Value, error) {
	newValue, err := TrimSpace(value.Interface().(string))
	return reflect.ValueOf(newValue), err
}

// makeTrimSpace returns the trim space normalizer function.
func makeTrimSpace(_ string) CheckFunc[reflect.Value] {
	return reflectTrimSpace
}
