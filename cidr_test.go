// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package checker_test

import (
	"fmt"
	"testing"

	"github.com/cinar/checker"
)

func ExampleIsCIDR() {
	_, err := checker.IsCIDR("2001:db8::/32")
	if err != nil {
		fmt.Println(err)
	}
}

func TestIsCIDRInvalid(t *testing.T) {
	_, err := checker.IsCIDR("900.800.200.100//24")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestIsCIDRValid(t *testing.T) {
	_, err := checker.IsCIDR("2001:db8::/32")
	if err != nil {
		t.Fatal(err)
	}
}

func TestCheckCidrNonString(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Network struct {
		Subnet int `checkers:"cidr"`
	}

	network := &Network{}

	checker.CheckStruct(network)
}

func TestCheckCIDRInvalid(t *testing.T) {
	type Network struct {
		Subnet string `checkers:"cidr"`
	}

	network := &Network{
		Subnet: "900.800.200.100//24",
	}

	_, valid := checker.CheckStruct(network)
	if valid {
		t.Fail()
	}
}

func TestCheckCIDRValid(t *testing.T) {
	type Network struct {
		Subnet string `checkers:"cidr"`
	}

	network := &Network{
		Subnet: "192.0.2.0/24",
	}

	_, valid := checker.CheckStruct(network)
	if !valid {
		t.Fail()
	}
}
