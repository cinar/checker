---
title: Catch Bad Validation Tags at Compile Time with checkerlint
published: true
description: Struct tags are just strings — a typo'd checker name, a wrong-typed field, or a renamed cross-field target all compile fine and fail silently at runtime. checkerlint catches all three before you ship.
tags: go, golang, staticanalysis, testing
canonical_url: https://dev.to/onurcinar/catch-bad-validation-tags-at-compile-time-with-checkerlint-2iaa
---

Struct tags are string literals. The Go compiler checks that your struct compiles — it has no idea what `checkers:"eq-field:Passwrd"` means, so a typo in a field name, a checker applied to a field of the wrong type, or a renamed field that a cross-field rule still points at all compile fine. They fail later, at runtime, sometimes silently, sometimes as a panic in the middle of handling a request.

```go
type Registration struct {
	Password        string `checkers:"trim required"`
	ConfirmPassword string `checkers:"required eq-field:Passwrd"` // typo: no such field
	Age             int    `checkers:"email"`                     // email is string-only
}
```

Nothing here trips `go build`, `go vet`, or a normal linter — they all treat `checkers:"..."` as an opaque string. The first bug only surfaces the moment someone submits a registration form and `eq-field` can't find a field called `Passwrd`. The second is worse: `email` assumes a string under the hood, so calling it on an `int` field panics at validation time instead of returning a normal error.

[`checkerlint`](https://github.com/cinar/checker/tree/main/checkerlint) is a `go/analysis`-based static analyzer, shipped as its own module in the [Checker](https://github.com/cinar/checker) repo, that reads these tags at build/lint time and catches exactly this class of bug before it ships:

```
./registration.go:3:2: checkerlint: eq-field references field "Passwrd", which doesn't exist on this struct
./registration.go:4:2: checkerlint: email requires a string, but the field's type is int
```

## What it actually checks

Three things, all specific to how `checkers`/`validate` tags can go wrong:

1. **Unknown checker names.** Every token in the tag has to be a registered checker, normalizer, field-relative checker, `omitempty`, or a name your own code registered via `RegisterMaker`/`RegisterFieldMaker` with a string literal. Typo `requird` instead of `required` and `checkerlint` flags it — nothing else in your toolchain will.
2. **Type compatibility.** A well-scoped set of built-ins panic at runtime on the wrong kind — string-only checkers (`email`, `trim`, `url`, and friends) and numeric-only ones (`gt`, `gte`, `lt`, `lte`). `checkerlint` checks the tagged field's actual type against what the checker expects, including through one level of pointer indirection and through a slice/array/map's element type for item-level tokens.
3. **Cross-field targets.** `eq-field`, `after-field`, `before-field`, `required-if`, and `required-unless` all name a sibling field by string. `checkerlint` confirms that field actually exists on the struct — including, on a best-effort basis, embedded/anonymous fields — so a rename doesn't silently leave a dangling reference.

It knows the built-in vocabulary from whichever version of `github.com/cinar/checker/v2` your code is built against, so it stays in sync automatically as new checkers ship. It can't see a checker registered only at runtime in a separate package, or one whose name is built from a non-literal expression — both are edge cases outside what static analysis can reach.

## Running it

Install the standalone binary:

```bash
go install github.com/cinar/checker/v2/checkerlint/cmd/checkerlint@latest
```

Run it like `go vet` — same flags, same package patterns:

```bash
checkerlint ./...
```

Or plug it into `go vet` directly as a vet tool:

```bash
go vet -vettool=$(which checkerlint) ./...
```

## Wiring it into golangci-lint

If your CI already runs `golangci-lint`, add `checkerlint` as a module plugin (`golangci-lint` v2's plugin system) instead of a separate step:

```yaml
# .golangci.yml
version: "2"
linters:
  settings:
    custom:
      checkerlint:
        type: module
        path: github.com/cinar/checker/v2/checkerlint
        settings: {}
```

Module-plugin build steps have shifted across `golangci-lint` releases, so check [its module-plugin docs](https://golangci-lint.run/plugins/module-plugins/) for the exact setup your version expects — the `.golangci.yml` shape above is the stable part.

## Why this matters more than it sounds like

Every struct-tag-driven validation library has this exact blind spot — `go-playground/validator`, `ozzo-validation`, all of them. The tags are strings; the compiler can't see inside them. Most teams find out about a typo'd checker name or a stale cross-field reference the way you'd expect: a bug report, or a support ticket about a validation rule that "just doesn't work," usually for a code path that isn't covered by a test with that exact malformed input.

`checkerlint` moves that failure from "someone hits it in production" to "the linter fails your PR." That's a category of bug — silently wrong or panicking struct tags — that the rest of the Go toolchain has no way to see, closed at the point where it's cheapest to fix.

## Try it

```bash
go install github.com/cinar/checker/v2/checkerlint/cmd/checkerlint@latest
checkerlint ./...
```

Full details, including exactly which checkers are type-checked, are in the [checkerlint README](https://github.com/cinar/checker/tree/main/checkerlint). It's a separate, independently versioned module — adding it to your CI doesn't add a single dependency to the core `checker` library your code actually imports.
