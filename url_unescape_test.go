// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package checker_test

import (
	"testing"

	"github.com/cinar/checker"
)

func TestNormalizeURLUnescapeNonString(t *testing.T) {
	defer checker.FailIfNoPanic(t)

	type Request struct {
		Query int `checkers:"url-unescape"`
	}

	request := &Request{}

	checker.Check(request)
}

func TestNormalizeURLUnescape(t *testing.T) {
	type Request struct {
		Query string `checkers:"url-unescape"`
	}

	request := &Request{
		Query: "param1%2Fparam2+%3D+1+%2B+2+%26+3+%2B+4",
	}

	_, valid := checker.Check(request)
	if !valid {
		t.Fail()
	}

	if request.Query != "param1/param2 = 1 + 2 & 3 + 4" {
		t.Fail()
	}
}
