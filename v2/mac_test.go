// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2_test

import (
	"fmt"
	"testing"

	v2 "github.com/cinar/checker/v2"
)

func ExampleIsMAC() {
	_, err := v2.IsMAC("00:1A:2B:3C:4D:5E")
	if err != nil {
		fmt.Println(err)
	}
}

func TestIsMACInvalid(t *testing.T) {
	_, err := v2.IsMAC("invalid-mac")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestIsMACValid(t *testing.T) {
	_, err := v2.IsMAC("00:1A:2B:3C:4D:5E")
	if err != nil {
		t.Fatal(err)
	}
}

func TestCheckMACNonString(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Device struct {
		MAC int `checkers:"mac"`
	}

	device := &Device{}

	v2.CheckStruct(device)
}

func TestCheckMACInvalid(t *testing.T) {
	type Device struct {
		MAC string `checkers:"mac"`
	}

	device := &Device{
		MAC: "invalid-mac",
	}

	_, ok := v2.CheckStruct(device)
	if ok {
		t.Fatal("expected error")
	}
}

func TestCheckMACValid(t *testing.T) {
	type Device struct {
		MAC string `checkers:"mac"`
	}

	device := &Device{
		MAC: "00:1A:2B:3C:4D:5E",
	}

	_, ok := v2.CheckStruct(device)
	if !ok {
		t.Fatal("expected valid")
	}
}
