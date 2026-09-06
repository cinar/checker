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

const testJWT = "eyJhbGciOiJIUzI1NiJ9.eyJzdWIiOiIxMjM0In0.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"

func ExampleIsJWT() {
	_, err := v2.IsJWT(testJWT)
	if err != nil {
		fmt.Println(err)
	}
}

func TestIsJWTInvalid(t *testing.T) {
	_, err := v2.IsJWT("not.a-jwt")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestIsJWTValid(t *testing.T) {
	_, err := v2.IsJWT(testJWT)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCheckJWTNonString(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Session struct {
		Token int `checkers:"jwt"`
	}

	session := &Session{}

	v2.CheckStruct(session)
}

func TestCheckJWTInvalid(t *testing.T) {
	type Session struct {
		Token string `checkers:"jwt"`
	}

	session := &Session{
		Token: "not.a-jwt",
	}

	_, ok := v2.CheckStruct(session)
	if ok {
		t.Fatal("expected error")
	}
}

func TestCheckJWTValid(t *testing.T) {
	type Session struct {
		Token string `checkers:"jwt"`
	}

	session := &Session{
		Token: testJWT,
	}

	_, ok := v2.CheckStruct(session)
	if !ok {
		t.Fatal("expected valid")
	}
}
