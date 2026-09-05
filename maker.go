// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

import (
	"fmt"
	"reflect"
	"strings"
	"sync"
)

// MakeCheckFunc is a function that returns a check function using the given params.
type MakeCheckFunc func(params string) CheckFunc[reflect.Value]

// MakeCheckFieldFunc is a function that returns a field-relative check function
// using the given params.
type MakeCheckFieldFunc func(params string) CheckFieldFunc

// makersMu guards both fieldMakers and makers below. They're populated at
// init time and then only mutated through RegisterMaker/RegisterFieldMaker,
// but those can run concurrently with CheckStruct calls (e.g. a custom
// checker registered during server startup while requests are already being
// handled), so every read and write goes through this mutex.
var makersMu sync.RWMutex

// fieldMakers provides a mapping of maker functions for field-relative checks
// keyed by the check name.
var fieldMakers = map[string]MakeCheckFieldFunc{
	nameAfterField:     makeAfterField,
	nameBeforeField:    makeBeforeField,
	nameEqField:        makeEqField,
	nameRequiredIf:     makeRequiredIf,
	nameRequiredUnless: makeRequiredUnless,
}

// makers provides a mapping of maker functions keyed by the check name.
var makers = map[string]MakeCheckFunc{
	nameAfter:          makeAfter,
	nameAlpha:          makeAlpha,
	nameAlphanumeric:   makeAlphanumeric,
	nameASCII:          makeASCII,
	nameBefore:         makeBefore,
	nameCIDR:           makeCIDR,
	nameContains:       makeContains,
	nameCreditCard:     makeCreditCard,
	nameDigits:         makeDigits,
	nameEmail:          makeEmail,
	nameEndsWith:       makeEndsWith,
	nameEOA:            makeEOA,
	nameEq:             makeEq,
	nameFQDN:           makeFQDN,
	nameGt:             makeGt,
	nameGte:            makeGte,
	nameHash:           makeHash,
	nameHex:            makeHex,
	nameHTMLEscape:     makeHTMLEscape,
	nameHTMLUnescape:   makeHTMLUnescape,
	nameIP:             makeIP,
	nameIPv4:           makeIPv4,
	nameIPv6:           makeIPv6,
	nameISBN:           makeISBN,
	nameISO31661Alpha2: makeISO31661Alpha2,
	nameISO31661Alpha3: makeISO31661Alpha3,
	nameISO6391:        makeISO6391,
	nameLen:            makeLen,
	nameLower:          makeLower,
	nameLt:             makeLt,
	nameLte:            makeLte,
	nameLUHN:           makeLUHN,
	nameMAC:            makeMAC,
	nameMaxLen:         makeMaxLen,
	nameMinLen:         makeMinLen,
	nameNe:             makeNe,
	nameNumeric:        makeNumeric,
	nameOneOf:          makeOneOf,
	nameRegexp:         makeRegexp,
	nameRequired:       makeRequired,
	nameStartsWith:     makeStartsWith,
	nameTime:           makeTime,
	nameTitle:          makeTitle,
	nameTrimLeft:       makeTrimLeft,
	nameTrimRight:      makeTrimRight,
	nameTrimSpace:      makeTrimSpace,
	nameUpper:          makeUpper,
	nameURL:            makeURL,
	nameURLEscape:      makeURLEscape,
	nameURLUnescape:    makeURLUnescape,
	nameUUID:           makeUUID,
}

// RegisterMaker registers a new maker function with the given name.
func RegisterMaker(name string, maker MakeCheckFunc) {
	makersMu.Lock()
	defer makersMu.Unlock()

	makers[name] = maker
}

// RegisterFieldMaker registers a new field-relative maker function with the given name.
func RegisterFieldMaker(name string, maker MakeCheckFieldFunc) {
	makersMu.Lock()
	defer makersMu.Unlock()

	fieldMakers[name] = maker
}

// RegisteredMakerNames returns the name of every currently registered
// non-field-relative checker/normalizer maker, including built-ins and any
// custom makers added via RegisterMaker. The order is not significant.
// Intended for tooling built on top of checker, such as the checkerlint
// static analyzer, that needs to know the current checker vocabulary.
func RegisteredMakerNames() []string {
	makersMu.RLock()
	defer makersMu.RUnlock()

	names := make([]string, 0, len(makers))
	for name := range makers {
		names = append(names, name)
	}

	return names
}

// RegisteredFieldMakerNames returns the name of every currently registered
// field-relative checker maker, including built-ins and any custom makers
// added via RegisterFieldMaker. The order is not significant. Intended for
// tooling built on top of checker, such as the checkerlint static
// analyzer, that needs to know the current checker vocabulary.
func RegisteredFieldMakerNames() []string {
	makersMu.RLock()
	defer makersMu.RUnlock()

	names := make([]string, 0, len(fieldMakers))
	for name := range fieldMakers {
		names = append(names, name)
	}

	return names
}

// makeChecks take a checker config and the parent struct's reflect.Value (invalid
// unless called from CheckStruct), and returns the check functions.
func makeChecks(config string, parent reflect.Value) []CheckFunc[reflect.Value] {
	fields := strings.Fields(config)

	checks := make([]CheckFunc[reflect.Value], len(fields))

	for i, field := range fields {
		name, params, _ := strings.Cut(field, ":")

		makersMu.RLock()
		fieldMaker, isFieldMaker := fieldMakers[name]
		maker, isMaker := makers[name]
		makersMu.RUnlock()

		if isFieldMaker {
			fieldCheck := fieldMaker(params)

			checks[i] = func(value reflect.Value) (reflect.Value, error) {
				return fieldCheck(parent, value)
			}

			continue
		}

		if !isMaker {
			panic(fmt.Sprintf("check %s not found", name))
		}

		checks[i] = maker(params)
	}

	return checks
}
