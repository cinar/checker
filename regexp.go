// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package checker

import (
	"errors"
	"reflect"
	"regexp"
)

// tagRegexp is the tag of the checker.
const tagRegexp = "regexp"

// ErrNotMatch indicates that the given string does not match the regexp pattern.
var ErrNotMatch = errors.New("please enter a valid input")

// MakeRegexpMaker makes a regexp checker maker for the given regexp expression with the given invalid result.
func MakeRegexpMaker(expression string, invalidError error) MakeFunc {
	return func(_ string) CheckFunc {
		return MakeRegexpChecker(expression, invalidError)
	}
}

// MakeRegexpChecker makes a regexp checker for the given regexp expression with the given invalid result.
func MakeRegexpChecker(expression string, invalidError error) CheckFunc {
	pattern := regexp.MustCompile(expression)

	return func(value, parent reflect.Value) error {
		return checkRegexp(value, pattern, invalidError)
	}
}

// makeRegexp makes a checker function for the regexp.
func makeRegexp(config string) CheckFunc {
	return MakeRegexpChecker(config, ErrNotMatch)
}

// checkRegexp checks if the given string matches the regexp pattern.
func checkRegexp(value reflect.Value, pattern *regexp.Regexp, invalidError error) error {
	if value.Kind() != reflect.String {
		panic("string expected")
	}

	if !pattern.MatchString(value.String()) {
		return invalidError
	}

	return nil
}
