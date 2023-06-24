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

func TestIsRequired(t *testing.T) {
	s := "valid"

	if checker.IsRequired(s) != checker.ResultValid {
		t.Fail()
	}
}

func TestIsRequiredUninitializedString(t *testing.T) {
	var s string

	if checker.IsRequired(s) != checker.ResultRequired {
		t.Fail()
	}
}

func TestIsRequiredEmptyString(t *testing.T) {
	s := ""

	if checker.IsRequired(s) == checker.ResultValid {
		t.Fail()
	}
}

func TestIsRequiredUninitializedNumber(t *testing.T) {
	var n int

	if checker.IsRequired(n) != checker.ResultRequired {
		t.Fail()
	}
}

func TestIsRequiredValidSlice(t *testing.T) {
	s := []int{1}

	if checker.IsRequired(s) != checker.ResultValid {
		t.Fail()
	}
}

func TestIsRequiredUninitializedSlice(t *testing.T) {
	var s []int

	if checker.IsRequired(s) != checker.ResultRequired {
		t.Fail()
	}
}

func TestIsRequiredEmptySlice(t *testing.T) {
	s := make([]int, 0)

	if checker.IsRequired(s) != checker.ResultRequired {
		t.Fail()
	}
}

func TestIsRequiredValidArray(t *testing.T) {
	s := [1]int{1}

	if checker.IsRequired(s) != checker.ResultValid {
		t.Fail()
	}
}

func TestIsRequiredEmptyArray(t *testing.T) {
	s := [1]int{}

	if checker.IsRequired(s) != checker.ResultRequired {
		t.Fail()
	}
}

func TestIsRequiredValidMap(t *testing.T) {
	m := map[string]string{
		"a": "b",
	}

	if checker.IsRequired(m) != checker.ResultValid {
		t.Fail()
	}
}

func TestIsRequiredUninitializedMap(t *testing.T) {
	var m map[string]string

	if checker.IsRequired(m) != checker.ResultRequired {
		t.Fail()
	}
}

func TestCheckRequiredEmptyMap(t *testing.T) {
	m := map[string]string{}

	if checker.IsRequired(m) != checker.ResultRequired {
		t.Fail()
	}
}

func TestCheckRequiredValid(t *testing.T) {
	type User struct {
		Username string `checkers:"required"`
	}

	user := &User{
		Username: "checker",
	}

	_, valid := checker.Check(user)
	if !valid {
		t.Fail()
	}
}

func TestCheckRequiredInvalid(t *testing.T) {
	type User struct {
		Username string `checkers:"required"`
	}

	user := &User{}

	_, valid := checker.Check(user)
	if valid {
		t.Fail()
	}
}
