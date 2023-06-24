// Package checker is a Go library for validating user input through struct tags.
//
// Copyright (c) 2023 Onur Cinar. All Rights Reserved.
// The source code is provided under MIT License.
//
// https://github.com/cinar/checker
package checker

import "testing"

// FailIfNoPanic fails if test didn't panic. Use this function with the defer.
func FailIfNoPanic(t *testing.T) {
	if r := recover(); r == nil {
		t.Fail()
	}
}
