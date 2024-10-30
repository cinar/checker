// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package checker

import (
	"net/url"
	"reflect"
)

// tagURLEscape is the tag of the normalizer.
const tagURLEscape = "url-escape"

// makeURLEscape makes a normalizer function for the URL escape normalizer.
func makeURLEscape(_ string) CheckFunc {
	return normalizeURLEscape
}

// normalizeURLEscape applies URL escaping to special characters.
// Uses net.url.QueryEscape for the actual escape operation.
func normalizeURLEscape(value, _ reflect.Value) error {
	if value.Kind() != reflect.String {
		panic("string expected")
	}

	value.SetString(url.QueryEscape(value.String()))

	return nil
}
