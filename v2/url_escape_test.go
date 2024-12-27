// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2_test

import (
	"testing"

	v2 "github.com/cinar/checker/v2"
)

func TestNormalizeURLEscapeNonString(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Request struct {
		Query int `checkers:"url-escape"`
	}

	request := &Request{}

	v2.CheckStruct(request)
}

func TestNormalizeURLEscape(t *testing.T) {
	type Request struct {
		Query string `checkers:"url-escape"`
	}

	request := &Request{
		Query: "param1/param2 = 1 + 2 & 3 + 4",
	}

	_, valid := v2.CheckStruct(request)
	if !valid {
		t.Fail()
	}

	if request.Query != "param1%2Fparam2+%3D+1+%2B+2+%26+3+%2B+4" {
		t.Fail()
	}
}
