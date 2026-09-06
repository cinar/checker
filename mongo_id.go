// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

import (
	"reflect"
)

const (
	// nameMongoID is the name of the mongo-id check.
	nameMongoID = "mongo-id"

	// mongoIDExpression matches a 24-character hex-encoded MongoDB ObjectID.
	mongoIDExpression = `^[0-9a-fA-F]{24}$`
)

var (
	// ErrNotMongoID indicates that the given string is not a valid MongoDB ObjectID.
	ErrNotMongoID = NewCheckError("NOT_MONGO_ID")
)

// IsMongoID checks if the given string is a valid MongoDB ObjectID, e.g. "507f1f77bcf86cd799439011".
func IsMongoID(value string) (string, error) {
	if _, err := IsRegexp(mongoIDExpression, value); err != nil {
		return value, ErrNotMongoID
	}

	return value, nil
}

// checkMongoID checks if the given string is a valid MongoDB ObjectID.
func checkMongoID(value reflect.Value) (reflect.Value, error) {
	_, err := IsMongoID(reflectString(value))
	return value, err
}

// makeMongoID makes a checker function for the mongo-id checker.
func makeMongoID(_ string) CheckFunc[reflect.Value] {
	return checkMongoID
}
