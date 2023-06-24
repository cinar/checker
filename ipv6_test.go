// Package checker is a Go library for validating user input through struct tags.
//
// https://github.com/cinar/checker
//
// Copyright 2023 Onur Cinar. All rights reserved. 
// Use of this source code is governed by a MIT-style 
// license that can be found in the LICENSE file. 
//
package checker_test

import (
	"testing"

	"github.com/cinar/checker"
)

func TestIsIPV6Invalid(t *testing.T) {
	if checker.IsIPV6("900.800.200.100") == checker.ResultValid {
		t.Fail()
	}
}

func TestIsIPV6InvalidV4(t *testing.T) {
	if checker.IsIPV6("192.168.1.1") == checker.ResultValid {
		t.Fail()
	}
}

func TestIsIPV6Valid(t *testing.T) {
	if checker.IsIPV6("2001:db8::68") != checker.ResultValid {
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
