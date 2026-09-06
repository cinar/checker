// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

import (
	"reflect"
	"strings"
)

const (
	// nameStripInvisible is the name of the strip-invisible normalizer.
	nameStripInvisible = "strip-invisible"
)

// invisibleRunes is the set of zero-width and bidirectional control
// characters StripInvisible removes: zero-width space, zero-width
// non-joiner, zero-width joiner, word joiner, and the zero-width no-break
// space (BOM), plus the bidirectional embedding/override/isolate controls
// behind the "Trojan Source" spoofing technique (CVE-2021-42574). An
// attacker can use these to split a keyword-filtered word invisibly, or to
// make displayed text diverge from its logical character order.
//
// Some of these runes have legitimate uses in ordinary text -- zero-width
// joiner in emoji sequences, zero-width non-joiner in Persian and other
// scripts -- so apply this normalizer only to fields where an invisible
// character is never expected, such as a handle, username, or search
// keyword, not general free-text content.
//
// Every key is written as a \u escape, not a literal rune, so the
// characters this file removes stay visible in source and diffs instead of
// disappearing into them.
var invisibleRunes = map[rune]bool{
	'\u200b': true, // zero width space
	'\u200c': true, // zero width non-joiner
	'\u200d': true, // zero width joiner
	'\u2060': true, // word joiner
	'\ufeff': true, // zero width no-break space (BOM)
	'\u202a': true, // left-to-right embedding
	'\u202b': true, // right-to-left embedding
	'\u202c': true, // pop directional formatting
	'\u202d': true, // left-to-right override
	'\u202e': true, // right-to-left override
	'\u2066': true, // left-to-right isolate
	'\u2067': true, // right-to-left isolate
	'\u2068': true, // first strong isolate
	'\u2069': true, // pop directional isolate
}

// StripInvisible returns the value with zero-width and bidirectional
// control characters removed.
func StripInvisible(value string) (string, error) {
	return strings.Map(func(r rune) rune {
		if invisibleRunes[r] {
			return -1
		}

		return r
	}, value), nil
}

// reflectStripInvisible returns the value with zero-width and
// bidirectional control characters removed.
func reflectStripInvisible(value reflect.Value) (reflect.Value, error) {
	newValue, err := StripInvisible(reflectString(value))
	return reflect.ValueOf(newValue).Convert(value.Type()), err
}

// makeStripInvisible returns the strip-invisible normalizer function.
func makeStripInvisible(_ string) CheckFunc[reflect.Value] {
	return reflectStripInvisible
}
