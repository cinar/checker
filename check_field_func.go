// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

import (
	"fmt"
	"reflect"
)

// CheckFieldFunc is a function that takes the parent struct's reflect.Value along with
// the current field's reflect.Value, performs a check that may depend on a sibling
// field, and returns the resulting value and any error that occurred. The parent is
// only valid when the check is run through CheckStruct.
type CheckFieldFunc func(parent, value reflect.Value) (reflect.Value, error)

// lookupParentField returns the dereferenced value of the named field on the parent
// struct. It panics if the parent is not a valid struct context, or if the named
// field cannot be found, as both indicate a configuration error.
func lookupParentField(checkerName string, parent reflect.Value, name string) reflect.Value {
	if !parent.IsValid() {
		panic(fmt.Sprintf("%s requires a struct context, use CheckStruct", checkerName))
	}

	field := parent.FieldByName(name)
	if !field.IsValid() {
		panic(fmt.Sprintf("field %s not found", name))
	}

	return reflect.Indirect(field)
}
