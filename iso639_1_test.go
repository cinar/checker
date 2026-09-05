// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2_test

import (
	"errors"
	"testing"

	v2 "github.com/cinar/checker/v2"
)

func TestIsISO6391Success(t *testing.T) {
	for _, value := range []string{"en", "tr", "fr", "de", "zh", "ja"} {
		result, err := v2.IsISO6391(value)
		if err != nil {
			t.Fatalf("%s: unexpected error %v", value, err)
		}

		if result != value {
			t.Fatalf("actual %s expected %s", result, value)
		}
	}
}

func TestIsISO6391Unknown(t *testing.T) {
	_, err := v2.IsISO6391("xx")

	if !errors.Is(err, v2.ErrNotISO6391) {
		t.Fatalf("got unexpected error %v", err)
	}
}

func TestIsISO6391CaseSensitive(t *testing.T) {
	_, err := v2.IsISO6391("EN")

	if !errors.Is(err, v2.ErrNotISO6391) {
		t.Fatalf("got unexpected error %v", err)
	}
}

func TestCheckStructISO6391(t *testing.T) {
	type Document struct {
		Language string `checkers:"lower iso639-1"`
	}

	document := &Document{
		Language: "xx",
	}

	errs, ok := v2.CheckStruct(document)
	if ok {
		t.Fatal("expected errors")
	}

	if !errors.Is(errs["Language"], v2.ErrNotISO6391) {
		t.Fatalf("expected language not iso639-1 %v", errs)
	}

	document.Language = "EN"

	errs, ok = v2.CheckStruct(document)
	if !ok {
		t.Fatalf("got unexpected errors %v", errs)
	}

	if document.Language != "en" {
		t.Fatalf("expected the lower normalizer to be applied, got %q", document.Language)
	}
}
