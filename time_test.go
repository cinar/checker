// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2_test

import (
	"errors"
	"testing"

	v2 "github.com/cinar/checker/v2"
)

func ExampleIsTime() {
	value := "2024-12-31"

	_, err := v2.IsTime("DateOnly", value)
	if err != nil {
		panic(err)
	}
}

func ExampleIsTime_custom() {
	rfc3339Layout := "2006-01-02T15:04:05Z07:00"

	value := "2024-12-31T10:20:00Z07:00"

	_, err := v2.IsTime(rfc3339Layout, value)
	if err != nil {
		panic(err)
	}
}

func TestIsTimeSuccess(t *testing.T) {
	value := "2024-12-31"

	result, err := v2.IsTime("DateOnly", value)
	if result != value {
		t.Fatalf("result (%s) is not the original value (%s)", result, value)
	}

	if err != nil {
		t.Fatal(err)
	}
}

func TestIsTimeError(t *testing.T) {
	value := "2024-12-31"

	result, err := v2.IsTime("2006-02-01", value)
	if result != value {
		t.Fatalf("result (%s) is not the original value (%s)", result, value)
	}

	if !errors.Is(err, v2.ErrTime) {
		t.Fatalf("expected error %s actual %s", v2.ErrTime, err)
	}

	message := "Not a valid time."

	if err.Error() != message {
		t.Fatalf("expected %s actual %s", message, err.Error())
	}
}

func TestStructTimeError(t *testing.T) {
	type Person struct {
		Birthday string `checkers:"time:DateOnly"`
	}

	person := &Person{
		Birthday: "2024-31-12",
	}

	errs, ok := v2.CheckStruct(person)
	if ok {
		t.Fatalf("expected errors")
	}

	if !errors.Is(errs["Birthday"], v2.ErrTime) {
		t.Fatalf("expected time error")
	}
}
