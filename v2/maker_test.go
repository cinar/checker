// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2_test

import (
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
