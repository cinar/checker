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
	// nameIP is the name of the IP check.
	nameIP = "ip"
)

var (
	// ErrNotIP indicates that the given value is not a valid IP address.
	ErrNotIP = NewCheckError("IP")
)

// IsIP checks if the value is a valid IP address.
func IsIP(value string) (string, error) {
	if net.ParseIP(value) == nil {
		return value, ErrNotIP
	}
	return value, nil
}

// checkIP checks if the value is a valid IP address.
func checkIP(value reflect.Value) (reflect.Value, error) {
	_, err := IsIP(value.Interface().(string))
	return value, err
}

// makeIP makes a checker function for the IP checker.
func makeIP(_ string) CheckFunc[reflect.Value] {
	return checkIP
}
