// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package checker

import (
	"reflect"
	"strings"
	"unicode"
)

// NormalizerTitle is the name of the normalizer.
const NormalizerTitle = "title"

// makeTitle makes a normalizer function for the title normalizer.
func makeTitle(_ string) CheckFunc {
	return normalizeTitle
}

// normalizeTitle maps the first letter of each word to their upper case.
func normalizeTitle(value, _ reflect.Value) Result {
	if value.Kind() != reflect.String {
		panic("string expected")
	}

	var sb strings.Builder
	begin := true

	for _, c := range value.String() {
		if unicode.IsLetter(c) {
			if begin {
				c = unicode.ToUpper(c)
				begin = false
			} else {
				c = unicode.ToLower(c)
			}
		} else {
			begin = true
		}

		sb.WriteRune(c)
	}

	value.SetString(sb.String())

	return ResultValid
}
