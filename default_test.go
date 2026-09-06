// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2_test

import (
	"testing"

	v2 "github.com/cinar/checker/v2"
)

func TestDefaultStringZero(t *testing.T) {
	result, err := v2.Default("guest")("")
	if err != nil {
		t.Fatal(err)
	}

	if result != "guest" {
		t.Fatalf("actual %s expected %s", result, "guest")
	}
}

func TestDefaultStringNonZero(t *testing.T) {
	result, err := v2.Default("guest")("alice")
	if err != nil {
		t.Fatal(err)
	}

	if result != "alice" {
		t.Fatalf("actual %s expected %s", result, "alice")
	}
}

func TestDefaultIntZero(t *testing.T) {
	result, err := v2.Default(3)(0)
	if err != nil {
		t.Fatal(err)
	}

	if result != 3 {
		t.Fatalf("actual %d expected %d", result, 3)
	}
}

func TestReflectDefaultStringZero(t *testing.T) {
	type Person struct {
		Role string `checkers:"default:guest"`
	}

	person := &Person{}

	if _, ok := v2.CheckStruct(person); !ok {
		t.Fatal("expected valid")
	}

	if person.Role != "guest" {
		t.Fatalf("actual %s expected %s", person.Role, "guest")
	}
}

func TestReflectDefaultStringNonZero(t *testing.T) {
	type Person struct {
		Role string `checkers:"default:guest"`
	}

	person := &Person{Role: "admin"}

	if _, ok := v2.CheckStruct(person); !ok {
		t.Fatal("expected valid")
	}

	if person.Role != "admin" {
		t.Fatalf("actual %s expected %s", person.Role, "admin")
	}
}

func TestReflectDefaultBool(t *testing.T) {
	type Settings struct {
		Enabled bool `checkers:"default:true"`
	}

	settings := &Settings{}

	if _, ok := v2.CheckStruct(settings); !ok {
		t.Fatal("expected valid")
	}

	if !settings.Enabled {
		t.Fatal("expected Enabled to default to true")
	}
}

func TestReflectDefaultInt(t *testing.T) {
	type Order struct {
		Retries int `checkers:"default:3"`
	}

	order := &Order{}

	if _, ok := v2.CheckStruct(order); !ok {
		t.Fatal("expected valid")
	}

	if order.Retries != 3 {
		t.Fatalf("actual %d expected %d", order.Retries, 3)
	}
}

func TestReflectDefaultUint(t *testing.T) {
	type Order struct {
		Retries uint `checkers:"default:3"`
	}

	order := &Order{}

	if _, ok := v2.CheckStruct(order); !ok {
		t.Fatal("expected valid")
	}

	if order.Retries != 3 {
		t.Fatalf("actual %d expected %d", order.Retries, 3)
	}
}

func TestReflectDefaultFloat(t *testing.T) {
	type Order struct {
		Price float64 `checkers:"default:9.99"`
	}

	order := &Order{}

	if _, ok := v2.CheckStruct(order); !ok {
		t.Fatal("expected valid")
	}

	if order.Price != 9.99 {
		t.Fatalf("actual %f expected %f", order.Price, 9.99)
	}
}

func TestReflectDefaultPointerNil(t *testing.T) {
	type Order struct {
		Timeout *int `checkers:"default:30"`
	}

	order := &Order{}

	if _, ok := v2.CheckStruct(order); !ok {
		t.Fatal("expected valid")
	}

	if order.Timeout == nil || *order.Timeout != 30 {
		t.Fatalf("expected Timeout to default to 30, got %v", order.Timeout)
	}
}

func TestReflectDefaultPointerNonNil(t *testing.T) {
	type Order struct {
		Timeout *int `checkers:"default:30"`
	}

	given := 5
	order := &Order{Timeout: &given}

	if _, ok := v2.CheckStruct(order); !ok {
		t.Fatal("expected valid")
	}

	if order.Timeout != &given || *order.Timeout != 5 {
		t.Fatalf("expected Timeout to stay 5, got %v", order.Timeout)
	}
}

func TestReflectDefaultOrderingWithTrimAndRequired(t *testing.T) {
	type Person struct {
		Role string `checkers:"trim default:guest required"`
	}

	person := &Person{Role: "   "}

	errs, ok := v2.CheckStruct(person)
	if !ok {
		t.Fatalf("expected valid, got %v", errs)
	}

	if person.Role != "guest" {
		t.Fatalf("actual %s expected %s", person.Role, "guest")
	}
}

func TestReflectDefaultInvalidBool(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Settings struct {
		Enabled bool `checkers:"default:abcd"`
	}

	v2.CheckStruct(&Settings{})
}

func TestReflectDefaultInvalidInt(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Order struct {
		Retries int `checkers:"default:abcd"`
	}

	v2.CheckStruct(&Order{})
}

func TestReflectDefaultInvalidUint(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Order struct {
		Retries uint `checkers:"default:abcd"`
	}

	v2.CheckStruct(&Order{})
}

func TestReflectDefaultInvalidFloat(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Order struct {
		Price float64 `checkers:"default:abcd"`
	}

	v2.CheckStruct(&Order{})
}

func TestReflectDefaultUnsupportedKind(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Order struct {
		Tags []string `checkers:"@default:abcd"`
	}

	v2.CheckStruct(&Order{})
}
