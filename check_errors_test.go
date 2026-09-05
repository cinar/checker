// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2_test

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	v2 "github.com/cinar/checker/v2"
	"github.com/cinar/checker/v2/locales"
)

func TestCheckErrorsErrorEmpty(t *testing.T) {
	errs := v2.CheckErrors{}

	if errs.Error() != "" {
		t.Fatalf("expected empty message, got %q", errs.Error())
	}
}

func TestCheckErrorsErrorSortedByFieldName(t *testing.T) {
	type Person struct {
		Name    string `checkers:"required"`
		Address string `checkers:"required"`
	}

	person := &Person{}

	errs, ok := v2.CheckStruct(person)
	if ok {
		t.Fatal("expected errors")
	}

	expected := "Address: Required value is missing.; Name: Required value is missing."

	if errs.Error() != expected {
		t.Fatalf("actual %q expected %q", errs.Error(), expected)
	}
}

func ExampleCheckErrors_JSON() {
	type Person struct {
		Name string `checkers:"required"`
	}

	person := &Person{}

	errs, ok := v2.CheckStruct(person)
	if !ok {
		data, err := errs.JSON()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(string(data))
	}
	// Output: {"Name":{"code":"REQUIRED","message":"Required value is missing."}}
}

func TestCheckErrorsJSON(t *testing.T) {
	type Person struct {
		Name  string `checkers:"required"`
		Email string `checkers:"required"`
	}

	person := &Person{}

	errs, ok := v2.CheckStruct(person)
	if ok {
		t.Fatal("expected errors")
	}

	data, err := errs.JSON()
	if err != nil {
		t.Fatal(err)
	}

	var fields map[string]v2.FieldError

	if err := json.Unmarshal(data, &fields); err != nil {
		t.Fatal(err)
	}

	if len(fields) != 2 {
		t.Fatalf("expected 2 fields, got %d", len(fields))
	}

	if fields["Name"].Code != "REQUIRED" || fields["Name"].Message != "Required value is missing." {
		t.Fatalf("unexpected field error for Name: %+v", fields["Name"])
	}
}

func TestCheckErrorsJSONWithLocale(t *testing.T) {
	locale := "de-DE"
	code := "REQUIRED"
	message := "Erforderlich."

	v2.RegisterLocale(locale, map[string]string{
		code: message,
	})

	type Person struct {
		Name string `checkers:"required"`
	}

	person := &Person{}

	errs, ok := v2.CheckStruct(person)
	if ok {
		t.Fatal("expected errors")
	}

	data, err := errs.JSONWithLocale(locale)
	if err != nil {
		t.Fatal(err)
	}

	var fields map[string]v2.FieldError

	if err := json.Unmarshal(data, &fields); err != nil {
		t.Fatal(err)
	}

	if fields["Name"].Message != message {
		t.Fatalf("actual %q expected %q", fields["Name"].Message, message)
	}
}

func TestCheckErrorsJSONNonCheckError(t *testing.T) {
	locales.EnUSMessages["NOT_FRUIT"] = "Not a fruit name."

	v2.RegisterMaker("is-fruit-plain-error", func(_ string) v2.CheckFunc[reflect.Value] {
		return func(value reflect.Value) (reflect.Value, error) {
			return value, fmt.Errorf("not a fruit")
		}
	})

	type Item struct {
		Name string `checkers:"is-fruit-plain-error"`
	}

	item := &Item{
		Name: "onur",
	}

	errs, ok := v2.CheckStruct(item)
	if ok {
		t.Fatal("expected errors")
	}

	data, err := errs.JSON()
	if err != nil {
		t.Fatal(err)
	}

	var fields map[string]v2.FieldError

	if err := json.Unmarshal(data, &fields); err != nil {
		t.Fatal(err)
	}

	if fields["Name"].Code != "" {
		t.Fatalf("expected empty code for a plain error, got %q", fields["Name"].Code)
	}

	if fields["Name"].Message != "not a fruit" {
		t.Fatalf("actual %q expected %q", fields["Name"].Message, "not a fruit")
	}
}

func TestCheckErrorsAsError(t *testing.T) {
	type Person struct {
		Name string `checkers:"required"`
	}

	person := &Person{}

	errs, ok := v2.CheckStruct(person)
	if ok {
		t.Fatal("expected errors")
	}

	var err error = errs

	if err.Error() != errs.Error() {
		t.Fatalf("expected CheckErrors to satisfy the error interface")
	}
}
