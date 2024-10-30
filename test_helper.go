// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package checker

import "testing"

// FailIfNoPanic fails if test didn't panic. Use this function with the defer.
func FailIfNoPanic(t *testing.T) {
	if r := recover(); r == nil {
		t.Fail()
	}
}
