[![GoDoc](https://godoc.org/github.com/cinar/checker?status.svg)](https://godoc.org/github.com/cinar/checker)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/cinar/checker)](https://goreportcard.com/report/github.com/cinar/checker)
![Go CI](https://github.com/cinar/checker/actions/workflows/ci.yml/badge.svg)
[![codecov](https://codecov.io/gh/cinar/checker/branch/main/graph/badge.svg?token=VO9BYBHJHE)](https://codecov.io/gh/cinar/checker)

# Checker

Checker is a Go library that helps you validate user input. It can be used to validate user input stored in a struct, or to validate individual pieces of input.

There are many validation libraries available, but I prefer to build my own tools and avoid pulling in unnecessary dependencies. That's why I created Checker, a simple validation library with no dependencies. It's easy to use and gets the job done.

## Usage

To get started, install the Checker library with the following command:

```bash
go get github.com/cinar/checker/v2
```

Next, you will need to import the library into your source file. You can do this by following the example below:

```golang
import (
    checker "github.com/cinar/checker/v2"
)
```

### Validating User Input Stored in a Struct

Checker can be used in two ways. The first way is to validate user input stored in a struct. To do this, you can list the checkers through the struct tag for each field. Here is an example:

```golang
type Person struct {
    Name string `checkers:"trim required"`
}

person := &Person{}

errors, valid := checker.CheckStruct(person)
if !valid {
    // Send the errors back to the user
}
```

### Validating Individual User Input with Multiple Checkers

If you do not want to validate user input stored in a struct, you can individually call the checker functions to validate the user input. Here is an example:

```golang
var name

name, err := v2.Check(name, v2.TrimSpace, v2.Required)
if err != nil {
    // Send the result back to the user
}
```

### Validating Individual User Input

If you do not want to validate user input stored in a struct, you can individually call the checker functions to validate the user input. Here is an example:

```golang
var name

err := checker.IsRequired(name)
if err != nil {
    // Send the result back to the user
}
```

## Normalizers and Checkers

Checkers are used to check for problems in user input, while normalizers are used to transform user input into a preferred format. For example, a normalizer could be used to trim spaces from the beginning and end of a string, or to convert a string to title case.

I am not entirely happy with the decision to combine checkers and normalizers into a single library, but using them together can be useful. Normalizers and checkers can be mixed in any order when defining the validation steps for user data. For example, the trim normalizer can be used in conjunction with the required checker to first trim the user input and then check if the user provided the required information. Here is an example:

```golang
type Person struct {
    Name string `checkers:"trim required"`
}
```

# Checkers Provided

This package currently provides the following checkers:

- `alphanumeric` checks if the given string consists of only alphanumeric characters.
- `ascii` checks if the given string consists of only ASCII characters.
- `cidr` checker checks if the value is a valid CIDR notation IP address and prefix length.
- `credit-card` checks if the given value is a valid credit card number.
- `digits` checks if the given string consists of only digit characters.
- `email` checks if the given string is an email address.
- `fqdn` checks if the given string is a fully qualified domain name.
- `ip` checks if the given value is an IP address.
- `ipv4` checks if the given value is an IPv4 address.
- `ipv6` checks if the given value is an IPv6 address.
- `isbn` checks if the given value is a valid ISBN number.
- `luhn` checks if the given number is valid based on the Luhn algorithm.
- `mac` checks if the given value is a valid an IEEE 802 MAC-48, EUI-48, EUI-64, or a 20-octet IP over InfiniBand link-layer address.
- `max` checks if the given value is less than the given maximum.
- `max-length` checks if the length of the given value is less than the given maximum length.
- `min` checks if the given value is greather than the given minimum.
- `min-length` checks if the length of the given value is greather than the given minimum length.
- `regexp` checks if the given string matches the regexp pattern.
- `required` checks if the required value is provided.
- `same` checks if the given value is equal to the value of the field with the given name.
- `url` checks if the given value is a valid URL.

# Normalizers Provided

This package currently provides the following normalizers. They can be mixed with the checkers when defining the validation steps for user data.

- `html-escape` applies HTML escaping to special characters.
- `html-unescape` applies HTML unescaping to special characters.
- `lower` maps all Unicode letters in the given value to their lower case.
- `upper` maps all Unicode letters in the given value to their upper case.
- `title` maps the first letter of each word to their upper case.
- `trim` removes the whitespaces at the beginning and at the end of the given value.
- `trim-left` removes the whitespaces at the beginning of the given value.
- `trim-right` removes the whitespaces at the end of the given value.
- `url-escape` applies URL escaping to special characters.
- `url-unescape` applies URL unescaping to special characters.

# Custom Checkers

You can define a custom checker or normalizer and register it as shown below:

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

Once you have registered your custom checker, you can use it by simply specifying its name.

```golang
type Item struct {
	Name string `checkers:"is-fruit"`
}

person := &Item{
	Name: "banana",
}

err, ok := v2.CheckStruct(person)
if !ok {
	fmt.Println(err)
}
```

# Contributing to the Project

Anyone can contribute to Checkers library. Please make sure to read our [Contributor Covenant Code of Conduct](./CODE_OF_CONDUCT.md) guide first. Follow the [How to Contribute to Checker](./CONTRIBUTING.md) to contribute.

# License

This library is free to use, modify, and distribute under the terms of the MIT license. The full license text can be found in the [LICENSE](./LICENSE) file.

The MIT license is a permissive license that allows you to do almost anything with the library, as long as you retain the copyright notice and the license text. This means that you can use the library in commercial products, modify it, and redistribute it without having to ask for permission from the authors.

The [LICENSE](./LICENSE) file is located in the root directory of the library. You can open it in a text editor to read the full license text.
