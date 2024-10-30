// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package checker

import (
	"net"
	"reflect"
)

// CheckerIP is the name of the checker.
const CheckerIP = "ip"

// ResultNotIP indicates that the given value is not an IP address.
const ResultNotIP = "NOT_IP"

// IsIP checks if the given value is an IP address.
func IsIP(value string) Result {
	ip := net.ParseIP(value)
	if ip == nil {
		return ResultNotIP
	}

	return ResultValid
}

// makeIP makes a checker function for the ip checker.
func makeIP(_ string) CheckFunc {
	return checkIP
}

// checkIP checks if the given value is an IP address.
func checkIP(value, _ reflect.Value) Result {
	if value.Kind() != reflect.String {
		panic("string expected")
	}

	return IsIP(value.String())
}
