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
	// nameIPv4 is the name of the IPv4 check.
	nameIPv4 = "ipv4"
)

var (
	// ErrNotIPv4 indicates that the given value is not a valid IPv4 address.
	ErrNotIPv4 = NewCheckError("NOT_IPV4")
)

// IsIPv4 checks if the value is a valid IPv4 address.
func IsIPv4(value string) (string, error) {
	ip := net.ParseIP(value)
	if ip == nil || ip.To4() == nil {
		return value, ErrNotIPv4
	}
	return value, nil
}

// checkIPv4 checks if the value is a valid IPv4 address.
func checkIPv4(value reflect.Value) (reflect.Value, error) {
	_, err := IsIPv4(value.Interface().(string))
	return value, err
}

// makeIPv4 makes a checker function for the IPv4 checker.
func makeIPv4(_ string) CheckFunc[reflect.Value] {
	return checkIPv4
}
