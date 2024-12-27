// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

import (
	"net"
	"reflect"
)

const (
	// nameIPv6 is the name of the IPv6 check.
	nameIPv6 = "ipv6"
)

var (
	// ErrNotIPv6 indicates that the given value is not a valid IPv6 address.
	ErrNotIPv6 = NewCheckError("IPv6")
)

// IsIPv6 checks if the value is a valid IPv6 address.
func IsIPv6(value string) (string, error) {
	if net.ParseIP(value) == nil || net.ParseIP(value).To4() != nil {
		return value, ErrNotIPv6
	}
	return value, nil
}

// checkIPv6 checks if the value is a valid IPv6 address.
func checkIPv6(value reflect.Value) (reflect.Value, error) {
	_, err := IsIPv6(value.Interface().(string))
	return value, err
}

// makeIPv6 makes a checker function for the IPv6 checker.
func makeIPv6(_ string) CheckFunc[reflect.Value] {
	return checkIPv6
}
