# checker (CLI)

`checker` is a standalone command-line interface to [Checker](https://github.com/cinar/checker), for running its checkers and normalizers from shell scripts, CI pipelines, and Git hooks — no Go code required. It ships as a single, dependency-free static binary: no Node, Python, or JVM runtime to install alongside it.

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

It doesn't hardcode or reimplement the checker vocabulary: every subcommand goes through the same `checker.CheckWithConfig`, `checker.RegisteredMakerNames`, and `checker.RegisteredFieldMakerNames` calls a Go caller would use, so it automatically supports every checker and normalizer the linked core module knows about — built-in or custom-registered, today or in a future release — with no CLI changes required.

## Install

```bash
go install github.com/cinar/checker/v2/cli/cmd/checker@latest
```

## Usage

### `checker check`

Runs a `checkers`/`validate` tag config string — the exact same syntax as a struct field's `checkers:"..."` tag — against a value, or against stdin if the value is omitted:

```bash
checker check email "user@example.com"
echo "  Test@Example.com  " | checker check "trim lower email"
```

Checkers and normalizers share one pipeline, so this also works for normalization alone, or a mix:

```bash
$ checker check "trim lower" "  Test@Example.com  "
test@example.com
```

Flags:

- `--locale=<tag>` — render the error message in one of the 23 shipped locales (e.g. `de-DE`, `ja-JP`; see the root [README](../README.md#localized-error-messages) for the full list). Defaults to `en-US`.
- `--json` — print a JSON result object instead of plain text, for scripting:

  ```bash
  $ checker check --json email "not-an-email"
  {"valid":false,"error":{"code":"NOT_EMAIL","message":"Not a valid email address."}}
  ```

Exit codes: `0` if every check passes, `1` if the value fails a check *or* the configuration string itself is invalid (an unknown checker name, or a field-relative checker used standalone), `2` for a command-line usage error (bad flags, wrong number of arguments).

Field-relative checkers (`eq-field`, `after-field`, `before-field`, `required-if`, `required-unless`) need an enclosing struct to compare against, which a single command-line value doesn't have. Using one with `check` reports a clean "invalid check configuration" error and exits `1`, rather than the panic the core module raises when the same thing happens outside `CheckStruct`.

### `checker list`

Lists every registered checker/normalizer name, plus, separately, the field-relative ones that `check` can't run standalone:

```bash
checker list
```

### `checker version`

Prints the version of the `github.com/cinar/checker/v2` module the binary was built against:

```bash
checker version
```

## Example: a Git hook

```bash
#!/bin/sh
# .git/hooks/pre-push -- reject tags that aren't valid semver
tag=$(git describe --tags --exact-match 2>/dev/null) || exit 0
checker check semver "$tag" >/dev/null || {
	echo "refusing to push non-semver tag: $tag" >&2
	exit 1
}
```

## License

This library is free to use, modify, and distribute under the terms of the MIT license found in the [LICENSE](../LICENSE) file of the parent [Checker](https://github.com/cinar/checker) repository.
