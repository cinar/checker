// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package checker_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/cinar/checker"
)

func ExampleCheck() {
	name := "    Onur Cinar    "

	name, err := checker.Check(name, checker.TrimSpace, checker.Required)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(name)
	// Output: Onur Cinar
}

func ExampleCheckStruct() {
	type Person struct {
		Name string `checkers:"trim required"`
	}

	person := &Person{
		Name: "    Onur Cinar    ",
	}

	errs, ok := checker.CheckStruct(person)
	if !ok {
		fmt.Println(errs)
		return
	}

	fmt.Println(person.Name)
	// Output: Onur Cinar
}

func TestCheckTrimSpaceRequiredSuccess(t *testing.T) {
	input := "    test    "
	expected := "test"

	actual, err := checker.Check(input, checker.TrimSpace, checker.Required)
	if err != nil {
		t.Fatal(err)
	}

	if actual != expected {
		t.Fatalf("actual %s expected %s", actual, expected)
	}
}

func TestCheckTrimSpaceRequiredMissing(t *testing.T) {
	input := "    "
	expected := ""

	actual, err := checker.Check(input, checker.TrimSpace, checker.Required)
	if !errors.Is(err, checker.ErrRequired) {
		t.Fatalf("got unexpected error %v", err)
	}

	if actual != expected {
		t.Fatalf("actual %s expected %s", actual, expected)
	}
}

func TestCheckWithConfigSuccess(t *testing.T) {
	input := "    test    "
	expected := "test"

	actual, err := checker.CheckWithConfig(input, "trim required")
	if err != nil {
		t.Fatal(err)
	}

	if actual != expected {
		t.Fatalf("actual %s expected %s", actual, expected)
	}
}

func TestCheckWithConfigRequiredMissing(t *testing.T) {
	input := "    "
	expected := ""

	actual, err := checker.CheckWithConfig(input, "trim required")
	if !errors.Is(err, checker.ErrRequired) {
		t.Fatalf("got unexpected error %v", err)
	}

	if actual != expected {
		t.Fatalf("actual %s expected %s", actual, expected)
	}
}

func TestCheckStructSuccess(t *testing.T) {
	type Address struct {
		Street string `checkers:"required"`
	}

	type Person struct {
		Name    string `checkers:"required"`
		Address *Address
	}

	person := &Person{
		Name: "Onur Cinar",
		Address: &Address{
			Street: "1234 Main",
		},
	}

	errs, ok := checker.CheckStruct(person)
	if !ok {
		t.Fatalf("got unexpected errs %v", errs)
	}
}

func TestCheckStructRequiredMissing(t *testing.T) {
	type Address struct {
		Street string `checkers:"required"`
	}

	type Person struct {
		Name    string `checkers:"required"`
		Address *Address
	}

	person := &Person{
		Name: "",
		Address: &Address{
			Street: "",
		},
	}

	errs, ok := checker.CheckStruct(person)
	if ok {
		t.Fatal("expected errors")
	}

	if !errors.Is(errs["Name"], checker.ErrRequired) {
		t.Fatalf("expected name required %v", errs)
	}

	if !errors.Is(errs["Address.Street"], checker.ErrRequired) {
		t.Fatalf("expected streed required %v", errs)
	}
}

func TestCheckStructCustomName(t *testing.T) {
	type Person struct {
		Name string `json:"name" checkers:"required"`
	}

	person := &Person{
		Name: "",
	}

	errs, ok := checker.CheckStruct(person)
	if ok {
		t.Fatal("expected errors")
	}

	if !errors.Is(errs["name"], checker.ErrRequired) {
		t.Fatalf("expected name required %v", errs)
	}
}

func TestCheckStructSlice(t *testing.T) {
	type Person struct {
		Name   string   `checkers:"required"`
		Emails []string `checkers:"@max-len:1 max-len:4"`
	}

	person := &Person{
		Name: "Onur Cinar",
		Emails: []string{
			"onur.cinar",
		},
	}

	errs, ok := checker.CheckStruct(person)
	if ok {
		t.Fatal("expected errors")
	}

	if !errors.Is(errs["Emails[0]"], checker.ErrMaxLen) {
		t.Fatalf("expected email max len")
	}
}
