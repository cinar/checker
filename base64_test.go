// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2_test

import (
	"fmt"
	"testing"

	v2 "github.com/cinar/checker/v2"
)

func ExampleIsBase64() {
	_, err := v2.IsBase64("aGVsbG8=")
	if err != nil {
		fmt.Println(err)
	}
}

func TestIsBase64Invalid(t *testing.T) {
	_, err := v2.IsBase64("not-base64!!")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestIsBase64EmptyInvalid(t *testing.T) {
	_, err := v2.IsBase64("")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestIsBase64Valid(t *testing.T) {
	_, err := v2.IsBase64("aGVsbG8=")
	if err != nil {
		t.Fatal(err)
	}
}

func TestCheckBase64NonString(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Blob struct {
		Data int `checkers:"base64"`
	}

	blob := &Blob{}

	v2.CheckStruct(blob)
}

func TestCheckBase64Invalid(t *testing.T) {
	type Blob struct {
		Data string `checkers:"base64"`
	}

	blob := &Blob{
		Data: "not-base64!!",
	}

	_, ok := v2.CheckStruct(blob)
	if ok {
		t.Fatal("expected error")
	}
}

func TestCheckBase64Valid(t *testing.T) {
	type Blob struct {
		Data string `checkers:"base64"`
	}

	blob := &Blob{
		Data: "aGVsbG8=",
	}

	_, ok := v2.CheckStruct(blob)
	if !ok {
		t.Fatal("expected valid")
	}
}
