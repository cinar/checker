// Package checker is a Go library for validating user input through struct tags.
//
// https://github.com/cinar/checker
//
// Copyright 2023 Onur Cinar. All rights reserved. 
// Use of this source code is governed by a MIT-style 
// license that can be found in the LICENSE file. 
//
package checker

import (
	"net"
	"reflect"
)

// CheckerCidr is the name of the checker.
const CheckerCidr = "cidr"

// ResultNotCidr indicates that the given value is not a valid CIDR.
const ResultNotCidr = "NOT_CIDR"

// IsCidr checker checks if the value is a valid CIDR notation IP address and prefix length.
func IsCidr(value string) Result {
	_, _, err := net.ParseCIDR(value)
	if err != nil {
		return ResultNotCidr
	}

	return ResultValid
}

// makeCidr makes a checker function for the ip checker.
func makeCidr(_ string) CheckFunc {
	return checkCidr
}

// checkCidr checker checks if the value is a valid CIDR notation IP address and prefix length.
func checkCidr(value, _ reflect.Value) Result {
	if value.Kind() != reflect.String {
		panic("string expected")
	}

	return IsCidr(value.String())
}
