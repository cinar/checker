// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package checker_test

import (
	"testing"

	"github.com/cinar/checker"
)

func ExampleIsIPV6() {
	err := checker.IsIPV6("2001:db8::68")

	if err != nil {
		// Send the mistakes back to the user
	}
}

func TestIsIPV6Invalid(t *testing.T) {
	if checker.IsIPV6("900.800.200.100") == nil {
		t.Fail()
	}
}

func TestIsIPV6InvalidV4(t *testing.T) {
	if checker.IsIPV6("192.168.1.1") == nil {
		t.Fail()
	}
}

func TestIsIPV6Valid(t *testing.T) {
	if checker.IsIPV6("2001:db8::68") != nil {
		t.Fail()
	}
}

func TestCheckIPV6NonString(t *testing.T) {
	defer checker.FailIfNoPanic(t)

	type Request struct {
		RemoteIP int `checkers:"ipv6"`
	}

	request := &Request{}

	checker.Check(request)
}

func TestCheckIPV6Invalid(t *testing.T) {
	type Request struct {
		RemoteIP string `checkers:"ipv6"`
	}

	request := &Request{
		RemoteIP: "900.800.200.100",
	}

	_, valid := checker.Check(request)
	if valid {
		t.Fail()
	}
}

func TestCheckIPV6Valid(t *testing.T) {
	type Request struct {
		RemoteIP string `checkers:"ipv6"`
	}

	request := &Request{
		RemoteIP: "2001:db8::68",
	}

	_, valid := checker.Check(request)
	if !valid {
		t.Fail()
	}
}
