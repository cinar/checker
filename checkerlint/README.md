# checkerlint

`checkerlint` is a static analyzer for [Checker](https://github.com/cinar/checker)'s `checkers` (and `validate` fallback) struct tags. Struct tags are opaque string literals, so a typo'd checker name, a checker applied to a field of the wrong type, or a cross-field checker (`eq-field`, `after-field`, `before-field`, `required-if`, `required-unless`) whose target field was renamed all compile fine and only fail — sometimes silently — when that code path runs. `checkerlint` catches all three at build/lint time instead.

```golang
type Registration struct {
	Password        string `checkers:"trim required"`
	ConfirmPassword string `checkers:"required eq-field:Passwrd"` // typo: no such field
	Age             int    `checkers:"email"`                     // email is string-only
}
```

```
./registration.go:3:2: checkerlint: eq-field references field "Passwrd", which doesn't exist on this struct
./registration.go:4:2: checkerlint: email requires a string, but the field's type is int
```

It knows the built-in checker vocabulary from the version of `github.com/cinar/checker/v2` it's built against, plus any custom checker registered in the analyzed packages via a `checker.RegisterMaker`/`RegisterFieldMaker` call with a string-literal name. It can't see a checker registered only at runtime in a different package, or under a name built from a non-literal expression.

## Install

```bash
go install github.com/cinar/checker/v2/checkerlint/cmd/checkerlint@latest
```

## Usage

Standalone, with the same flags and package patterns as `go vet`:

```bash
checkerlint ./...
```

As a `go vet` tool:

```bash
go vet -vettool=$(which checkerlint) ./...
```

As a custom analyzer in [`golangci-lint`](https://golangci-lint.run/) (v2's module plugin system):

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
linters-settings:
  gomodguard: {}
```

Consult `golangci-lint`'s [module plugin documentation](https://golangci-lint.run/plugins/module-plugins/) for the exact build step your `golangci-lint` version expects; module-plugin support and its build workflow have changed across `golangci-lint` releases.

## What it checks

- **Unknown checker names** — every token in a `checkers`/`validate` tag must be a registered checker/normalizer, a registered field-relative checker, `omitempty`, or a name your own code registers via `RegisterMaker`/`RegisterFieldMaker` with a string literal.
- **Type compatibility** — a well-scoped set of built-in checkers that panic at runtime on the wrong kind (string-only checkers like `email`, `trim`, `url`, ...; numeric-only checkers `gt`/`gte`/`lt`/`lte`) are checked against the tagged field's type, including through one level of pointer indirection and through a slice/array/map's element type for item-level tokens.
- **Cross-field targets** — `eq-field`, `after-field`, `before-field`, `required-if`, and `required-unless` must reference a field that actually exists on the same struct (including, on a best-effort basis, embedded/anonymous fields).

## License

This library is free to use, modify, and distribute under the terms of the MIT license found in the [LICENSE](../LICENSE) file of the parent [Checker](https://github.com/cinar/checker) repository.
