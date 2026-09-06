# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What this is

Checker is a zero-dependency Go library (`github.com/cinar/checker/v2`) for validating and normalizing user
input, driven by struct tags (`checkers:"..."`) or plain function calls. It ships 23 translated locales, JSON
Schema generation from checker tags, separately-versioned `gin`/`echo`/`nethttp`/`fiber` framework adapter
modules, an opt-in `nfkc` checker module, a `checkerlint` static analyzer, a `checkergen` reflection-free
code generator, and a standalone `cli` command-line interface.

## Repository layout

This is a multi-module repo:
- `.` — the core `checker` module (`go.mod`, go 1.23.2). No external dependencies.
- `gin/` — `github.com/cinar/checker/v2/gin`, its own `go.mod`/`taskfile.yml`, `replace`s the core module with `../`.
- `echo/` — `github.com/cinar/checker/v2/echo`, same pattern as `gin/`.
- `checkerlint/` — `github.com/cinar/checker/v2/checkerlint`, a `go/analysis` static analyzer for `checkers`/
  `validate` struct tags, plus its `cmd/checkerlint` standalone binary. Same `replace ../` pattern as `gin`/`echo`,
  but pins newer linter versions in its own `taskfile.yml` (it needs a newer Go toolchain than the core module's
  `go 1.23.2` for its `golang.org/x/tools` dependency). Its `testdata/src/...` tree stands in for the real
  `github.com/cinar/checker/v2` import path so `analysistest` fixtures can exercise custom-checker detection
  without pulling in the real module — don't confuse it with an actual dependency.
- `cli/` — `github.com/cinar/checker/v2/cli`, a standalone command-line interface (`cmd/checker` binary) that runs
  any registered checker/normalizer via `CheckWithConfig` against a value from a shell script, CI pipeline, or Git
  hook. Same `replace ../` pattern as `gin`/`echo`/`checkerlint`; has no external dependencies at all (no `go.sum`),
  since it only depends on the core module.
- `nethttp/` — `github.com/cinar/checker/v2/nethttp`, a `net/http` adapter (`Bind`/`Check`) built only on
  `encoding/json` + `net/http`. Same `replace ../` pattern as `gin`/`echo`; no external dependencies at all (no
  `go.sum`), same as `cli/`.
- `fiber/` — `github.com/cinar/checker/v2/fiber`, a Fiber v3 adapter (`Bind`/`Check`). Same `replace ../` pattern
  as `gin`/`echo`.
- `nfkc/` — `github.com/cinar/checker/v2/nfkc`, registers an `nfkc` checkers-tag normalizer (Unicode
  Normalization Form KC) via `checker.RegisterMaker` in its `init` function — a blank import is enough to use
  it. Its own module, not part of the core, because NFKC needs `golang.org/x/text/unicode/norm`; keeping it
  isolated keeps the core module's zero-dependency promise intact for callers who don't opt in. Same
  `replace ../` pattern as the other adapters.
- `checkergen/` — `github.com/cinar/checker/v2/checkergen`, generates a `Check<Type>(v *Type)
  (checker.CheckErrors, bool)` function per eligible struct, calling the same checker/normalizer plain
  functions (`checker.IsEmail`, `checker.MinLen[string](8)`, ...) directly via `checker.Check` instead of
  walking the struct with `reflect` at runtime — coexists with `CheckStruct`, not a replacement for it. Same
  `replace ../` pattern as `checkerlint`, and needs the same newer Go toolchain for `golang.org/x/tools`
  (`go/packages`, for type info). Its `testdata/fixture/` and `testdata/ineligible/` hold committed,
  regenerated-by-test `checkergen_generated.go` fixtures (see `generate_test.go`'s
  `TestGenerateMatchesCommittedFixture` for the drift check, and `differential_test.go` for the
  behavioral-parity-with-`CheckStruct` check) — regenerate them with
  `go run ./cmd/checkergen ./testdata/<dir>` after a `callSpecs` change, same as any other generated-and-committed
  code. A test that calls `Generate` with anything other than a plain unfiltered rerun (e.g. a `-type` filter)
  must target a scratch copy (`copyToScratch` in `generate_test.go`), never a real `testdata` package directly,
  since `Generate` overwrites that package's committed, shared output file as a side effect.
- `locales/` — the `locales` package, one file per locale (e.g. `en_us.go`) plus `locales.go` and `locales_test.go`.

Each module has its own `taskfile.yml` and is built/linted/tested independently (see `.github/workflows/ci.yml`,
which runs parallel jobs per module: root, `gin/`, `echo/`, `nethttp/`, `fiber/`, `nfkc/`, `checkerlint/`,
`checkergen/`, `cli/`).

## Commands

All commands use [Task](https://taskfile.dev) (`go run github.com/go-task/task/v3/cmd/task@v3.38.0`), run from
the module's own directory (root, `gin/`, `echo/`, `nethttp/`, `fiber/`, `nfkc/`, `checkerlint/`, `checkergen/`,
or `cli/`):

- `task` (default) — runs `fmt`, `lint`, then `test`, in that order.
- `task fmt` — `go fix ./...`
- `task lint` — `go vet`, `gosec` (excludes
  `gin`/`echo`/`examples`/`checkerlint`/`cli`/`nethttp`/`fiber`/`nfkc`/`checkergen` dirs when run from root —
  `gosec` doesn't respect Go module boundaries or the `testdata` convention the way `go build`/`go vet`/`go test`
  do, so any subdirectory with its own `go.mod`, or a `testdata` tree with fixtures that don't type-check
  against the root module, needs an explicit `-exclude-dir` — **a new module here needs this list updated in
  both the root `taskfile.yml` and this file, or the root `build` CI job breaks**, as happened after the
  `nethttp`/`fiber` modules landed), `staticcheck`, and `revive` (config in `revive.toml`).
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
additionally implement `CheckFieldFunc func(parent, value reflect.Value) (reflect.Value, error)`
(`check_field_func.go`) and a `makeCheckFieldFunc`; the `CheckFieldFunc` variant panics if invoked outside
`CheckStruct` (no parent struct context), since it resolves the sibling by name via `parent`. Its own plain
function (item 1 above, e.g. `IsEqField[T comparable](value, other T, name string) (T, error)`) instead takes
the sibling's already-resolved value directly, so it's usable standalone; the `reflect.Value`-based variant
(item 2, e.g. `checkEqField`) resolves the sibling via `parent` and, where the semantics match exactly
(`after-field`/`before-field`), delegates to the plain function. `eq-field`'s `reflect.Value` variant keeps
`reflect.DeepEqual` instead of the plain function's `==`, since a struct-tag field's type isn't known until
runtime and might not satisfy `comparable`. All maker functions are registered by name in the
`makers`/`fieldMakers` maps in `maker.go`; new ones can be added
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

### Framework adapters (`gin/`, `echo/`, `nethttp/`, `fiber/`)

Each adapter is a thin, separately-versioned module exposing `Bind`/`Check` functions that bind a request body
then call `checker.CheckStruct`, writing a JSON 400 response (via `CheckErrors.JSON()`) automatically on
failure. Keep these modules dependency-isolated from the core — they `replace` the core module with `../` for
local development/CI only; that replace has no effect on external consumers. `nethttp/` has no dependency
beyond the core module (`encoding/json` + `net/http`, both stdlib); `gin/`, `echo/`, and `fiber/` each pull in
their respective framework.

### Command-line interface (`cli/`)

`cli/cli.go`'s `Run(args, stdin, stdout, stderr) int` is the whole CLI's entry point, exercised directly from
tests (no `os.Exit` inside it) and wrapped by a two-line `cmd/checker/main.go`. It never hardcodes the checker
vocabulary — `check` goes through `checker.CheckWithConfig`, `list` through `RegisteredMakerNames`/
`RegisteredFieldMakerNames` — so it can't drift out of sync with the core module's registered checkers. It
recovers panics from `CheckWithConfig` (an unknown checker name, or a field-relative checker used outside a
struct) into a normal returned error instead of crashing, since a CLI's checker-config argument is closer to
user input than the compile-time struct tag the core module assumes. `locales.go` eagerly registers all 23
shipped locales at `init` so `--locale=<tag>` works out of the box, unlike the core module's opt-in
`RegisterLocale` (importing the core module should never force-pull translations a caller doesn't use, but a
standalone CLI process has no such caller to defer to).

## Conventions

- Every new file needs the standard copyright header (see `header.txt` / any existing file for the exact form).
- Don't edit `CHANGELOG.md` for a PR. It no longer accumulates a running `## [Unreleased]` section — that pattern
  made it a near-guaranteed merge conflict whenever two PRs landed close together (the same problem `DOC.md` used
  to cause). Unreleased changes are covered by GitHub's auto-generated release notes instead; see the note at the
  top of `CHANGELOG.md`.
- `revive.toml` enforces `package-comments`, `exported`, etc. — exported identifiers need doc comments.
