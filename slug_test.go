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

func ExampleIsSlug() {
	_, err := v2.IsSlug("hello-world")
	if err != nil {
		fmt.Println(err)
	}
}

func TestIsSlugInvalid(t *testing.T) {
	_, err := v2.IsSlug("Hello World")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestIsSlugConsecutiveHyphensInvalid(t *testing.T) {
	_, err := v2.IsSlug("hello--world")
	if err == nil {
		t.Fatal("expected error for consecutive hyphens")
	}
}

func TestIsSlugLeadingHyphenInvalid(t *testing.T) {
	_, err := v2.IsSlug("-hello")
	if err == nil {
		t.Fatal("expected error for leading hyphen")
	}
}

func TestIsSlugValid(t *testing.T) {
	_, err := v2.IsSlug("hello-world")
	if err != nil {
		t.Fatal(err)
	}
}

func TestCheckSlugNonString(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Post struct {
		Slug int `checkers:"slug"`
	}

	post := &Post{}

	v2.CheckStruct(post)
}

func TestCheckSlugInvalid(t *testing.T) {
	type Post struct {
		Slug string `checkers:"slug"`
	}

	post := &Post{
		Slug: "Hello World",
	}

	_, ok := v2.CheckStruct(post)
	if ok {
		t.Fatal("expected error")
	}
}

func TestCheckSlugValid(t *testing.T) {
	type Post struct {
		Slug string `checkers:"slug"`
	}

	post := &Post{
		Slug: "hello-world",
	}

	_, ok := v2.CheckStruct(post)
	if !ok {
		t.Fatal("expected valid")
	}
}
