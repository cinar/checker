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

// tagCidr is the tag of the checker.
const tagCidr = "cidr"

// ErrNotCidr indicates that the given value is not a valid CIDR.
var ErrNotCidr = errors.New("please enter a valid CIDR")

// IsCidr checker checks if the value is a valid CIDR notation IP address and prefix length.
func IsCidr(value string) error {
	_, _, err := net.ParseCIDR(value)
	if err != nil {
		return ErrNotCidr
	}

	return nil
}

// makeCidr makes a checker function for the ip checker.
func makeCidr(_ string) CheckFunc {
	return checkCidr
}

// checkCidr checker checks if the value is a valid CIDR notation IP address and prefix length.
func checkCidr(value, _ reflect.Value) error {
	if value.Kind() != reflect.String {
		panic("string expected")
	}

	return IsCidr(value.String())
}
