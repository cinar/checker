// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

import (
	"math/big"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

const (
	// nameIBAN is the name of the IBAN check.
	nameIBAN = "iban"
)

var (
	// ErrNotIBAN indicates that the given string is not a valid IBAN.
	ErrNotIBAN = NewCheckError("NOT_IBAN")

	// ibanShapeExpression matches the general IBAN shape: a two-letter
	// country code, two check digits, and up to 30 alphanumeric BBAN
	// characters. It does not enforce the per-country BBAN length.
	ibanShapeExpression = regexp.MustCompile(`^[A-Z]{2}\d{2}[A-Z0-9]{11,30}$`)

	// ibanCheckModulus is the modulus used by the ISO 7064 mod 97-10
	// check digit scheme.
	ibanCheckModulus = big.NewInt(97)
)

// IsIBAN checks if the given string is a valid IBAN, verifying the ISO
// 7064 mod 97-10 check digits rather than just the shape. Spaces are
// ignored, so "GB29 NWBK 6016 1331 9268 19" and "GB29NWBK60161331926819"
// are equivalent.
func IsIBAN(value string) (string, error) {
	cleaned := strings.ToUpper(strings.ReplaceAll(value, " ", ""))

	if !ibanShapeExpression.MatchString(cleaned) {
		return value, ErrNotIBAN
	}

	rearranged := cleaned[4:] + cleaned[:4]

	var numeric strings.Builder
	for _, r := range rearranged {
		if r >= 'A' && r <= 'Z' {
			numeric.WriteString(strconv.Itoa(int(r-'A') + 10))
		} else {
			numeric.WriteRune(r)
		}
	}

	total, ok := new(big.Int).SetString(numeric.String(), 10)
	if !ok || new(big.Int).Mod(total, ibanCheckModulus).Int64() != 1 {
		return value, ErrNotIBAN
	}

	return value, nil
}

// checkIBAN checks if the given string is a valid IBAN.
func checkIBAN(value reflect.Value) (reflect.Value, error) {
	_, err := IsIBAN(reflectString(value))
	return value, err
}

// makeIBAN makes a checker function for the IBAN checker.
func makeIBAN(_ string) CheckFunc[reflect.Value] {
	return checkIBAN
}
