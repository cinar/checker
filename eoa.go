// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

import (
	"reflect"
	"regexp"
)

const (
	// nameEOA is the name of the EOA check.
	nameEOA = "eoa"
)

var (
	// ErrNotEOA indicates that the given value is not a valid externally owned address (EOA).
	ErrNotEOA = NewCheckError("NOT_EOA")

	// eoaRegex is the regular expression for validating an EOA, i.e. an Ethereum address.
	eoaRegex = regexp.MustCompile(`^0x[0-9a-fA-F]{40}$`)
)

// IsEOA checks if the value is a valid externally owned address (EOA), i.e. an
// Ethereum address: the "0x" prefix followed by 40 hexadecimal characters. It
// does not check the EIP-55 mixed-case checksum, only the address's shape.
func IsEOA(value string) (string, error) {
	if !eoaRegex.MatchString(value) {
		return value, ErrNotEOA
	}

	return value, nil
}

// checkEOA checks if the value is a valid externally owned address (EOA).
func checkEOA(value reflect.Value) (reflect.Value, error) {
	_, err := IsEOA(value.Interface().(string))
	return value, err
}

// makeEOA makes a checker function for the EOA checker.
func makeEOA(_ string) CheckFunc[reflect.Value] {
	return checkEOA
}
