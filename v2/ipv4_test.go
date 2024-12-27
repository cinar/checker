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

func ExampleIsIPv4() {
	_, err := v2.IsIPv4("192.168.1.1")
	if err != nil {
		fmt.Println(err)
	}
}

func TestIsIPv4Invalid(t *testing.T) {
	_, err := v2.IsIPv4("2001:db8::1")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestIsIPv4Valid(t *testing.T) {
	_, err := v2.IsIPv4("192.168.1.1")
	if err != nil {
		t.Fatal(err)
	}
}

func TestCheckIPv4NonString(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Network struct {
		Address int `checkers:"ipv4"`
	}

	network := &Network{}

	v2.CheckStruct(network)
}

func TestCheckIPv4Invalid(t *testing.T) {
	type Network struct {
		Address string `checkers:"ipv4"`
	}

	network := &Network{
		Address: "2001:db8::1",
	}

	_, ok := v2.CheckStruct(network)
	if ok {
		t.Fatal("expected error")
	}
}

func TestCheckIPv4Valid(t *testing.T) {
	type Network struct {
		Address string `checkers:"ipv4"`
	}

	network := &Network{
		Address: "192.168.1.1",
	}

	_, ok := v2.CheckStruct(network)
	if !ok {
		t.Fatal("expected valid")
	}
}
