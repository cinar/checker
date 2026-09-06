// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

// Package fiber integrates Checker with the Fiber web framework, binding and
// validating a request body in a single call, and writing a JSON 400
// response automatically when binding or validation fails.
package fiber

import (
	"github.com/gofiber/fiber/v3"

	checker "github.com/cinar/checker/v2"
)

// jsonContentType is the Content-Type header Check sets on a validation
// failure response; Bind's own failure path uses c.JSON, which sets its own
// Content-Type header.
const jsonContentType = "application/json; charset=utf-8"

// Bind binds the request body into target using Fiber's content-type aware
// Bind().Body, then runs Check on it. If binding fails, it writes a 400 JSON
// response describing the bind error to c and returns false. Otherwise, it
// returns the result of Check.
func Bind(c fiber.Ctx, target any) bool {
	if err := c.Bind().Body(target); err != nil {
		_ = c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		return false
	}

	return Check(c, target)
}

// Check runs checker.CheckStruct on target. If validation fails, it writes a
// 400 JSON response to c, whose body is target's checker.CheckErrors
// marshaled with CheckErrors.JSON, and returns false. Otherwise, it returns
// true, leaving target normalized by any checker normalizers for the
// handler to use.
func Check(c fiber.Ctx, target any) bool {
	errs, ok := checker.CheckStruct(target)
	if !ok {
		// json.Marshal cannot fail for FieldError's string fields.
		data, _ := errs.JSON()

		c.Set("Content-Type", jsonContentType)
		_ = c.Status(fiber.StatusBadRequest).Send(data)

		return false
	}

	return true
}
