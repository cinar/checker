// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

// Package gin integrates Checker with the Gin web framework, binding and
// validating a request body in a single call, and writing a JSON 400
// response automatically when binding or validation fails.
package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"

	checker "github.com/cinar/checker/v2"
)

// Bind binds the request body into obj using Gin's content-type aware
// ShouldBind, then runs Check on it. If binding fails, it aborts the context
// with a 400 JSON response describing the bind error and returns false.
// Otherwise, it returns the result of Check.
func Bind(c *gin.Context, obj any) bool {
	if err := c.ShouldBind(obj); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return false
	}

	return Check(c, obj)
}

// Check runs checker.CheckStruct on obj. If validation fails, it aborts the
// context with a 400 JSON response, whose body is obj's checker.CheckErrors
// marshaled with CheckErrors.JSON, and returns false. Otherwise, it returns
// true, leaving obj normalized by any checker normalizers for the handler
// to use.
func Check(c *gin.Context, obj any) bool {
	errs, ok := checker.CheckStruct(obj)
	if !ok {
		// json.Marshal cannot fail for FieldError's string fields.
		data, _ := errs.JSON()

		c.Abort()
		c.Data(http.StatusBadRequest, "application/json; charset=utf-8", data)

		return false
	}

	return true
}
