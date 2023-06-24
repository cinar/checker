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

// CheckerIPV4 is the name of the checker.
const CheckerIPV4 = "ipv4"

// ResultNotIPV4 indicates that the given value is not an IPv4 address.
const ResultNotIPV4 = "NOT_IP_V4"

// IsIPV4 checks if the given value is an IPv4 address.
func IsIPV4(value string) Result {
	ip := net.ParseIP(value)
	if ip == nil {
		return ResultNotIPV4
	}

	if ip.To4() == nil {
		return ResultNotIPV4
	}

	return ResultValid
}

// makeIPV4 makes a checker function for the ipV4 checker.
func makeIPV4(_ string) CheckFunc {
	return checkIPV4
}

// checkIPV4 checks if the given value is an IPv4 address.
func checkIPV4(value, _ reflect.Value) Result {
	if value.Kind() != reflect.String {
		panic("string expected")
	}

	return IsIPV4(value.String())
}
