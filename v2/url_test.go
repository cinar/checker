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

func ExampleIsURL() {
	_, err := v2.IsURL("https://example.com")
	if err != nil {
		fmt.Println(err)
	}
}

func TestIsURLInvalid(t *testing.T) {
	_, err := v2.IsURL("invalid-url")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestIsURLValid(t *testing.T) {
	_, err := v2.IsURL("https://example.com")
	if err != nil {
		t.Fatal(err)
	}
}

func TestCheckURLNonString(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Website struct {
		Link int `checkers:"url"`
	}

	website := &Website{}

	v2.CheckStruct(website)
}

func TestCheckURLInvalid(t *testing.T) {
	type Website struct {
		Link string `checkers:"url"`
	}

	website := &Website{
		Link: "invalid-url",
	}

	_, ok := v2.CheckStruct(website)
	if ok {
		t.Fatal("expected error")
	}
}

func TestCheckURLValid(t *testing.T) {
	type Website struct {
		Link string `checkers:"url"`
	}

	website := &Website{
		Link: "https://example.com",
	}

	_, ok := v2.CheckStruct(website)
	if !ok {
		t.Fatal("expected valid")
	}
}
