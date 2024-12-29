// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2_test

import (
	"testing"

	v2 "github.com/cinar/checker/v2"
)

func TestURLEscape(t *testing.T) {
	input := "param1/param2 = 1 + 2 & 3 + 4"
	expected := "param1%2Fparam2+%3D+1+%2B+2+%26+3+%2B+4"

	actual, err := v2.URLEscape(input)
	if err != nil {
		t.Fatal(err)
	}

	if actual != expected {
		t.Fatalf("actual %s expected %s", actual, expected)
	}
}

func TestReflectURLEscape(t *testing.T) {
	type Request struct {
		Query string `checkers:"url-escape"`
	}

	request := &Request{
		Query: "param1/param2 = 1 + 2 & 3 + 4",
	}

	expected := "param1%2Fparam2+%3D+1+%2B+2+%26+3+%2B+4"

	errs, ok := v2.CheckStruct(request)
	if !ok {
		t.Fatalf("got unexpected errors %v", errs)
	}

	if request.Query != expected {
		t.Fatalf("actual %s expected %s", request.Query, expected)
	}
}
