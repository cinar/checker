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
	// nameMAC is the name of the MAC check.
	nameMAC = "mac"
)

var (
	// ErrNotMAC indicates that the given value is not a valid MAC address.
	ErrNotMAC = NewCheckError("NOT_MAC")
)

// IsMAC checks if the value is a valid MAC address.
func IsMAC(value string) (string, error) {
	_, err := net.ParseMAC(value)
	if err != nil {
		return value, ErrNotMAC
	}
	return value, nil
}

// checkMAC checks if the value is a valid MAC address.
func checkMAC(value reflect.Value) (reflect.Value, error) {
	_, err := IsMAC(value.Interface().(string))
	return value, err
}

// makeMAC makes a checker function for the MAC checker.
func makeMAC(_ string) CheckFunc[reflect.Value] {
	return checkMAC
}
