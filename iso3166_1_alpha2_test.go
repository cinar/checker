// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2_test

import (
	"errors"
	"testing"

	v2 "github.com/cinar/checker/v2"
)

func TestIsISO31661Alpha2Success(t *testing.T) {
	for _, value := range []string{"US", "TR", "FR", "DE", "CN", "JP"} {
		result, err := v2.IsISO31661Alpha2(value)
		if err != nil {
			t.Fatalf("%s: unexpected error %v", value, err)
		}

		if result != value {
			t.Fatalf("actual %s expected %s", result, value)
		}
	}
}

func TestIsISO31661Alpha2Unknown(t *testing.T) {
	_, err := v2.IsISO31661Alpha2("XX")

	if !errors.Is(err, v2.ErrNotISO31661Alpha2) {
		t.Fatalf("got unexpected error %v", err)
	}
}

func TestIsISO31661Alpha2CaseSensitive(t *testing.T) {
	_, err := v2.IsISO31661Alpha2("us")

	if !errors.Is(err, v2.ErrNotISO31661Alpha2) {
		t.Fatalf("got unexpected error %v", err)
	}
}

func TestCheckStructISO31661Alpha2(t *testing.T) {
	type Address struct {
		Country string `checkers:"upper iso3166-1-alpha-2"`
	}

	address := &Address{
		Country: "XX",
	}

	errs, ok := v2.CheckStruct(address)
	if ok {
		t.Fatal("expected errors")
	}

	if !errors.Is(errs["Country"], v2.ErrNotISO31661Alpha2) {
		t.Fatalf("expected country not iso3166-1-alpha-2 %v", errs)
	}

	address.Country = "us"

	errs, ok = v2.CheckStruct(address)
	if !ok {
		t.Fatalf("got unexpected errors %v", errs)
	}

	if address.Country != "US" {
		t.Fatalf("expected the upper normalizer to be applied, got %q", address.Country)
	}
}
