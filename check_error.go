// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

import (
	"strings"
	"sync"
	"text/template"

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

// errorMessagesMu guards errorMessages, which can be written concurrently
// with error rendering through RegisterLocale (e.g. registered during
// server startup while other goroutines are already rendering errors).
var errorMessagesMu sync.RWMutex

// errorMessages is the map of localized error messages.
var errorMessages = map[string]map[string]string{
	locales.EnUS: locales.EnUSMessages,
}

// errorTemplateCache holds parsed *template.Template values keyed by their
// source message string, so a given message is only ever parsed once. Like
// regexpCache, the key space is bounded by the (locale, code) message pairs
// baked into the code, not by request data.
var errorTemplateCache sync.Map

// compileErrorTemplate returns the parsed *template.Template for message,
// parsing and caching it on first use.
func compileErrorTemplate(message string) (*template.Template, error) {
	if cached, ok := errorTemplateCache.Load(message); ok {
		return cached.(*template.Template), nil
	}

	tmpl, err := template.New("error").Parse(message)
	if err != nil {
		return nil, err
	}

	cached, _ := errorTemplateCache.LoadOrStore(message, tmpl)

	return cached.(*template.Template), nil
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
	message := getLocalizedErrorMessage(locale, c.Code)

	// Fast path: most messages ("Required value is missing.", "Not a valid
	// email address.", ...) have no {{ }} placeholders at all, so they need
	// no template parsing or execution.
	if !strings.Contains(message, "{{") {
		return message
	}

	tmpl, err := compileErrorTemplate(message)
	if err != nil {
		return c.Code
	}

	var rendered strings.Builder
	if err := tmpl.Execute(&rendered, c.Data); err != nil {
		return c.Code
	}

	return rendered.String()
}

// RegisterLocale registers the localized error messages for the given locale.
func RegisterLocale(locale string, messages map[string]string) {
	errorMessagesMu.Lock()
	defer errorMessagesMu.Unlock()

	errorMessages[locale] = messages
}

// getLocalizedErrorMessage returns the localized error message for the given locale and code.
func getLocalizedErrorMessage(locale, code string) string {
	errorMessagesMu.RLock()
	defer errorMessagesMu.RUnlock()

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
