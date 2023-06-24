// Package checker is a Go library for validating user input through struct tags.
//
// Copyright (c) 2023 Onur Cinar. All Rights Reserved.
// The source code is provided under MIT License.
//
// https://github.com/cinar/checker
package checker

import (
	"net"
	"reflect"
)

// CheckerIPV6 is the name of the checker.
const CheckerIPV6 = "ipv6"

// ResultNotIPV6 indicates that the given value is not an IPv6 address.
const ResultNotIPV6 = "NOT_IP_V6"

// IsIPV6 checks if the given value is an IPv6 address.
func IsIPV6(value string) Result {
	ip := net.ParseIP(value)
	if ip == nil {
		return ResultNotIPV6
	}

	if ip.To4() != nil {
		return ResultNotIPV6
	}

	return ResultValid
}

// makeIPV6 makes a checker function for the ipV6 checker.
func makeIPV6(_ string) CheckFunc {
	return checkIPV6
}

// checkIPV6 checks if the given value is an IPv6 address.
func checkIPV6(value, _ reflect.Value) Result {
	if value.Kind() != reflect.String {
		panic("string expected")
	}

	return IsIPV6(value.String())
}
