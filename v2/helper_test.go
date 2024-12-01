// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2_test

import "testing"

// FailIfNoPanic fails the test if there were no panic.
func FailIfNoPanic(t *testing.T, message string) {
	t.Helper()

	if r := recover(); r == nil {
		t.Fatal(message)
	}
}
