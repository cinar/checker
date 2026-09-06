// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

import (
	"reflect"
	"regexp"
	"strings"
)

const (
	// nameISBN is the name of the ISBN check.
	nameISBN = "isbn"
)

var (
	// ErrNotISBN indicates that the given value is not a valid ISBN.
	ErrNotISBN = NewCheckError("NOT_ISBN")

	// isbnRegex is the regular expression for validating the shape of an
	// ISBN-10 or ISBN-13, after hyphen/space separators have been stripped.
	isbnRegex = regexp.MustCompile(`^(97(8|9))?\d{9}(\d|X)$`)

	// isbnSeparators strips the hyphens and spaces real-world ISBNs are
	// usually printed with, e.g. "978-0-306-40615-7".
	isbnSeparators = strings.NewReplacer("-", "", " ", "")
)

// IsISBN checks if the value is a valid ISBN-10 or ISBN-13, verifying the
// check digit rather than just the shape.
func IsISBN(value string) (string, error) {
	cleaned := isbnSeparators.Replace(value)

	if !isbnRegex.MatchString(cleaned) {
		return value, ErrNotISBN
	}

	valid := isValidISBN10(cleaned)
	if len(cleaned) != 10 {
		valid = isValidISBN13(cleaned)
	}

	if !valid {
		return value, ErrNotISBN
	}

	return value, nil
}

// isValidISBN10 verifies the ISBN-10 mod-11 check digit, treating a
// trailing "X" as the value 10.
func isValidISBN10(value string) bool {
	sum := 0

	for i := 0; i < 10; i++ {
		d := 10
		if value[i] != 'X' {
			d = int(value[i] - '0')
		}

		sum += d * (10 - i)
	}

	return sum%11 == 0
}

// isValidISBN13 verifies the ISBN-13 mod-10 check digit, alternating
// weights of 1 and 3.
func isValidISBN13(value string) bool {
	sum := 0

	for i := 0; i < 13; i++ {
		d := int(value[i] - '0')
		if i%2 != 0 {
			d *= 3
		}

		sum += d
	}

	return sum%10 == 0
}

// checkISBN checks if the value is a valid ISBN-10 or ISBN-13.
func checkISBN(value reflect.Value) (reflect.Value, error) {
	_, err := IsISBN(reflectString(value))
	return value, err
}

// makeISBN makes a checker function for the ISBN checker.
func makeISBN(_ string) CheckFunc[reflect.Value] {
	return checkISBN
}
