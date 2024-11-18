// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

import "strings"

// TrimSpace returns the value of the string with whitespace removed from both ends.
func TrimSpace(value string) (string, error) {
	return strings.TrimSpace(value), nil
}
