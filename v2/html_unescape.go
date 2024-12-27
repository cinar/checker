// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

import (
	"html"
	"reflect"
)

// nameHTMLUnescape is the name of the HTML unescape normalizer.
const nameHTMLUnescape = "html-unescape"

// HTMLUnescape applies HTML unescaping to special characters.
func HTMLUnescape(value string) (string, error) {
	return html.UnescapeString(value), nil
}

// reflectHTMLUnescape applies HTML unescaping to special characters.
func reflectHTMLUnescape(value reflect.Value) (reflect.Value, error) {
	newValue, err := HTMLUnescape(value.Interface().(string))
	return reflect.ValueOf(newValue), err
}

// makeHTMLUnescape returns the HTML unescape normalizer function.
func makeHTMLUnescape(_ string) CheckFunc[reflect.Value] {
	return reflectHTMLUnescape
}
