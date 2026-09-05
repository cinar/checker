// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

import (
	"reflect"
	"regexp"
	"sync"
)

// nameRegexp is the name of the regexp check.
const nameRegexp = "regexp"

// ErrNotMatch indicates that the given string does not match the regexp pattern.
var ErrNotMatch = NewCheckError("REGEXP")

// regexpCache holds compiled *regexp.Regexp values keyed by their source
// expression, so a given expression is only ever compiled once. Checker
// tags are static and developer-authored, so the set of distinct
// expressions a process ever compiles is bounded by the code, not by
// request input -- unlike, say, caching on a raw user-supplied string.
var regexpCache sync.Map

// compileRegexp returns the compiled *regexp.Regexp for expression, compiling
// and caching it on first use. Panics if expression doesn't compile, same as
// regexp.MustCompile.
func compileRegexp(expression string) *regexp.Regexp {
	if compiled, ok := regexpCache.Load(expression); ok {
		return compiled.(*regexp.Regexp)
	}

	compiled, _ := regexpCache.LoadOrStore(expression, regexp.MustCompile(expression))

	return compiled.(*regexp.Regexp)
}

// IsRegexp checks if the given string matches the given regexp expression.
func IsRegexp(expression, value string) (string, error) {
	if !compileRegexp(expression).MatchString(value) {
		return value, ErrNotMatch
	}

	return value, nil
}

// MakeRegexpChecker makes a regexp checker for the given regexp expression with the given invalid result.
func MakeRegexpChecker(expression string, invalidError error) CheckFunc[reflect.Value] {
	return func(value reflect.Value) (reflect.Value, error) {
		_, err := IsRegexp(expression, reflectString(value))
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
