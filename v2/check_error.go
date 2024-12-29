// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

import (
	"html/template"
	"strings"

	"github.com/cinar/checker/v2/locales"
)

// CheckError defines the check error.
type CheckError struct {
	// Code is the error code.
	Code string

	// data is the error data.
	Data map[string]interface{}
}

const (
	// DefaultLocale is the default locale.
	DefaultLocale = locales.EnUS
)

// errorMessages is the map of localized error messages.
var errorMessages = map[string]map[string]string{
	locales.EnUS: locales.EnUsMessages,
}

// NewCheckError creates a new check error with the given code.
func NewCheckError(code string) *CheckError {
	return NewCheckErrorWithData(
		code,
		make(map[string]interface{}),
	)
}

// NewCheckErrorWithData creates a new check error with the given code and data.
func NewCheckErrorWithData(code string, data map[string]interface{}) *CheckError {
	return &CheckError{
		Code: code,
		Data: data,
	}
}

// Error returns the error message for the check.
func (c *CheckError) Error() string {
	return c.ErrorWithLocale(DefaultLocale)
}

// Is reports whether the check error is the same as the target error.
func (c *CheckError) Is(target error) bool {
	if other, ok := target.(*CheckError); ok {
		return c.Code == other.Code
	}

	return false
}

// ErrorWithLocale returns the localized error message for the check with the given locale.
func (c *CheckError) ErrorWithLocale(locale string) string {
	tmpl, err := template.New("error").Parse(getLocalizedErrorMessage(locale, c.Code))
	if err != nil {
		return c.Code
	}

	var message strings.Builder
	if err := tmpl.Execute(&message, c.Data); err != nil {
		return c.Code
	}

	return message.String()
}

// RegisterLocale registers the localized error messages for the given locale.
func RegisterLocale(locale string, messages map[string]string) {
	errorMessages[locale] = messages
}

// getLocalizedErrorMessage returns the localized error message for the given locale and code.
func getLocalizedErrorMessage(locale, code string) string {
	if messages, found := errorMessages[locale]; found {
		if message, exists := messages[code]; exists {
			return message
		}
	}

	if messages, found := errorMessages[DefaultLocale]; found {
		if message, exists := messages[code]; exists {
			return message
		}
	}

	return code
}
