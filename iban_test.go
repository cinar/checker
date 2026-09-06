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

func ExampleIsIBAN() {
	_, err := v2.IsIBAN("DE89370400440532013000")
	if err != nil {
		fmt.Println(err)
	}
}

func TestIsIBANInvalidShape(t *testing.T) {
	_, err := v2.IsIBAN("not-an-iban")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestIsIBANBadChecksumInvalid(t *testing.T) {
	_, err := v2.IsIBAN("GB29NWBK60161331926818")
	if err == nil {
		t.Fatal("expected error for bad check digits")
	}
}

func TestIsIBANValid(t *testing.T) {
	_, err := v2.IsIBAN("DE89370400440532013000")
	if err != nil {
		t.Fatal(err)
	}
}

func TestIsIBANSpacedValid(t *testing.T) {
	_, err := v2.IsIBAN("GB29 NWBK 6016 1331 9268 19")
	if err != nil {
		t.Fatal(err)
	}
}

func TestCheckIBANNonString(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Account struct {
		IBAN int `checkers:"iban"`
	}

	account := &Account{}

	v2.CheckStruct(account)
}

func TestCheckIBANInvalid(t *testing.T) {
	type Account struct {
		IBAN string `checkers:"iban"`
	}

	account := &Account{
		IBAN: "not-an-iban",
	}

	_, ok := v2.CheckStruct(account)
	if ok {
		t.Fatal("expected error")
	}
}

func TestCheckIBANValid(t *testing.T) {
	type Account struct {
		IBAN string `checkers:"iban"`
	}

	account := &Account{
		IBAN: "DE89370400440532013000",
	}

	_, ok := v2.CheckStruct(account)
	if !ok {
		t.Fatal("expected valid")
	}
}
