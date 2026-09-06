// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2_test

import (
	"reflect"
	"testing"

	v2 "github.com/cinar/checker/v2"
)

func TestCheckStructMessagesOverridesScalarField(t *testing.T) {
	type Person struct {
		Name string `checkers:"required" checkersMsg:"required=Name is required"`
	}

	person := &Person{}

	errs, valid := v2.CheckStruct(person)
	if valid {
		t.Fatalf("actual %t expected %t", valid, false)
	}

	expected := "Name is required"

	if errs["Name"].Error() != expected {
		t.Fatalf("actual %s expected %s", errs["Name"].Error(), expected)
	}
}

func TestCheckStructMessagesTemplatePlaceholder(t *testing.T) {
	type Person struct {
		Name string `checkers:"min-len:8" checkersMsg:"min-len=Must be at least {{ .min }} characters"`
	}

	person := &Person{Name: "short"}

	errs, valid := v2.CheckStruct(person)
	if valid {
		t.Fatalf("actual %t expected %t", valid, false)
	}

	expected := "Must be at least 8 characters"

	if errs["Name"].Error() != expected {
		t.Fatalf("actual %s expected %s", errs["Name"].Error(), expected)
	}
}

func TestCheckStructMessagesNoTagFallsBackToLocale(t *testing.T) {
	type Person struct {
		Name string `checkers:"required"`
	}

	person := &Person{}

	errs, valid := v2.CheckStruct(person)
	if valid {
		t.Fatalf("actual %t expected %t", valid, false)
	}

	if errs["Name"].Error() != v2.ErrRequired.Error() {
		t.Fatalf("actual %s expected %s", errs["Name"].Error(), v2.ErrRequired.Error())
	}
}

func TestCheckStructMessagesUnmatchedNameIsNoOp(t *testing.T) {
	type Person struct {
		Name string `checkers:"required" checkersMsg:"email=Not used"`
	}

	person := &Person{}

	errs, valid := v2.CheckStruct(person)
	if valid {
		t.Fatalf("actual %t expected %t", valid, false)
	}

	if errs["Name"].Error() != v2.ErrRequired.Error() {
		t.Fatalf("actual %s expected %s", errs["Name"].Error(), v2.ErrRequired.Error())
	}
}

func TestCheckStructMessagesContainerAndItem(t *testing.T) {
	type Tags struct {
		Values []string `checkers:"@min-len:2 required" checkersMsg:"@min-len=Need at least 2 tags;required=Tag cannot be blank"`
	}

	tags := &Tags{Values: []string{"a", ""}}

	errs, valid := v2.CheckStruct(tags)
	if valid {
		t.Fatalf("actual %t expected %t", valid, false)
	}

	if _, ok := errs["Values"]; ok {
		t.Fatalf("expected no container-level error, got %v", errs["Values"])
	}

	expectedItem := "Tag cannot be blank"

	if errs["Values[1]"].Error() != expectedItem {
		t.Fatalf("actual %s expected %s", errs["Values[1]"].Error(), expectedItem)
	}

	tags = &Tags{Values: []string{"a"}}

	errs, valid = v2.CheckStruct(tags)
	if valid {
		t.Fatalf("actual %t expected %t", valid, false)
	}

	expectedContainer := "Need at least 2 tags"

	if errs["Values"].Error() != expectedContainer {
		t.Fatalf("actual %s expected %s", errs["Values"].Error(), expectedContainer)
	}
}

func TestReflectCheckWithConfigIgnoresMessages(t *testing.T) {
	value := reflect.ValueOf("")

	_, err := v2.ReflectCheckWithConfig(value, "required")
	if err == nil {
		t.Fatal("expected an error")
	}

	if err.Error() != v2.ErrRequired.Error() {
		t.Fatalf("actual %s expected %s", err.Error(), v2.ErrRequired.Error())
	}
}
