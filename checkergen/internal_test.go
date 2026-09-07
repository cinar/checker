// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package checkergen

import (
	"go/ast"
	"testing"
)

func TestDirPattern(t *testing.T) {
	tests := []struct {
		in, want string
	}{
		{".", "."},
		{"..", ".."},
		{"./sub", "./sub"},
		{"../sub", "../sub"},
		{"/abs/path", "/abs/path"},
		{"sub", "./sub"},
	}

	for _, test := range tests {
		if got := dirPattern(test.in); got != test.want {
			t.Errorf("dirPattern(%q) = %q, want %q", test.in, got, test.want)
		}
	}
}

func TestZeroLiteral(t *testing.T) {
	tests := []struct {
		goType, want string
	}{
		{"string", `""`},
		{"bool", "false"},
		{"int", "0"},
		{"float64", "0"},
	}

	for _, test := range tests {
		if got := zeroLiteral(test.goType); got != test.want {
			t.Errorf("zeroLiteral(%q) = %q, want %q", test.goType, got, test.want)
		}
	}
}

func TestTagConfigNoTag(t *testing.T) {
	field := &ast.Field{}

	if _, ok := tagConfig(field); ok {
		t.Fatal("expected no config for a field with no tag at all")
	}
}

func TestTagConfigMalformedLiteral(t *testing.T) {
	field := &ast.Field{Tag: &ast.BasicLit{Value: "not-a-quoted-string"}}

	if _, ok := tagConfig(field); ok {
		t.Fatal("expected no config for an unparsable tag literal")
	}
}

func TestTagConfigValidateFallback(t *testing.T) {
	field := &ast.Field{Tag: &ast.BasicLit{Value: "`validate:\"required\"`"}}

	config, ok := tagConfig(field)
	if !ok || config != "required" {
		t.Fatalf("actual (%q, %v) expected (\"required\", true)", config, ok)
	}
}

func TestJSONNameNoTagAtAll(t *testing.T) {
	field := &ast.Field{Names: []*ast.Ident{{Name: "Name"}}}

	if got := jsonName(field); got != "Name" {
		t.Fatalf("actual %q expected %q", got, "Name")
	}
}

func TestJSONNameEmptyJSONTag(t *testing.T) {
	field := &ast.Field{
		Names: []*ast.Ident{{Name: "Name"}},
		Tag:   &ast.BasicLit{Value: "`json:\"\"`"},
	}

	if got := jsonName(field); got != "Name" {
		t.Fatalf("actual %q expected %q", got, "Name")
	}
}

func TestJSONNameMalformedLiteral(t *testing.T) {
	field := &ast.Field{
		Names: []*ast.Ident{{Name: "Name"}},
		Tag:   &ast.BasicLit{Value: "not-a-quoted-string"},
	}

	if got := jsonName(field); got != "Name" {
		t.Fatalf("actual %q expected %q", got, "Name")
	}
}

func TestJSONNameNoJSONTag(t *testing.T) {
	field := &ast.Field{
		Names: []*ast.Ident{{Name: "Name"}},
		Tag:   &ast.BasicLit{Value: "`checkers:\"required\"`"},
	}

	if got := jsonName(field); got != "Name" {
		t.Fatalf("actual %q expected %q", got, "Name")
	}
}

func TestJSONNameCommaOnly(t *testing.T) {
	field := &ast.Field{
		Names: []*ast.Ident{{Name: "Name"}},
		Tag:   &ast.BasicLit{Value: "`json:\",omitempty\"`"},
	}

	if got := jsonName(field); got != "Name" {
		t.Fatalf("actual %q expected %q", got, "Name")
	}
}

func TestJSONNameDash(t *testing.T) {
	field := &ast.Field{
		Names: []*ast.Ident{{Name: "Name"}},
		Tag:   &ast.BasicLit{Value: "`json:\"-\"`"},
	}

	if got := jsonName(field); got != "Name" {
		t.Fatalf("actual %q expected %q", got, "Name")
	}
}

func TestAnyFieldTaggedNone(t *testing.T) {
	st := &ast.StructType{
		Fields: &ast.FieldList{
			List: []*ast.Field{
				{Names: []*ast.Ident{{Name: "Name"}}},
			},
		},
	}

	if anyFieldTagged(st) {
		t.Fatal("expected no tagged fields")
	}
}

// TestOneNumberOrStringArgAfterNumeric confirms eq/ne's arg builder emits
// an unquoted numeric literal when the tag's parameter parses as a number
// (for a numeric field), even though it's untestable differentially: the
// core module's reflect-based eq/ne checker (checkEq/checkNe) hardcodes
// reflectString, so CheckStruct only ever supports eq/ne on a string
// field via a struct tag, unlike the fully generic exported IsEq/IsNe
// this generates a call to.
func TestOneNumberOrStringArgAfterNumeric(t *testing.T) {
	args, err := oneNumberOrStringArgAfter("eq")("42")
	if err != nil {
		t.Fatal(err)
	}

	want := []string{"value", "42"}
	if len(args) != len(want) || args[0] != want[0] || args[1] != want[1] {
		t.Fatalf("actual %v expected %v", args, want)
	}
}

// TestOneOfArgsNumeric is the oneof equivalent of
// TestOneNumberOrStringArgAfterNumeric, same reasoning: checkOneOf also
// hardcodes reflectString.
func TestOneOfArgsNumeric(t *testing.T) {
	args, err := oneOfArgs("1,2,3")
	if err != nil {
		t.Fatal(err)
	}

	want := []string{"value", "1", "2", "3"}

	if len(args) != len(want) {
		t.Fatalf("actual %v expected %v", args, want)
	}

	for i := range want {
		if args[i] != want[i] {
			t.Fatalf("actual %v expected %v", args, want)
		}
	}
}
