<p align="center">
    <a href="https://pkg.go.dev/github.com/cinar/checker/v2"><img src="https://img.shields.io/badge/Go_Reference-007D9C?style=for-the-badge&logo=go&logoColor=white" alt="Go Reference" /></a>
    <a href="LICENSE"><img src="https://img.shields.io/github/license/cinar/checker?style=for-the-badge" alt="License" /></a>
    <a href="https://github.com/cinar/checker/actions/workflows/ci.yml"><img src="https://img.shields.io/github/actions/workflow/status/cinar/checker/ci.yml?branch=main&style=for-the-badge&logo=github&label=CI" alt="Go CI" /></a>
    <a href="https://codecov.io/gh/cinar/checker"><img src="https://img.shields.io/codecov/c/github/cinar/checker?style=for-the-badge&logo=codecov&logoColor=white" alt="Codecov" /></a>
    <a href="https://github.com/cinar/checker/stargazers"><img src="https://img.shields.io/github/stars/cinar/checker?style=for-the-badge&logo=github&logoColor=white" alt="GitHub Stars" /></a>
</p>

<h1 align="center">
    <img src="logo.svg" alt="Checker logo: a tag with a punched checkmark" width="128" height="128" /><br />
    Checker
</h1>

<p align="center">Checker is a lightweight Go library for validating and normalizing user input, driven by struct tags or plain function calls, with zero external dependencies.</p>

<p align="center">
    <a href="#quickstart-example">Quickstart</a> &middot;
    <a href="#normalizers-provided">Normalizers</a> &middot;
    <a href="#checkers-provided">Checkers</a> &middot;
    <a href="#json-schema-generation">JSON Schema</a> &middot;
    <a href="#localized-error-messages">Locales</a> &middot;
    <a href="#framework-integration">Frameworks</a>
</p>

- **Zero dependencies** — the core module imports nothing beyond the Go standard library.
- **Struct tags or plain functions** — validate a whole struct declaratively, or call checkers directly for one-off values.
- **Checkers and normalizers together** — trim, then require; lowercase, then validate — mixed in any order, in one pass.
- **Cross-field and conditional rules** — compare fields against each other, or require a field only when another has a given value.
- **23 built-in locales** — opt-in, translated error messages, matching the set go-playground/validator ships.
- **JSON Schema generation** — turn a struct's checker tags into a JSON Schema document, for API docs or frontend validation, without hand-maintaining a second copy of your rules.
- **Framework adapters** — thin, separately-versioned `gin`, `echo`, `nethttp`, and `fiber` modules bind a request and validate it in one call.
- **Trojan Source & Unicode normalization** — the built-in `strip-invisible` normalizer strips zero-width/bidi spoofing characters, and the opt-in `nfkc` module folds compatibility characters (fullwidth digits, ligatures, ...) into their canonical form.
- **Context-aware pipelines** — `Pipeline[T]` reuses the same checkers for domain rules that need a `context.Context` (a DB lookup, a tenant claim) struct tags can't carry.
- **Static analysis** — the separate `checkerlint` module catches unknown checker names, type mismatches, and dangling cross-field targets at build time, not runtime.
- **Code generation** — the separate `checkergen` module turns `checkers` tags into reflection-free Go validation code, ~3x faster with far fewer allocations, for structs on a hot path.
- **Command-line interface** — the separate `checker` binary runs any checker or normalizer from a shell script, CI pipeline, or Git hook, with no Go code and no runtime dependency of its own.

## Table of Contents

- [Why Checker?](#why-checker)
- [Usage](#usage)
  - [Quickstart Example](#quickstart-example)
  - [Interactive Playgrounds & Examples](#interactive-playgrounds--examples)
  - [Validating Structs](#validating-structs)
  - [Validating Individual Values](#validating-individual-values)
- [Normalizers and Checkers](#normalizers-and-checkers)
- [Normalizers Provided](#normalizers-provided)
- [Checkers Provided](#checkers-provided)
- [Custom Checkers and Normalizers](#custom-checkers-and-normalizers)
- [Slice and Item Level Checkers](#slice-and-item-level-checkers)
- [Optional Fields with `omitempty`](#optional-fields-with-omitempty)
- [Field-Relative and Conditional Checkers](#field-relative-and-conditional-checkers)
- [Programmatic Pipelines (Context-Aware Validation)](#programmatic-pipelines-context-aware-validation)
  - [Context-Aware Struct Tags with `CheckStructWithContext`](#context-aware-struct-tags-with-checkstructwithcontext)
- [Localized Error Messages](#localized-error-messages)
- [Custom Error Messages](#custom-error-messages)
- [Structured Errors](#structured-errors)
- [JSON Schema Generation](#json-schema-generation)
- [Framework Integration](#framework-integration)
- [Unicode Normalization (NFKC)](#unicode-normalization-nfkc)
- [Static Analysis](#static-analysis)
- [Code Generation](#code-generation)
- [Command-Line Interface](#command-line-interface)
- [Performance](#performance)
- [Changelog](#changelog)
- [Contributing to the Project](#contributing-to-the-project)
- [License](#license)

## Why Checker?

Validating user input in Go often requires piecing together separate libraries: one for struct validation, another for string trimming and normalization, and custom code or extra tooling for JSON Schemas.

Checker provides a **single, zero-dependency pipeline** designed around three core pillars:

1. **Zero External Dependencies** — The core module imports nothing beyond the Go standard library. No supply-chain bloat or dependency drift.
2. **Unified Normalization & Validation** — Clean, transform, and validate input in one atomic pass (`trim lower required email`). Input is mutated in-place.
3. **Living JSON Schemas** — Struct tags generate Draft 2020-12 JSON Schema documents for frontend contract sharing and API docs without duplicating rules.

### Feature Comparison

| Feature | Checker (`v2`) | go-playground/validator | ozzo-validation |
| :--- | :---: | :---: | :---: |
| **External Dependencies** | **0 (Standard library only)** | 4+ packages | Standard library only |
| **In-place Normalization** | **Built-in (`trim`, `lower`, etc.)** | Not supported | Not supported |
| **JSON Schema Generation** | **Built-in (`checker.JSONSchema`)** | Requires external tooling | Not supported |
| **Cross-field Validation** | **Built-in (`eq-field`, `required-if`, ...)** | Built-in (`eqfield`, etc.) | Custom rules |
| **Slice / Container Rules** | **Container (`@max-len`) + Items** | Items only | Custom loops |
| **Internationalization** | **23 Locales (Opt-in import)** | 23 Locales (Opt-in, per-language package) | Manual translation |
| **API Error Payloads** | **Built-in (`errs.JSON()`)** | Manual formatting | Manual formatting |

Checker's own test coverage is enforced at 100% in CI (see the badge above); third-party coverage numbers change too often to be worth restating here — check each project's own CI badge for its current figure.

## Usage

To begin using the Checker library, install it with the following command:

```bash
go get github.com/cinar/checker/v2
```

Then, import the library into your source file:

```golang
import (
	checker "github.com/cinar/checker/v2"
)
```

### Quickstart Example

Here is a real-world registration request showing in-place normalization, cross-field comparison, slice validation, and structured error responses:

```golang
package main

import (
	"fmt"

	checker "github.com/cinar/checker/v2"
)

type SignupRequest struct {
	// 1. Normalizes in-place (trims spaces, converts to lowercase) and validates:
	Email           string   `json:"email" checkers:"trim lower required email"`
	Password        string   `json:"password" checkers:"required min-len:8"`
	ConfirmPassword string   `json:"confirm_password" checkers:"required eq-field:Password"`
	Roles           []string `json:"roles" checkers:"@max-len:3 trim alphanumeric"`
}

func main() {
	req := &SignupRequest{
		Email:           "  ALICE@EXAMPLE.COM  ",
		Password:        "supersecret123",
		ConfirmPassword: "supersecret123",
		Roles:           []string{"  admin  ", "editor"},
	}

	// CheckStruct normalizes fields in-place and validates all rules:
	errs, valid := checker.CheckStruct(req)
	if !valid {
		// Marshals errors directly to JSON for HTTP API responses:
		data, _ := errs.JSON()
		fmt.Println(string(data))
		return
	}

	// Values are sanitized in-place and ready for database persistence:
	fmt.Println(req.Email) // "alice@example.com"
	fmt.Println(req.Roles) // ["admin", "editor"]
}
```

> 🎮 **[Try this example live on the Go Playground](https://go.dev/play/p/FfkXm5oC9ii)**

### Interactive Playgrounds & Examples

Try Checker immediately in your browser via the official Go Playground:

- 🎮 [**Quickstart & Normalization Playground**](https://go.dev/play/p/5X8ukfSOnZ1) — In-place normalization, validation, and JSON error responses.
- 🎮 [**JSON Schema Generation Playground**](https://go.dev/play/p/MkBT8QVY5-c) — Generate Draft 2020-12 schemas from struct tags.
- 🎮 [**Locales (i18n) Playground**](https://go.dev/play/p/9j9L8nI10MR) — Multilingual error messages in German, Spanish, French, and Japanese.
- 🎮 [**Standard net/http Handler Playground**](https://go.dev/play/p/M5cPYl4eDoJ) — Zero-dependency HTTP request validation.

Explore the [**examples/**](examples/) directory for full standalone project templates, including [Gin](examples/gin/) and [Echo](examples/echo/) integrations.

### Validating Structs

Checker can validate user input stored in a struct by listing the checkers in the struct tags for each field. Here is an example:

```golang
type Person struct {
	Name string `checkers:"trim required"`
}

person := &Person{
	Name: " Onur Cinar ",
}

errors, valid := checker.CheckStruct(person)
if !valid {
	// Handle validation errors
}
```

If a field has no `checkers` tag, `CheckStruct` and `JSONSchema` both fall back to reading a `validate` tag instead — the conventional tag name from `go-playground/validator`. This doesn't give Checker any understanding of `validator`'s own tag syntax (only the tag *name* is a fallback, not its contents), but it means a codebase already tagged `validate:"required"` with Checker-compatible rules picks up validation without renaming every tag. An explicit, empty `checkers:""` tag is respected as "no checks" rather than falling back.

### Validating Individual Values

You can also validate individual user input by calling checker functions directly:

```golang
name := " Onur Cinar "

name, err := checker.Check(name, checker.Trim, checker.Required)
if err != nil {
	// Handle validation error
}
```

The checkers and normalizers can also be provided through a config string:

```golang
name := " Onur Cinar "

name, err := checker.CheckWithConfig(name, "trim required")
if err != nil {
	// Handle validation error
}
```

For simpler validation, you can call individual checker functions:

```golang
name := "Onur Cinar"

err := checker.IsRequired(name)
if err != nil {
	// Handle validation error
}
```

## Normalizers and Checkers

Checkers validate user input, while normalizers transform it into a preferred format. For example, a normalizer can trim spaces from a string or convert it to title case.

Although combining checkers and normalizers into a single library might seem unconventional, using them together can be beneficial. They can be mixed in any order when defining validation steps. For instance, you can use the `trim` normalizer with the `required` checker to first trim the input and then ensure it is provided:

```golang
type Person struct {
	Name string `checkers:"trim required"`
}
```

## Normalizers Provided

Normalizers mutate string values in-place and pass them to the next checker in the chain:

| Category | Normalizer | Tag | Description |
| :--- | :--- | :--- | :--- |
| **Whitespace** | [`TrimSpace`](https://pkg.go.dev/github.com/cinar/checker/v2#TrimSpace) | `trim` | Trims whitespace from both sides of the string |
| | [`TrimLeft`](https://pkg.go.dev/github.com/cinar/checker/v2#TrimLeft) | `trim-left` | Trims whitespace from the left side of the string |
| | [`TrimRight`](https://pkg.go.dev/github.com/cinar/checker/v2#TrimRight) | `trim-right` | Trims whitespace from the right side of the string |
| **Case** | [`Lower`](https://pkg.go.dev/github.com/cinar/checker/v2#Lower) | `lower` | Converts the string to lowercase |
| | [`Upper`](https://pkg.go.dev/github.com/cinar/checker/v2#Upper) | `upper` | Converts the string to uppercase |
| | [`Title`](https://pkg.go.dev/github.com/cinar/checker/v2#Title) | `title` | Converts the string to title case |
| **Sanitization** | [`HTMLEscape`](https://pkg.go.dev/github.com/cinar/checker/v2#HTMLEscape) | `html-escape` | Escapes special characters in the string for HTML |
| | [`HTMLUnescape`](https://pkg.go.dev/github.com/cinar/checker/v2#HTMLUnescape) | `html-unescape` | Unescapes special characters in the string for HTML |
| | [`URLEscape`](https://pkg.go.dev/github.com/cinar/checker/v2#URLEscape) | `url-escape` | Escapes special characters in the string for URLs |
| | [`URLUnescape`](https://pkg.go.dev/github.com/cinar/checker/v2#URLUnescape) | `url-unescape` | Unescapes special characters in the string for URLs |
| **Security** | [`StripInvisible`](https://pkg.go.dev/github.com/cinar/checker/v2#StripInvisible) | `strip-invisible` | Removes zero-width and bidirectional control characters (the "Trojan Source" spoofing technique, CVE-2021-42574), which an attacker can use to split a keyword-filtered word invisibly or make displayed text diverge from its logical order |
| **Fallback** | [`Default`](https://pkg.go.dev/github.com/cinar/checker/v2#Default) | `default:<value>` | Replaces the value with `<value>` if it's currently its zero value, otherwise leaves it untouched |

Unlike the other normalizers, `default` isn't string-only: it also works on `bool`, the `int`/`uint`/`float` kinds, and a pointer to any of those (a nil pointer gets a freshly allocated default; a non-nil one is left alone). It's a poor fit alongside `omitempty` on the same field — `omitempty` skips every remaining check exactly when the value is zero, which is exactly the case `default` exists to handle, so combining the two on one field means `default` never runs.

`strip-invisible` removes zero-width space, zero-width non-joiner, zero-width joiner, word joiner, the zero-width no-break space (BOM), and the bidirectional embedding/override/isolate control characters. Some of these have legitimate uses in ordinary text — zero-width joiner in emoji sequences, zero-width non-joiner in Persian and other scripts — so apply it only to fields where an invisible character is never expected, such as a handle, username, or search keyword, not general free-text content.

`strip-invisible` handles invisible spoofing characters without pulling in a dependency. For *visible* lookalikes — fullwidth digits, ligatures, and other compatibility characters a naive keyword filter or uniqueness check would treat as distinct from their canonical form — see the opt-in [`nfkc`](#unicode-normalization-nfkc) module instead.

## Checkers Provided

Checkers validate that a value conforms to expected criteria, returning an error if validation fails.

### Strings & Characters

| Checker | Tag Syntax | Description |
| :--- | :--- | :--- |
| [`IsASCII`](https://pkg.go.dev/github.com/cinar/checker/v2#IsASCII) | `ascii` | Ensures the string contains only ASCII characters |
| [`IsAlphanumeric`](https://pkg.go.dev/github.com/cinar/checker/v2#IsAlphanumeric) | `alphanumeric` | Ensures the string contains only letters and numbers |
| [`IsAlpha`](https://pkg.go.dev/github.com/cinar/checker/v2#IsAlpha) | `alpha` | Ensures the string contains only letters |
| [`IsNumeric`](https://pkg.go.dev/github.com/cinar/checker/v2#IsNumeric) | `numeric` | Ensures the string is a valid numeric string, e.g. `"-3.14"` (unlike `digits`, allows a leading sign and a decimal point) |
| [`IsDigits`](https://pkg.go.dev/github.com/cinar/checker/v2#IsDigits) | `digits` | Ensures the string contains only digits |
| [`IsHex`](https://pkg.go.dev/github.com/cinar/checker/v2#IsHex) | `hex` | Ensures the string contains only hexadecimal digits |
| [`IsHexColor`](https://pkg.go.dev/github.com/cinar/checker/v2#IsHexColor) | `hex-color` | Ensures the string is a valid `#`-prefixed hex color code (3, 4, 6, or 8 digits) |
| [`IsBase64`](https://pkg.go.dev/github.com/cinar/checker/v2#IsBase64) | `base64` | Ensures the string is a valid standard (RFC 4648) base64-encoded string |
| [`IsBase64URL`](https://pkg.go.dev/github.com/cinar/checker/v2#IsBase64URL) | `base64-url` | Ensures the string is a valid base64url-encoded string |
| [`IsSlug`](https://pkg.go.dev/github.com/cinar/checker/v2#IsSlug) | `slug` | Ensures the string is a valid URL slug, e.g. `"hello-world"` |
| [`IsContains`](https://pkg.go.dev/github.com/cinar/checker/v2#IsContains) | `contains:<substr>` | Ensures the string contains the given substring |
| [`IsStartsWith`](https://pkg.go.dev/github.com/cinar/checker/v2#IsStartsWith) | `starts-with:<prefix>` | Ensures the string starts with the given prefix |
| [`IsEndsWith`](https://pkg.go.dev/github.com/cinar/checker/v2#IsEndsWith) | `ends-with:<suffix>` | Ensures the string ends with the given suffix |
| [`MakeRegexpChecker`](https://pkg.go.dev/github.com/cinar/checker/v2#MakeRegexpChecker) | `regexp:<pattern>` | Ensures the string matches the pattern |

### Presence, Sizes & Numeric Bounds

| Checker | Tag Syntax | Description |
| :--- | :--- | :--- |
| [`Required`](https://pkg.go.dev/github.com/cinar/checker/v2#Required) | `required` | Ensures the value is provided |
| [`MinLen`](https://pkg.go.dev/github.com/cinar/checker/v2#MinLen) | `min-len:<n>` | Ensures the length of the given value (string, slice, or map) is at least n |
| [`MaxLen`](https://pkg.go.dev/github.com/cinar/checker/v2#MaxLen) | `max-len:<n>` | Ensures the length of the given value (string, slice, or map) is at most n |
| [`Len`](https://pkg.go.dev/github.com/cinar/checker/v2#Len) | `len:<n>` | Ensures the length of the given value (string, slice, or map) is exactly n |
| [`IsGte`](https://pkg.go.dev/github.com/cinar/checker/v2#IsGte) | `gte:<n>` | Ensures the value is greater than or equal to the specified number |
| [`IsLte`](https://pkg.go.dev/github.com/cinar/checker/v2#IsLte) | `lte:<n>` | Ensures the value is less than or equal to the specified number |
| [`IsGt`](https://pkg.go.dev/github.com/cinar/checker/v2#IsGt) | `gt:<n>` | Ensures the value is strictly greater than the specified number |
| [`IsLt`](https://pkg.go.dev/github.com/cinar/checker/v2#IsLt) | `lt:<n>` | Ensures the value is strictly less than the specified number |
| [`IsOneOf`](https://pkg.go.dev/github.com/cinar/checker/v2#IsOneOf) | `oneof:<v1>,<v2>,...` | Ensures the value equals one of a comma-separated list of allowed values |
| [`IsEq`](https://pkg.go.dev/github.com/cinar/checker/v2#IsEq) | `eq:<value>` | Ensures the value equals the given literal value |
| [`IsNe`](https://pkg.go.dev/github.com/cinar/checker/v2#IsNe) | `ne:<value>` | Ensures the value does not equal the given literal value |
| [`IsInt`](https://pkg.go.dev/github.com/cinar/checker/v2#IsInt) | `int` | Ensures the value is a whole number, with no fractional part |
| [`IsPositive`](https://pkg.go.dev/github.com/cinar/checker/v2#IsPositive) | `positive` | Ensures the value is strictly greater than zero |
| [`IsNegative`](https://pkg.go.dev/github.com/cinar/checker/v2#IsNegative) | `negative` | Ensures the value is strictly less than zero |
| [`IsNonnegative`](https://pkg.go.dev/github.com/cinar/checker/v2#IsNonnegative) | `nonnegative` | Ensures the value is greater than or equal to zero |
| [`IsMultipleOf`](https://pkg.go.dev/github.com/cinar/checker/v2#IsMultipleOf) | `multiple-of:<n>` | Ensures the value is a multiple of n, within a small floating-point tolerance |
| [`IsFinite`](https://pkg.go.dev/github.com/cinar/checker/v2#IsFinite) | `finite` | Ensures the value is neither `NaN` nor an infinity |

### Field-Relative & Conditional (Structs)

| Checker | Tag Syntax | Description |
| :--- | :--- | :--- |
| [`IsEqField`](https://pkg.go.dev/github.com/cinar/checker/v2#IsEqField) | `eq-field:<field>` | Ensures the value is equal to the value of another field on the struct |
| [`IsRequiredIf`](https://pkg.go.dev/github.com/cinar/checker/v2#IsRequiredIf) | `required-if:<field>:<val>` | Ensures the value is provided when another field is equal to a given value |
| [`IsRequiredUnless`](https://pkg.go.dev/github.com/cinar/checker/v2#IsRequiredUnless) | `required-unless:<field>:<val>` | Ensures the value is provided unless another field is equal to a given value |
| [`IsBeforeField`](https://pkg.go.dev/github.com/cinar/checker/v2#IsBeforeField) | `before-field:<layout>:<field>` | Ensures the value is a time before another field on the struct |
| [`IsAfterField`](https://pkg.go.dev/github.com/cinar/checker/v2#IsAfterField) | `after-field:<layout>:<field>` | Ensures the value is a time after another field on the struct |

### Network & Web

| Checker | Tag Syntax | Description |
| :--- | :--- | :--- |
| [`IsEmail`](https://pkg.go.dev/github.com/cinar/checker/v2#IsEmail) | `email` | Ensures the string is a valid email address |
| [`IsURL`](https://pkg.go.dev/github.com/cinar/checker/v2#IsURL) | `url` | Ensures the string is a valid URL |
| [`IsFQDN`](https://pkg.go.dev/github.com/cinar/checker/v2#IsFQDN) | `fqdn` | Ensures the string is a valid fully qualified domain name |
| [`IsIP`](https://pkg.go.dev/github.com/cinar/checker/v2#IsIP) | `ip` | Ensures the string is a valid IP address |
| [`IsIPv4`](https://pkg.go.dev/github.com/cinar/checker/v2#IsIPv4) | `ipv4` | Ensures the string is a valid IPv4 address |
| [`IsIPv6`](https://pkg.go.dev/github.com/cinar/checker/v2#IsIPv6) | `ipv6` | Ensures the string is a valid IPv6 address |
| [`IsCIDR`](https://pkg.go.dev/github.com/cinar/checker/v2#IsCIDR) | `cidr` | Ensures the string is a valid CIDR notation |
| [`IsMAC`](https://pkg.go.dev/github.com/cinar/checker/v2#IsMAC) | `mac` | Ensures the string is a valid MAC address |
| [`IsE164`](https://pkg.go.dev/github.com/cinar/checker/v2#IsE164) | `e164` | Ensures the string is a valid E.164 phone number, e.g. `"+14155552671"` |

### Dates & Times

| Checker | Tag Syntax | Description |
| :--- | :--- | :--- |
| [`IsTime`](https://pkg.go.dev/github.com/cinar/checker/v2#IsTime) | `time:<layout>` | Ensures the string matches the provided time layout |
| [`IsAfter`](https://pkg.go.dev/github.com/cinar/checker/v2#IsAfter) | `after:<layout>:<time>` | Ensures the value is a time after the given reference time |
| [`IsBefore`](https://pkg.go.dev/github.com/cinar/checker/v2#IsBefore) | `before:<layout>:<time>` | Ensures the value is a time before the given reference time |

### Identifiers, Cryptography & Standards

| Checker | Tag Syntax | Description |
| :--- | :--- | :--- |
| [`IsAnyCreditCard`](https://pkg.go.dev/github.com/cinar/checker/v2#IsAnyCreditCard) | `credit-card` | Ensures the string is a valid credit card number |
| [`IsLUHN`](https://pkg.go.dev/github.com/cinar/checker/v2#IsLUHN) | `luhn` | Ensures the string is a valid LUHN number |
| [`IsHash`](https://pkg.go.dev/github.com/cinar/checker/v2#IsHash) | `hash:<algo>` | Ensures the string is a valid hex hash (`md5`, `sha1`, `sha256`, `sha384`, `sha512`) |
| [`IsEOA`](https://pkg.go.dev/github.com/cinar/checker/v2#IsEOA) | `eoa` | Ensures the string is a valid Ethereum externally owned address (EOA) |
| [`IsISBN`](https://pkg.go.dev/github.com/cinar/checker/v2#IsISBN) | `isbn` | Ensures the string is a valid ISBN |
| [`IsISO31661Alpha2`](https://pkg.go.dev/github.com/cinar/checker/v2#IsISO31661Alpha2) | `iso3166-1-alpha-2` | Ensures the string is a valid 2-letter ISO 3166-1 alpha-2 country code |
| [`IsISO31661Alpha3`](https://pkg.go.dev/github.com/cinar/checker/v2#IsISO31661Alpha3) | `iso3166-1-alpha-3` | Ensures the string is a valid 3-letter ISO 3166-1 alpha-3 country code |
| [`IsISO6391`](https://pkg.go.dev/github.com/cinar/checker/v2#IsISO6391) | `iso639-1` | Ensures the string is a valid 2-letter ISO 639-1 language code |
| [`IsPostalCode`](https://pkg.go.dev/github.com/cinar/checker/v2#IsPostalCode) | `postal-code:<country>` | Ensures the string matches the postal code format for the given ISO 3166-1 alpha-2 country code (e.g. `postal-code:US`); covers a curated set of common countries, not every country's postal system |
| [`IsUUID`](https://pkg.go.dev/github.com/cinar/checker/v2#IsUUID) | `uuid` | Ensures the string is a valid RFC 4122 UUID (any version), case-insensitive |
| [`IsULID`](https://pkg.go.dev/github.com/cinar/checker/v2#IsULID) | `ulid` | Ensures the string is a valid ULID |
| [`IsIBAN`](https://pkg.go.dev/github.com/cinar/checker/v2#IsIBAN) | `iban` | Ensures the string is a valid IBAN, verifying the ISO 7064 mod 97-10 check digits |
| [`IsMongoID`](https://pkg.go.dev/github.com/cinar/checker/v2#IsMongoID) | `mongo-id` | Ensures the string is a valid 24-character MongoDB ObjectID |
| [`IsSemver`](https://pkg.go.dev/github.com/cinar/checker/v2#IsSemver) | `semver` | Ensures the string is a valid Semantic Versioning 2.0.0 version |
| [`IsJWT`](https://pkg.go.dev/github.com/cinar/checker/v2#IsJWT) | `jwt` | Ensures the string has the structural shape of a JWT (does not verify the signature) |

## Custom Checkers and Normalizers

You can define custom checkers or normalizers and register them for use in your validation logic. Here is an example of how to create and register a custom checker:

```golang
checker.RegisterMaker("is-fruit", func(params string) checker.CheckFunc[reflect.Value] {
	return func(value reflect.Value) (reflect.Value, error) {
		stringValue := value.Interface().(string)

		if stringValue == "apple" || stringValue == "banana" {
			return value, nil
		}

		return value, checker.NewCheckError("NOT_FRUIT")
	}
})
```

In this example, the custom checker `is-fruit` checks if the input value is either "apple" or "banana". If the value is not one of these, it returns an error.

Once registered, you can use your custom checker in struct tags just like the built-in checkers:

```golang
type Item struct {
	Name string `checkers:"is-fruit"`
}

item := &Item{
	Name: "banana",
}

errors, valid := checker.CheckStruct(item)
if !valid {
	fmt.Println(errors)
}
```

In this example, the `is-fruit` checker is used to validate that the `Name` field of the `Item` struct is either "apple" or "banana".

## Slice and Item Level Checkers

When adding checker struct tags to a slice, you can use the `@` prefix to indicate that the checker should be applied to the slice itself. Checkers without the `@` prefix will be applied to the individual items within the slice. Here is an example:

```golang
type Person struct {
	Name   string   `checkers:"required"`
	Emails []string `checkers:"@max-len:2 max-len:64"`
}
```

In this example:
- `@max-len:2` ensures that the `Emails` slice itself has at most two items.
- `max-len:64` ensures that each email string within the `Emails` slice has a maximum length of 64 characters.

Maps are supported the same way, with the `@` prefix applying to the map itself and item-level checkers applying to each value in the map. Nested structs and pointers held in a map are also checked, and normalizers are written back into the map.

```golang
type Person struct {
	Name   string            `checkers:"required"`
	Emails map[string]string `checkers:"@max-len:2 trim max-len:64"`
}
```

## Optional Fields with `omitempty`

Add `omitempty` to a field's tag to skip every other checker in that tag when the field is its zero value (`""`, `0`, `nil`, an empty slice/map, ...), while still validating it normally when a value is provided:

```golang
type Profile struct {
	Website string `checkers:"omitempty url"`
}
```

An empty `Website` is left alone; a non-empty one still has to be a valid URL. `omitempty` looks at the field's original value, not a value already transformed by an earlier normalizer in the same tag, so `trim omitempty required` on `"   "` still fails `required` after trimming — a whitespace-only string isn't the zero value to begin with.

With the `@` prefix, `omitempty` applies to a slice or map container itself rather than its items, e.g. `@omitempty @min-len:1` skips the container-level check on a nil slice.

Pairing `omitempty` with `required` on the same field is a contradiction — "optional" and "required" can't both hold — so pipeline-wise `omitempty` wins and `required` never runs on a zero value; avoid writing the two together.

## Field-Relative and Conditional Checkers

Some checkers compare a field's value against another field on the same struct. As a `checkers:"..."` tag, these are only available through `CheckStruct`, since resolving the named sibling field relies on the parent struct being known at runtime. Each one also has a plain, non-reflection function — [`IsEqField`](https://pkg.go.dev/github.com/cinar/checker/v2#IsEqField), [`IsRequiredIf`](https://pkg.go.dev/github.com/cinar/checker/v2#IsRequiredIf), [`IsRequiredUnless`](https://pkg.go.dev/github.com/cinar/checker/v2#IsRequiredUnless), [`IsBeforeField`](https://pkg.go.dev/github.com/cinar/checker/v2#IsBeforeField), [`IsAfterField`](https://pkg.go.dev/github.com/cinar/checker/v2#IsAfterField) — that takes the sibling's already-resolved value directly, the same shape as every other `IsX` checker (`IsEmail`, `IsGte`, ...); it's for calling standalone or from generated code, not for use in a `checkers` tag.

- `eq-field:FieldName` ensures the value is equal to the named sibling field's value. This is useful for a password confirmation field.
- `required-if:FieldName:Value` ensures the value is provided when the named sibling field is equal to the given value.
- `required-unless:FieldName:Value` ensures the value is provided unless the named sibling field is equal to the given value.
- `before-field:Layout:FieldName` / `after-field:Layout:FieldName` ensure the value, parsed using the given time layout, is before/after the named sibling field's value, parsed using the same layout — the field-relative equivalent of `before`/`after`. Unlike `before`/`after`, an unparsable sibling value returns an error rather than panicking, since it's struct data rather than checker configuration.

```golang
type Registration struct {
	Password        string `checkers:"required"`
	ConfirmPassword string `checkers:"eq-field:Password"`

	Country string `checkers:"required"`
	State   string `checkers:"required-if:Country:US"`
}

type Trip struct {
	ReturnAt string `checkers:"required"`
	DepartAt string `checkers:"before-field:DateOnly:ReturnAt"`
}
```

## Programmatic Pipelines (Context-Aware Validation)

Struct tags are great for a flat DTO's shape, but they can't express rules that need request-scoped state — a database uniqueness check, a tenant boundary, an auth claim — since there's no way to pass a `context.Context` through a tag. `Pipeline[T]` is a small, fully opt-in, generic builder for exactly that: it reuses the same checker/normalizer functions as struct tags via `Field`, and adds `Rule` for whole-value, context-aware domain checks. Combine it with `CheckStruct` freely on the same type — one validates shape, the other validates domain rules that need `ctx`.

```golang
type User struct {
	Email         string
	Role          string
	HasMFAEnabled bool
}

pipeline := checker.NewPipeline[User]().Step(
	checker.Field("Email", func(u *User) *string { return &u.Email },
		checker.TrimSpace, checker.Lower, checker.Required, checker.IsEmail,
	),
	checker.Rule("MFA", func(ctx context.Context, u *User) error {
		if u.Role == "admin" && !u.HasMFAEnabled {
			return checker.NewCheckError("MFA_REQUIRED_FOR_ADMIN")
		}
		return nil
	}),
)

errs, ok := pipeline.Validate(ctx, user)
```

`Field` normalizes and validates one field in place, exactly like a `checkers` tag chain: checks run in order and stop at the first error, and any normalizer among them writes its result back before the next check runs. `Rule` runs against the whole value with `ctx`, for anything a single field's tag can't express. `Validate` runs every step regardless of earlier failures — matching `CheckStruct`'s field-independent error collection — and returns the same `CheckErrors` type, so both validation styles produce API-ready errors the same way.

### Context-Aware Struct Tags with `CheckStructWithContext`

`Pipeline[T]` is the right tool when validation is mostly programmatic. If a struct is otherwise validated entirely through `checkers` tags and only one or two fields need `ctx`, a `checkersCtx` tag avoids splitting that struct's rules across two places. `CheckStructWithContext` runs exactly like `CheckStruct`, and additionally runs each field's `checkersCtx` tag against a `context.Context`, using a checker registered with [RegisterCtxMaker](https://pkg.go.dev/github.com/cinar/checker/v2#RegisterCtxMaker):

```golang
checker.RegisterCtxMaker("unique-email", func(_ string) checker.CheckFuncCtx[reflect.Value] {
	return func(ctx context.Context, value reflect.Value) (reflect.Value, error) {
		if db.EmailTaken(ctx, value.String()) {
			return value, checker.NewCheckError("EMAIL_TAKEN")
		}
		return value, nil
	}
})

type SignupRequest struct {
	Email string `checkers:"required email" checkersCtx:"unique-email"`
}

errs, ok := checker.CheckStructWithContext(ctx, req)
```

A field's `checkersCtx` checks only run if its `checkers`/`validate` tag didn't already fail, the same way a single checker chain stops at its first error. `checkersCtx` is silently ignored by `CheckStruct` and `CheckWithConfig`, since neither has a `context.Context` to run it with, so the two tags coexist on the same field without conflict. `CheckWithContext` is also available as the context-aware counterpart of `Check`, for running a sequence of `CheckFuncCtx[T]` against a single value directly.

## Localized Error Messages

When validation fails, Checker returns an error. By default, the [Error()](https://pkg.go.dev/github.com/cinar/checker/v2#CheckError.Error) function provides a human-readable error message in `en-US` locale.

```golang
_, err := checker.IsEmail("abcd")
if err != nil {
	fmt.Println(err)
	// Output: Not a valid email address.
}
```

To get error messages in other languages, use the [ErrorWithLocale()](https://pkg.go.dev/github.com/cinar/checker/v2#CheckError.ErrorWithLocale) function. By default, only `en-US` is registered; the rest ship as data in the `locales` package and are opt-in via [RegisterLocale](https://pkg.go.dev/github.com/cinar/checker/v2#RegisterLocale), so importing `checker` never pulls in translations you don't use. Checker ships translations for:

`ar-SA`, `de-DE`, `en-US`, `es-ES`, `fa-IR`, `fr-FR`, `hy-AM`, `id-ID`, `it-IT`, `ja-JP`, `ko-KR`, `lv-LV`, `nl-NL`, `pl-PL`, `pt-BR`, `pt-PT`, `ru-RU`, `th-TH`, `tr-TR`, `uk-UA`, `vi-VN`, `zh-CN`, `zh-TW`

Register any of these the same way:

```golang
// Register de-DE localized error messages.
checker.RegisterLocale(locales.DeDE, locales.DeDEMessages)

_, err := checker.IsEmail("abcd")
if err != nil {
	fmt.Println(err.ErrorWithLocale(locales.DeDE))
	// Output: Keine gültige E-Mail-Adresse.
}
```

You can also customize existing error messages or add new ones to `locales.EnUSMessages` and other locale maps.

```golang
// Register the en-US localized error message for the custom NOT_FRUIT error code.
locales.EnUSMessages["NOT_FRUIT"] = "Not a fruit name."

errors, valid := checker.CheckStruct(item)
if !valid {
	fmt.Println(errors)
	// Output: map[Name:Not a fruit name.]
}
```

Error messages are generated using Go template functions, allowing them to include variables.

```golang
// Custom checker error containing the item name.
err := checker.NewCheckErrorWithData(
	"NOT_FRUIT",
	map[string]interface{}{
		"name": "abcd",
	},
)

// Register the en-US localized error message for the custom NOT_FRUIT error code.
locales.EnUSMessages["NOT_FRUIT"] = "Name {{ .name }} is not a fruit name."

errors, valid := checker.CheckStruct(item)
if !valid {
	fmt.Println(errors)
	// Output: map[Name:Name abcd is not a fruit name.]
}
```

## Custom Error Messages

Locale messages are shared by every field that uses a given checker. To override the wording for one field's checks without touching a locale or registering a whole new checker, add a `checkersMsg` tag alongside `checkers`: a semicolon-separated list of `name=message` pairs, keyed by the bare checker name (no `:params`).

```golang
type Signup struct {
	Email string `checkers:"required email" checkersMsg:"required=Email is required;email=Enter a valid email address"`
}
```

A custom message is rendered as a Go template against the failing check's own data, the same way a locale message is, so placeholders still work:

```golang
type Signup struct {
	Password string `checkers:"min-len:8" checkersMsg:"min-len=Must be at least {{ .min }} characters"`
}
```

For a slice or map field, prefix a container-level override with `@`, matching the `checkers` tag's own container-vs-item convention:

```golang
type Signup struct {
	Roles []string `checkers:"@min-len:1 required" checkersMsg:"@min-len=Pick at least one role;required=Role cannot be blank"`
}
```

A name with no matching check in `checkers`/`validate` is simply unused. A custom message bypasses locale lookup entirely for that occurrence — it's a literal string you wrote, not a translation key, so `ErrorWithLocale()`/`JSONWithLocale()` return it unchanged regardless of locale. It also can't contain a literal `;` (the pair separator). This is struct-tag-only for now; `Check`/`CheckWithConfig` for standalone values are unaffected.

## Structured Errors

`CheckStruct` returns [CheckErrors](https://pkg.go.dev/github.com/cinar/checker/v2#CheckErrors), a `map[string]error` that also implements the `error` interface, so it can be returned or wrapped directly like any other error.

```golang
errs, valid := checker.CheckStruct(person)
if !valid {
	return errs // errs.Error() joins every field's message
}
```

To respond to an API request, use [JSON()](https://pkg.go.dev/github.com/cinar/checker/v2#CheckErrors.JSON) to marshal the errors into a field name to `{code, message}` object, ready to be written directly as the response body:

```golang
errs, valid := checker.CheckStruct(person)
if !valid {
	data, _ := errs.JSON()
	w.WriteHeader(http.StatusBadRequest)
	w.Write(data)
	// {"Name":{"code":"REQUIRED","message":"Required value is missing."}}
	return
}
```

Use [JSONWithLocale()](https://pkg.go.dev/github.com/cinar/checker/v2#CheckErrors.JSONWithLocale) to localize the messages in the response, the same way `ErrorWithLocale()` works for a single error.

```golang
data, _ := errs.JSONWithLocale(locales.DeDE)
```

For an [RFC 9457](https://www.rfc-editor.org/rfc/rfc9457) `application/problem+json` response instead, use [ProblemDetails()](https://pkg.go.dev/github.com/cinar/checker/v2#CheckErrors.ProblemDetails):

```golang
errs, valid := checker.CheckStruct(person)
if !valid {
	data, _ := json.Marshal(errs.ProblemDetails())
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(http.StatusBadRequest)
	w.Write(data)
	// {"type":"about:blank","title":"Your request parameters failed validation.","status":400,
	//  "invalid-params":[{"name":"Name","reason":"Required value is missing.","code":"REQUIRED"}]}
	return
}
```

`Type`, `Title`, and `Status` are plain exported fields on the returned `*ProblemDetails` — RFC 9457 leaves their exact values to the API producer, so set them directly to override the defaults (`"about:blank"`, a generic title, `400`). `ProblemDetailsWithLocale()` localizes the messages, matching `JSONWithLocale()`.

## JSON Schema Generation

`JSONSchema` generates a [JSON Schema](https://json-schema.org/) document describing the shape and validation rules declared in a struct's `checkers` tags — useful for documenting an API, or handing a frontend enough information to mirror your validation rules without duplicating them by hand. It's a static, type-level operation: it reads a struct's type and tags, never its field values, so a zero value works as well as a populated one.

```golang
type Person struct {
	Name  string `json:"name" checkers:"trim required"`
	Email string `json:"email" checkers:"required email"`
}

schema := checker.JSONSchema(&Person{})

data, _ := json.MarshalIndent(schema, "", "  ")
fmt.Println(string(data))
```

```json
{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "title": "Person",
  "type": "object",
  "properties": {
    "email": { "type": "string", "format": "email" },
    "name": { "type": "string" }
  },
  "required": ["email", "name"]
}
```

Most checkers map directly onto a JSON Schema keyword: `required` becomes an entry in `required`, `min-len`/`max-len` become `minLength`/`maxLength` (or `minItems`/`maxItems` for a slice, `minProperties`/`maxProperties` for a map), `gte`/`lte` become `minimum`/`maximum`, `email`/`url`/`ipv4`/`ipv6`/`ip`/`cidr`/`mac`/`fqdn` become a `format`, `regexp:pattern`/`hex`/`alphanumeric`/`ascii`/`digits`/`hash:algorithm`/`postal-code:country` become `pattern`, and `oneof:a,b,c` becomes `enum`. Normalizers (`trim`, `lower`, `upper`, ...) are skipped, since they transform data rather than constrain its shape. A checker with no JSON Schema equivalent — a custom checker, or one like `eq-field` that compares against another field — is recorded in an `x-checker` vendor extension instead of being silently dropped:

```golang
type Registration struct {
	Password        string `checkers:"required"`
	ConfirmPassword string `checkers:"eq-field:Password"`
}
```

```json
{
  "ConfirmPassword": { "type": "string", "x-checker": ["eq-field:Password"] }
}
```

A `description` struct tag, independent of `checkers`, becomes the field's `description` keyword:

```golang
type Person struct {
	Email string `json:"email" checkers:"required email" description:"User's primary email address"`
}
```

```json
{ "email": { "type": "string", "format": "email", "description": "User's primary email address" } }
```

Register a [SchemaMakeFunc](https://pkg.go.dev/github.com/cinar/checker/v2#SchemaMakeFunc) with [RegisterSchemaMaker](https://pkg.go.dev/github.com/cinar/checker/v2#RegisterSchemaMaker) to teach `JSONSchema` how to translate a custom checker instead:

```golang
checker.RegisterSchemaMaker("is-fruit", func(schema *checker.Schema, _ string) {
	schema.Enum = []string{"apple", "banana"}
})
```

## Framework Integration

Checker ships thin, separately-versioned adapter modules that bind a request and run `CheckStruct` in a single call, writing a JSON `400` response automatically when binding or validation fails. Each adapter is its own Go module, so the framework it wraps is only pulled in if you actually `go get` that adapter; the core `checker` module stays dependency-free either way.

### Gin

```bash
go get github.com/cinar/checker/v2/gin
```

```golang
import checkergin "github.com/cinar/checker/v2/gin"

type Registration struct {
	Name  string `json:"name" checkers:"trim required"`
	Email string `json:"email" checkers:"required email"`
}

router.POST("/register", func(c *gin.Context) {
	var registration Registration

	if !checkergin.Bind(c, &registration) {
		// The 400 response has already been written by Bind.
		return
	}

	c.JSON(http.StatusOK, registration)
})
```

See [gin/README.md](gin/README.md) for the full example, including how to call `checkergin.Check` directly when the struct is assembled from more than just the request body.

### Echo

```bash
go get github.com/cinar/checker/v2/echo
```

```golang
import checkerecho "github.com/cinar/checker/v2/echo"

type Registration struct {
	Name  string `json:"name" checkers:"trim required"`
	Email string `json:"email" checkers:"required email"`
}

e.POST("/register", func(c echo.Context) error {
	var registration Registration

	if !checkerecho.Bind(c, &registration) {
		// The 400 response has already been written by Bind.
		return nil
	}

	return c.JSON(http.StatusOK, registration)
})
```

See [echo/README.md](echo/README.md) for the full example, including how to call `checkerecho.Check` directly when the struct is assembled from more than just the request body.

### net/http

```bash
go get github.com/cinar/checker/v2/nethttp
```

```golang
import checkernethttp "github.com/cinar/checker/v2/nethttp"

type Registration struct {
	Name  string `json:"name" checkers:"trim required"`
	Email string `json:"email" checkers:"required email"`
}

http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
	var registration Registration

	if !checkernethttp.Bind(w, r, &registration) {
		// The 400 response has already been written by Bind.
		return
	}

	json.NewEncoder(w).Encode(registration)
})
```

Unlike `gin`/`echo`, this adapter has no external dependency at all beyond the core `checker` module — only `encoding/json` and `net/http`, both standard library. See [nethttp/README.md](nethttp/README.md) for the full example, including how to call `checkernethttp.Check` directly when the struct is assembled from more than just the request body.

### Fiber

```bash
go get github.com/cinar/checker/v2/fiber
```

```golang
import checkerfiber "github.com/cinar/checker/v2/fiber"

type Registration struct {
	Name  string `json:"name" checkers:"trim required"`
	Email string `json:"email" checkers:"required email"`
}

app.Post("/register", func(c fiber.Ctx) error {
	var registration Registration

	if !checkerfiber.Bind(c, &registration) {
		// The 400 response has already been written by Bind.
		return nil
	}

	return c.JSON(registration)
})
```

See [fiber/README.md](fiber/README.md) for the full example, including how to call `checkerfiber.Check` directly when the struct is assembled from more than just the request body.

## Unicode Normalization (NFKC)

[`nfkc`](nfkc/), a separate module, registers an `nfkc` normalizer applying [Unicode Normalization Form KC](https://unicode.org/reports/tr15/): it composes combining character sequences and replaces compatibility characters — fullwidth digits, ligatures, and other stylistic variants — with their canonical equivalents, so `"ＡＬＩＣＥ"` and `"ALICE"` compare equal after normalization. It's kept out of the core module because it needs `golang.org/x/text/unicode/norm`; a blank import is enough to opt in.

```bash
go get github.com/cinar/checker/v2/nfkc
```

```golang
import (
	checker "github.com/cinar/checker/v2"
	_ "github.com/cinar/checker/v2/nfkc"
)

type Handle struct {
	Name string `checkers:"trim nfkc required min-len:3"`
}
```

See [nfkc/README.md](nfkc/README.md) for the full example, including calling `nfkc.Normalize` directly on a one-off value.

## Static Analysis

`checkers`/`validate` struct tags are string literals, so a typo'd checker name, a checker applied to a field of the wrong type, or a renamed cross-field target (`eq-field`, `after-field`, ...) all compile fine and only fail at runtime. [`checkerlint`](checkerlint/), a separate module, is a `go/analysis`-based static analyzer that catches all three at build/lint time instead:

```bash
go install github.com/cinar/checker/v2/checkerlint/cmd/checkerlint@latest
checkerlint ./...
```

See [checkerlint/README.md](checkerlint/README.md) for `go vet -vettool` and `golangci-lint` module-plugin integration.

## Code Generation

`CheckStruct` walks a struct with `reflect` on every call (a cached execution plan keeps this fast — see [Performance](#performance) below — but it's still reflection). [`checkergen`](checkergen/), a separate module, instead generates a `Check<Type>` function that calls the same checkers directly, with no reflection at runtime at all:

```bash
go get github.com/cinar/checker/v2/checkergen
```

```golang
//go:generate go run github.com/cinar/checker/v2/checkergen/cmd/checkergen

type SignupRequest struct {
	Email    string `json:"email" checkers:"trim lower required email"`
	Password string `json:"password" checkers:"required min-len:8"`
}
```

```bash
go generate ./...
```

Benchmarked against the equivalent `CheckStruct` call, on the same struct and input:

| Struct | `CheckStruct` | Generated | Speedup |
| :--- | ---: | ---: | :---: |
| 5-field signup form | 1830 ns/op, 1304 B/op, 24 allocs/op | 562 ns/op, 168 B/op, 7 allocs/op | **~3.3x faster, ~8x less memory** |
| 25-field mixed checker coverage | 8200 ns/op, 4928 B/op, 59 allocs/op | 2730 ns/op, 760 B/op, 9 allocs/op | **~3.0x faster, ~6.5x less memory** |

`checkergen` and `CheckStruct` are meant to coexist — generate code for the structs on your hot paths, leave everything else on `CheckStruct`. A struct with a field outside `checkergen`'s scope (a nested struct, a slice/map field, or a named type like `type Email string`, for now) is skipped with a clear reason instead of generated incorrectly. See [checkergen/README.md](checkergen/README.md) for the full scope, the exact benchmark setup, and how to reproduce the numbers above.

## Command-Line Interface

[`checker`](cli/), a separate module, is a standalone command-line interface to the library, for running any checker or normalizer from a shell script, CI pipeline, or Git hook without writing Go code. It ships as a single, dependency-free static binary.

```bash
go install github.com/cinar/checker/v2/cli/cmd/checker@latest
```

```bash
$ checker check email "user@example.com"
user@example.com
$ echo $?
0

$ echo "  Test@Example.com  " | checker check "trim lower email"
test@example.com

$ checker check --json email "not-an-email"
{"valid":false,"error":{"code":"NOT_EMAIL","message":"Not a valid email address."}}
```

`checker check <config> [value]` takes the exact same `checkers`/`validate` tag config string syntax used in a struct tag, so it never falls behind the checker vocabulary — it doesn't hardcode a checker list, it calls the same `CheckWithConfig`, `RegisteredMakerNames`, and `RegisteredFieldMakerNames` functions a Go caller would. `checker list` prints every registered name, and `--locale=<tag>` renders the error message in any of the 23 shipped locales. See [cli/README.md](cli/README.md) for the full command reference, flags, and a Git hook example.

## Performance

Checker is designed for low memory allocations and high throughput in HTTP request pipelines with zero external dependencies. `CheckStruct` compiles each struct type's `checkers`/`validate` tags into a resolved execution plan (field metadata, container/item splits, and maker-resolved checker chains) on first use and caches it, so repeated validation of the same struct type — or of any field sharing an identical tag string — doesn't re-parse tag strings or re-resolve checker names on every call. [`benchmark_test.go`](benchmark_test.go) covers `CheckStruct` on a small struct (success and failure paths), a larger struct (nested struct, slice, map, cross-field and conditional checkers), and static `JSONSchema` generation.

Run the benchmarks on your machine:

```bash
go test -bench=. -benchmem -run '^$'
```

Measured with `go test -bench=. -benchmem -benchtime=2s` on Linux x86_64 (Intel N100, Go 1.23.2). These are real, reproducible measurements from the benchmark file linked above, not hand-typed estimates — re-run the command yourself to confirm on your own hardware:

| Benchmark | Iterations | Time / Op | Memory / Op | Allocs / Op |
| :--- | :---: | :---: | :---: | :---: |
| `BenchmarkCheckStruct_Simple_Success` (3 fields, all valid) | ~1,940,000 | **1.2 µs/op** | 760 B/op | 19 allocs/op |
| `BenchmarkCheckStruct_Simple_Failure` (3 fields, one invalid) | ~1,760,000 | **1.4 µs/op** | 1,032 B/op | 19 allocs/op |
| `BenchmarkCheckStruct_Complex` (10 fields incl. nested struct, slice, map) | ~320,000 | **7.8 µs/op** | 3,576 B/op | 78 allocs/op |
| `BenchmarkJSONSchema` (static schema generation) | ~315,000 | **7.3 µs/op** | 6,716 B/op | 72 allocs/op |

## Changelog

See [CHANGELOG.md](CHANGELOG.md) for a history of notable changes to this project.

## Contributing to the Project

Anyone can contribute to Checkers library. Please make sure to read our [Contributor Covenant Code of Conduct](./CODE_OF_CONDUCT.md) guide first. Follow the [How to Contribute to Checker](./CONTRIBUTING.md) to contribute.

## License

Checker is provided under the MIT License, reproduced below and also available in the [LICENSE](./LICENSE) file.

```
Copyright (c) 2023-2026 Onur Cinar.
The source code is provided under MIT License.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
