// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package checker_test

import (
	"testing"

	"github.com/cinar/checker"
)

func TestCheckErrorError(t *testing.T) {
	code := "CODE"

	err := checker.NewCheckError(code)

	if err.Error() != code {
		t.Fatalf("actaul %s expected %s", err.Error(), code)
	}
}
