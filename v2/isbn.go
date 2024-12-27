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
	// nameISBN is the name of the ISBN check.
	nameISBN = "isbn"
)

var (
	// ErrNotISBN indicates that the given value is not a valid ISBN.
	ErrNotISBN = NewCheckError("ISBN")

	// isbnRegex is the regular expression for validating ISBN-10 and ISBN-13.
	isbnRegex = regexp.MustCompile(`^(97(8|9))?\d{9}(\d|X)$`)
)

// IsISBN checks if the value is a valid ISBN-10 or ISBN-13.
func IsISBN(value string) (string, error) {
	if !isbnRegex.MatchString(value) {
		return value, ErrNotISBN
	}
	return value, nil
}

// checkISBN checks if the value is a valid ISBN-10 or ISBN-13.
func checkISBN(value reflect.Value) (reflect.Value, error) {
	_, err := IsISBN(value.Interface().(string))
	return value, err
}

// makeISBN makes a checker function for the ISBN checker.
func makeISBN(_ string) CheckFunc[reflect.Value] {
	return checkISBN
}
