// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

import (
	"net/url"
	"reflect"
)

// nameURLUnescape is the name of the URL unescape normalizer.
const nameURLUnescape = "url-unescape"

// URLUnescape applies URL unescaping to special characters.
func URLUnescape(value string) (string, error) {
	unescaped, err := url.QueryUnescape(value)
	return unescaped, err
}

// reflectURLUnescape applies URL unescaping to special characters.
func reflectURLUnescape(value reflect.Value) (reflect.Value, error) {
	newValue, err := URLUnescape(value.Interface().(string))
	return reflect.ValueOf(newValue), err
}

// makeURLUnescape returns the URL unescape normalizer function.
func makeURLUnescape(_ string) CheckFunc[reflect.Value] {
	return reflectURLUnescape
}
