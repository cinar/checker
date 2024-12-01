// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

import (
	"fmt"
	"reflect"
	"strings"
)

// ReflectCheckFunc is a function that takes reflect value and performs
// a check on it. It updates the value and it returns any error that
// occurred during the check.
type ReflectCheckFunc func(value reflect.Value) error

// MakeCheckFunc is a function that returns a check function using the given params.
type MakeCheckFunc func(params string) ReflectCheckFunc

// makers provides a mapping of maker functions keyed by the check name.
var makers = map[string]MakeCheckFunc{
	nameMaxLen:    makeMaxLen,
	nameMinLen:    makeMinLen,
	nameRequired:  makeRequired,
	nameTrimSpace: makeTrimSpace,
}

// makeChecks take a checker config and returns the check functions.
func makeChecks(config string) []ReflectCheckFunc {
	fields := strings.Fields(config)

	checks := make([]ReflectCheckFunc, len(fields))

	for i, field := range fields {
		name, params, _ := strings.Cut(field, ":")

		maker, ok := makers[name]
		if !ok {
			panic(fmt.Sprintf("check %s not found", name))
		}

		checks[i] = maker(params)
	}

	return checks
}
