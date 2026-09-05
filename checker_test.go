// Copyright (c) 2023-2026 Onur Cinar.
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

func TestCheckStructNilPointerFieldSkipsChildFields(t *testing.T) {
	type Address struct {
		Street string `checkers:"required"`
	}

	type Person struct {
		Name    string `checkers:"required"`
		Address *Address
	}

	person := &Person{
		Name: "Onur Cinar",
	}

	if _, ok := v2.CheckStruct(person); !ok {
		t.Fatal("expected valid, nil Address has nothing to check")
	}
}

func TestCheckStructNilPointerFieldRequired(t *testing.T) {
	type Address struct {
		Street string `checkers:"required"`
	}

	type Person struct {
		Name    string   `checkers:"required"`
		Address *Address `checkers:"required"`
	}

	person := &Person{
		Name: "Onur Cinar",
	}

	errs, ok := v2.CheckStruct(person)
	if ok {
		t.Fatal("expected errors")
	}

	if !errors.Is(errs["Address"], v2.ErrRequired) {
		t.Fatalf("expected Address required %v", errs)
	}
}

func TestCheckStructPassedByValueDoesNotPanic(t *testing.T) {
	type Person struct {
		Name string `checkers:"trim required"`
	}

	person := Person{
		Name: "  Onur Cinar  ",
	}

	//nolint:staticcheck // deliberately passing by value to exercise the unaddressable path
	if _, ok := v2.CheckStruct(person); !ok {
		t.Fatal("expected valid")
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

func TestCheckStructCustomNameWithOptions(t *testing.T) {
	type Person struct {
		Name string `json:"name,omitempty" checkers:"required"`
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

func TestCheckStructJSONIgnoredFieldFallsBackToFieldName(t *testing.T) {
	type Person struct {
		Name string `json:"-" checkers:"required"`
	}

	person := &Person{
		Name: "",
	}

	errs, ok := v2.CheckStruct(person)
	if ok {
		t.Fatal("expected errors")
	}

	if !errors.Is(errs["Name"], v2.ErrRequired) {
		t.Fatalf("expected Name required %v", errs)
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

func TestCheckStructSliceConfigSplitIsCached(t *testing.T) {
	// Checking a second value of the same struct type exercises the
	// warm-cache path of the slice/map "@"-prefix container/item config
	// split, not just the cold, first-parse path.
	type Person struct {
		Name   string   `checkers:"required"`
		Emails []string `checkers:"@max-len:1 max-len:4"`
	}

	people := []*Person{
		{Name: "Onur Cinar", Emails: []string{"abcd"}},
		{Name: "Jane Doe", Emails: []string{"efgh"}},
	}

	for _, person := range people {
		if _, ok := v2.CheckStruct(person); !ok {
			t.Fatalf("expected valid for %+v", person)
		}
	}
}

func TestCheckStructMapNormalizesAndValidates(t *testing.T) {
	type Person struct {
		Name   string            `checkers:"required"`
		Emails map[string]string `checkers:"@min-len:1 trim max-len:3"`
	}

	person := &Person{
		Name: "Onur Cinar",
		Emails: map[string]string{
			"work": " onur ",
		},
	}

	errs, ok := v2.CheckStruct(person)
	if ok {
		t.Fatal("expected errors")
	}

	if !errors.Is(errs["Emails[work]"], v2.ErrMaxLen) {
		t.Fatalf("expected email max len %v", errs)
	}

	if person.Emails["work"] != "onur" {
		t.Fatalf("expected trimmed value to be written back, got %q", person.Emails["work"])
	}
}

func TestCheckStructMapEmpty(t *testing.T) {
	type Person struct {
		Name   string            `checkers:"required"`
		Emails map[string]string `checkers:"@min-len:1"`
	}

	person := &Person{
		Name: "Onur Cinar",
	}

	errs, ok := v2.CheckStruct(person)
	if ok {
		t.Fatal("expected errors")
	}

	if !errors.Is(errs["Emails"], v2.ErrMinLen) {
		t.Fatalf("expected emails min len %v", errs)
	}
}

func TestCheckStructMapOfStructPointers(t *testing.T) {
	type Address struct {
		Street string `checkers:"required"`
	}

	type Person struct {
		Name      string `checkers:"required"`
		Addresses map[string]*Address
	}

	person := &Person{
		Name: "Onur Cinar",
		Addresses: map[string]*Address{
			"home": {Street: ""},
		},
	}

	errs, ok := v2.CheckStruct(person)
	if ok {
		t.Fatal("expected errors")
	}

	if !errors.Is(errs["Addresses[home].Street"], v2.ErrRequired) {
		t.Fatalf("expected street required %v", errs)
	}

	if person.Addresses["home"] == nil {
		t.Fatal("expected the map entry pointer to be preserved")
	}
}

func TestCheckStructMapOfPointersNormalizesInPlace(t *testing.T) {
	type Person struct {
		Name  string             `checkers:"required"`
		Notes map[string]*string `checkers:"trim"`
	}

	note := "  hello  "

	person := &Person{
		Name: "Onur Cinar",
		Notes: map[string]*string{
			"first": &note,
		},
	}

	errs, ok := v2.CheckStruct(person)
	if !ok {
		t.Fatalf("got unexpected errors %v", errs)
	}

	if *person.Notes["first"] != "hello" {
		t.Fatalf("expected trimmed value to be written back through the pointer, got %q", *person.Notes["first"])
	}
}
