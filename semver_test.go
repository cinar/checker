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

func ExampleIsSemver() {
	_, err := v2.IsSemver("1.2.3-alpha.1+build.123")
	if err != nil {
		fmt.Println(err)
	}
}

func TestIsSemverInvalid(t *testing.T) {
	_, err := v2.IsSemver("1.0")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestIsSemverLeadingVInvalid(t *testing.T) {
	_, err := v2.IsSemver("v1.0.0")
	if err == nil {
		t.Fatal("expected error for leading 'v'")
	}
}

func TestIsSemverValid(t *testing.T) {
	_, err := v2.IsSemver("1.2.3-alpha.1+build.123")
	if err != nil {
		t.Fatal(err)
	}
}

func TestCheckSemverNonString(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Release struct {
		Version int `checkers:"semver"`
	}

	release := &Release{}

	v2.CheckStruct(release)
}

func TestCheckSemverInvalid(t *testing.T) {
	type Release struct {
		Version string `checkers:"semver"`
	}

	release := &Release{
		Version: "1.0",
	}

	_, ok := v2.CheckStruct(release)
	if ok {
		t.Fatal("expected error")
	}
}

func TestCheckSemverValid(t *testing.T) {
	type Release struct {
		Version string `checkers:"semver"`
	}

	release := &Release{
		Version: "1.2.3",
	}

	_, ok := v2.CheckStruct(release)
	if !ok {
		t.Fatal("expected valid")
	}
}
