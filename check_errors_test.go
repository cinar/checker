// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2_test

import (
	"encoding/json"
	"fmt"
	"net/http"
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

func TestCheckErrorsJSONWithCustomMessage(t *testing.T) {
	type Person struct {
		Name string `checkers:"required" checkersMsg:"required=Name is required"`
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

	if fields["Name"].Code != "REQUIRED" || fields["Name"].Message != "Name is required" {
		t.Fatalf("unexpected field error for Name: %+v", fields["Name"])
	}
}

func TestCheckErrorsProblemDetailsWithCustomMessage(t *testing.T) {
	type Person struct {
		Name string `checkers:"required" checkersMsg:"required=Name is required"`
	}

	person := &Person{}

	errs, ok := v2.CheckStruct(person)
	if ok {
		t.Fatal("expected errors")
	}

	pd := errs.ProblemDetails()

	if len(pd.InvalidParams) != 1 || pd.InvalidParams[0].Code != "REQUIRED" || pd.InvalidParams[0].Reason != "Name is required" {
		t.Fatalf("expected custom reason, got %+v", pd.InvalidParams)
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

func ExampleCheckErrors_ProblemDetails() {
	type Person struct {
		Name string `checkers:"required"`
	}

	person := &Person{}

	errs, ok := v2.CheckStruct(person)
	if !ok {
		data, err := json.Marshal(errs.ProblemDetails())
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(string(data))
	}
	// Output: {"type":"about:blank","title":"Your request parameters failed validation.","status":400,"invalid-params":[{"name":"Name","reason":"Required value is missing.","code":"REQUIRED"}]}
}

func TestCheckErrorsProblemDetailsDefaults(t *testing.T) {
	type Person struct {
		Name    string `checkers:"required"`
		Address string `checkers:"required"`
	}

	person := &Person{}

	errs, ok := v2.CheckStruct(person)
	if ok {
		t.Fatal("expected errors")
	}

	pd := errs.ProblemDetails()

	if pd.Type != "about:blank" {
		t.Fatalf("expected default Type, got %q", pd.Type)
	}

	if pd.Status != http.StatusBadRequest {
		t.Fatalf("expected default Status 400, got %d", pd.Status)
	}

	if len(pd.InvalidParams) != 2 {
		t.Fatalf("expected 2 invalid params, got %d", len(pd.InvalidParams))
	}

	// Sorted by field name, matching CheckErrors.Error()'s ordering.
	if pd.InvalidParams[0].Name != "Address" || pd.InvalidParams[1].Name != "Name" {
		t.Fatalf("expected sorted invalid params, got %+v", pd.InvalidParams)
	}

	if pd.InvalidParams[1].Code != "REQUIRED" || pd.InvalidParams[1].Reason != "Required value is missing." {
		t.Fatalf("unexpected invalid param for Name: %+v", pd.InvalidParams[1])
	}
}

func TestCheckErrorsProblemDetailsOverrides(t *testing.T) {
	type Person struct {
		Name string `checkers:"required"`
	}

	person := &Person{}

	errs, ok := v2.CheckStruct(person)
	if ok {
		t.Fatal("expected errors")
	}

	pd := errs.ProblemDetails()
	pd.Type = "https://example.com/probs/validation-error"
	pd.Title = "Custom title"
	pd.Status = http.StatusUnprocessableEntity

	if pd.Type != "https://example.com/probs/validation-error" || pd.Title != "Custom title" || pd.Status != http.StatusUnprocessableEntity {
		t.Fatalf("expected overrides to stick, got %+v", pd)
	}
}

func TestCheckErrorsProblemDetailsWithLocale(t *testing.T) {
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

	pd := errs.ProblemDetailsWithLocale(locale)

	if len(pd.InvalidParams) != 1 || pd.InvalidParams[0].Reason != message {
		t.Fatalf("expected localized reason %q, got %+v", message, pd.InvalidParams)
	}
}

func TestCheckErrorsProblemDetailsNonCheckError(t *testing.T) {
	v2.RegisterMaker("is-fruit-plain-error-problem-details", func(_ string) v2.CheckFunc[reflect.Value] {
		return func(value reflect.Value) (reflect.Value, error) {
			return value, fmt.Errorf("not a fruit")
		}
	})

	type Item struct {
		Name string `checkers:"is-fruit-plain-error-problem-details"`
	}

	item := &Item{
		Name: "onur",
	}

	errs, ok := v2.CheckStruct(item)
	if ok {
		t.Fatal("expected errors")
	}

	pd := errs.ProblemDetails()

	if len(pd.InvalidParams) != 1 || pd.InvalidParams[0].Code != "" || pd.InvalidParams[0].Reason != "not a fruit" {
		t.Fatalf("expected empty code and plain message, got %+v", pd.InvalidParams)
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
