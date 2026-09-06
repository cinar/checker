// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2_test

import (
	"testing"

	v2 "github.com/cinar/checker/v2"
)

func TestStripInvisible(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "zero width space splitting a filtered word",
			input:    "s\u200bex",
			expected: "sex",
		},
		{
			name:     "zero width non-joiner and joiner",
			input:    "a\u200cb\u200dc",
			expected: "abc",
		},
		{
			name:     "word joiner and BOM",
			input:    "a\u2060b\ufeffc",
			expected: "abc",
		},
		{
			name:     "bidi embedding and override controls",
			input:    "a\u202ab\u202bc\u202cd\u202de\u202ef",
			expected: "abcdef",
		},
		{
			name:     "bidi isolate controls",
			input:    "a\u2066b\u2067c\u2068d\u2069e",
			expected: "abcde",
		},
		{
			name:     "plain string is left untouched",
			input:    "Onur Cinar",
			expected: "Onur Cinar",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := v2.StripInvisible(test.input)
			if err != nil {
				t.Fatal(err)
			}

			if actual != test.expected {
				t.Fatalf("actual %q expected %q", actual, test.expected)
			}
		})
	}
}

func TestReflectStripInvisible(t *testing.T) {
	type Person struct {
		Name string `checkers:"strip-invisible"`
	}

	person := &Person{
		Name: "O\u200bnur",
	}

	expected := "Onur"

	errs, ok := v2.CheckStruct(person)
	if !ok {
		t.Fatalf("got unexpected errors %v", errs)
	}

	if person.Name != expected {
		t.Fatalf("actual %q expected %q", person.Name, expected)
	}
}
