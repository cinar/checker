// Package checker is a Go library for validating user input through struct tags.
//
// https://github.com/cinar/checker
//
// Copyright 2023 Onur Cinar. All rights reserved. 
// Use of this source code is governed by a MIT-style 
// license that can be found in the LICENSE file. 
//
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
