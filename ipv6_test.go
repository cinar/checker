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

func ExampleIsIPv6() {
	_, err := v2.IsIPv6("2001:db8::1")
	if err != nil {
		fmt.Println(err)
	}
}

func TestIsIPv6Invalid(t *testing.T) {
	_, err := v2.IsIPv6("192.168.1.1")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestIsIPv6Valid(t *testing.T) {
	_, err := v2.IsIPv6("2001:db8::1")
	if err != nil {
		t.Fatal(err)
	}
}

func TestCheckIPv6NonString(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Network struct {
		Address int `checkers:"ipv6"`
	}

	network := &Network{}

	v2.CheckStruct(network)
}

func TestCheckIPv6Invalid(t *testing.T) {
	type Network struct {
		Address string `checkers:"ipv6"`
	}

	network := &Network{
		Address: "192.168.1.1",
	}

	_, ok := v2.CheckStruct(network)
	if ok {
		t.Fatal("expected error")
	}
}

func TestCheckIPv6Valid(t *testing.T) {
	type Network struct {
		Address string `checkers:"ipv6"`
	}

	network := &Network{
		Address: "2001:db8::1",
	}

	_, ok := v2.CheckStruct(network)
	if !ok {
		t.Fatal("expected valid")
	}
}
