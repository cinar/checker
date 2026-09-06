// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

import (
	"fmt"
	"regexp"
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
	// JSON numbers can't be NaN or an infinity by spec, so "finite" has no
	// keyword to emit: it's a pure in-memory Go-value guard, not a shape
	// constraint on data that ever arrives as JSON.
	nameFinite:       true,
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
	nameAlphanumeric: schemaPatternConst("^[0-9A-Za-z]+$"),
	nameASCII:        schemaPatternConst("^[\x00-\x7f]*$"),
	nameCIDR:         schemaFormat("cidr"),
	nameContains:     schemaContains,
	nameDefault:      schemaDefault,
	nameDigits:       schemaPatternConst("^[0-9]+$"),
	nameEmail:        schemaFormat("email"),
	nameEndsWith:     schemaEndsWith,
	nameEq:           schemaConst,
	nameFQDN:         schemaFormat("hostname"),
	nameGt:           schemaExclusiveMinimum,
	nameGte:          schemaMinimum,
	nameHash:         schemaHash,
	nameHex:          schemaPatternConst("^[0-9a-fA-F]+$"),
	nameInt:          schemaInt,
	nameIP:           schemaFormat("ip"),
	nameIPv4:         schemaFormat("ipv4"),
	nameIPv6:         schemaFormat("ipv6"),
	nameLen:          schemaLen,
	nameLt:           schemaExclusiveMaximum,
	nameLte:          schemaMaximum,
	nameMAC:          schemaFormat("mac"),
	nameMaxLen:       schemaMaxLen,
	nameMinLen:       schemaMinLen,
	nameMultipleOf:   schemaMultipleOf,
	nameNegative:     schemaNegative,
	nameNonnegative:  schemaNonnegative,
	nameOneOf:        schemaOneOf,
	namePositive:     schemaPositive,
	namePostalCode:   schemaPostalCode,
	nameRegexp:       schemaPattern,
	nameStartsWith:   schemaStartsWith,
	nameURL:          schemaFormat("uri"),
	nameUUID:         schemaFormat("uuid"),
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

// schemaDefault sets the Schema's Default to the normalizer's fallback value.
func schemaDefault(schema *Schema, params string) {
	schema.Default = &params
}

// schemaContains sets the Schema's Pattern so the value must contain the
// checker's substring parameter anywhere in the string.
func schemaContains(schema *Schema, params string) {
	schema.Pattern = ".*" + regexp.QuoteMeta(params) + ".*"
}

// schemaStartsWith sets the Schema's Pattern so the value must start with
// the checker's prefix parameter.
func schemaStartsWith(schema *Schema, params string) {
	schema.Pattern = "^" + regexp.QuoteMeta(params) + ".*"
}

// schemaEndsWith sets the Schema's Pattern so the value must end with the
// checker's suffix parameter.
func schemaEndsWith(schema *Schema, params string) {
	schema.Pattern = ".*" + regexp.QuoteMeta(params) + "$"
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

// schemaPatternConst returns a SchemaMakeFunc that sets the Schema's Pattern
// to a fixed regular expression, for checkers whose shape doesn't depend on
// any parameter.
func schemaPatternConst(pattern string) SchemaMakeFunc {
	return func(schema *Schema, _ string) {
		schema.Pattern = pattern
	}
}

// schemaHash sets the Schema's Pattern to match a hex-encoded digest of the
// exact length expected for the hash checker's algorithm parameter. Panics
// if the algorithm isn't one of the hash checker's supported algorithms,
// same as IsHash.
func schemaHash(schema *Schema, params string) {
	length, ok := hashLengths[params]
	if !ok {
		panic(fmt.Sprintf("unknown hash algorithm %s", params))
	}

	schema.Pattern = fmt.Sprintf("^[0-9a-fA-F]{%d}$", length)
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

// schemaInt sets the Schema's Type to "integer", overriding the "number" a
// float-kind field otherwise gets from jsonTypeForKind.
func schemaInt(schema *Schema, _ string) {
	schema.Type = "integer"
}

// schemaPositive sets the Schema's ExclusiveMinimum to 0.
func schemaPositive(schema *Schema, _ string) {
	zero := 0.0
	schema.ExclusiveMinimum = &zero
}

// schemaNegative sets the Schema's ExclusiveMaximum to 0.
func schemaNegative(schema *Schema, _ string) {
	zero := 0.0
	schema.ExclusiveMaximum = &zero
}

// schemaNonnegative sets the Schema's Minimum to 0.
func schemaNonnegative(schema *Schema, _ string) {
	zero := 0.0
	schema.Minimum = &zero
}

// schemaMultipleOf sets the Schema's MultipleOf from the checker's numeric
// parameter. Panics if the parameter cannot be parsed as a number.
func schemaMultipleOf(schema *Schema, params string) {
	n, err := strconv.ParseFloat(params, 64)
	if err != nil {
		panic("unable to parse params as float")
	}

	schema.MultipleOf = &n
}

// schemaPostalCode sets the Schema's Pattern to the postal code regular
// expression for the checker's country parameter. Panics if the country
// isn't one of the supported codes, same as IsPostalCode.
func schemaPostalCode(schema *Schema, params string) {
	pattern, ok := postalCodePatterns[strings.ToUpper(params)]
	if !ok {
		panic(fmt.Sprintf("unsupported postal code country %s", params))
	}

	schema.Pattern = pattern.String()
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

// schemaLen sets the Schema's Min/Max Length, Items, or Properties (both
// bounds equal, since len is an exact length), depending on whether it
// describes a string, array, or object, from the checker's integer
// parameter. Panics if the parameter cannot be parsed as an integer.
func schemaLen(schema *Schema, params string) {
	n, err := strconv.Atoi(params)
	if err != nil {
		panic("unable to parse len")
	}

	m := n

	switch schema.Type {
	case "array":
		schema.MinItems, schema.MaxItems = &n, &m
	case "object":
		schema.MinProperties, schema.MaxProperties = &n, &m
	default:
		schema.MinLength, schema.MaxLength = &n, &m
	}
}
