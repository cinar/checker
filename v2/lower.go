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
	// nameLower is the name of the lower normalizer.
	nameLower = "lower"
)

// Lower maps all Unicode letters in the given value to their lower case.
func Lower(value string) (string, error) {
	return strings.ToLower(value), nil
}

// reflectLower maps all Unicode letters in the given value to their lower case.
func reflectLower(value reflect.Value) (reflect.Value, error) {
	newValue, err := Lower(value.Interface().(string))
	return reflect.ValueOf(newValue), err
}

// makeLower returns the lower normalizer function.
func makeLower(_ string) CheckFunc[reflect.Value] {
	return reflectLower
}
