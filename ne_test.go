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

func ExampleIsNe() {
	_, err := v2.IsNe("user", "admin")
	if err != nil {
		fmt.Println(err)
	}
}

func TestIsNeValid(t *testing.T) {
	_, err := v2.IsNe("user", "admin")
	if err != nil {
		t.Fatal(err)
	}
}

func TestIsNeInvalid(t *testing.T) {
	_, err := v2.IsNe("admin", "admin")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestIsNeInt(t *testing.T) {
	_, err := v2.IsNe(5, 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNeErrorMessage(t *testing.T) {
	_, err := v2.IsNe("admin", "admin")

	expected := "Value must not equal admin."

	if err.Error() != expected {
		t.Fatalf("actual %q expected %q", err.Error(), expected)
	}
}

func TestCheckNeNonString(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type User struct {
		Role int `checkers:"ne:0"`
	}

	user := &User{}

	v2.CheckStruct(user)
}

func TestCheckNeInvalid(t *testing.T) {
	type User struct {
		Role string `checkers:"ne:admin"`
	}

	user := &User{
		Role: "admin",
	}

	errs, ok := v2.CheckStruct(user)
	if ok {
		t.Fatal("expected error")
	}

	if !errors.Is(errs["Role"], v2.ErrEq) {
		t.Fatalf("expected ErrEq, got %v", errs)
	}
}

func TestCheckNeValid(t *testing.T) {
	type User struct {
		Role string `checkers:"ne:admin"`
	}

	user := &User{
		Role: "user",
	}

	errs, ok := v2.CheckStruct(user)
	if !ok {
		t.Fatal(errs)
	}
}
