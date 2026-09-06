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

func ExampleIsBase64URL() {
	_, err := v2.IsBase64URL("aGVsbG8_d29ybGQ=")
	if err != nil {
		fmt.Println(err)
	}
}

func TestIsBase64URLInvalid(t *testing.T) {
	_, err := v2.IsBase64URL("not base64url!!")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestIsBase64URLEmptyInvalid(t *testing.T) {
	_, err := v2.IsBase64URL("")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestIsBase64URLRejectsStandardAlphabet(t *testing.T) {
	_, err := v2.IsBase64URL("aGVsbG8/d29ybGQ=")
	if err == nil {
		t.Fatal("expected error for standard-alphabet base64 with '/'")
	}
}

func TestIsBase64URLValid(t *testing.T) {
	_, err := v2.IsBase64URL("aGVsbG8_d29ybGQ=")
	if err != nil {
		t.Fatal(err)
	}
}

func TestCheckBase64URLNonString(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Blob struct {
		Data int `checkers:"base64-url"`
	}

	blob := &Blob{}

	v2.CheckStruct(blob)
}

func TestCheckBase64URLInvalid(t *testing.T) {
	type Blob struct {
		Data string `checkers:"base64-url"`
	}

	blob := &Blob{
		Data: "not base64url!!",
	}

	_, ok := v2.CheckStruct(blob)
	if ok {
		t.Fatal("expected error")
	}
}

func TestCheckBase64URLValid(t *testing.T) {
	type Blob struct {
		Data string `checkers:"base64-url"`
	}

	blob := &Blob{
		Data: "aGVsbG8_d29ybGQ=",
	}

	_, ok := v2.CheckStruct(blob)
	if !ok {
		t.Fatal("expected valid")
	}
}
