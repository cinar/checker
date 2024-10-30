// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package checker

import (
	"reflect"
	"strings"
	"testing"
)

func TestInitCheckersUnknown(t *testing.T) {
	defer FailIfNoPanic(t)

	initCheckers("unknown")
}

func TestInitCheckersKnwon(t *testing.T) {
	checkers := initCheckers("required")

	if len(checkers) != 1 {
		t.Fail()
	}

	if reflect.ValueOf(checkers[0]).Pointer() != reflect.ValueOf(checkRequired).Pointer() {
		t.Fail()
	}
}

func TestRegister(t *testing.T) {
	var checker CheckFunc = func(_, _ reflect.Value) error {
		return nil
	}

	var maker MakeFunc = func(_ string) CheckFunc {
		return checker
	}

	name := "test"

	Register(name, maker)

	checkers := initCheckers(name)

	if len(checkers) != 1 {
		t.Fail()
	}

	if reflect.ValueOf(checkers[0]).Pointer() != reflect.ValueOf(checker).Pointer() {
		t.Fail()
	}
}

func TestCheckInvalid(t *testing.T) {
	type Person struct {
		Name string `checkers:"required"`
	}

	person := &Person{}

	errors, valid := Check(person)
	if valid {
		t.Fail()
	}

	if len(errors) != 1 {
		t.Fail()
	}

	if errors["Name"] != ErrRequired {
		t.Fail()
	}
}

func TestCheckValid(t *testing.T) {
	type Person struct {
		Name string `checkers:"required"`
	}

	person := &Person{
		Name: "Onur",
	}

	errors, valid := Check(person)
	if !valid {
		t.Fail()
	}

	if len(errors) != 0 {
		t.Fail()
	}
}

func TestCheckNoStruct(t *testing.T) {
	defer FailIfNoPanic(t)

	s := "unknown"
	Check(s)
}

func TestCheckNestedStruct(t *testing.T) {
	type Address struct {
		Street string `checkers:"required"`
	}

	type Person struct {
		Name string `checkers:"required"`
		Home Address
	}

	person := &Person{}

	errors, valid := Check(person)
	if valid {
		t.Fail()
	}

	if len(errors) != 2 {
		t.Fail()
	}

	if errors["Name"] != ErrRequired {
		t.Fail()
	}

	if errors["Home.Street"] != ErrRequired {
		t.Fail()
	}
}

func TestNumberOfInvalid(t *testing.T) {
	defer FailIfNoPanic(t)

	s := "invalid"

	numberOf(reflect.ValueOf(s))
}

func TestNumberOfInt(t *testing.T) {
	n := 10

	if numberOf(reflect.ValueOf(n)) != float64(n) {
		t.Fail()
	}
}

func TestNumberOfFloat(t *testing.T) {
	n := float64(10.10)

	if numberOf(reflect.ValueOf(n)) != n {
		t.Fail()
	}
}

func TestCheckerNamesAreLowerCase(t *testing.T) {
	for checker := range makers {
		if strings.ToLower(checker) != checker {
			t.Fail()
		}
	}
}

func BenchmarkCheck(b *testing.B) {
	type Address struct {
		Street string `checkers:"required"`
	}

	type Person struct {
		Name string `checkers:"required"`
		Home Address
	}

	person := &Person{}

	for n := 0; n < b.N; n++ {
		Check(person)
	}
}
