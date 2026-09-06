// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

import (
	"reflect"
)

const (
	// nameJWT is the name of the JWT check.
	nameJWT = "jwt"

	// jwtExpression matches the structural shape of a JWT: three
	// base64url segments separated by dots. It does not decode the
	// segments or verify the signature.
	jwtExpression = `^[A-Za-z0-9_-]+\.[A-Za-z0-9_-]+\.[A-Za-z0-9_-]+$`
)

var (
	// ErrNotJWT indicates that the given string is not a structurally valid JWT.
	ErrNotJWT = NewCheckError("NOT_JWT")
)

// IsJWT checks if the given string has the structural shape of a JWT (three
// dot-separated base64url segments). It does not verify the signature.
func IsJWT(value string) (string, error) {
	if _, err := IsRegexp(jwtExpression, value); err != nil {
		return value, ErrNotJWT
	}

	return value, nil
}

// checkJWT checks if the given string has the structural shape of a JWT.
func checkJWT(value reflect.Value) (reflect.Value, error) {
	_, err := IsJWT(reflectString(value))
	return value, err
}

// makeJWT makes a checker function for the JWT checker.
func makeJWT(_ string) CheckFunc[reflect.Value] {
	return checkJWT
}
