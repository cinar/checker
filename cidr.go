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
	// nameCIDR is the name of the CIDR check.
	nameCIDR = "cidr"
)

var (
	// ErrNotCIDR indicates that the given value is not a valid CIDR.
	ErrNotCIDR = NewCheckError("NOT_CIDR")
)

// IsCIDR checks if the value is a valid CIDR notation IP address and prefix length.
func IsCIDR(value string) (string, error) {
	_, _, err := net.ParseCIDR(value)
	if err != nil {
		return value, ErrNotCIDR
	}

	return value, nil
}

// checkCIDR checks if the value is a valid CIDR notation IP address and prefix length.
func checkCIDR(value reflect.Value) (reflect.Value, error) {
	_, err := IsCIDR(value.Interface().(string))
	return value, err
}

// makeCIDR makes a checker function for the CIDR checker.
func makeCIDR(_ string) CheckFunc[reflect.Value] {
	return checkCIDR
}
