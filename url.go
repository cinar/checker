// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package checker

import (
	"errors"
	"net/url"
	"reflect"
)

// CheckerURL is the name of the checker.
const CheckerURL = "url"

// ErrNotURL indicates that the given value is not a valid URL.
var ErrNotURL = errors.New("please enter a valid URL")

// IsURL checks if the given value is a valid URL.
func IsURL(value string) error {
	url, err := url.ParseRequestURI(value)
	if err != nil {
		return ErrNotURL
	}

	if url.Scheme == "" {
		return ErrNotURL
	}

	if url.Host == "" {
		return ErrNotURL
	}

	return nil
}

// makeURL makes a checker function for the URL checker.
func makeURL(_ string) CheckFunc {
	return checkURL
}

// checkURL checks if the given value is a valid URL.
func checkURL(value, _ reflect.Value) error {
	if value.Kind() != reflect.String {
		panic("string expected")
	}

	return IsURL(value.String())
}
