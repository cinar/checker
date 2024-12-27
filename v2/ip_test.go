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

func ExampleIsIP() {
	_, err := v2.IsIP("192.168.1.1")
	if err != nil {
		fmt.Println(err)
	}
}

func TestIsIPInvalid(t *testing.T) {
	_, err := v2.IsIP("invalid-ip")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestIsIPValid(t *testing.T) {
	_, err := v2.IsIP("192.168.1.1")
	if err != nil {
		t.Fatal(err)
	}
}

func TestCheckIPNonString(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Network struct {
		Address int `checkers:"ip"`
	}

	network := &Network{}

	v2.CheckStruct(network)
}

func TestCheckIPInvalid(t *testing.T) {
	type Network struct {
		Address string `checkers:"ip"`
	}

	network := &Network{
		Address: "invalid-ip",
	}

	_, ok := v2.CheckStruct(network)
	if ok {
		t.Fatal("expected error")
	}
}

func TestCheckIPValid(t *testing.T) {
	type Network struct {
		Address string `checkers:"ip"`
	}

	network := &Network{
		Address: "192.168.1.1",
	}

	_, ok := v2.CheckStruct(network)
	if !ok {
		t.Fatal("expected valid")
	}
}
