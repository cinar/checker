// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package checker_test

import (
	"testing"

	"github.com/cinar/checker"
)

func ExampleIsIP() {
	result := checker.IsIP("2001:db8::68")

	if result != checker.ResultValid {
		// Send the mistakes back to the user
	}
}

func TestIsIPInvalid(t *testing.T) {
	if checker.IsIP("900.800.200.100") == checker.ResultValid {
		t.Fail()
	}
}

func TestIsIPValid(t *testing.T) {
	if checker.IsIP("2001:db8::68") != checker.ResultValid {
		t.Fail()
	}
}

func TestCheckIpNonString(t *testing.T) {
	defer checker.FailIfNoPanic(t)

	type Request struct {
		RemoteIP int `checkers:"ip"`
	}

	request := &Request{}

	checker.Check(request)
}

func TestCheckIpInvalid(t *testing.T) {
	type Request struct {
		RemoteIP string `checkers:"ip"`
	}

	request := &Request{
		RemoteIP: "900.800.200.100",
	}

	_, valid := checker.Check(request)
	if valid {
		t.Fail()
	}
}

func TestCheckIPValid(t *testing.T) {
	type Request struct {
		RemoteIP string `checkers:"ip"`
	}

	request := &Request{
		RemoteIP: "192.168.1.1",
	}

	_, valid := checker.Check(request)
	if !valid {
		t.Fail()
	}
}
