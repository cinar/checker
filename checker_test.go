// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2_test

import (
	"errors"
	"fmt"
	"testing"

	v2 "github.com/cinar/checker/v2"
)

func ExampleCheck() {
	name := "    Onur Cinar    "

	name, err := v2.Check(name, v2.TrimSpace, v2.Required)
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

	errs, ok := v2.CheckStruct(person)
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

	actual, err := v2.Check(input, v2.TrimSpace, v2.Required)
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

	actual, err := v2.Check(input, v2.TrimSpace, v2.Required)
	if !errors.Is(err, v2.ErrRequired) {
		t.Fatalf("got unexpected error %v", err)
	}

	if actual != expected {
		t.Fatalf("actual %s expected %s", actual, expected)
	}
}

func TestCheckWithConfigSuccess(t *testing.T) {
	input := "    test    "
	expected := "test"

	actual, err := v2.CheckWithConfig(input, "trim required")
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

	actual, err := v2.CheckWithConfig(input, "trim required")
	if !errors.Is(err, v2.ErrRequired) {
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

	errors, ok := v2.CheckStruct(person)
	if !ok {
		t.Fatalf("got unexpected errors %v", errors)
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

	errs, ok := v2.CheckStruct(person)
	if ok {
		t.Fatal("expected errors")
	}

	if !errors.Is(errs["Name"], v2.ErrRequired) {
		t.Fatalf("expected name required %v", errs)
	}

	if !errors.Is(errs["Address.Street"], v2.ErrRequired) {
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

	errs, ok := v2.CheckStruct(person)
	if ok {
		t.Fatal("expected errors")
	}

	if !errors.Is(errs["name"], v2.ErrRequired) {
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

	errs, ok := v2.CheckStruct(person)
	if ok {
		t.Fatal("expected errors")
	}

	if !errors.Is(errs["Emails[0]"], v2.ErrMaxLen) {
		t.Fatalf("expected email max len")
	}
}
