// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

// CheckFunc is a function that takes a value of type T and performs
// a check on it. It returns the resulting value and any error that
// occurred during the check.
type CheckFunc[T any] func(value T) (T, error)
