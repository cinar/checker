[![GoDoc](https://godoc.org/github.com/cinar/checker?status.svg)](https://godoc.org/github.com/cinar/checker)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
![Go CI](https://github.com/cinar/checker/actions/workflows/go.yml/badge.svg)

# Checker

Checker is a Go library that helps you validate user input. It can be used to validate user input stored in a struct, or to validate individual pieces of input.

There are many validation libraries available, but I prefer to build my own tools and avoid pulling in unnecessary dependencies. That's why I created Checker, a simple validation library with no dependencies. It's easy to use and gets the job done.

## Usage

To get started, install the Checker library with the following command:

```bash
go get github.com/cinar/checker
```

Next, you will need to import the library into your source file. You can do this by following the example below:

```golang
import (
    "github.com/cinar/checker"
)
```

### Validating User Input Stored in a Struct

Checker can be used in two ways. The first way is to validate user input stored in a struct. To do this, you can list the checkers through the struct tag for each field. Here is an example:

```golang
type Person struct {
    Name string `checkers:"required"`
}

person := &Person{}

mistakes, valid := checker.Check(person)
if !valid {
    // Send the mistakes back to the user
}
```

### Validating Individual User Input

If you do not want to validate user input stored in a struct, you can individually call the checker functions to validate the user input. Here is an example:

```golang
var name

result := checker.IsRequired(name)
if result != ResultValid {
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

- [alphanumeric](doc/checkers/alphanumeric.md) checks if the given string consists of only alphanumeric characters.
- [ascii](doc/checkers/ascii.md) checks if the given string consists of only ASCII characters.
- [cidr](doc/checkers/cidr.md) checker checks if the value is a valid CIDR notation IP address and prefix length.
- [digits](doc/checkers/digits.md) checks if the given string consists of only digit characters.
- [email](doc/checkers/email.md) checks if the given string is an email address.
- [fqdn](doc/checkers/fqdn.md) checks if the given string is a fully qualified domain name.
- [ip](doc/checkers/ip.md) checks if the given value is an IP address.
- [ipv4](doc/checkers/ipv4.md) checks if the given value is an IPv4 address.
- [ipv6](doc/checkers/ipv6.md) checks if the given value is an IPv6 address.
- [mac](doc/checkers/mac.md) checks if the given value is a valid an IEEE 802 MAC-48, EUI-48, EUI-64, or a 20-octet IP over InfiniBand link-layer address.
- [max](doc/checkers/max.md) checks if the given value is less than the given maximum.
- [max-length](doc/checkers/maxlength.md) checks if the length of the given value is less than the given maximum length.
- [min](doc/checkers/min.md) checks if the given value is greather than the given minimum.
- [min-length](doc/checkers/minlength.md) checks if the length of the given value is greather than the given minimum length.
- [regexp](doc/checkers/regexp.md) checks if the given string matches the regexp pattern.
- [required](doc/checkers/required.md) checks if the required value is provided.
- [same](doc/checkers/same.md) checks if the given value is equal to the value of the field with the given name.

# Normalizers Provided

This package currently provides the following normalizers. They can be mixed with the checkers when defining the validation steps for user data.

- [lower](doc/normalizers/lower.md) maps all Unicode letters in the given value to their lower case.
- [upper](doc/normalizers/upper.md) maps all Unicode letters in the given value to their upper case.
- [title](doc/normalizers/title.md) maps the first letter of each word to their upper case.
- [trim](doc/normalizers/trim.md) removes the whitespaces at the beginning and at the end of the given value.
- [trim-left](doc/normalizers/trim_left.md) removes the whitespaces at the beginning of the given value.
- [trim-right](doc/normalizers/trim_right.md) removes the whitespaces at the end of the given value.

# Custom Checkers

To define a custom checker, you need to create a new function with the following parameters:

```golang
func CustomChecker(value, parent reflect.Value) Result {
    return ResultValid
}
```
type MakeFunc 
You also need to create a make function that takes the checker configuration and returns a reference to the checker function.

```golang
func CustomMaker(params string) CheckFunc {
    return CustomChecker
}
```

Finally, you need to call the ```Register``` function to register your custom checker.

```golang
checker.Register("custom-checker", CustomMaker)
```

Once you have registered your custom checker, you can use it by simply specifying its name.

```golang
type User struct {
    Username string `checkers:"custom-checker"`
}
```

# Code Quality

User input validation is a critical task that must be performed correctly in order to ensure that user data is handled correctly. This is why it is important to have extensive unit testing in place for any user input validation library. 

The Checker library has a code coverage threshold of 100%, which means that all of the code in the library has been tested. This ensures that the library is extremely reliable and that it will not fail to validate user input correctly.

The test cases for the library can be found in the _test.go files. These files contain a comprehensive set of tests that cover all of the possible scenarios for user input validation.

If you are planning to make a pull request to this project, please make sure to add enough test cases to ensure that the code coverage remains at 100%. This will help to ensure that the library remains reliable and that user data is handled correctly.

# License

This library is free to use, modify, and distribute under the terms of the MIT license. The full license text can be found in the [LICENSE](./LICENSE) file.

The MIT license is a permissive license that allows you to do almost anything with the library, as long as you retain the copyright notice and the license text. This means that you can use the library in commercial products, modify it, and redistribute it without having to ask for permission from the authors.

The [LICENSE](./LICENSE) file is located in the root directory of the library. You can open it in a text editor to read the full license text.
