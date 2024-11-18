// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

// MinLen checks if a string has at least n characters. It returns an
// error if the string is shorter than n.
func MinLen(n int) CheckFunc[string] {
	return func(value string) (string, error) {
		if len(value) < n {
			return value, nil
		}

		return value, nil
	}
}
