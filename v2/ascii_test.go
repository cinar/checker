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

func ExampleIsASCII() {
	_, err := v2.IsASCII("Checker")
	if err != nil {
		fmt.Println(err)
	}
}

func TestIsASCIIInvalid(t *testing.T) {
	_, err := v2.IsASCII("𝄞 Music!")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestIsASCIIValid(t *testing.T) {
	_, err := v2.IsASCII("Checker")
	if err != nil {
		t.Fatal(err)
	}
}

func TestCheckASCIINonString(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type User struct {
		Username int `checkers:"ascii"`
	}

	user := &User{}

	v2.CheckStruct(user)
}

func TestCheckASCIIInvalid(t *testing.T) {
	type User struct {
		Username string `checkers:"ascii"`
	}

	user := &User{
		Username: "𝄞 Music!",
	}

	_, ok := v2.CheckStruct(user)
	if ok {
		t.Fatal("expected error")
	}
}

func TestCheckASCIIValid(t *testing.T) {
	type User struct {
		Username string `checkers:"ascii"`
	}

	user := &User{
		Username: "checker",
	}

	_, ok := v2.CheckStruct(user)
	if !ok {
		t.Fatal("expected valid")
	}
}
