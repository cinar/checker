# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What this is

Checker is a zero-dependency Go library (`github.com/cinar/checker/v2`) for validating and normalizing user
input, driven by struct tags (`checkers:"..."`) or plain function calls. It ships 23 translated locales, JSON
Schema generation from checker tags, and separately-versioned `gin`/`echo` framework adapter modules.

## Repository layout

This is a multi-module repo:
- `.` — the core `checker` module (`go.mod`, go 1.23.2). No external dependencies.
- `gin/` — `github.com/cinar/checker/v2/gin`, its own `go.mod`/`taskfile.yml`, `replace`s the core module with `../`.
- `echo/` — `github.com/cinar/checker/v2/echo`, same pattern as `gin/`.
- `locales/` — the `locales` package, one file per locale (e.g. `en_us.go`) plus `locales.go` and `locales_test.go`.

Each module has its own `taskfile.yml` and is built/linted/tested independently (see `.github/workflows/ci.yml`,
which runs three parallel jobs: root, `gin/`, `echo/`).

## Commands

All commands use [Task](https://taskfile.dev) (`go run github.com/go-task/task/v3/cmd/task@v3.38.0`), run from
the module's own directory (root, `gin/`, or `echo/`):

- `task` (default) — runs `fmt`, `lint`, then `test`, in that order.
- `task fmt` — `go fix ./...`
- `task lint` — `go vet`, `gosec` (excludes `gin`/`echo` dirs when run from root), `staticcheck`, and `revive`
  (config in `revive.toml`).
- `task test` — `go test -cover -coverprofile=coverage.out ./...` then prints per-func coverage.

Equivalent plain `go` commands work too, e.g.:
- Run a single test: `go test -run TestRequired ./...` (or `go test -run TestRequired .` from root).
- Run all tests with coverage: `go test -cover ./...`

The project enforces **100% test coverage** (see `CONTRIBUTING.md`) — any new checker, normalizer, or branch
needs tests to match, in the corresponding `_test.go` file.

## Core architecture

### Checker pipeline

A "checker" is a `CheckFunc[T] func(value T) (T, error)` (`check_func.go`). Checkers and normalizers are the same
type — a normalizer just transforms the value and returns a nil error (e.g. `Trim`, `Lower`). `Check` /
`CheckWithConfig` (`checker.go`) run a sequence of these against one value, short-circuiting on the first error.

Each built-in checker lives in its own file at the repo root (e.g. `required.go`, `email.go`, `after_field.go`),
paired with a `_test.go`. A checker file typically defines:
1. A generic exported function usable standalone, e.g. `Required[T any](value T) (T, error)`.
2. A `reflect.Value`-based variant used internally by struct-tag checking.
3. A `make<Name>(params string) CheckFunc[reflect.Value]` "maker" function that parses the checker's tag
   parameters (e.g. `"layout:field"` for `after-field`) into a closure.
4. A `name<Name>` constant (the tag keyword) and any exported `Err*` sentinel `*CheckError` values.

Field-relative checkers (compare against a sibling struct field, e.g. `eq-field`, `after-field`, `required-if`)
instead implement `CheckFieldFunc func(parent, value reflect.Value) (reflect.Value, error)` (`check_field_func.go`)
and a `makeCheckFieldFunc`; they panic if invoked outside `CheckStruct` (no parent struct context). All maker
functions are registered by name in the `makers`/`fieldMakers` maps in `maker.go`; new ones can be added
externally via `RegisterMaker`/`RegisterFieldMaker`.

### Struct-tag validation

`CheckStruct` (`checker.go`) walks a struct via `reflect`, breadth-first, handling nested structs, slices, and
maps. A `checkStructJob` queue entry carries the value, its parent (for field-relative checks), its tag config,
and a `SetFunc` to write the normalized value back (map values need special handling since a non-pointer map
value isn't addressable). Slice/map tags split into element-level vs. container-level configs using an `@`
prefix (`splitSliceConfig`), so e.g. `checkers:"@min-len:1 required"` validates the slice itself while `required`
applies to each element.

### Errors

`*CheckError` (`check_error.go`) carries a machine-readable `Code` plus template `Data`, and renders a localized
message via `html/template` against `errorMessages[locale]`. `CheckErrors` (`check_errors.go`) is the
`map[string]error` returned by `CheckStruct`, keyed by fully-qualified field name (dotted for nested structs,
bracketed for slice/map indices); it implements `error` itself and has `JSON()`/`JSONWithLocale()` for producing
an HTTP-API-ready error body.

### Locales

`locales/` holds one file per locale, each exporting a `map[string]string` of error code → templated message
(same `{{ .placeholder }}` variables as `en_us.go` for a given code). `locales_test.go` enforces every locale
defines every code with matching placeholders — adding a checker (new error code) or changing a message's
placeholders requires updating every locale file, not just `en_us.go`.

### JSON Schema generation

`schema.go`'s `JSONSchema(st any) *Schema` walks a struct's type (not its values — this is purely static) and
translates checker tags into a JSON Schema document. `schema_maker.go` maps checker names to `SchemaMakeFunc`s
that refine a field's `*Schema` (e.g. `min-len` → `MinLength`); checkers without a mapping are recorded verbatim
in `Schema.XChecker` instead of being silently dropped, and normalizers are ignored entirely (tracked in
`ignoredForSchema`). New checkers that constrain shape should register a schema maker via `RegisterSchemaMaker`.

### Framework adapters (`gin/`, `echo/`)

Each adapter is a thin, separately-versioned module exposing `Bind`/`Check` functions that bind a request body
then call `checker.CheckStruct`, writing a JSON 400 response (via `CheckErrors.JSON()`) automatically on
failure. Keep these modules dependency-isolated from the core — they `replace` the core module with `../` for
local development/CI only; that replace has no effect on external consumers.

## Conventions

- Every new file needs the standard copyright header (see `header.txt` / any existing file for the exact form).
- Don't edit `CHANGELOG.md` for a PR. It no longer accumulates a running `## [Unreleased]` section — that pattern
  made it a near-guaranteed merge conflict whenever two PRs landed close together (the same problem `DOC.md` used
  to cause). Unreleased changes are covered by GitHub's auto-generated release notes instead; see the note at the
  top of `CHANGELOG.md`.
- `revive.toml` enforces `package-comments`, `exported`, etc. — exported identifiers need doc comments.
