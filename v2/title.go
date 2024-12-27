// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

import (
	"reflect"
	"strings"
	"unicode"
)

const (
	// nameTitle is the name of the title normalizer.
	nameTitle = "title"
)

// Title returns the value of the string with the first letter of each word in upper case.
func Title(value string) (string, error) {
	var sb strings.Builder
	begin := true

	for _, c := range value {
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

	return sb.String(), nil
}

// reflectTitle returns the value of the string with the first letter of each word in upper case.
func reflectTitle(value reflect.Value) (reflect.Value, error) {
	newValue, err := Title(value.Interface().(string))
	return reflect.ValueOf(newValue), err
}

// makeTitle returns the title normalizer function.
func makeTitle(_ string) CheckFunc[reflect.Value] {
	return reflectTitle
}
