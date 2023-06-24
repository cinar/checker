// Package checker is a Go library for validating user input through struct tags.
//
// Copyright (c) 2023 Onur Cinar. All Rights Reserved.
// The source code is provided under MIT License.
//
// https://github.com/cinar/checker
package checker_test

import (
	"testing"

	"github.com/cinar/checker"
)

func TestFailIfNoPanicValid(t *testing.T) {
	defer checker.FailIfNoPanic(t)
	panic("")
}

func TestFailIfNoPanicInvalid(t *testing.T) {
	defer checker.FailIfNoPanic(t)
	defer checker.FailIfNoPanic(nil)
}
