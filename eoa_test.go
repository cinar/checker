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

func TestIsEOASuccess(t *testing.T) {
	value := "0x71C7656EC7ab88b098defB751B7401B5f6d8976F"

	result, err := v2.IsEOA(value)
	if err != nil {
		t.Fatal(err)
	}

	if result != value {
		t.Fatalf("actual %s expected %s", result, value)
	}
}

func TestIsEOAMissingPrefix(t *testing.T) {
	_, err := v2.IsEOA("71C7656EC7ab88b098defB751B7401B5f6d8976")

	if !errors.Is(err, v2.ErrNotEOA) {
		t.Fatalf("got unexpected error %v", err)
	}
}

func TestIsEOAWrongLength(t *testing.T) {
	_, err := v2.IsEOA("0x71C7656EC7ab88b098defB751B7401B5f6d897")

	if !errors.Is(err, v2.ErrNotEOA) {
		t.Fatalf("got unexpected error %v", err)
	}
}

func TestIsEOANotHex(t *testing.T) {
	_, err := v2.IsEOA("0xzzC7656EC7ab88b098defB751B7401B5f6d8976")

	if !errors.Is(err, v2.ErrNotEOA) {
		t.Fatalf("got unexpected error %v", err)
	}
}

func TestCheckStructEOA(t *testing.T) {
	type Wallet struct {
		Address string `checkers:"eoa"`
	}

	wallet := &Wallet{
		Address: "not-an-address",
	}

	errs, ok := v2.CheckStruct(wallet)
	if ok {
		t.Fatal("expected errors")
	}

	if !errors.Is(errs["Address"], v2.ErrNotEOA) {
		t.Fatalf("expected address not EOA %v", errs)
	}

	wallet.Address = "0x71C7656EC7ab88b098defB751B7401B5f6d8976F"

	errs, ok = v2.CheckStruct(wallet)
	if !ok {
		t.Fatalf("got unexpected errors %v", errs)
	}
}
