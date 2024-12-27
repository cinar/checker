// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2_test

import (
	"fmt"
	"reflect"
	"testing"

	v2 "github.com/cinar/checker/v2"
)

func TestMakeCheckersUnknown(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Person struct {
		Name string `checkers:"unknown"`
	}

	person := &Person{
		Name: "Onur",
	}

	v2.CheckStruct(person)
}

func ExampleRegisterMaker() {
	v2.RegisterMaker("is-fruit", func(params string) v2.CheckFunc[reflect.Value] {
		return func(value reflect.Value) (reflect.Value, error) {
			stringValue := value.Interface().(string)

			if stringValue == "apple" || stringValue == "banana" {
				return value, nil
			}

			return value, v2.NewCheckError("NOT_FRUIT")
		}
	})

	type Item struct {
		Name string `checkers:"is-fruit"`
	}

	person := &Item{
		Name: "banana",
	}

	err, ok := v2.CheckStruct(person)
	if !ok {
		fmt.Println(err)
	}
}

func TestRegisterMaker(t *testing.T) {
	v2.RegisterMaker("unknown", func(params string) v2.CheckFunc[reflect.Value] {
		return func(value reflect.Value) (reflect.Value, error) {
			return value, nil
		}
	})

	type Person struct {
		Name string `checkers:"unknown"`
	}

	person := &Person{
		Name: "Onur",
	}

	_, ok := v2.CheckStruct(person)
	if !ok {
		t.Fatal("expected valid")
	}
}
