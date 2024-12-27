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
	// nameUpper is the name of the upper normalizer.
	nameUpper = "upper"
)

// Upper maps all Unicode letters in the given value to their upper case.
func Upper(value string) (string, error) {
	return strings.ToUpper(value), nil
}

// reflectUpper maps all Unicode letters in the given value to their upper case.
func reflectUpper(value reflect.Value) (reflect.Value, error) {
	newValue, err := Upper(value.Interface().(string))
	return reflect.ValueOf(newValue), err
}

// makeUpper returns the upper normalizer function.
func makeUpper(_ string) CheckFunc[reflect.Value] {
	return reflectUpper
}
