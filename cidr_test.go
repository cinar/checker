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

func ExampleIsCIDR() {
	_, err := v2.IsCIDR("2001:db8::/32")
	if err != nil {
		fmt.Println(err)
	}
}

func TestIsCIDRInvalid(t *testing.T) {
	_, err := v2.IsCIDR("900.800.200.100//24")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestIsCIDRValid(t *testing.T) {
	_, err := v2.IsCIDR("2001:db8::/32")
	if err != nil {
		t.Fatal(err)
	}
}

func TestCheckCIDRNonString(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Network struct {
		Subnet int `checkers:"cidr"`
	}

	network := &Network{}

	v2.CheckStruct(network)
}

func TestCheckCIDRInvalid(t *testing.T) {
	type Network struct {
		Subnet string `checkers:"cidr"`
	}

	network := &Network{
		Subnet: "900.800.200.100//24",
	}

	_, ok := v2.CheckStruct(network)
	if ok {
		t.Fatal("expected error")
	}
}

func TestCheckCIDRValid(t *testing.T) {
	type Network struct {
		Subnet string `checkers:"cidr"`
	}

	network := &Network{
		Subnet: "192.0.2.0/24",
	}

	_, ok := v2.CheckStruct(network)
	if !ok {
		t.Fatal("expected valid")
	}
}
