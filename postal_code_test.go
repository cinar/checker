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

func TestIsPostalCodeSuccess(t *testing.T) {
	tests := []struct {
		country string
		value   string
	}{
		{"US", "94103"},
		{"US", "94103-1234"},
		{"us", "94103"},
		{"CA", "K1A 0B1"},
		{"GB", "SW1A 1AA"},
		{"DE", "10115"},
		{"JP", "100-0001"},
	}

	for _, test := range tests {
		value, err := v2.IsPostalCode(test.country, test.value)
		if err != nil {
			t.Fatalf("%s: unexpected error %v", test.country, err)
		}

		if value != test.value {
			t.Fatalf("%s: actual %s expected %s", test.country, value, test.value)
		}
	}
}

func TestIsPostalCodeFailure(t *testing.T) {
	_, err := v2.IsPostalCode("US", "not-a-zip")

	if !errors.Is(err, v2.ErrNotPostalCode) {
		t.Fatalf("got unexpected error %v", err)
	}
}

func TestIsPostalCodeUnsupportedCountry(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	v2.IsPostalCode("XX", "12345")
}

func TestCheckStructPostalCode(t *testing.T) {
	type Address struct {
		Zip string `checkers:"postal-code:US"`
	}

	address := &Address{
		Zip: "not-a-zip",
	}

	errs, ok := v2.CheckStruct(address)
	if ok {
		t.Fatal("expected errors")
	}

	if !errors.Is(errs["Zip"], v2.ErrNotPostalCode) {
		t.Fatalf("expected zip not postal code %v", errs)
	}

	address.Zip = "94103"

	errs, ok = v2.CheckStruct(address)
	if !ok {
		t.Fatalf("got unexpected errors %v", errs)
	}
}
