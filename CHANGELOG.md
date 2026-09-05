# Changelog

All notable changes to this project are documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/).

Changes prior to v2.0.0 are not individually documented here; see the
[Git history](https://github.com/cinar/checker/commits/main) and
[GitHub Releases](https://github.com/cinar/checker/releases) instead.

## [Unreleased]

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

### Changed

- README links to API docs now point to [pkg.go.dev](https://pkg.go.dev/github.com/cinar/checker/v2)
  instead of a committed `DOC.md`, and the GoDoc badge now points to
  pkg.go.dev instead of the retired godoc.org.

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
