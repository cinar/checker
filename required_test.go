// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package checker_test

import (
	"errors"
	"testing"

	"github.com/cinar/checker"
)

func TestRequiredSuccess(t *testing.T) {
	value := "test"

	result, err := checker.Required(value)

	if result != value {
		t.Fatalf("result (%s) is not the original value (%s)", result, value)
	}

	if err != nil {
		t.Fatal(err)
	}
}

func TestRequiredMissing(t *testing.T) {
	var value string

	result, err := checker.Required(value)

	if result != value {
		t.Fatalf("result (%s) is not the original value (%s)", result, value)
	}

	if !errors.Is(err, checker.ErrRequired) {
		t.Fatalf("got unexpected error %v", err)
	}
}
