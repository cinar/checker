// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/cinar/checker/v2/locales"

	v2 "github.com/cinar/checker/v2"
)

func TestRegisteredMakerNamesIncludesBuiltins(t *testing.T) {
	names := v2.RegisteredMakerNames()

	found := false
	for _, name := range names {
		if name == "required" {
			found = true
			break
		}
	}

	if !found {
		t.Fatalf("expected \"required\" among registered maker names, got %v", names)
	}
}

func TestRegisteredFieldMakerNamesIncludesBuiltins(t *testing.T) {
	names := v2.RegisteredFieldMakerNames()

	found := false
	for _, name := range names {
		if name == "eq-field" {
			found = true
			break
		}
	}

	if !found {
		t.Fatalf("expected \"eq-field\" among registered field maker names, got %v", names)
	}
}

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
	locales.EnUSMessages["NOT_FRUIT"] = "Not a fruit name."

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

func TestRegisterMakerInvalidatesConfigCache(t *testing.T) {
	const name = "register-maker-invalidation-test"

	type Person struct {
		Name string `checkers:"register-maker-invalidation-test"`
	}

	v2.RegisterMaker(name, func(params string) v2.CheckFunc[reflect.Value] {
		return func(value reflect.Value) (reflect.Value, error) {
			return value, v2.NewCheckError("REJECTED_BY_OLD_MAKER")
		}
	})

	// Validating once here, before the maker is replaced below, is the
	// whole point of the test: it forces the "register-maker-invalidation-test"
	// checkers tag config to be resolved and cached against the *old* maker.
	if _, ok := v2.CheckStruct(&Person{Name: "Onur"}); ok {
		t.Fatal("expected the old maker to reject the value")
	}

	v2.RegisterMaker(name, func(params string) v2.CheckFunc[reflect.Value] {
		return func(value reflect.Value) (reflect.Value, error) {
			return value, nil
		}
	})

	// If re-registering the same name didn't invalidate the cached,
	// already-compiled config, this would still run the old maker and fail.
	if _, ok := v2.CheckStruct(&Person{Name: "Onur"}); !ok {
		t.Fatal("expected the new maker to take effect after re-registration")
	}
}

func TestRegisterFieldMaker(t *testing.T) {
	v2.RegisterFieldMaker("eq-const", func(params string) v2.CheckFieldFunc {
		return func(parent, value reflect.Value) (reflect.Value, error) {
			if value.Interface().(string) != params {
				return value, v2.NewCheckError("NOT_EQ_CONST")
			}

			return value, nil
		}
	})

	type Person struct {
		Name string `checkers:"eq-const:Onur"`
	}

	person := &Person{
		Name: "Onur",
	}

	_, ok := v2.CheckStruct(person)
	if !ok {
		t.Fatal("expected valid")
	}
}
