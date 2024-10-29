// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package checker_test

import (
	"testing"

	"github.com/cinar/checker"
)

func ExampleIsRequired() {
	var name string

	err := checker.IsRequired(name)
	if err != nil {
		// Send the err back to the user
	}
}

func TestIsRequired(t *testing.T) {
	s := "valid"

	if checker.IsRequired(s) != nil {
		t.Fail()
	}
}

func TestIsRequiredUninitializedString(t *testing.T) {
	var s string

	if checker.IsRequired(s) != checker.ErrRequired {
		t.Fail()
	}
}

func TestIsRequiredEmptyString(t *testing.T) {
	s := ""

	if checker.IsRequired(s) == nil {
		t.Fail()
	}
}

func TestIsRequiredUninitializedNumber(t *testing.T) {
	var n int

	if checker.IsRequired(n) != checker.ErrRequired {
		t.Fail()
	}
}

func TestIsRequiredValidSlice(t *testing.T) {
	s := []int{1}

	if checker.IsRequired(s) != nil {
		t.Fail()
	}
}

func TestIsRequiredUninitializedSlice(t *testing.T) {
	var s []int

	if checker.IsRequired(s) != checker.ErrRequired {
		t.Fail()
	}
}

func TestIsRequiredEmptySlice(t *testing.T) {
	s := make([]int, 0)

	if checker.IsRequired(s) != checker.ErrRequired {
		t.Fail()
	}
}

func TestIsRequiredValidArray(t *testing.T) {
	s := [1]int{1}

	if checker.IsRequired(s) != nil {
		t.Fail()
	}
}

func TestIsRequiredEmptyArray(t *testing.T) {
	s := [1]int{}

	if checker.IsRequired(s) != checker.ErrRequired {
		t.Fail()
	}
}

func TestIsRequiredValidMap(t *testing.T) {
	m := map[string]string{
		"a": "b",
	}

	if checker.IsRequired(m) != nil {
		t.Fail()
	}
}

func TestIsRequiredUninitializedMap(t *testing.T) {
	var m map[string]string

	if checker.IsRequired(m) != checker.ErrRequired {
		t.Fail()
	}
}

func TestCheckRequiredEmptyMap(t *testing.T) {
	m := map[string]string{}

	if checker.IsRequired(m) != checker.ErrRequired {
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
