// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
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
