// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

// Command checkerlint runs the checkerlint static analyzer standalone,
// with the same flags and usage as go vet.
package main

import (
	"github.com/cinar/checker/v2/checkerlint"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(checkerlint.Analyzer)
}
