// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package checker

import (
	"html"
	"reflect"
)

// tagHTMLEscape is the tag of the normalizer.
const tagHTMLEscape = "html-escape"

// makeHTMLEscape makes a normalizer function for the HTML escape normalizer.
func makeHTMLEscape(_ string) CheckFunc {
	return normalizeHTMLEscape
}

// normalizeHTMLEscape applies HTML escaping to special characters.
// Uses html.EscapeString for the actual escape operation.
func normalizeHTMLEscape(value, _ reflect.Value) error {
	if value.Kind() != reflect.String {
		panic("string expected")
	}

	value.SetString(html.EscapeString(value.String()))

	return nil
}
