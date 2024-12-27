// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2_test

import (
	"reflect"
	"testing"

	v2 "github.com/cinar/checker/v2"
)

func TestCheckRegexpNonString(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type User struct {
		Username int `checkers:"regexp:^[A-Za-z]$"`
	}

	user := &User{}

	v2.CheckStruct(user)
}

func TestCheckRegexpInvalid(t *testing.T) {
	type User struct {
		Username string `checkers:"regexp:^[A-Za-z]+$"`
	}

	user := &User{
		Username: "abcd1234",
	}

	_, ok := v2.CheckStruct(user)
	if ok {
		t.Fatal("expected error")
	}
}

func TestCheckRegexpValid(t *testing.T) {
	type User struct {
		Username string `checkers:"regexp:^[A-Za-z]+$"`
	}

	user := &User{
		Username: "abcd",
	}

	_, ok := v2.CheckStruct(user)
	if !ok {
		t.Fatal("expected valid")
	}
}

func TestMakeRegexpChecker(t *testing.T) {
	checkHex := v2.MakeRegexpChecker("^[A-Fa-f0-9]+$", v2.ErrNotMatch)

	_, err := checkHex(reflect.ValueOf("f0f0f0"))
	if err != nil {
		t.Fail()
	}
}
