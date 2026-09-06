// Copyright (c) 2023-2026 Onur Cinar.
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
	nameBase64:         makeBase64,
	nameBase64URL:      makeBase64URL,
	nameBefore:         makeBefore,
	nameCIDR:           makeCIDR,
	nameContains:       makeContains,
	nameCreditCard:     makeCreditCard,
	nameDefault:        makeDefault,
	nameDigits:         makeDigits,
	nameE164:           makeE164,
	nameEmail:          makeEmail,
	nameEndsWith:       makeEndsWith,
	nameEOA:            makeEOA,
	nameEq:             makeEq,
	nameFinite:         makeFinite,
	nameFQDN:           makeFQDN,
	nameGt:             makeGt,
	nameGte:            makeGte,
	nameHash:           makeHash,
	nameHex:            makeHex,
	nameHexColor:       makeHexColor,
	nameHTMLEscape:     makeHTMLEscape,
	nameHTMLUnescape:   makeHTMLUnescape,
	nameIBAN:           makeIBAN,
	nameInt:            makeInt,
	nameIP:             makeIP,
	nameIPv4:           makeIPv4,
	nameIPv6:           makeIPv6,
	nameISBN:           makeISBN,
	nameISO31661Alpha2: makeISO31661Alpha2,
	nameISO31661Alpha3: makeISO31661Alpha3,
	nameISO6391:        makeISO6391,
	nameJWT:            makeJWT,
	nameLen:            makeLen,
	nameLower:          makeLower,
	nameLt:             makeLt,
	nameLte:            makeLte,
	nameLUHN:           makeLUHN,
	nameMAC:            makeMAC,
	nameMaxLen:         makeMaxLen,
	nameMinLen:         makeMinLen,
	nameMongoID:        makeMongoID,
	nameMultipleOf:     makeMultipleOf,
	nameNe:             makeNe,
	nameNegative:       makeNegative,
	nameNonnegative:    makeNonnegative,
	nameNumeric:        makeNumeric,
	nameOneOf:          makeOneOf,
	namePositive:       makePositive,
	namePostalCode:     makePostalCode,
	nameRegexp:         makeRegexp,
	nameRequired:       makeRequired,
	nameSemver:         makeSemver,
	nameSlug:           makeSlug,
	nameStartsWith:     makeStartsWith,
	nameStripInvisible: makeStripInvisible,
	nameTime:           makeTime,
	nameTitle:          makeTitle,
	nameTrimLeft:       makeTrimLeft,
	nameTrimRight:      makeTrimRight,
	nameTrimSpace:      makeTrimSpace,
	nameULID:           makeULID,
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
	configCache.Clear()
}

// RegisterFieldMaker registers a new field-relative maker function with the given name.
func RegisterFieldMaker(name string, maker MakeCheckFieldFunc) {
	makersMu.Lock()
	defer makersMu.Unlock()

	fieldMakers[name] = maker
	configCache.Clear()
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

// compiledCheck is a checkers/validate tag token already resolved to its
// checker/normalizer closure. fieldFn is set instead of fn for a
// field-relative check, since that still needs the enclosing struct's
// reflect.Value bound in at call time, which varies per CheckStruct call
// and so can't be baked in at compile time like fn can.
type compiledCheck struct {
	fn      CheckFunc[reflect.Value]
	fieldFn CheckFieldFunc
	name    string
}

// run executes the compiled check against value, threading parent through
// only if the check is field-relative.
func (c compiledCheck) run(parent, value reflect.Value) (reflect.Value, error) {
	if c.fieldFn != nil {
		return c.fieldFn(parent, value)
	}

	return c.fn(value)
}

// compiledConfig is the parsed, maker-resolved form of a checkers/validate
// tag config string: the omitempty modifier extracted, and every remaining
// token already resolved against the makers/fieldMakers registries.
type compiledConfig struct {
	omitEmpty bool
	checks    []compiledCheck
}

// configCache caches *compiledConfig by its exact source config string, so
// a given tag value (e.g. "trim required email") is only ever tokenized
// and resolved once, no matter how many struct instances, fields, or types
// share that exact string. RegisterMaker and RegisterFieldMaker clear it,
// since a newly (re-)registered name can change how an already-cached
// string resolves.
var configCache sync.Map

// getCompiledConfig returns the compiledConfig for config, compiling and
// caching it on first use.
func getCompiledConfig(config string) *compiledConfig {
	if cached, ok := configCache.Load(config); ok {
		return cached.(*compiledConfig)
	}

	// Held for the rest of the function, including the store into
	// configCache below: RegisterMaker/RegisterFieldMaker take the write
	// lock and then clear configCache, so as long as a compile-in-flight
	// keeps the read lock until its result is stored, the writer can't
	// proceed (and clear the cache) until after that store lands. That
	// ordering guarantees a registration can never be immediately
	// followed by a stale, pre-registration entry silently reappearing.
	makersMu.RLock()
	defer makersMu.RUnlock()

	remaining, omitEmpty := extractOmitEmpty(config)

	compiled := &compiledConfig{
		omitEmpty: omitEmpty,
		checks:    compileChecksLocked(remaining),
	}

	actual, _ := configCache.LoadOrStore(config, compiled)

	return actual.(*compiledConfig)
}

// compileChecksLocked resolves a checkers tag config (with any omitempty
// token already removed) into compiled checks. The caller must hold at
// least makersMu.RLock() for the duration of the call.
func compileChecksLocked(config string) []compiledCheck {
	fields := strings.Fields(config)

	checks := make([]compiledCheck, len(fields))

	for i, field := range fields {
		name, params, _ := strings.Cut(field, ":")

		if fieldMaker, ok := fieldMakers[name]; ok {
			checks[i] = compiledCheck{fieldFn: fieldMaker(params), name: name}
			continue
		}

		maker, ok := makers[name]
		if !ok {
			panic(fmt.Sprintf("check %s not found", name))
		}

		checks[i] = compiledCheck{fn: maker(params), name: name}
	}

	return checks
}

// runCompiledChecks runs compiled checks against value in order, binding
// parent for any field-relative check, and short-circuits on the first
// error, mirroring Check's semantics. If the failing check's name has an
// entry in messages, the *CheckError is replaced with a copy carrying that
// message (see CheckError.WithMessage), so it overrides the locale-based
// message for this specific field. messages may be nil.
func runCompiledChecks(value, parent reflect.Value, checks []compiledCheck, messages map[string]string) (reflect.Value, error) {
	var err error

	for _, check := range checks {
		value, err = check.run(parent, value)
		if err != nil {
			if message, ok := messages[check.name]; ok {
				if checkErr, ok := err.(*CheckError); ok {
					err = checkErr.WithMessage(message)
				}
			}

			break
		}
	}

	return value, err
}
