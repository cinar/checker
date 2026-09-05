// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2_test

import (
	"testing"

	v2 "github.com/cinar/checker/v2"
)

// definedEmail is a defined type whose underlying kind is string, used to
// verify that checkers work with custom string types, not just the
// built-in string type.
type definedEmail string

func TestCheckStructDefinedStringTypeChecker(t *testing.T) {
	type Account struct {
		Email definedEmail `checkers:"required email"`
	}

	account := &Account{
		Email: "test@example.com",
	}

	if _, ok := v2.CheckStruct(account); !ok {
		t.Fatal("expected valid")
	}
}

func TestCheckStructDefinedStringTypeNormalizer(t *testing.T) {
	type Account struct {
		Name definedEmail `checkers:"trim lower"`
	}

	account := &Account{
		Name: "  TEST@EXAMPLE.COM  ",
	}

	if _, ok := v2.CheckStruct(account); !ok {
		t.Fatal("expected valid")
	}

	if account.Name != "test@example.com" {
		t.Fatalf("expected normalized value, got %q", account.Name)
	}
}
