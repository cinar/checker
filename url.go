// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package checker

import (
	"net/url"
	"reflect"
)

// CheckerURL is the name of the checker.
const CheckerURL = "url"

// ResultNotURL indicates that the given value is not a valid URL.
const ResultNotURL = "NOT_URL"

// IsURL checks if the given value is a valid URL.
func IsURL(value string) Result {
	url, err := url.ParseRequestURI(value)
	if err != nil {
		return ResultNotURL
	}

	if url.Scheme == "" {
		return ResultNotURL
	}

	if url.Host == "" {
		return ResultNotURL
	}

	return ResultValid
}

// makeURL makes a checker function for the URL checker.
func makeURL(_ string) CheckFunc {
	return checkURL
}

// checkURL checks if the given value is a valid URL.
func checkURL(value, _ reflect.Value) Result {
	if value.Kind() != reflect.String {
		panic("string expected")
	}

	return IsURL(value.String())
}
