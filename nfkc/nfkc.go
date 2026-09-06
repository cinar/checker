// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

// Package nfkc registers an "nfkc" checkers-tag normalizer applying Unicode
// Normalization Form KC (NFKC): it composes combining character sequences
// and replaces compatibility characters -- many homoglyphs and stylistic
// variants a naive keyword filter, uniqueness check, or handle comparison
// would otherwise treat as distinct -- with their canonical equivalents.
//
// Importing this package, even with a blank import, registers the
// normalizer through checker.RegisterMaker in its init function, so
// checkers:"nfkc" becomes usable without any further setup:
//
//	import (
//		checker "github.com/cinar/checker/v2"
//		_ "github.com/cinar/checker/v2/nfkc"
//	)
//
// It's a separate module, not part of the core checker package, because
// NFKC normalization needs golang.org/x/text/unicode/norm; keeping it
// isolated means the core module's zero-dependency promise holds regardless
// of whether a caller opts into this normalizer.
package nfkc

import (
	"reflect"

	"golang.org/x/text/unicode/norm"

	checker "github.com/cinar/checker/v2"
)

// name is the checkers tag name this package registers.
const name = "nfkc"

func init() {
	checker.RegisterMaker(name, makeNFKC)

	// A no-op SchemaMakeFunc is the only public way to tell JSONSchema this
	// is a normalizer: the core module's own ignoredForSchema set (used for
	// its built-in normalizers) is unexported, so an external package has
	// no way to add to it directly. Registering one here, instead of
	// leaving "nfkc" to land in the generated Schema's XChecker list, keeps
	// this normalizer's behavior consistent with every built-in one.
	checker.RegisterSchemaMaker(name, func(_ *checker.Schema, _ string) {})
}

// Normalize returns value converted to Unicode Normalization Form KC
// (NFKC).
func Normalize(value string) (string, error) {
	return norm.NFKC.String(value), nil
}

// reflectNormalize returns value converted to NFKC, preserving value's
// underlying defined string type (e.g. `type Handle string`).
func reflectNormalize(value reflect.Value) (reflect.Value, error) {
	newValue, err := Normalize(value.String())
	return reflect.ValueOf(newValue).Convert(value.Type()), err
}

// makeNFKC returns the nfkc normalizer function.
func makeNFKC(_ string) checker.CheckFunc[reflect.Value] {
	return reflectNormalize
}
