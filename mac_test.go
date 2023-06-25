// Package checker is a Go library for validating user input through struct tags.
//
// https://github.com/cinar/checker
//
// Copyright 2023 Onur Cinar. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package checker_test

import (
	"testing"

	"github.com/cinar/checker"
)

func ExampleIsMac() {
	result := checker.IsMac("00:00:5e:00:53:01")

	if result != checker.ResultValid {
		// Send the mistakes back to the user
	}
}

func TestIsMacInvalid(t *testing.T) {
	if checker.IsMac("1234") == checker.ResultValid {
		t.Fail()
	}
}

func TestIsMacValid(t *testing.T) {
	if checker.IsMac("00:00:5e:00:53:01") != checker.ResultValid {
		t.Fail()
	}
}

func TestCheckMacNonString(t *testing.T) {
	defer checker.FailIfNoPanic(t)

	type Network struct {
		HardwareAddress int `checkers:"mac"`
	}

	network := &Network{}

	checker.Check(network)
}

func TestCheckMacInvalid(t *testing.T) {
	type Network struct {
		HardwareAddress string `checkers:"mac"`
	}

	network := &Network{
		HardwareAddress: "1234",
	}

	_, valid := checker.Check(network)
	if valid {
		t.Fail()
	}
}

func TestCheckMacValid(t *testing.T) {
	type Network struct {
		HardwareAddress string `checkers:"mac"`
	}

	network := &Network{
		HardwareAddress: "00:00:5e:00:53:01",
	}

	_, valid := checker.Check(network)
	if !valid {
		t.Fail()
	}
}
