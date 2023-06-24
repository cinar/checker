// Package checker is a Go library for validating user input through struct tags.
//
// https://github.com/cinar/checker
//
// Copyright 2023 Onur Cinar. All rights reserved. 
// Use of this source code is governed by a MIT-style 
// license that can be found in the LICENSE file. 
//
package checker

import (
	"reflect"
	"regexp"
)

// CheckerRegexp is the name of the checker.
const CheckerRegexp = "regexp"

// ResultNotMatch indicates that the given string does not match the regexp pattern.
const ResultNotMatch = "NOT_MATCH"

// MakeRegexpMaker makes a regexp checker maker for the given regexp expression with the given invalid result.
func MakeRegexpMaker(expression string, invalidResult Result) MakeFunc {
	return func(_ string) CheckFunc {
		return MakeRegexpChecker(expression, invalidResult)
	}
}

// MakeRegexpChecker makes a regexp checker for the given regexp expression with the given invalid result.
func MakeRegexpChecker(expression string, invalidResult Result) CheckFunc {
	pattern := regexp.MustCompile(expression)

	return func(value, parent reflect.Value) Result {
		return checkRegexp(value, pattern, invalidResult)
	}
}

// makeRegexp makes a checker function for the regexp.
func makeRegexp(config string) CheckFunc {
	return MakeRegexpChecker(config, ResultNotMatch)
}

// checkRegexp checks if the given string matches the regexp pattern.
func checkRegexp(value reflect.Value, pattern *regexp.Regexp, invalidResult Result) Result {
	if value.Kind() != reflect.String {
		panic("string expected")
	}

	if !pattern.MatchString(value.String()) {
		return invalidResult
	}

	return ResultValid
}
