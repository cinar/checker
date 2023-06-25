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

func ExampleIsIPV4() {
	result := checker.IsIPV4("192.168.1.1")

	if result != checker.ResultValid {
		// Send the mistakes back to the user
	}
}

func TestIsIPV4Invalid(t *testing.T) {
	if checker.IsIPV4("900.800.200.100") == checker.ResultValid {
		t.Fail()
	}
}

func TestIsIPV4InvalidV6(t *testing.T) {
	if checker.IsIPV4("2001:db8::68") == checker.ResultValid {
		t.Fail()
	}
}

func TestIsIPV4Valid(t *testing.T) {
	if checker.IsIPV4("192.168.1.1") != checker.ResultValid {
		t.Fail()
	}
}

func TestCheckIPV4NonString(t *testing.T) {
	defer checker.FailIfNoPanic(t)

	type Request struct {
		RemoteIP int `checkers:"ipv4"`
	}

	request := &Request{}

	checker.Check(request)
}

func TestCheckIPV4Invalid(t *testing.T) {
	type Request struct {
		RemoteIP string `checkers:"ipv4"`
	}

	request := &Request{
		RemoteIP: "900.800.200.100",
	}

	_, valid := checker.Check(request)
	if valid {
		t.Fail()
	}
}

func TestCheckIPV4Valid(t *testing.T) {
	type Request struct {
		RemoteIP string `checkers:"ipv4"`
	}

	request := &Request{
		RemoteIP: "192.168.1.1",
	}

	_, valid := checker.Check(request)
	if !valid {
		t.Fail()
	}
}
