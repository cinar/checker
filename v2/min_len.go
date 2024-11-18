// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

import (
	"fmt"
)

// MinLenError defines the minimum length error.
type MinLenError struct {
	// minLen is the minimum number of characters required.
	minLen int
}

// newMinLenError initializes a new minimum length error based on the given minimum length.
func newMinLenError(minLen int) *MinLenError {
	return &MinLenError{
		minLen: minLen,
	}
}

// Error provides the string representation of the error.
func (m *MinLenError) Error() string {
	return fmt.Sprintf("must be at least %d characters", m.minLen)
}

// MinLen checks if a string has at least n characters. It returns an
// error if the string is shorter than n.
func MinLen(n int) CheckFunc[string] {
	return func(value string) (string, error) {
		if len(value) < n {
			return value, newMinLenError(n)
		}

		return value, nil
	}
}
