// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2_test

import (
	"errors"
	"fmt"
	"testing"

	v2 "github.com/cinar/checker/v2"
)

func ExampleIsContains() {
	_, err := v2.IsContains("@", "onur@example.com")
	if err != nil {
		fmt.Println(err)
	}
}

func TestIsContainsValid(t *testing.T) {
	_, err := v2.IsContains("@", "onur@example.com")
	if err != nil {
		t.Fatal(err)
	}
}

func TestIsContainsInvalid(t *testing.T) {
	_, err := v2.IsContains("@", "onur.example.com")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestCheckContainsNonString(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type User struct {
		Email int `checkers:"contains:@"`
	}

	user := &User{}

	v2.CheckStruct(user)
}

func TestCheckContainsInvalid(t *testing.T) {
	type User struct {
		Email string `checkers:"contains:@"`
	}

	user := &User{
		Email: "onur.example.com",
	}

	errs, ok := v2.CheckStruct(user)
	if ok {
		t.Fatal("expected error")
	}

	if !errors.Is(errs["Email"], v2.ErrNotContains) {
		t.Fatalf("expected ErrNotContains, got %v", errs)
	}
}

func TestCheckContainsValid(t *testing.T) {
	type User struct {
		Email string `checkers:"contains:@"`
	}

	user := &User{
		Email: "onur@example.com",
	}

	errs, ok := v2.CheckStruct(user)
	if !ok {
		t.Fatal(errs)
	}
}
