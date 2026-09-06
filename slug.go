// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

import (
	"reflect"
)

const (
	// nameSlug is the name of the slug check.
	nameSlug = "slug"

	// slugExpression matches a URL slug: lowercase letters and digits,
	// with single hyphens as separators, no leading, trailing, or
	// consecutive hyphens.
	slugExpression = `^[a-z0-9]+(?:-[a-z0-9]+)*$`
)

var (
	// ErrNotSlug indicates that the given string is not a valid slug.
	ErrNotSlug = NewCheckError("NOT_SLUG")
)

// IsSlug checks if the given string is a valid URL slug, e.g. "hello-world".
func IsSlug(value string) (string, error) {
	if _, err := IsRegexp(slugExpression, value); err != nil {
		return value, ErrNotSlug
	}

	return value, nil
}

// checkSlug checks if the given string is a valid slug.
func checkSlug(value reflect.Value) (reflect.Value, error) {
	_, err := IsSlug(reflectString(value))
	return value, err
}

// makeSlug makes a checker function for the slug checker.
func makeSlug(_ string) CheckFunc[reflect.Value] {
	return checkSlug
}
