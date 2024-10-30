// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package checker_test

import (
	"testing"

	"github.com/cinar/checker"
)

func ExampleIsFqdn() {
	err := checker.IsFqdn("zdo.com")
	if err != nil {
		// Send the errors back to the user
	}
}

func TestCheckFdqnWithoutTld(t *testing.T) {
	if checker.IsFqdn("abcd") == nil {
		t.Fail()
	}
}

func TestCheckFdqnShortTld(t *testing.T) {
	if checker.IsFqdn("abcd.c") == nil {
		t.Fail()
	}
}

func TestCheckFdqnNumericTld(t *testing.T) {
	if checker.IsFqdn("abcd.1234") == nil {
		t.Fail()
	}
}

func TestCheckFdqnLong(t *testing.T) {
	if checker.IsFqdn("abcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcd.com") == nil {
		t.Fail()
	}
}

func TestCheckFdqnInvalidCharacters(t *testing.T) {
	if checker.IsFqdn("ab_cd.com") == nil {
		t.Fail()
	}
}

func TestCheckFdqnStaringWithHyphen(t *testing.T) {
	if checker.IsFqdn("-abcd.com") == nil {
		t.Fail()
	}
}

func TestCheckFdqnStaringEndingWithHyphen(t *testing.T) {
	if checker.IsFqdn("abcd-.com") == nil {
		t.Fail()
	}
}

func TestCheckFdqnStartingWithDot(t *testing.T) {
	if checker.IsFqdn(".abcd.com") == nil {
		t.Fail()
	}
}

func TestCheckFdqnEndingWithDot(t *testing.T) {
	if checker.IsFqdn("abcd.com.") == nil {
		t.Fail()
	}
}

func TestCheckFqdnNonString(t *testing.T) {
	defer checker.FailIfNoPanic(t)

	type Request struct {
		Domain int `checkers:"fqdn"`
	}

	request := &Request{}

	checker.Check(request)
}

func TestCheckFqdnValid(t *testing.T) {
	type Request struct {
		Domain string `checkers:"fqdn"`
	}

	request := &Request{
		Domain: "zdo.com",
	}

	_, valid := checker.Check(request)
	if !valid {
		t.Fail()
	}
}
