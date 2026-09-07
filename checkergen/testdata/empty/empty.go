// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

// Package empty has no struct with a checkers/validate tag at all, to
// exercise Generate's no-op path: no output file is written.
package empty

type Plain struct {
	Name string
}
