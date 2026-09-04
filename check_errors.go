// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

import (
	"encoding/json"
	"sort"
	"strings"
)

// CheckErrors is a map of field names to their corresponding check errors,
// as returned by CheckStruct.
type CheckErrors map[string]error

// FieldError is the JSON representation of a single field's check error. Code
// is the machine-readable check error code, and Message is the localized
// human-readable error message.
type FieldError struct {
	// Code is the check error code, such as "REQUIRED". It is empty when the
	// field's error is not a *CheckError.
	Code string `json:"code"`

	// Message is the localized human-readable error message.
	Message string `json:"message"`
}

// Error joins the errors of all the fields into a single human-readable message,
// sorted by field name for a deterministic result. It allows CheckErrors to be
// used and propagated as a regular error.
func (errs CheckErrors) Error() string {
	names := make([]string, 0, len(errs))
	for name := range errs {
		names = append(names, name)
	}

	sort.Strings(names)

	var message strings.Builder

	for i, name := range names {
		if i > 0 {
			message.WriteString("; ")
		}

		message.WriteString(name)
		message.WriteString(": ")
		message.WriteString(errs[name].Error())
	}

	return message.String()
}

// JSON marshals the errors as a JSON object of field name to FieldError,
// localized using the default locale. It is suitable for use directly as
// an HTTP API error response body.
func (errs CheckErrors) JSON() ([]byte, error) {
	return errs.JSONWithLocale(DefaultLocale)
}

// JSONWithLocale marshals the errors as a JSON object of field name to
// FieldError, localized using the given locale.
func (errs CheckErrors) JSONWithLocale(locale string) ([]byte, error) {
	fields := make(map[string]FieldError, len(errs))

	for name, err := range errs {
		fields[name] = newFieldError(err, locale)
	}

	return json.Marshal(fields)
}

// newFieldError converts an error into its JSON field error representation.
// An error that is not a *CheckError only carries a message, since it has no code.
func newFieldError(err error, locale string) FieldError {
	checkErr, ok := err.(*CheckError)
	if !ok {
		return FieldError{Message: err.Error()}
	}

	return FieldError{
		Code:    checkErr.Code,
		Message: checkErr.ErrorWithLocale(locale),
	}
}
