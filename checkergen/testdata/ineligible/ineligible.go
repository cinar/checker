// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package ineligible

// Eligible has a valid, generatable field, to prove one struct being
// skipped doesn't prevent generation for others in the same package.
type Eligible struct {
	Name string `checkers:"required"`
}

// UnknownChecker has a checkers tag naming a checker checkergen has no
// mapping for.
type UnknownChecker struct {
	Name string `checkers:"is-fruit"`
}

// NamedType has a tagged field whose Go type is a named string type, not
// the plain string type itself -- outside v1's scope (see the checkergen
// package doc).
type NamedType struct {
	Email EmailAddress `checkers:"required"`
}

// EmailAddress is a named string type, used only to make NamedType's field
// ineligible for generation.
type EmailAddress string

// SliceField has a tagged slice field -- outside v1's scope.
type SliceField struct {
	Roles []string `checkers:"required"`
}

// BadIntParam has a curried checker (min-len) with a non-integer parameter.
type BadIntParam struct {
	Name string `checkers:"min-len:abc"`
}

// BadNumberParam has a numeric checker (gte) with a non-numeric parameter.
type BadNumberParam struct {
	Age int `checkers:"gte:abc"`
}

// MalformedTwoPart has a checker (after) requiring two ":"-separated
// parameters, given only one -- this panics at generate time, the same way
// the equivalent runtime maker panics on a malformed tag, recovered into a
// skip rather than crashing generation for the rest of the package.
type MalformedTwoPart struct {
	BornAt string `checkers:"after:DateOnly"`
}

// WrongNumericType has "multiple-of" (hardcoded to float64) applied to an
// int field.
type WrongNumericType struct {
	Weight int `checkers:"multiple-of:5"`
}

// MissingSibling has a field-relative checker referencing a sibling field
// that doesn't exist on the struct.
type MissingSibling struct {
	ConfirmPassword string `checkers:"eq-field:Password"`
}

// WithEmbedded has an embedded field (out of scope for v1, silently
// skipped rather than treated as ineligible, since it has no tag of its
// own) alongside a normal eligible field, to prove the embedded field
// doesn't block generation for the rest of the struct.
type WithEmbedded struct {
	Eligible
	Name string `checkers:"required"`
}
