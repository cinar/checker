// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

// Package echo integrates Checker with the Echo web framework, binding and
// validating a request in a single call, and writing a JSON 400 response
// automatically when binding or validation fails.
package echo

import (
	"net/http"

	"github.com/labstack/echo/v4"

	checker "github.com/cinar/checker/v2"
)

// Bind binds the request into obj using Echo's default binder, then runs
// Check on it. If binding fails, it writes a 400 JSON response describing
// the bind error to c and returns false. Otherwise, it returns the result
// of Check.
func Bind(c echo.Context, obj any) bool {
	if err := c.Bind(obj); err != nil {
		_ = c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		return false
	}

	return Check(c, obj)
}

// Check runs checker.CheckStruct on obj. If validation fails, it writes a
// 400 JSON response to c, whose body is obj's checker.CheckErrors marshaled
// with CheckErrors.JSON, and returns false. Otherwise, it returns true,
// leaving obj normalized by any checker normalizers for the handler to use.
func Check(c echo.Context, obj any) bool {
	errs, ok := checker.CheckStruct(obj)
	if !ok {
		// json.Marshal cannot fail for FieldError's string fields.
		data, _ := errs.JSON()

		_ = c.JSONBlob(http.StatusBadRequest, data)

		return false
	}

	return true
}
