// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

import (
	"html"
	"reflect"
)

const (
	// nameHTMLEscape is the name of the HTML escape normalizer.
	nameHTMLEscape = "html-escape"
)

// HTMLEscape applies HTML escaping to special characters.
func HTMLEscape(value string) (string, error) {
	return html.EscapeString(value), nil
}

// reflectHTMLEscape applies HTML escaping to special characters.
func reflectHTMLEscape(value reflect.Value) (reflect.Value, error) {
	newValue, err := HTMLEscape(value.Interface().(string))
	return reflect.ValueOf(newValue), err
}

// makeHTMLEscape returns the HTML escape normalizer function.
func makeHTMLEscape(_ string) CheckFunc[reflect.Value] {
	return reflectHTMLEscape
}
