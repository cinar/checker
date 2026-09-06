// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

const (
	// namePostalCode is the name of the postal code check.
	namePostalCode = "postal-code"
)

var (
	// ErrNotPostalCode indicates that the given value is not a valid postal
	// code for the given country.
	ErrNotPostalCode = NewCheckError("NOT_POSTAL_CODE")

	// postalCodePatterns maps an ISO 3166-1 alpha-2 country code to the
	// regular expression its postal codes match. Coverage isn't exhaustive:
	// it's a curated set of commonly validated countries rather than every
	// country's postal system.
	postalCodePatterns = map[string]*regexp.Regexp{
		"AU": regexp.MustCompile(`^\d{4}$`),
		"BR": regexp.MustCompile(`^\d{5}-?\d{3}$`),
		"CA": regexp.MustCompile(`^[A-Za-z]\d[A-Za-z] ?\d[A-Za-z]\d$`),
		"CH": regexp.MustCompile(`^\d{4}$`),
		"CN": regexp.MustCompile(`^\d{6}$`),
		"DE": regexp.MustCompile(`^\d{5}$`),
		"ES": regexp.MustCompile(`^\d{5}$`),
		"FR": regexp.MustCompile(`^\d{5}$`),
		"GB": regexp.MustCompile(`^[A-Za-z]{1,2}\d[A-Za-z\d]? ?\d[A-Za-z]{2}$`),
		"IN": regexp.MustCompile(`^\d{6}$`),
		"IT": regexp.MustCompile(`^\d{5}$`),
		"JP": regexp.MustCompile(`^\d{3}-?\d{4}$`),
		"KR": regexp.MustCompile(`^\d{5}$`),
		"MX": regexp.MustCompile(`^\d{5}$`),
		"NL": regexp.MustCompile(`^\d{4} ?[A-Za-z]{2}$`),
		"PL": regexp.MustCompile(`^\d{2}-\d{3}$`),
		"PT": regexp.MustCompile(`^\d{4}-\d{3}$`),
		"RU": regexp.MustCompile(`^\d{6}$`),
		"SE": regexp.MustCompile(`^\d{3} ?\d{2}$`),
		"TR": regexp.MustCompile(`^\d{5}$`),
		"US": regexp.MustCompile(`^\d{5}(-\d{4})?$`),
	}
)

// IsPostalCode checks if the value matches the postal code format for the
// given ISO 3166-1 alpha-2 country code (case-insensitive). Panics if
// country isn't one of the supported codes, since that indicates a struct
// tag or call-site typo rather than a data problem — same as IsHash
// panicking on an unknown algorithm.
func IsPostalCode(country, value string) (string, error) {
	pattern, ok := postalCodePatterns[strings.ToUpper(country)]
	if !ok {
		panic(fmt.Sprintf("unsupported postal code country %s", country))
	}

	if !pattern.MatchString(value) {
		return value, newPostalCodeError(country)
	}

	return value, nil
}

// makePostalCode makes a postal code check function for the given country parameter.
func makePostalCode(params string) CheckFunc[reflect.Value] {
	return func(value reflect.Value) (reflect.Value, error) {
		_, err := IsPostalCode(params, reflectString(value))
		return value, err
	}
}

// newPostalCodeError creates a new postal code error for the given country.
func newPostalCodeError(country string) error {
	return NewCheckErrorWithData(
		ErrNotPostalCode.Code,
		map[string]interface{}{
			"country": country,
		},
	)
}
