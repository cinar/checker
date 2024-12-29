[![GoDoc](https://godoc.org/github.com/cinar/checker?status.svg)](https://godoc.org/github.com/cinar/checker)
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

- [`ascii`](DOC.md#IsASCII): Ensures the string contains only ASCII characters.
- [`alphanumeric`](DOC.md#IsAlphanumeric): Ensures the string contains only letters and numbers.
- [`credit-card`](DOC.md#IsAnyCreditCard): Ensures the string is a valid credit card number.
- [`cidr`](DOC.md#IsCIDR): Ensures the string is a valid CIDR notation.
- [`digits`](DOC.md#IsDigits): Ensures the string contains only digits.
- [`email`](DOC.md#IsEmail): Ensures the string is a valid email address.
- [`fqdn`](DOC.md#IsFQDN): Ensures the string is a valid fully qualified domain name.
- [`hex`](DOC.md#IsHex): Ensures the string contains only hex digits.
- [`ip`](DOC.md#IsIP): Ensures the string is a valid IP address.
- [`ipv4`](DOC.md#IsIPv4): Ensures the string is a valid IPv4 address.
- [`ipv6`](DOC.md#IsIPv6): Ensures the string is a valid IPv6 address.
- [`isbn`](DOC.md#IsISBN): Ensures the string is a valid ISBN.
- [`luhn`](DOC.md#IsLUHN): Ensures the string is a valid LUHN number.
- [`mac`](DOC.md#IsMAC): Ensures the string is a valid MAC address.
- [`max-len`](DOC.md#func-maxlen): Ensures the length of the given value (string, slice, or map) is at most n.
- [`min-len`](DOC.md#func-minlen): Ensures the length of the given value (string, slice, or map) is at least n.
- [`required`](DOC.md#func-required) Ensures the value is provided.
- [`regexp`](DOC.md#func-makeregexpchecker) Ensured the string matches the pattern.
- [`url`](DOC.md#IsURL): Ensures the string is a valid URL.

# Normalizers Provided

- [`lower`](DOC.md#Lower): Converts the string to lowercase.
- [`title`](DOC.md#Title): Converts the string to title case.
- [`trim-left`](DOC.md#TrimLeft): Trims whitespace from the left side of the string.
- [`trim-right`](DOC.md#TrimRight): Trims whitespace from the right side of the string.
- [`trim`](DOC.md#TrimSpace): Trims whitespace from both sides of the string.
- [`upper`](DOC.md#Upper): Converts the string to uppercase.
- [`html-escape`](DOC.md#HTMLEscape): Escapes special characters in the string for HTML.
- [`html-unescape`](DOC.md#HTMLUnescape): Unescapes special characters in the string for HTML.
- [`url-escape`](DOC.md#URLEscape): Escapes special characters in the string for URLs.
- [`url-unescape`](DOC.md#URLUnescape): Unescapes special characters in the string for URLs.

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

# Localized Error Messages

When validation fails, Checker returns an error. By default, the [Error()](DOC.md#CheckError.Error) function provides a human-readable error message in `en-US` locale.

```golang
_, err := checker.IsEmail("abcd")
if err != nil {
	fmt.Println(err)
	// Output: Not a valid email address.
}
```

To get error messages in other languages, use the [ErrorWithLocale()](DOC.md#CheckError.ErrorWithLocale) function. By default, only `en-US` is registered. You can register additional languages by calling [RegisterLocale](DOC.md#RegisterLocale).

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
err := NewCheckErrorWithData(
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

# Contributing to the Project

Anyone can contribute to Checkers library. Please make sure to read our [Contributor Covenant Code of Conduct](./CODE_OF_CONDUCT.md) guide first. Follow the [How to Contribute to Checker](./CONTRIBUTING.md) to contribute.

# License

This library is free to use, modify, and distribute under the terms of the MIT license. The full license text can be found in the [LICENSE](./LICENSE) file.

The MIT license is a permissive license that allows you to do almost anything with the library, as long as you retain the copyright notice and the license text. This means that you can use the library in commercial products, modify it, and redistribute it without having to ask for permission from the authors.

The [LICENSE](./LICENSE) file is located in the root directory of the library. You can open it in a text editor to read the full license text.
