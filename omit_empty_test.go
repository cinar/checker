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

func TestCheckStructOmitEmptySkipsZeroValue(t *testing.T) {
	type Profile struct {
		Website string `checkers:"omitempty url"`
	}

	profile := &Profile{}

	if _, ok := v2.CheckStruct(profile); !ok {
		t.Fatal("expected valid: empty value should be skipped by omitempty")
	}
}

func TestCheckStructOmitEmptyStillChecksNonZeroValue(t *testing.T) {
	type Profile struct {
		Website string `checkers:"omitempty url"`
	}

	profile := &Profile{
		Website: "not a url",
	}

	errs, ok := v2.CheckStruct(profile)
	if ok {
		t.Fatal("expected errors: non-empty value must still be validated")
	}

	if !errors.Is(errs["Website"], v2.ErrNotURL) {
		t.Fatalf("expected ErrNotURL, got %v", errs)
	}
}

func TestCheckStructOmitEmptyValidNonZeroValue(t *testing.T) {
	type Profile struct {
		Website string `checkers:"omitempty url"`
	}

	profile := &Profile{
		Website: "https://example.com",
	}

	if _, ok := v2.CheckStruct(profile); !ok {
		t.Fatal("expected valid")
	}
}

func TestCheckStructOmitEmptyRequiredIsANoOpOnZeroValue(t *testing.T) {
	// omitempty always looks at the field's own zero-ness, so pairing it
	// with required on a zero value skips required too -- documented,
	// intentional behavior, not a bug: "optional, but must be valid if
	// present" and "required" are contradictory requirements to begin with.
	type Profile struct {
		Name string `checkers:"omitempty required"`
	}

	profile := &Profile{}

	if _, ok := v2.CheckStruct(profile); !ok {
		t.Fatal("expected valid: omitempty skips required on a zero value")
	}
}

func TestCheckStructOmitEmptyLooksAtOriginalValueNotNormalized(t *testing.T) {
	// "   " is non-zero as a string, so omitempty does not trigger; trim
	// then reduces it to "", and required correctly fails on that.
	type Profile struct {
		Name string `checkers:"trim omitempty required"`
	}

	profile := &Profile{
		Name: "   ",
	}

	errs, ok := v2.CheckStruct(profile)
	if ok {
		t.Fatal("expected errors")
	}

	if !errors.Is(errs["Name"], v2.ErrRequired) {
		t.Fatalf("expected ErrRequired, got %v", errs)
	}
}

func TestCheckStructOmitEmptySliceContainer(t *testing.T) {
	type Person struct {
		Emails []string `checkers:"@omitempty @min-len:1"`
	}

	person := &Person{}

	if _, ok := v2.CheckStruct(person); !ok {
		t.Fatal("expected valid: nil slice should be skipped by @omitempty")
	}
}

func TestCheckStructOmitEmptySliceContainerNonEmptyStillChecked(t *testing.T) {
	type Person struct {
		Emails []string `checkers:"@omitempty @max-len:1"`
	}

	person := &Person{
		Emails: []string{"a", "b"},
	}

	errs, ok := v2.CheckStruct(person)
	if ok {
		t.Fatal("expected errors")
	}

	if !errors.Is(errs["Emails"], v2.ErrMaxLen) {
		t.Fatalf("expected ErrMaxLen, got %v", errs)
	}
}

func TestCheckWithConfigOmitEmpty(t *testing.T) {
	actual, err := v2.CheckWithConfig("", "omitempty required")
	if err != nil {
		t.Fatal(err)
	}

	if actual != "" {
		t.Fatalf("expected empty string, got %q", actual)
	}
}
