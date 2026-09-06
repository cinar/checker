---
title: A Command-Line Interface for Go's Checker Validation Library
published: false
description: checker is a new, dependency-free standalone binary that runs any of Checker's Go validation checkers and normalizers from a shell script, CI pipeline, or Git hook -- no Go code required.
tags: go, golang, cli, devops
---

Struct-tag validation is great when you're writing Go. It's useless the moment the thing you need to validate is a shell variable, a CI job input, or a Git tag about to get pushed. Until now, the answer for "is this a valid email/UUID/semver string" in a shell script has been a hand-rolled `grep -P` regex, copied between scripts and slowly drifting out of sync with whatever the real validation logic does elsewhere in the codebase.

[Checker](https://github.com/cinar/checker) — the zero-dependency Go validation library — now ships a standalone CLI that closes that gap:

```bash
$ checker check email "user@example.com"
user@example.com
$ echo $?
0

$ checker check email "not-an-email"
Not a valid email address.
$ echo $?
1
```

## Why a CLI, and why now

Checker's whole pitch is that validation and normalization rules live in one place — a `checkers:"..."` struct tag — instead of being reimplemented every time someone needs to check the same thing in a different context. A shell script is exactly one of those other contexts, and until now it was left out: whatever regex you wrote for `IsSemver` in Go had no equivalent you could drop into a `pre-push` hook.

The CLI fixes that by not reimplementing anything. `checker check <config> <value>` takes the *exact same* config string syntax as a struct's `checkers:"..."` tag, and runs it through the same `checker.CheckWithConfig` function a Go caller would use:

```bash
checker check semver "1.4.2"
checker check iban "DE89370400440532013000"
checker check "min-len:8" "$PASSWORD"
```

Because it calls straight through to `CheckWithConfig` and reads the checker vocabulary from `checker.RegisteredMakerNames()` at runtime instead of hardcoding a list, the CLI automatically supports every checker and normalizer the linked core module knows about — including ones that don't exist yet. Ship a new checker in the core module, rebuild the CLI against it, and it's usable from the shell with zero CLI code changes.

## Checkers and normalizers, same as in Go

Checkers and normalizers share one pipeline in Checker — `trim`, `lower`, `default`, and friends transform a value instead of just judging it — and that carries straight over to the CLI. Feed `check` a value from stdin instead of an argument, and it prints the *normalized* result:

```bash
$ echo "  Test@Example.com  " | checker check "trim lower email"
test@example.com
```

That's a real, if small, example of the "validate and normalize in one pass" idea working from a shell pipeline, not just inside a Go struct.

## Machine-readable output and locales

Two flags make it easier to wire into other tooling:

```bash
$ checker check --json email "not-an-email"
{"valid":false,"error":{"code":"NOT_EMAIL","message":"Not a valid email address."}}

$ checker check --locale=ja-JP email "not-an-email"
有効なメールアドレスではありません。
```

`--json` gives you a stable, parseable result instead of scraping stdout/stderr, and `--locale` reaches into the same 23 translated locales the Go library ships — the CLI registers all of them at startup, since a short-lived process has no caller to defer the cost to the way a library import does.

## No runtime to install

This is the part that actually matters for the CI/shell-script use case: `checker` is a single static binary. No Node.js for `ajv-cli`, no Python for `check-jsonschema` — `go install` it once, or drop a prebuilt binary into a minimal container image, and it runs. For a validation step in a Docker build or an Alpine-based CI runner, that's the difference between "one binary" and "install a whole language runtime just to check a string."

## A Git hook example

Reject a tag that isn't valid semver before it gets pushed:

```bash
#!/bin/sh
# .git/hooks/pre-push
tag=$(git describe --tags --exact-match 2>/dev/null) || exit 0
checker check semver "$tag" >/dev/null || {
	echo "refusing to push non-semver tag: $tag" >&2
	exit 1
}
```

Or gate a CI job on an environment variable actually being a URL before a deploy step runs:

```bash
checker check url "$DEPLOY_TARGET" || { echo "DEPLOY_TARGET is not a valid URL"; exit 1; }
```

## What it doesn't do

To be upfront about the boundary: `checker check` validates one value against one config string. It doesn't walk a whole JSON document against a schema, and it doesn't run `CheckStruct`-style cross-field or conditional rules (`eq-field`, `required-if`, and the rest) standalone, since those need an enclosing struct to compare against — using one reports a clean error instead of the panic the core library raises in the same situation outside `CheckStruct`. If you need whole-document validation today, [`JSONSchema`](https://pkg.go.dev/github.com/cinar/checker/v2#JSONSchema) generation from your Go structs is still the way to get there, from Go.

## Try it

```bash
go install github.com/cinar/checker/v2/cli/cmd/checker@latest
checker list   # see every checker/normalizer name
checker help   # full usage
```

It's a separate, independently versioned module — installing the CLI doesn't add a single dependency to the core `checker` library your Go code imports. Full command reference in the [cli README](https://github.com/cinar/checker/tree/main/cli).
