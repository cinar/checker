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

func ExampleIsMongoID() {
	_, err := v2.IsMongoID("507f1f77bcf86cd799439011")
	if err != nil {
		fmt.Println(err)
	}
}

func TestIsMongoIDInvalid(t *testing.T) {
	_, err := v2.IsMongoID("not-a-mongo-id")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestIsMongoIDWrongLengthInvalid(t *testing.T) {
	_, err := v2.IsMongoID("507f1f77bcf86cd79943901")
	if err == nil {
		t.Fatal("expected error for 23-character value")
	}
}

func TestIsMongoIDValid(t *testing.T) {
	_, err := v2.IsMongoID("507f1f77bcf86cd799439011")
	if err != nil {
		t.Fatal(err)
	}
}

func TestCheckMongoIDNonString(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Document struct {
		ID int `checkers:"mongo-id"`
	}

	document := &Document{}

	v2.CheckStruct(document)
}

func TestCheckMongoIDInvalid(t *testing.T) {
	type Document struct {
		ID string `checkers:"mongo-id"`
	}

	document := &Document{
		ID: "not-a-mongo-id",
	}

	_, ok := v2.CheckStruct(document)
	if ok {
		t.Fatal("expected error")
	}
}

func TestCheckMongoIDValid(t *testing.T) {
	type Document struct {
		ID string `checkers:"mongo-id"`
	}

	document := &Document{
		ID: "507f1f77bcf86cd799439011",
	}

	_, ok := v2.CheckStruct(document)
	if !ok {
		t.Fatal("expected valid")
	}
}
