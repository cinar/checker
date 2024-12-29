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

func ExampleIsEmail() {
	_, err := v2.IsEmail("test@example.com")
	if err != nil {
		fmt.Println(err)
	}
}

func TestIsEmailInvalid(t *testing.T) {
	_, err := v2.IsEmail("invalid-email")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestIsEmailValid(t *testing.T) {
	_, err := v2.IsEmail("test@example.com")
	if err != nil {
		t.Fatal(err)
	}
}

func TestCheckEmailNonString(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type User struct {
		Email int `checkers:"email"`
	}

	user := &User{}

	v2.CheckStruct(user)
}

func TestCheckEmailInvalid(t *testing.T) {
	type User struct {
		Email string `checkers:"email"`
	}

	user := &User{
		Email: "invalid-email",
	}

	_, ok := v2.CheckStruct(user)
	if ok {
		t.Fatal("expected error")
	}
}

func TestCheckEmailValid(t *testing.T) {
	type User struct {
		Email string `checkers:"email"`
	}

	user := &User{
		Email: "test@example.com",
	}

	_, ok := v2.CheckStruct(user)
	if !ok {
		t.Fatal("expected valid")
	}
}
