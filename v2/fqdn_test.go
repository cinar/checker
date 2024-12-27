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

func ExampleIsFQDN() {
	_, err := v2.IsFQDN("example.com")
	if err != nil {
		fmt.Println(err)
	}
}

func TestIsFQDNInvalid(t *testing.T) {
	_, err := v2.IsFQDN("invalid_fqdn")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestIsFQDNValid(t *testing.T) {
	_, err := v2.IsFQDN("example.com")
	if err != nil {
		t.Fatal(err)
	}
}

func TestCheckFQDNNonString(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Domain struct {
		Name int `checkers:"fqdn"`
	}

	domain := &Domain{}

	v2.CheckStruct(domain)
}

func TestCheckFQDNInvalid(t *testing.T) {
	type Domain struct {
		Name string `checkers:"fqdn"`
	}

	domain := &Domain{
		Name: "invalid_fqdn",
	}

	_, ok := v2.CheckStruct(domain)
	if ok {
		t.Fatal("expected error")
	}
}

func TestCheckFQDNValid(t *testing.T) {
	type Domain struct {
		Name string `checkers:"fqdn"`
	}

	domain := &Domain{
		Name: "example.com",
	}

	_, ok := v2.CheckStruct(domain)
	if !ok {
		t.Fatal("expected valid")
	}
}
