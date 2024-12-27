// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

import (
	"reflect"
	"regexp"
)

// nameRegexp is the name of the regexp check.
const nameRegexp = "regexp"

// ErrNotMatch indicates that the given string does not match the regexp pattern.
var ErrNotMatch = NewCheckError("REGEXP")

// IsRegexp checks if the given string matches the given regexp expression.
func IsRegexp(expression, value string) (string, error) {
	if !regexp.MustCompile(expression).MatchString(value) {
		return value, ErrNotMatch
	}

	return value, nil
}

// MakeRegexpChecker makes a regexp checker for the given regexp expression with the given invalid result.
func MakeRegexpChecker(expression string, invalidError error) CheckFunc[reflect.Value] {
	return func(value reflect.Value) (reflect.Value, error) {
		if value.Kind() != reflect.String {
			panic("string expected")
		}

		_, err := IsRegexp(expression, value.String())
		if err != nil {
			return value, invalidError
		}

		return value, nil
	}
}

// makeRegexp makes a checker function for the regexp.
func makeRegexp(config string) CheckFunc[reflect.Value] {
	return MakeRegexpChecker(config, ErrNotMatch)
}
