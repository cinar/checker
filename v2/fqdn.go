// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

import (
    "reflect"
    "regexp"
)

const (
    // nameFQDN is the name of the FQDN check.
    nameFQDN = "fqdn"
)

var (
    // ErrNotFQDN indicates that the given value is not a valid FQDN.
    ErrNotFQDN = NewCheckError("FQDN")

    // fqdnRegex is the regular expression for validating FQDN.
    fqdnRegex = regexp.MustCompile(`^(?:[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z]{2,}$`)
)

// IsFQDN checks if the value is a valid fully qualified domain name (FQDN).
func IsFQDN(value string) (string, error) {
    if !fqdnRegex.MatchString(value) {
        return value, ErrNotFQDN
    }
    return value, nil
}

// checkFQDN checks if the value is a valid fully qualified domain name (FQDN).
func checkFQDN(value reflect.Value) (reflect.Value, error) {
    _, err := IsFQDN(value.Interface().(string))
    return value, err
}

// makeFQDN makes a checker function for the FQDN checker.
func makeFQDN(_ string) CheckFunc[reflect.Value] {
    return checkFQDN
}