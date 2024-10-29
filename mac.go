// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package checker

import (
	"net"
	"reflect"
)

// CheckerMac is the name of the checker.
const CheckerMac = "mac"

// ResultNotMac indicates that the given value is not an MAC address.
const ResultNotMac = "NOT_MAC"

// IsMac checks if the given value is a valid an IEEE 802 MAC-48, EUI-48, EUI-64, or a 20-octet IP over InfiniBand link-layer address.
func IsMac(value string) Result {
	_, err := net.ParseMAC(value)
	if err != nil {
		return ResultNotMac
	}

	return ResultValid
}

// makeMac makes a checker function for the ip checker.
func makeMac(_ string) CheckFunc {
	return checkMac
}

// checkMac checks if the given value is a valid an IEEE 802 MAC-48, EUI-48, EUI-64, or a 20-octet IP over InfiniBand link-layer address.
func checkMac(value, _ reflect.Value) Result {
	if value.Kind() != reflect.String {
		panic("string expected")
	}

	return IsMac(value.String())
}
