// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package checker_test

import (
	"testing"

	"github.com/cinar/checker"
)

func ExampleIsCidr() {
	err := checker.IsCidr("2001:db8::/32")
	if err != nil {
		// Send the mistakes back to the user
	}
}

func TestIsCidrInvalid(t *testing.T) {
	if checker.IsCidr("900.800.200.100//24") == nil {
		t.Fail()
	}
}

func TestIsCidrValid(t *testing.T) {
	if checker.IsCidr("2001:db8::/32") != nil {
		t.Fail()
	}
}

func TestCheckCidrNonString(t *testing.T) {
	defer checker.FailIfNoPanic(t)

	type Network struct {
		Subnet int `checkers:"cidr"`
	}

	network := &Network{}

	checker.Check(network)
}

func TestCheckCidrInvalid(t *testing.T) {
	type Network struct {
		Subnet string `checkers:"cidr"`
	}

	network := &Network{
		Subnet: "900.800.200.100//24",
	}

	_, valid := checker.Check(network)
	if valid {
		t.Fail()
	}
}

func TestCheckCidrValid(t *testing.T) {
	type Network struct {
		Subnet string `checkers:"cidr"`
	}

	network := &Network{
		Subnet: "192.0.2.0/24",
	}

	_, valid := checker.Check(network)
	if !valid {
		t.Fail()
	}
}
