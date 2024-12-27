// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

import (
	"net/mail"
	"reflect"
)

const (
	// nameEmail is the name of the email check.
	nameEmail = "email"
)

var (
	// ErrNotEmail indicates that the given value is not a valid email address.
	ErrNotEmail = NewCheckError("Email")
)

// IsEmail checks if the value is a valid email address.
func IsEmail(value string) (string, error) {
	_, err := mail.ParseAddress(value)
	if err != nil {
		return value, ErrNotEmail
	}
	return value, nil
}

// checkEmail checks if the value is a valid email address.
func checkEmail(value reflect.Value) (reflect.Value, error) {
	_, err := IsEmail(value.Interface().(string))
	return value, err
}

// makeEmail makes a checker function for the email checker.
func makeEmail(_ string) CheckFunc[reflect.Value] {
	return checkEmail
}
