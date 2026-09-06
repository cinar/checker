// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

import (
	"net/url"
	"reflect"
)

const (
	// nameURL is the name of the URL check.
	nameURL = "url"
)

var (
	// ErrNotURL indicates that the given value is not a valid URL.
	ErrNotURL = NewCheckError("NOT_URL")
)

// IsURL checks if the value is a valid absolute HTTP or HTTPS URL with a
// host, e.g. "https://example.com". Relative paths, opaque URIs (e.g.
// "mailto:a@b.com"), and other schemes are rejected.
func IsURL(value string) (string, error) {
	u, err := url.ParseRequestURI(value)
	if err != nil || (u.Scheme != "http" && u.Scheme != "https") || u.Host == "" {
		return value, ErrNotURL
	}
	return value, nil
}

// checkURL checks if the value is a valid URL.
func checkURL(value reflect.Value) (reflect.Value, error) {
	_, err := IsURL(reflectString(value))
	return value, err
}

// makeURL makes a checker function for the URL checker.
func makeURL(_ string) CheckFunc[reflect.Value] {
	return checkURL
}
