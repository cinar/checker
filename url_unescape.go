// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package checker

import (
	"net/url"
	"reflect"
)

// NormalizerURLUnescape is the name of the normalizer.
const NormalizerURLUnescape = "url-unescape"

// makeURLUnescape makes a normalizer function for the URL unscape normalizer.
func makeURLUnescape(_ string) CheckFunc {
	return normalizeURLUnescape
}

// normalizeURLUnescape applies URL unescaping to special characters.
// Uses url.QueryUnescape for the actual unescape operation.
func normalizeURLUnescape(value, _ reflect.Value) error {
	if value.Kind() != reflect.String {
		panic("string expected")
	}

	unescaped, err := url.QueryUnescape(value.String())
	if err == nil {
		value.SetString(unescaped)
	}

	return nil
}
