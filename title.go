// Package checker is a Go library for validating user input through struct tags.
//
// https://github.com/cinar/checker
//
// Copyright 2023 Onur Cinar. All rights reserved. 
// Use of this source code is governed by a MIT-style 
// license that can be found in the LICENSE file. 
//
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
