// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package nfkc_test

import (
	"testing"

	checker "github.com/cinar/checker/v2"
	checkernfkc "github.com/cinar/checker/v2/nfkc"
)

func TestNormalize(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "fullwidth digits compose to ASCII digits",
			input:    "１２３",
			expected: "123",
		},
		{
			name:     "ligature compatibility decomposition",
			input:    "ﬁx",
			expected: "fix",
		},
		{
			name:     "plain ASCII is left untouched",
			input:    "Onur Cinar",
			expected: "Onur Cinar",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := checkernfkc.Normalize(test.input)
			if err != nil {
				t.Fatal(err)
			}

			if actual != test.expected {
				t.Fatalf("actual %q expected %q", actual, test.expected)
			}
		})
	}
}

func TestCheckStructNFKC(t *testing.T) {
	type Handle struct {
		Name string `checkers:"nfkc"`
	}

	handle := &Handle{
		Name: "ＡＬＩＣＥ",
	}

	expected := "ALICE"

	errs, ok := checker.CheckStruct(handle)
	if !ok {
		t.Fatalf("got unexpected errors %v", errs)
	}

	if handle.Name != expected {
		t.Fatalf("actual %q expected %q", handle.Name, expected)
	}
}

func TestJSONSchemaIgnoresNFKC(t *testing.T) {
	type Handle struct {
		Name string `checkers:"nfkc required"`
	}

	schema := checker.JSONSchema(&Handle{})

	name := schema.Properties["Name"]

	if len(name.XChecker) != 0 {
		t.Fatalf("expected nfkc to be ignored, got %v", name.XChecker)
	}
}
