// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

import (
	"reflect"
	"unicode"
)

const (
	// nameLUHN is the name of the LUHN check.
	nameLUHN = "luhn"
)

var (
	// ErrNotLUHN indicates that the given value is not a valid LUHN number.
	ErrNotLUHN = NewCheckError("NOT_LUHN")
)

// IsLUHN checks if the value is a valid LUHN number.
func IsLUHN(value string) (string, error) {
	var sum int
	var alt bool

	for i := len(value) - 1; i >= 0; i-- {
		r := rune(value[i])
		if !unicode.IsDigit(r) {
			return value, ErrNotLUHN
		}

		n := int(r - '0')
		if alt {
			n *= 2
			if n > 9 {
				n -= 9
			}
		}
		sum += n
		alt = !alt
	}

	if sum%10 != 0 {
		return value, ErrNotLUHN
	}

	return value, nil
}

// checkLUHN checks if the value is a valid LUHN number.
func checkLUHN(value reflect.Value) (reflect.Value, error) {
	_, err := IsLUHN(value.Interface().(string))
	return value, err
}

// makeLUHN makes a checker function for the LUHN checker.
func makeLUHN(_ string) CheckFunc[reflect.Value] {
	return checkLUHN
}
