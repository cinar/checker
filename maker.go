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

// MakeCheckFieldFunc is a function that returns a field-relative check function
// using the given params.
type MakeCheckFieldFunc func(params string) CheckFieldFunc

// fieldMakers provides a mapping of maker functions for field-relative checks
// keyed by the check name.
var fieldMakers = map[string]MakeCheckFieldFunc{
	nameEqField:        makeEqField,
	nameRequiredIf:     makeRequiredIf,
	nameRequiredUnless: makeRequiredUnless,
}

// makers provides a mapping of maker functions keyed by the check name.
var makers = map[string]MakeCheckFunc{
	nameAfter:        makeAfter,
	nameAlphanumeric: makeAlphanumeric,
	nameASCII:        makeASCII,
	nameBefore:       makeBefore,
	nameCIDR:         makeCIDR,
	nameCreditCard:   makeCreditCard,
	nameDigits:       makeDigits,
	nameEmail:        makeEmail,
	nameEOA:          makeEOA,
	nameFQDN:         makeFQDN,
	nameGte:          makeGte,
	nameHash:         makeHash,
	nameHex:          makeHex,
	nameHTMLEscape:   makeHTMLEscape,
	nameHTMLUnescape: makeHTMLUnescape,
	nameIP:           makeIP,
	nameIPv4:         makeIPv4,
	nameIPv6:         makeIPv6,
	nameISBN:         makeISBN,
	nameLower:        makeLower,
	nameLte:          makeLte,
	nameLUHN:         makeLUHN,
	nameMAC:          makeMAC,
	nameMaxLen:       makeMaxLen,
	nameMinLen:       makeMinLen,
	nameRegexp:       makeRegexp,
	nameRequired:     makeRequired,
	nameTime:         makeTime,
	nameTitle:        makeTitle,
	nameTrimLeft:     makeTrimLeft,
	nameTrimRight:    makeTrimRight,
	nameTrimSpace:    makeTrimSpace,
	nameUpper:        makeUpper,
	nameURL:          makeURL,
	nameURLEscape:    makeURLEscape,
	nameURLUnescape:  makeURLUnescape,
}

// RegisterMaker registers a new maker function with the given name.
func RegisterMaker(name string, maker MakeCheckFunc) {
	makers[name] = maker
}

// RegisterFieldMaker registers a new field-relative maker function with the given name.
func RegisterFieldMaker(name string, maker MakeCheckFieldFunc) {
	fieldMakers[name] = maker
}

// makeChecks take a checker config and the parent struct's reflect.Value (invalid
// unless called from CheckStruct), and returns the check functions.
func makeChecks(config string, parent reflect.Value) []CheckFunc[reflect.Value] {
	fields := strings.Fields(config)

	checks := make([]CheckFunc[reflect.Value], len(fields))

	for i, field := range fields {
		name, params, _ := strings.Cut(field, ":")

		if fieldMaker, ok := fieldMakers[name]; ok {
			fieldCheck := fieldMaker(params)

			checks[i] = func(value reflect.Value) (reflect.Value, error) {
				return fieldCheck(parent, value)
			}

			continue
		}

		maker, ok := makers[name]
		if !ok {
			panic(fmt.Sprintf("check %s not found", name))
		}

		checks[i] = maker(params)
	}

	return checks
}
