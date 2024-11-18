// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2_test

import (
	"errors"
	"testing"

	v2 "github.com/cinar/checker/v2"
)

func TestCheckTrimeSpaceRequiredSuccess(t *testing.T) {
	input := "    test    "
	expected := "test"

	actual, err := v2.Check(input, v2.TrimSpace, v2.Required)
	if err != nil {
		t.Fatal(err)
	}

	if actual != expected {
		t.Fatalf("actual %s expected %s", actual, expected)
	}
}

func TestCheckTrimSpaceRequiredMissing(t *testing.T) {
	input := "    "
	expected := ""

	actual, err := v2.Check(input, v2.TrimSpace, v2.Required)
	if !errors.Is(err, v2.ErrRequired) {
		t.Fatalf("got unexpected error %v", err)
	}

	if actual != expected {
		t.Fatalf("actual %s expected %s", actual, expected)
	}
}
