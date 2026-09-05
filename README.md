<p align="center">
    <a href="https://pkg.go.dev/github.com/cinar/checker/v2"><img src="https://img.shields.io/badge/Go_Reference-007D9C?style=for-the-badge&logo=go&logoColor=white" alt="Go Reference" /></a>
    <a href="LICENSE"><img src="https://img.shields.io/github/license/cinar/checker?style=for-the-badge" alt="License" /></a>
    <a href="https://github.com/cinar/checker/actions/workflows/ci.yml"><img src="https://img.shields.io/github/actions/workflow/status/cinar/checker/ci.yml?branch=main&style=for-the-badge&logo=github&label=CI" alt="Go CI" /></a>
    <a href="https://codecov.io/gh/cinar/checker"><img src="https://img.shields.io/codecov/c/github/cinar/checker?style=for-the-badge&logo=codecov&logoColor=white" alt="Codecov" /></a>
    <a href="https://github.com/cinar/checker/stargazers"><img src="https://img.shields.io/github/stars/cinar/checker?style=for-the-badge&logo=github&logoColor=white" alt="GitHub Stars" /></a>
</p>

<p align="center">
    <img src="logo.svg" alt="Checker logo: a tag with a punched checkmark" width="96" height="96" />
</p>

<h1 align="center">Checker</h1>

<p align="center">Checker is a lightweight Go library for validating and normalizing user input, driven by struct tags or plain function calls, with zero external dependencies.</p>

- **Zero dependencies** — the core module imports nothing beyond the Go standard library.
- **Struct tags or plain functions** — validate a whole struct declaratively, or call checkers directly for one-off values.
- **Checkers and normalizers together** — trim, then require; lowercase, then validate — mixed in any order, in one pass.
- **Cross-field and conditional rules** — compare fields against each other, or require a field only when another has a given value.
- **23 built-in locales** — opt-in, translated error messages, matching the set go-playground/validator ships.
- **JSON Schema generation** — turn a struct's checker tags into a JSON Schema document, for API docs or frontend validation, without hand-maintaining a second copy of your rules.
- **Framework adapters** — thin, separately-versioned `gin` and `echo` modules bind a request and validate it in one call.

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
- [Localized Error Messages](#localized-error-messages)
- [Structured Errors](#structured-errors)
- [JSON Schema Generation](#json-schema-generation)
- [Framework Integration](#framework-integration)
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
| [`IsUUID`](https://pkg.go.dev/github.com/cinar/checker/v2#IsUUID) | `uuid` | Ensures the string is a valid RFC 4122 UUID (any version), case-insensitive |

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

Some checkers compare a field's value against another field on the same struct. These are only available through `CheckStruct`, since they rely on the parent struct being known.

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

Most checkers map directly onto a JSON Schema keyword: `required` becomes an entry in `required`, `min-len`/`max-len` become `minLength`/`maxLength` (or `minItems`/`maxItems` for a slice, `minProperties`/`maxProperties` for a map), `gte`/`lte` become `minimum`/`maximum`, `email`/`url`/`ipv4`/`ipv6`/`fqdn` become a `format`, and `regexp:pattern` becomes `pattern`. Normalizers (`trim`, `lower`, `upper`, ...) are skipped, since they transform data rather than constrain its shape. A checker with no JSON Schema equivalent — a custom checker, or one like `eq-field` that compares against another field — is recorded in an `x-checker` vendor extension instead of being silently dropped:

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

## Performance

Checker is designed for low memory allocations and high throughput in HTTP request pipelines with zero external dependencies. [`benchmark_test.go`](benchmark_test.go) covers `CheckStruct` on a small struct (success and failure paths), a larger struct (nested struct, slice, map, cross-field and conditional checkers), and static `JSONSchema` generation.

Run the benchmarks on your machine:

```bash
go test -bench=. -benchmem -run '^$'
```

Measured with `go test -bench=. -benchmem -benchtime=2s` on Linux x86_64 (Intel N100, Go 1.23.2). These are real, reproducible measurements from the benchmark file linked above, not hand-typed estimates — re-run the command yourself to confirm on your own hardware:

| Benchmark | Iterations | Time / Op | Memory / Op | Allocs / Op |
| :--- | :---: | :---: | :---: | :---: |
| `BenchmarkCheckStruct_Simple_Success` (3 fields, all valid) | ~870,000 | **2.6 µs/op** | 1,176 B/op | 37 allocs/op |
| `BenchmarkCheckStruct_Simple_Failure` (3 fields, one invalid) | ~880,000 | **2.7 µs/op** | 1,448 B/op | 37 allocs/op |
| `BenchmarkCheckStruct_Complex` (10 fields incl. nested struct, slice, map) | ~160,000 | **15.0 µs/op** | 5,993 B/op | 181 allocs/op |
| `BenchmarkJSONSchema` (static schema generation) | ~340,000 | **7.3 µs/op** | 6,716 B/op | 72 allocs/op |

## Changelog

See [CHANGELOG.md](CHANGELOG.md) for a history of notable changes to this project.

## Contributing to the Project

Anyone can contribute to Checkers library. Please make sure to read our [Contributor Covenant Code of Conduct](./CODE_OF_CONDUCT.md) guide first. Follow the [How to Contribute to Checker](./CONTRIBUTING.md) to contribute.

## License

This library is free to use, modify, and distribute under the terms of the MIT license. The full license text can be found in the [LICENSE](./LICENSE) file.

The MIT license is a permissive license that allows you to do almost anything with the library, as long as you retain the copyright notice and the license text. This means that you can use the library in commercial products, modify it, and redistribute it without having to ask for permission from the authors.

The [LICENSE](./LICENSE) file is located in the root directory of the library. You can open it in a text editor to read the full license text.
