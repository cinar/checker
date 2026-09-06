// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

import (
	"reflect"
)

const (
	// nameSemver is the name of the semver check.
	nameSemver = "semver"

	// semverExpression is the official semver.org regular expression for a
	// valid Semantic Versioning 2.0.0 version string.
	semverExpression = `^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)` +
		`(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?` +
		`(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`
)

var (
	// ErrNotSemver indicates that the given string is not a valid semantic version.
	ErrNotSemver = NewCheckError("NOT_SEMVER")
)

// IsSemver checks if the given string is a valid Semantic Versioning 2.0.0 version, e.g. "1.2.3-alpha+001".
func IsSemver(value string) (string, error) {
	if _, err := IsRegexp(semverExpression, value); err != nil {
		return value, ErrNotSemver
	}

	return value, nil
}

// checkSemver checks if the given string is a valid semantic version.
func checkSemver(value reflect.Value) (reflect.Value, error) {
	_, err := IsSemver(reflectString(value))
	return value, err
}

// makeSemver makes a checker function for the semver checker.
func makeSemver(_ string) CheckFunc[reflect.Value] {
	return checkSemver
}
