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

// MakeCheckFunc is a function that returns a check function using the given params.
type MakeCheckFunc func(params string) CheckFunc[reflect.Value]

// makers provides a mapping of maker functions keyed by the check name.
var makers = map[string]MakeCheckFunc{
	nameAlphanumeric: makeAlphanumeric,
	nameASCII:        makeASCII,
	nameCIDR:         makeCIDR,
	nameDigits:       makeDigits,
	nameEmail:        makeEmail,
	nameFQDN:         makeFQDN,
	nameIP:           makeIP,
	nameIPv4:         makeIPv4,
	nameIPv6:         makeIPv6,
	nameISBN:         makeISBN,
	nameLUHN:         makeLUHN,
	nameMAC:          makeMAC,
	nameMaxLen:       makeMaxLen,
	nameMinLen:       makeMinLen,
	nameRequired:     makeRequired,
	nameTrimSpace:    makeTrimSpace,
}

// makeChecks take a checker config and returns the check functions.
func makeChecks(config string) []CheckFunc[reflect.Value] {
	fields := strings.Fields(config)

	checks := make([]CheckFunc[reflect.Value], len(fields))

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
