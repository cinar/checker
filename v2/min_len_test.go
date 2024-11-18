// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2_test

import (
	"testing"

	v2 "github.com/cinar/checker/v2"
)

func TestMinLenSuccess(t *testing.T) {
	value := "test"

	check := v2.MinLen(4)

	result, err := check(value)
	if result != value {
		t.Fatalf("result (%s) is not the original value (%s)", result, value)
	}

	if err != nil {
		t.Fatal(err)
	}
}

func TestMinLenError(t *testing.T) {
	value := "test"

	check := v2.MinLen(5)

	result, err := check(value)
	if result != value {
		t.Fatalf("result (%s) is not the original value (%s)", result, value)
	}

	_, ok := err.(*v2.MinLenError)
	if !ok {
		t.Fatalf("got unexpected error %v", err)
	}
}
