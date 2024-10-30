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

// tagIPV4 is the tag of the checker.
const tagIPV4 = "ipv4"

// ErrNotIPV4 indicates that the given value is not an IPv4 address.
var ErrNotIPV4 = errors.New("please enter a valid IPv4 address")

// IsIPV4 checks if the given value is an IPv4 address.
func IsIPV4(value string) error {
	ip := net.ParseIP(value)
	if ip == nil {
		return ErrNotIPV4
	}

	if ip.To4() == nil {
		return ErrNotIPV4
	}

	return nil
}

// makeIPV4 makes a checker function for the ipV4 checker.
func makeIPV4(_ string) CheckFunc {
	return checkIPV4
}

// checkIPV4 checks if the given value is an IPv4 address.
func checkIPV4(value, _ reflect.Value) error {
	if value.Kind() != reflect.String {
		panic("string expected")
	}

	return IsIPV4(value.String())
}
