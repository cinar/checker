# Changelog

All notable changes to this project are documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/).

Changes prior to v2.0.0 are not individually documented here; see the
[Git history](https://github.com/cinar/checker/commits/main) and
[GitHub Releases](https://github.com/cinar/checker/releases) instead.

## [Unreleased]

### Fixed

- String checkers and normalizers (`email`, `alphanumeric`, `lower`, `trim`,
  `hex`, `url`, and others) no longer panic when applied to a defined type
  whose underlying kind is `string` (e.g. `type Email string`). They
  previously asserted `value.Interface().(string)` directly, which only
  matches the literal `string` type; normalizers also now convert the
  normalized value back to the field's original type before writing it,
  instead of panicking on `value.Set` with a plain `string`. Applying a
  string-only checker to a genuinely non-string field (e.g. `int`) still
  panics, as before, since that's a checker/field type mismatch.
- `gte`/`lte` no longer panic on unsigned integer fields (`uint`, `uint8`,
  `uint16`, `uint32`, `uint64`, `uintptr`). Previously only `CanInt()`/
  `CanFloat()` were handled, so e.g. `Quantity uint64 \`checkers:"gte:1"\``
  panicked with `"value is not numeric"` on any value. Applying `gte`/`lte`
  to a genuinely non-numeric field (`string`, `bool`, ...) still panics, as
  before — that's an intentional, tested checker/field type mismatch.
- The package-level checker/normalizer registries (`makers`, `fieldMakers`,
  `schemaMakers`, `errorMessages`) are now guarded by `sync.RWMutex`.
  `RegisterMaker`/`RegisterFieldMaker`/`RegisterSchemaMaker`/`RegisterLocale`
  wrote to these maps with no synchronization, so registering a custom
  checker or locale concurrently with `CheckStruct`/`JSONSchema` calls
  already in flight (e.g. during server startup) could trip Go's fatal,
  unrecoverable `concurrent map read and map write` error.

### Added

- `CheckErrors`, a structured error type returned by `CheckStruct`. It
  implements the `error` interface and adds `JSON()`/`JSONWithLocale()` to
  marshal validation errors into a `field name -> {code, message}` object,
  suitable for use directly as an HTTP API error response body.
- Field-relative and conditional checkers: `eq-field`, `required-if`, and
  `required-unless`, available through `CheckStruct` via a new
  `CheckFieldFunc`/`RegisterFieldMaker` mechanism.
- `CheckStruct` now traverses `map` fields the same way it already traverses
  slices: an `@`-prefixed tag applies to the map itself, a plain tag applies
  to each value.
- `checker/v2/gin` and `checker/v2/echo`, thin adapter modules that bind a
  request and run `CheckStruct` in one call, writing a JSON `400` response
  automatically when binding or validation fails. Each is its own Go module,
  so Gin/Echo are only pulled in by projects that `go get` that adapter.
- New checkers: `hash` (md5/sha1/sha256/sha384/sha512), `eoa` (Ethereum
  address shape), `before`/`after` (compare against a fixed reference time),
  `iso639-1` (language codes), `iso3166-1-alpha-2`/`iso3166-1-alpha-3`
  (country codes).
- 22 new locales, matching the set go-playground/validator ships:
  `ar-SA`, `de-DE`, `es-ES`, `fa-IR`, `fr-FR`, `hy-AM`, `id-ID`, `it-IT`,
  `ja-JP`, `ko-KR`, `lv-LV`, `nl-NL`, `pl-PL`, `pt-BR`, `pt-PT`, `ru-RU`,
  `th-TH`, `tr-TR`, `uk-UA`, `vi-VN`, `zh-CN`, `zh-TW`. A new
  `locales/locales_test.go` verifies every locale defines the same message
  codes and `{{ .placeholder }}` variables as `en-US`.
- `before-field`/`after-field` checkers, the field-relative equivalent of
  `before`/`after`: compares a value against a named sibling field's value
  instead of a fixed reference time.
- `JSONSchema`, which generates a JSON Schema document describing the shape
  and validation rules declared in a struct's `checkers` tags. Most checkers
  translate directly into a JSON Schema keyword (`required`, `minLength`/
  `maxLength`, `minimum`/`maximum`, `format`, `pattern`, ...); a checker with
  no equivalent is recorded in an `x-checker` vendor extension instead of
  being silently dropped. `RegisterSchemaMaker` lets a custom checker
  register its own translation.

### Changed

- README links to API docs now point to [pkg.go.dev](https://pkg.go.dev/github.com/cinar/checker/v2)
  instead of a committed `DOC.md`, and the GoDoc badge now points to
  pkg.go.dev instead of the retired godoc.org.
- README face lift: normalized heading levels (previously most major
  sections were `#`, the same level as the page title itself), added a
  table of contents, added a short feature-highlight list up top, and
  switched the badge row to larger `for-the-badge`-style shields (plus a
  GitHub stars badge), matching cinar/indicator's README.
- Removed the Go Report Card badge: the service was sunset and its badge
  now literally renders "go report: retired".

### Removed

- `DOC.md` and `locales/DOC.md`, the committed `gomarkdoc`-generated API
  reference. It embedded source-line links that shifted on any unrelated
  code change, which made it a near-guaranteed merge conflict on every
  concurrent pull request; pkg.go.dev already serves the same reference
  live from source, with no regeneration step and no committed file to
  conflict on.

## [2.0.1] - 2024-12-30

### Added

- `time` checker for validating a value against a named or literal Go
  reference-time layout.

## [2.0.0] - 2024-12-29

### Added

- `gte`/`lte` checkers.
- Localized error message support (`RegisterLocale`, `ErrorWithLocale`).
- `RegisterMaker` for registering custom checkers.
- `regexp` checker and `title` normalizer.

### Changed

- Major version bump to v2; see the [README](README.md) for the current API.
