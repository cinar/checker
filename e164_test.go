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

func ExampleIsE164() {
	_, err := v2.IsE164("+14155552671")
	if err != nil {
		fmt.Println(err)
	}
}

func TestIsE164Invalid(t *testing.T) {
	_, err := v2.IsE164("14155552671")
	if err == nil {
		t.Fatal("expected error for missing leading +")
	}
}

func TestIsE164LeadingZeroInvalid(t *testing.T) {
	_, err := v2.IsE164("+0123456789")
	if err == nil {
		t.Fatal("expected error for leading zero after +")
	}
}

func TestIsE164Valid(t *testing.T) {
	_, err := v2.IsE164("+14155552671")
	if err != nil {
		t.Fatal(err)
	}
}

func TestCheckE164NonString(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Contact struct {
		Phone int `checkers:"e164"`
	}

	contact := &Contact{}

	v2.CheckStruct(contact)
}

func TestCheckE164Invalid(t *testing.T) {
	type Contact struct {
		Phone string `checkers:"e164"`
	}

	contact := &Contact{
		Phone: "14155552671",
	}

	_, ok := v2.CheckStruct(contact)
	if ok {
		t.Fatal("expected error")
	}
}

func TestCheckE164Valid(t *testing.T) {
	type Contact struct {
		Phone string `checkers:"e164"`
	}

	contact := &Contact{
		Phone: "+14155552671",
	}

	_, ok := v2.CheckStruct(contact)
	if !ok {
		t.Fatal("expected valid")
	}
}
