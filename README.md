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
- [digits](doc/checkers/digits.md) checks if the given string consists of only digit characters.
- [ip](doc/checkers/ip.md) checks if the given value is an IP address.
- [ipv4](doc/checkers/ipv4.md) checks if the given value is an IPv4 address.
- [ipv6](doc/checkers/ipv6.md) checks if the given value is an IPv6 address.
- [max](doc/checkers/max.md) checks if the given value is less than the given maximum.
- [max-length](doc/checkers/maxlength.md) checks if the length of the given value is less than the given maximum length.
- [min](doc/checkers/min.md) checks if the given value is greather than the given minimum.
- [min-length](doc/checkers/minlength.md) checks if the length of the given value is greather than the given minimum length.
- [required](doc/checkers/required.md) checks if the required value is provided.
- [same](doc/checkers/same.md) checks if the given value is equal to the value of the field with the given name.
