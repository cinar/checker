// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

import (
	"net/url"
	"reflect"
)

const (
	// nameURLEscape is the name of the URL escape normalizer.
	nameURLEscape = "url-escape"
)

// normalizeURLEscape applies URL escaping to special characters.
// Uses net.url.QueryEscape for the actual escape operation.
func normalizeURLEscape(value string) (string, error) {
	return url.QueryEscape(value), nil
}

// checkURLEscape checks if the value is a valid URL escape string.
func checkURLEscape(value reflect.Value) (reflect.Value, error) {
	escaped, err := normalizeURLEscape(value.Interface().(string))
	if err != nil {
		return value, err
	}
	value.SetString(escaped)
	return value, nil
}

// makeURLEscape makes a normalizer function for the URL escape normalizer.
func makeURLEscape(_ string) CheckFunc[reflect.Value] {
	return checkURLEscape
}
