// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

// CheckError defines the check error.
type CheckError struct {
	// code is the error code.
	code string
}

// NewCheckError creates a new check error with the specified error code.
func NewCheckError(code string) *CheckError {
	return &CheckError{
		code: code,
	}
}

// Error returns the error message for the check.
func (c *CheckError) Error() string {
	return c.code
}
