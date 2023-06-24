// Package checker is a Go library for validating user input through struct tags.
//
// https://github.com/cinar/checker
//
// Copyright 2023 Onur Cinar. All rights reserved. 
// Use of this source code is governed by a MIT-style 
// license that can be found in the LICENSE file. 
//
package checker

import "testing"

// FailIfNoPanic fails if test didn't panic. Use this function with the defer.
func FailIfNoPanic(t *testing.T) {
	if r := recover(); r == nil {
		t.Fail()
	}
}
