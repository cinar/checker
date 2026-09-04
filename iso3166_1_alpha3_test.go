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

func TestIsISO31661Alpha3Success(t *testing.T) {
	for _, value := range []string{"USA", "TUR", "FRA", "DEU", "CHN", "JPN"} {
		result, err := v2.IsISO31661Alpha3(value)
		if err != nil {
			t.Fatalf("%s: unexpected error %v", value, err)
		}

		if result != value {
			t.Fatalf("actual %s expected %s", result, value)
		}
	}
}

func TestIsISO31661Alpha3Unknown(t *testing.T) {
	_, err := v2.IsISO31661Alpha3("XXX")

	if !errors.Is(err, v2.ErrNotISO31661Alpha3) {
		t.Fatalf("got unexpected error %v", err)
	}
}

func TestIsISO31661Alpha3CaseSensitive(t *testing.T) {
	_, err := v2.IsISO31661Alpha3("usa")

	if !errors.Is(err, v2.ErrNotISO31661Alpha3) {
		t.Fatalf("got unexpected error %v", err)
	}
}

func TestCheckStructISO31661Alpha3(t *testing.T) {
	type Address struct {
		Country string `checkers:"upper iso3166-1-alpha-3"`
	}

	address := &Address{
		Country: "XXX",
	}

	errs, ok := v2.CheckStruct(address)
	if ok {
		t.Fatal("expected errors")
	}

	if !errors.Is(errs["Country"], v2.ErrNotISO31661Alpha3) {
		t.Fatalf("expected country not iso3166-1-alpha-3 %v", errs)
	}

	address.Country = "usa"

	errs, ok = v2.CheckStruct(address)
	if !ok {
		t.Fatalf("got unexpected errors %v", errs)
	}

	if address.Country != "USA" {
		t.Fatalf("expected the upper normalizer to be applied, got %q", address.Country)
	}
}
