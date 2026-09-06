// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package fixture

// Coverage exercises the remaining checkergen-eligible checkers not
// already covered by SignupRequest: two-parameter checkers, "value after
// param" checkers, oneof, default, and the remaining field-relative
// checkers (required-if, required-unless, before-field, after-field).
type Coverage struct {
	Handle    string  `checkers:"contains:@ starts-with:@ ends-with:.com"`
	Code      string  `checkers:"regexp:^[A-Z]{3}$"`
	Sum       string  `checkers:"hash:sha256"`
	Zip       string  `checkers:"postal-code:US"`
	Status    string  `checkers:"eq:active"`
	Excluded  string  `checkers:"ne:banned"`
	Role      string  `checkers:"oneof:admin,editor,viewer"`
	Weight    float64 `checkers:"multiple-of:5"`
	Ratio     float64 `checkers:"finite int"`
	Greeting  string  `checkers:"default:hello"`
	Threshold int     `checkers:"default:10"`

	Country string `checkers:"required"`
	State   string `checkers:"required-if:Country:US"`
	Type    string `checkers:"required"`
	Email   string `checkers:"required-unless:Type:guest"`

	ReturnAt string `checkers:"required"`
	DepartAt string `checkers:"before-field:DateOnly:ReturnAt"`
	BornAt   string `checkers:"required"`
	DiesAt   string `checkers:"after-field:DateOnly:BornAt"`

	// Untouched has no checkers/validate tag at all: findStructs still
	// walks it (it's a field on an eligible struct), but there's nothing
	// to generate for it.
	Untouched string
}
