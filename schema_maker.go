// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

import (
	"strconv"
	"strings"
	"sync"
)

// SchemaMakeFunc refines a Schema for the given checker's parameter. It is
// called with the Schema of the value the checker is attached to: the field
// itself for a scalar, or the array/map Schema for an "@"-prefixed check.
type SchemaMakeFunc func(schema *Schema, params string)

// ignoredForSchema is the set of normalizer and pipeline-modifier names that
// don't constrain a Schema, so they're silently skipped rather than
// recorded in XChecker.
var ignoredForSchema = map[string]bool{
	nameHTMLEscape:   true,
	nameHTMLUnescape: true,
	nameLower:        true,
	nameOmitEmpty:    true,
	nameTitle:        true,
	nameTrimLeft:     true,
	nameTrimRight:    true,
	nameTrimSpace:    true,
	nameUpper:        true,
	nameURLEscape:    true,
	nameURLUnescape:  true,
}

// schemaMakersMu guards schemaMakers, which can be written concurrently with
// JSONSchema reads through RegisterSchemaMaker (e.g. registered during
// server startup while other goroutines are already generating schemas).
var schemaMakersMu sync.RWMutex

// schemaMakers provides a mapping of SchemaMakeFunc keyed by checker name,
// for checkers with a direct JSON Schema equivalent. A checker with no entry
// here is instead recorded in the Schema's XChecker list.
var schemaMakers = map[string]SchemaMakeFunc{
	nameEmail:  schemaFormat("email"),
	nameEq:     schemaConst,
	nameFQDN:   schemaFormat("hostname"),
	nameGt:     schemaExclusiveMinimum,
	nameGte:    schemaMinimum,
	nameIPv4:   schemaFormat("ipv4"),
	nameIPv6:   schemaFormat("ipv6"),
	nameLt:     schemaExclusiveMaximum,
	nameLte:    schemaMaximum,
	nameMaxLen: schemaMaxLen,
	nameMinLen: schemaMinLen,
	nameOneOf:  schemaOneOf,
	nameRegexp: schemaPattern,
	nameURL:    schemaFormat("uri"),
	nameUUID:   schemaFormat("uuid"),
}

// RegisterSchemaMaker registers a SchemaMakeFunc for the given checker or
// normalizer name, so JSONSchema can translate it into a JSON Schema
// keyword instead of recording it in XChecker.
func RegisterSchemaMaker(name string, maker SchemaMakeFunc) {
	schemaMakersMu.Lock()
	defer schemaMakersMu.Unlock()

	schemaMakers[name] = maker
}

// schemaOneOf sets the Schema's Enum from the checker's comma-separated
// list of allowed values.
func schemaOneOf(schema *Schema, params string) {
	schema.Enum = strings.Split(params, ",")
}

// schemaConst sets the Schema's Const to the checker's expected value.
func schemaConst(schema *Schema, params string) {
	schema.Const = &params
}

// schemaFormat returns a SchemaMakeFunc that sets the Schema's Format.
func schemaFormat(format string) SchemaMakeFunc {
	return func(schema *Schema, _ string) {
		schema.Format = format
	}
}

// schemaPattern sets the Schema's Pattern to the checker's regular expression parameter.
func schemaPattern(schema *Schema, params string) {
	schema.Pattern = params
}

// schemaMinimum sets the Schema's Minimum from the checker's numeric parameter.
// Panics if the parameter cannot be parsed as a number.
func schemaMinimum(schema *Schema, params string) {
	n, err := strconv.ParseFloat(params, 64)
	if err != nil {
		panic("unable to parse params as float")
	}

	schema.Minimum = &n
}

// schemaMaximum sets the Schema's Maximum from the checker's numeric parameter.
// Panics if the parameter cannot be parsed as a number.
func schemaMaximum(schema *Schema, params string) {
	n, err := strconv.ParseFloat(params, 64)
	if err != nil {
		panic("unable to parse params as float")
	}

	schema.Maximum = &n
}

// schemaExclusiveMinimum sets the Schema's ExclusiveMinimum from the
// checker's numeric parameter. Panics if the parameter cannot be parsed as
// a number.
func schemaExclusiveMinimum(schema *Schema, params string) {
	n, err := strconv.ParseFloat(params, 64)
	if err != nil {
		panic("unable to parse params as float")
	}

	schema.ExclusiveMinimum = &n
}

// schemaExclusiveMaximum sets the Schema's ExclusiveMaximum from the
// checker's numeric parameter. Panics if the parameter cannot be parsed as
// a number.
func schemaExclusiveMaximum(schema *Schema, params string) {
	n, err := strconv.ParseFloat(params, 64)
	if err != nil {
		panic("unable to parse params as float")
	}

	schema.ExclusiveMaximum = &n
}

// schemaMinLen sets the Schema's MinLength, MinItems, or MinProperties,
// depending on whether it describes a string, array, or object, from the
// checker's integer parameter. Panics if the parameter cannot be parsed as
// an integer.
func schemaMinLen(schema *Schema, params string) {
	n, err := strconv.Atoi(params)
	if err != nil {
		panic("unable to parse min length")
	}

	switch schema.Type {
	case "array":
		schema.MinItems = &n
	case "object":
		schema.MinProperties = &n
	default:
		schema.MinLength = &n
	}
}

// schemaMaxLen sets the Schema's MaxLength, MaxItems, or MaxProperties,
// depending on whether it describes a string, array, or object, from the
// checker's integer parameter. Panics if the parameter cannot be parsed as
// an integer.
func schemaMaxLen(schema *Schema, params string) {
	n, err := strconv.Atoi(params)
	if err != nil {
		panic("unable to parse max length")
	}

	switch schema.Type {
	case "array":
		schema.MaxItems = &n
	case "object":
		schema.MaxProperties = &n
	default:
		schema.MaxLength = &n
	}
}
