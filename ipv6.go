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

// tagIPV6 is the tag of the checker.
const tagIPV6 = "ipv6"

// ErrNotIPV6 indicates that the given value is not an IPv6 address.
var ErrNotIPV6 = errors.New("please enter a valid IPv6 address")

// IsIPV6 checks if the given value is an IPv6 address.
func IsIPV6(value string) error {
	ip := net.ParseIP(value)
	if ip == nil {
		return ErrNotIPV6
	}

	if ip.To4() != nil {
		return ErrNotIPV6
	}

	return nil
}

// makeIPV6 makes a checker function for the ipV6 checker.
func makeIPV6(_ string) CheckFunc {
	return checkIPV6
}

// checkIPV6 checks if the given value is an IPv6 address.
func checkIPV6(value, _ reflect.Value) error {
	if value.Kind() != reflect.String {
		panic("string expected")
	}

	return IsIPV6(value.String())
}
