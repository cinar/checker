// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package checkergen

import (
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"

	"golang.org/x/tools/go/packages"
)

// dirPattern turns dir into a go/packages directory pattern: it needs a
// "./" or "../" prefix (or to be absolute) to be recognized as a
// filesystem path rather than an import path.
func dirPattern(dir string) string {
	if filepath.IsAbs(dir) || strings.HasPrefix(dir, "."+string(filepath.Separator)) || strings.HasPrefix(dir, ".."+string(filepath.Separator)) || dir == "." || dir == ".." {
		return dir
	}

	return "." + string(filepath.Separator) + dir
}

const (
	checkerTag        = "checkers"
	validateTag       = "validate"
	omitEmptyName     = "omitempty"
	sliceConfigPrefix = "@"
)

// scalarGoTypes is the set of predeclared scalar type names a field's Go
// type must match exactly (see the package doc) for checkergen to generate
// code for it.
var scalarGoTypes = map[string]bool{
	"string": true, "bool": true,
	"int": true, "int8": true, "int16": true, "int32": true, "int64": true,
	"uint": true, "uint8": true, "uint16": true, "uint32": true, "uint64": true,
	"float32": true, "float64": true,
}

// zeroLiteral returns the Go zero-value literal for one of scalarGoTypes,
// used to generate an omitempty guard.
func zeroLiteral(goType string) string {
	switch goType {
	case "string":
		return `""`
	case "bool":
		return "false"
	default:
		return "0"
	}
}

// loadedStruct is one named struct type found in a loaded package, along
// with the type info needed to resolve its fields' Go types.
type loadedStruct struct {
	name   string
	fields []*ast.Field
	info   *types.Info
}

// findStructs loads pkgPattern (a go/packages pattern, e.g. "." for the
// current directory's package) and returns every named struct type
// declared in it, optionally restricted to typeFilter (struct names; a nil
// or empty filter means every struct).
func findStructs(dir string, typeFilter map[string]bool) ([]loadedStruct, string, error) {
	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedFiles | packages.NeedSyntax |
			packages.NeedTypes | packages.NeedTypesInfo,
	}

	// go/packages treats a bare relative path as an import path, not a
	// directory, unless it's prefixed with "./" or "../" -- normalize so
	// a caller passing a plain relative directory (as the checkergen CLI's
	// own default "." already does, but a caller-supplied one might not)
	// doesn't need to know that.
	pkgPattern := dirPattern(dir)

	pkgs, err := packages.Load(cfg, pkgPattern)
	if err != nil {
		return nil, "", fmt.Errorf("loading package %s: %w", pkgPattern, err)
	}

	if packages.PrintErrors(pkgs) > 0 {
		return nil, "", fmt.Errorf("package %s has errors, see above", pkgPattern)
	}

	if len(pkgs) != 1 {
		return nil, "", fmt.Errorf("pattern %s must resolve to exactly one package, got %d", pkgPattern, len(pkgs))
	}

	pkg := pkgs[0]

	var structs []loadedStruct

	for _, file := range pkg.Syntax {
		for _, decl := range file.Decls {
			genDecl, ok := decl.(*ast.GenDecl)
			if !ok || genDecl.Tok != token.TYPE {
				continue
			}

			for _, spec := range genDecl.Specs {
				typeSpec, ok := spec.(*ast.TypeSpec)
				if !ok {
					continue
				}

				structType, ok := typeSpec.Type.(*ast.StructType)
				if !ok {
					continue
				}

				if len(typeFilter) > 0 && !typeFilter[typeSpec.Name.Name] {
					continue
				}

				if !anyFieldTagged(structType) {
					continue
				}

				structs = append(structs, loadedStruct{
					name:   typeSpec.Name.Name,
					fields: structType.Fields.List,
					info:   pkg.TypesInfo,
				})
			}
		}
	}

	return structs, pkg.Name, nil
}

// anyFieldTagged reports whether at least one field of st has a non-empty
// checkers or validate tag, so a struct with no validation rules at all is
// skipped without a diagnostic (it's simply not a candidate, same as
// CheckStruct silently doing nothing for such a field).
func anyFieldTagged(st *ast.StructType) bool {
	for _, field := range st.Fields.List {
		if config, ok := tagConfig(field); ok && config != "" {
			return true
		}
	}

	return false
}

// tagConfig returns the effective checkers/validate tag config for field,
// and whether a "checkers" or "validate" tag was present at all -- mirroring
// the core module's fieldConfig, which prefers "checkers" and falls back to
// "validate".
func tagConfig(field *ast.Field) (string, bool) {
	if field.Tag == nil {
		return "", false
	}

	unquoted, err := strconv.Unquote(field.Tag.Value)
	if err != nil {
		return "", false
	}

	tag := reflect.StructTag(unquoted)

	if config, ok := tag.Lookup(checkerTag); ok {
		return config, true
	}

	config, ok := tag.Lookup(validateTag)
	return config, ok
}

// jsonName returns field's tag-facing name -- CheckStruct's own
// localFieldName logic: the json tag's property name if present (falling
// back to the Go field name for an empty or "-" json tag, or one with no
// name before the first comma), otherwise the Go field name.
func jsonName(field *ast.Field) string {
	goName := field.Names[0].Name

	if field.Tag == nil {
		return goName
	}

	unquoted, err := strconv.Unquote(field.Tag.Value)
	if err != nil {
		return goName
	}

	tag, ok := reflect.StructTag(unquoted).Lookup("json")
	if !ok || tag == "" || tag == "-" {
		return goName
	}

	name, _, _ := strings.Cut(tag, ",")
	if name == "" {
		return goName
	}

	return name
}

// goTypeName returns t's Go source-level type name, and whether it's one
// of scalarGoTypes: only a *types.Basic whose name is directly in that set
// qualifies, so a named/defined type (type Email string) or a pointer is
// reported as ineligible rather than silently treated as its underlying
// type -- see the package doc.
func goTypeName(t types.Type) (string, bool) {
	basic, ok := t.(*types.Basic)
	if !ok {
		return t.String(), false
	}

	return basic.Name(), scalarGoTypes[basic.Name()]
}
