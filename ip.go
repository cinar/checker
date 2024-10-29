// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package checker

import (
	"errors"
	"net"
	"reflect"
)

// CheckerIP is the name of the checker.
const CheckerIP = "ip"

// ErrNotIP indicates that the given value is not an IP address.
var ErrNotIP = errors.New("please enter a valid IP address")

// IsIP checks if the given value is an IP address.
func IsIP(value string) error {
	ip := net.ParseIP(value)
	if ip == nil {
		return ErrNotIP
	}

	return nil
}

// makeIP makes a checker function for the ip checker.
func makeIP(_ string) CheckFunc {
	return checkIP
}

// checkIP checks if the given value is an IP address.
func checkIP(value, _ reflect.Value) error {
	if value.Kind() != reflect.String {
		panic("string expected")
	}

	return IsIP(value.String())
}
