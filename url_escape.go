// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

import (
	"net/url"
	"reflect"
)

// nameURLEscape is the name of the URL escape normalizer.
const nameURLEscape = "url-escape"

// URLEscape applies URL escaping to special characters.
func URLEscape(value string) (string, error) {
	return url.QueryEscape(value), nil
}

// reflectURLEscape applies URL escaping to special characters.
func reflectURLEscape(value reflect.Value) (reflect.Value, error) {
	newValue, err := URLEscape(value.Interface().(string))
	return reflect.ValueOf(newValue), err
}

// makeURLEscape returns the URL escape normalizer function.
func makeURLEscape(_ string) CheckFunc[reflect.Value] {
	return reflectURLEscape
}
