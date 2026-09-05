[![Go Reference](https://pkg.go.dev/badge/github.com/cinar/checker/v2.svg)](https://pkg.go.dev/github.com/cinar/checker/v2)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/cinar/checker)](https://goreportcard.com/report/github.com/cinar/checker)
![Go CI](https://github.com/cinar/checker/actions/workflows/ci.yml/badge.svg)
[![codecov](https://codecov.io/gh/cinar/checker/branch/main/graph/badge.svg?token=VO9BYBHJHE)](https://codecov.io/gh/cinar/checker)

# Checker

Checker is a lightweight Go library designed to validate user input efficiently. It supports validation of both struct fields and individual input values.

While there are numerous validation libraries available, Checker stands out due to its simplicity and lack of external dependencies. This makes it an ideal choice for developers who prefer to minimize dependencies and maintain control over their tools. Checker is straightforward to use and effectively meets your validation needs.

## Usage

To begin using the Checker library, install it with the following command:

```bash
go get github.com/cinar/checker/v2
```

Then, import the library into your source file as shown below:

```golang
import (
	checker "github.com/cinar/checker/v2"
)
```

### Validating User Input Stored in a Struct

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

### Validating Individual User Input with Multiple Checkers

You can also validate individual user input by calling checker functions directly. Here is an example:

```golang
name := " Onur Cinar "

name, err := checker.Check(name, checker.Trim, checker.Required)
if err != nil {
	// Handle validation error
}
```

The checkers and normalizers can also be provided through a config string. Here is an example:

```golang
name := " Onur Cinar "

name, err := checker.CheckWithConfig(name, "trim requied")
if err != nil {
	// Handle validation error
}

```

### Validating Individual User Input

For simpler validation, you can call individual checker functions. Here is an example:

```golang
name := "Onur Cinar"

err := checker.IsRequired(name)
if err != nil {
	// Handle validation error
}
```

## Normalizers and Checkers

Checkers validate user input, while normalizers transform it into a preferred format. For example, a normalizer can trim spaces from a string or convert it to title case.

Although combining checkers and normalizers into a single library might seem unconventional, using them together can be beneficial. They can be mixed in any order when defining validation steps. For instance, you can use the `trim` normalizer with the `required` checker to first trim the input and then ensure it is provided. Here is an example:

```golang
type Person struct {
	Name string `checkers:"trim required"`
}
```

# Checkers Provided

- [`after`](https://pkg.go.dev/github.com/cinar/checker/v2#IsAfter): Ensures the value is a time after the given reference time, e.g. `after:DateOnly:2024-01-01`.
- [`after-field`](https://pkg.go.dev/github.com/cinar/checker/v2#IsAfterField): Ensures the value is a time after the value of another field on the struct, e.g. `after-field:DateOnly:BornAt`.
- [`ascii`](https://pkg.go.dev/github.com/cinar/checker/v2#IsASCII): Ensures the string contains only ASCII characters.
- [`alphanumeric`](https://pkg.go.dev/github.com/cinar/checker/v2#IsAlphanumeric): Ensures the string contains only letters and numbers.
- [`before`](https://pkg.go.dev/github.com/cinar/checker/v2#IsBefore): Ensures the value is a time before the given reference time, e.g. `before:DateOnly:2024-01-01`.
- [`before-field`](https://pkg.go.dev/github.com/cinar/checker/v2#IsBeforeField): Ensures the value is a time before the value of another field on the struct, e.g. `before-field:DateOnly:ReturnAt`.
- [`credit-card`](https://pkg.go.dev/github.com/cinar/checker/v2#IsAnyCreditCard): Ensures the string is a valid credit card number.
- [`cidr`](https://pkg.go.dev/github.com/cinar/checker/v2#IsCIDR): Ensures the string is a valid CIDR notation.
- [`digits`](https://pkg.go.dev/github.com/cinar/checker/v2#IsDigits): Ensures the string contains only digits.
- [`email`](https://pkg.go.dev/github.com/cinar/checker/v2#IsEmail): Ensures the string is a valid email address.
- [`eoa`](https://pkg.go.dev/github.com/cinar/checker/v2#IsEOA): Ensures the string is a valid externally owned address (EOA), i.e. an Ethereum address.
- [`eq-field`](https://pkg.go.dev/github.com/cinar/checker/v2#IsEqField): Ensures the value is equal to the value of another field on the struct.
- [`fqdn`](https://pkg.go.dev/github.com/cinar/checker/v2#IsFQDN): Ensures the string is a valid fully qualified domain name.
- [`gte`](https://pkg.go.dev/github.com/cinar/checker/v2#IsGte): Ensures the value is greater than or equal to the specified number.
- [`hash`](https://pkg.go.dev/github.com/cinar/checker/v2#IsHash): Ensures the string is a valid hex-encoded hash for the given algorithm (`md5`, `sha1`, `sha256`, `sha384`, or `sha512`), e.g. `hash:sha256`.
- [`hex`](https://pkg.go.dev/github.com/cinar/checker/v2#IsHex): Ensures the string contains only hexadecimal digits.
- [`ip`](https://pkg.go.dev/github.com/cinar/checker/v2#IsIP): Ensures the string is a valid IP address.
- [`ipv4`](https://pkg.go.dev/github.com/cinar/checker/v2#IsIPv4): Ensures the string is a valid IPv4 address.
- [`ipv6`](https://pkg.go.dev/github.com/cinar/checker/v2#IsIPv6): Ensures the string is a valid IPv6 address.
- [`isbn`](https://pkg.go.dev/github.com/cinar/checker/v2#IsISBN): Ensures the string is a valid ISBN.
- [`iso3166-1-alpha-2`](https://pkg.go.dev/github.com/cinar/checker/v2#IsISO31661Alpha2): Ensures the string is a valid two-letter ISO 3166-1 alpha-2 country code, e.g. `US`.
- [`iso3166-1-alpha-3`](https://pkg.go.dev/github.com/cinar/checker/v2#IsISO31661Alpha3): Ensures the string is a valid three-letter ISO 3166-1 alpha-3 country code, e.g. `USA`.
- [`iso639-1`](https://pkg.go.dev/github.com/cinar/checker/v2#IsISO6391): Ensures the string is a valid two-letter ISO 639-1 language code, e.g. `en`.
- [`lte`](https://pkg.go.dev/github.com/cinar/checker/v2#IsLte): Ensures the value is less than or equal to the specified number.
- [`luhn`](https://pkg.go.dev/github.com/cinar/checker/v2#IsLUHN): Ensures the string is a valid LUHN number.
- [`mac`](https://pkg.go.dev/github.com/cinar/checker/v2#IsMAC): Ensures the string is a valid MAC address.
- [`max-len`](https://pkg.go.dev/github.com/cinar/checker/v2#MaxLen): Ensures the length of the given value (string, slice, or map) is at most n.
- [`min-len`](https://pkg.go.dev/github.com/cinar/checker/v2#MinLen): Ensures the length of the given value (string, slice, or map) is at least n.
- [`required`](https://pkg.go.dev/github.com/cinar/checker/v2#Required) Ensures the value is provided.
- [`required-if`](https://pkg.go.dev/github.com/cinar/checker/v2#IsRequiredIf): Ensures the value is provided when another field is equal to a given value.
- [`required-unless`](https://pkg.go.dev/github.com/cinar/checker/v2#IsRequiredUnless): Ensures the value is provided unless another field is equal to a given value.
- [`regexp`](https://pkg.go.dev/github.com/cinar/checker/v2#MakeRegexpChecker) Ensured the string matches the pattern.
- [`time`](https://pkg.go.dev/github.com/cinar/checker/v2#IsTime) Ensured the string matches the provided time layout.
- [`url`](https://pkg.go.dev/github.com/cinar/checker/v2#IsURL): Ensures the string is a valid URL.

# Normalizers Provided

- [`lower`](https://pkg.go.dev/github.com/cinar/checker/v2#Lower): Converts the string to lowercase.
- [`title`](https://pkg.go.dev/github.com/cinar/checker/v2#Title): Converts the string to title case.
- [`trim-left`](https://pkg.go.dev/github.com/cinar/checker/v2#TrimLeft): Trims whitespace from the left side of the string.
- [`trim-right`](https://pkg.go.dev/github.com/cinar/checker/v2#TrimRight): Trims whitespace from the right side of the string.
- [`trim`](https://pkg.go.dev/github.com/cinar/checker/v2#TrimSpace): Trims whitespace from both sides of the string.
- [`upper`](https://pkg.go.dev/github.com/cinar/checker/v2#Upper): Converts the string to uppercase.
- [`html-escape`](https://pkg.go.dev/github.com/cinar/checker/v2#HTMLEscape): Escapes special characters in the string for HTML.
- [`html-unescape`](https://pkg.go.dev/github.com/cinar/checker/v2#HTMLUnescape): Unescapes special characters in the string for HTML.
- [`url-escape`](https://pkg.go.dev/github.com/cinar/checker/v2#URLEscape): Escapes special characters in the string for URLs.
- [`url-unescape`](https://pkg.go.dev/github.com/cinar/checker/v2#URLUnescape): Unescapes special characters in the string for URLs.

# Custom Checkers and Normalizers

You can define custom checkers or normalizers and register them for use in your validation logic. Here is an example of how to create and register a custom checker:

```golang
checker.RegisterMaker("is-fruit", func(params string) v2.CheckFunc[reflect.Value] {
	return func(value reflect.Value) (reflect.Value, error) {
		stringValue := value.Interface().(string)

		if stringValue == "apple" || stringValue == "banana" {
			return value, nil
		}

		return value, v2.NewCheckError("NOT_FRUIT")
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

errors, valid := v2.CheckStruct(item)
if !valid {
	fmt.Println(errors)
}
```

In this example, the `is-fruit` checker is used to validate that the `Name` field of the `Item` struct is either "apple" or "banana".

# Slice and Item Level Checkers

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

# Field-Relative and Conditional Checkers

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

# Localized Error Messages

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

errors, valid := v2.CheckStruct(item)
if !valid {
	fmt.Println(errors)
	// Output: map[Name:Not a fruit name.]
}
```

Error messages are generated using Golang template functions, allowing them to include variables.

```golang
// Custrom checker error containing the item name.
err := checker.NewCheckErrorWithData(
	"NOT_FRUIT",
	map[string]interface{}{
		"name": "abcd",
	},
)

// Register the en-US localized error message for the custom NOT_FRUIT error code.
locales.EnUSMessages["NOT_FRUIT"] = "Name {{ .name }} is not a fruit name."

errors, valid := v2.CheckStruct(item)
if !valid {
	fmt.Println(errors)
	// Output: map[Name:Name abcd is not a fruit name.]
}
```

# Structured Errors

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

# Framework Integration

Checker ships thin, separately-versioned adapter modules that bind a request and run `CheckStruct` in a single call, writing a JSON `400` response automatically when binding or validation fails. Each adapter is its own Go module, so the framework it wraps is only pulled in if you actually `go get` that adapter; the core `checker` module stays dependency-free either way.

## Gin

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

## Echo

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

# Changelog

See [CHANGELOG.md](CHANGELOG.md) for a history of notable changes to this project.

# Contributing to the Project

Anyone can contribute to Checkers library. Please make sure to read our [Contributor Covenant Code of Conduct](./CODE_OF_CONDUCT.md) guide first. Follow the [How to Contribute to Checker](./CONTRIBUTING.md) to contribute.

# License

This library is free to use, modify, and distribute under the terms of the MIT license. The full license text can be found in the [LICENSE](./LICENSE) file.

The MIT license is a permissive license that allows you to do almost anything with the library, as long as you retain the copyright notice and the license text. This means that you can use the library in commercial products, modify it, and redistribute it without having to ask for permission from the authors.

The [LICENSE](./LICENSE) file is located in the root directory of the library. You can open it in a text editor to read the full license text.
