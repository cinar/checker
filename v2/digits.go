// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

import (
    "reflect"
    "unicode"
)

const (
    // nameDigits is the name of the digits check.
    nameDigits = "digits"
)

var (
    // ErrNotDigits indicates that the given value is not a valid digits string.
    ErrNotDigits = NewCheckError("Digits")
)

// IsDigits checks if the value contains only digit characters.
func IsDigits(value string) (string, error) {
    for _, r := range value {
        if !unicode.IsDigit(r) {
            return value, ErrNotDigits
        }
    }
    return value, nil
}

// checkDigits checks if the value contains only digit characters.
func checkDigits(value reflect.Value) (reflect.Value, error) {
    _, err := IsDigits(value.Interface().(string))
    return value, err
}

// makeDigits makes a checker function for the digits checker.
func makeDigits(_ string) CheckFunc[reflect.Value] {
    return checkDigits
}