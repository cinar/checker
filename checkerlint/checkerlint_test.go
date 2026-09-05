// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package checkerlint_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/cinar/checker/v2/checkerlint"
)

func TestAnalyzer(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), checkerlint.Analyzer, "a")
}
