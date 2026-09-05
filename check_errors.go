// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

import (
	"encoding/json"
	"net/http"
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

// sortedNames returns the field names of errs, sorted for a deterministic
// iteration order.
func (errs CheckErrors) sortedNames() []string {
	names := make([]string, 0, len(errs))
	for name := range errs {
		names = append(names, name)
	}

	sort.Strings(names)

	return names
}

// Error joins the errors of all the fields into a single human-readable message,
// sorted by field name for a deterministic result. It allows CheckErrors to be
// used and propagated as a regular error.
func (errs CheckErrors) Error() string {
	names := errs.sortedNames()

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

// ProblemDetail is the RFC 9457 (Problem Details for HTTP APIs)
// representation of a single field's check error, listed under a
// ProblemDetails' InvalidParams.
type ProblemDetail struct {
	// Name is the field name.
	Name string `json:"name"`

	// Reason is the localized human-readable error message.
	Reason string `json:"reason"`

	// Code is the check error code, such as "REQUIRED". It is empty when
	// the field's error is not a *CheckError.
	Code string `json:"code,omitempty"`
}

// ProblemDetails is the RFC 9457 (Problem Details for HTTP APIs)
// representation of a CheckErrors value, suitable for use as the body of
// an "application/problem+json" HTTP response. RFC 9457 leaves Type,
// Title, and Status to the API producer; ProblemDetails/
// ProblemDetailsWithLocale fill in reasonable defaults, and both are
// plain exported fields so callers can override them directly on the
// returned value before serializing it.
type ProblemDetails struct {
	// Type is a URI reference identifying the problem type. Defaults to
	// "about:blank", the RFC 9457 default for a problem with no more
	// specific type of its own.
	Type string `json:"type"`

	// Title is a short, human-readable summary of the problem type.
	Title string `json:"title"`

	// Status is the HTTP status code for this occurrence of the problem.
	// Defaults to 400 Bad Request.
	Status int `json:"status"`

	// InvalidParams lists each failing field, sorted by name for a
	// deterministic result.
	InvalidParams []ProblemDetail `json:"invalid-params"`
}

// ProblemDetails builds an RFC 9457 Problem Details value from errs,
// localized using the default locale.
func (errs CheckErrors) ProblemDetails() *ProblemDetails {
	return errs.ProblemDetailsWithLocale(DefaultLocale)
}

// ProblemDetailsWithLocale builds an RFC 9457 Problem Details value from
// errs, localized using the given locale.
func (errs CheckErrors) ProblemDetailsWithLocale(locale string) *ProblemDetails {
	names := errs.sortedNames()

	invalidParams := make([]ProblemDetail, 0, len(names))

	for _, name := range names {
		field := newFieldError(errs[name], locale)

		invalidParams = append(invalidParams, ProblemDetail{
			Name:   name,
			Reason: field.Message,
			Code:   field.Code,
		})
	}

	return &ProblemDetails{
		Type:          "about:blank",
		Title:         "Your request parameters failed validation.",
		Status:        http.StatusBadRequest,
		InvalidParams: invalidParams,
	}
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
