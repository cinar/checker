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

errors, valid := checker.Check(person)
if !valid {
    // Send the errors back to the user
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

To define a custom checker, you need to create a new function with the following parameters:

```golang
func CustomChecker(value, parent reflect.Value) error {
    return nil
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

<!-- gomarkdoc:embed:start -->

<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# checker

```go
import "github.com/cinar/checker"
```

Package checker is a Go library for validating user input through struct tags.

Copyright \(c\) 2023\-2024 Onur Cinar. Use of this source code is governed by a MIT\-style license that can be found in the LICENSE file. https://github.com/cinar/checker

## Index

- [Variables](<#variables>)
- [func FailIfNoPanic\(t \*testing.T\)](<#FailIfNoPanic>)
- [func IsASCII\(value string\) error](<#IsASCII>)
- [func IsAlphanumeric\(value string\) error](<#IsAlphanumeric>)
- [func IsAmexCreditCard\(number string\) error](<#IsAmexCreditCard>)
- [func IsAnyCreditCard\(number string\) error](<#IsAnyCreditCard>)
- [func IsCidr\(value string\) error](<#IsCidr>)
- [func IsDigits\(value string\) error](<#IsDigits>)
- [func IsDinersCreditCard\(number string\) error](<#IsDinersCreditCard>)
- [func IsDiscoverCreditCard\(number string\) error](<#IsDiscoverCreditCard>)
- [func IsEmail\(email string\) error](<#IsEmail>)
- [func IsFqdn\(domain string\) error](<#IsFqdn>)
- [func IsIP\(value string\) error](<#IsIP>)
- [func IsIPV4\(value string\) error](<#IsIPV4>)
- [func IsIPV6\(value string\) error](<#IsIPV6>)
- [func IsISBN\(value string\) error](<#IsISBN>)
- [func IsISBN10\(value string\) error](<#IsISBN10>)
- [func IsISBN13\(value string\) error](<#IsISBN13>)
- [func IsJcbCreditCard\(number string\) error](<#IsJcbCreditCard>)
- [func IsLuhn\(number string\) error](<#IsLuhn>)
- [func IsMac\(value string\) error](<#IsMac>)
- [func IsMasterCardCreditCard\(number string\) error](<#IsMasterCardCreditCard>)
- [func IsMax\(value interface\{\}, max float64\) error](<#IsMax>)
- [func IsMaxLength\(value interface\{\}, maxLength int\) error](<#IsMaxLength>)
- [func IsMin\(value interface\{\}, min float64\) error](<#IsMin>)
- [func IsMinLength\(value interface\{\}, minLength int\) error](<#IsMinLength>)
- [func IsRequired\(v interface\{\}\) error](<#IsRequired>)
- [func IsURL\(value string\) error](<#IsURL>)
- [func IsUnionPayCreditCard\(number string\) error](<#IsUnionPayCreditCard>)
- [func IsVisaCreditCard\(number string\) error](<#IsVisaCreditCard>)
- [func Register\(name string, maker MakeFunc\)](<#Register>)
- [type CheckFunc](<#CheckFunc>)
  - [func MakeRegexpChecker\(expression string, invalidError error\) CheckFunc](<#MakeRegexpChecker>)
- [type Errors](<#Errors>)
  - [func Check\(s interface\{\}\) \(Errors, bool\)](<#Check>)
- [type MakeFunc](<#MakeFunc>)
  - [func MakeRegexpMaker\(expression string, invalidError error\) MakeFunc](<#MakeRegexpMaker>)


## Variables

<a name="ErrNotASCII"></a>ErrNotASCII indicates that the given string contains non\-ASCII characters.

```go
var ErrNotASCII = errors.New("please use standard English characters only")
```

<a name="ErrNotAlphanumeric"></a>ErrNotAlphanumeric indicates that the given string contains non\-alphanumeric characters.

```go
var ErrNotAlphanumeric = errors.New("please use only letters and numbers")
```

<a name="ErrNotCidr"></a>ErrNotCidr indicates that the given value is not a valid CIDR.

```go
var ErrNotCidr = errors.New("please enter a valid CIDR")
```

<a name="ErrNotCreditCard"></a>ErrNotCreditCard indicates that the given value is not a valid credit card number.

```go
var ErrNotCreditCard = errors.New("please enter a valid credit card number")
```

<a name="ErrNotDigits"></a>ErrNotDigits indicates that the given string contains non\-digit characters.

```go
var ErrNotDigits = errors.New("please enter a valid number")
```

<a name="ErrNotEmail"></a>ErrNotEmail indicates that the given string is not a valid email.

```go
var ErrNotEmail = errors.New("please enter a valid email address")
```

<a name="ErrNotFqdn"></a>ErrNotFqdn indicates that the given string is not a valid FQDN.

```go
var ErrNotFqdn = errors.New("please enter a valid domain name")
```

<a name="ErrNotIP"></a>ErrNotIP indicates that the given value is not an IP address.

```go
var ErrNotIP = errors.New("please enter a valid IP address")
```

<a name="ErrNotIPV4"></a>ErrNotIPV4 indicates that the given value is not an IPv4 address.

```go
var ErrNotIPV4 = errors.New("please enter a valid IPv4 address")
```

<a name="ErrNotIPV6"></a>ErrNotIPV6 indicates that the given value is not an IPv6 address.

```go
var ErrNotIPV6 = errors.New("please enter a valid IPv6 address")
```

<a name="ErrNotISBN"></a>ErrNotISBN indicates that the given value is not a valid ISBN.

```go
var ErrNotISBN = errors.New("please enter a valid ISBN number")
```

<a name="ErrNotLuhn"></a>ErrNotLuhn indicates that the given number is not valid based on the Luhn algorithm.

```go
var ErrNotLuhn = errors.New("please enter a valid LUHN")
```

<a name="ErrNotMac"></a>ErrNotMac indicates that the given value is not an MAC address.

```go
var ErrNotMac = errors.New("please enter a valid MAC address")
```

<a name="ErrNotMatch"></a>ErrNotMatch indicates that the given string does not match the regexp pattern.

```go
var ErrNotMatch = errors.New("please enter a valid input")
```

<a name="ErrNotSame"></a>ErrNotSame indicates that the given two values are not equal to each other.

```go
var ErrNotSame = errors.New("does not match the other")
```

<a name="ErrNotURL"></a>ErrNotURL indicates that the given value is not a valid URL.

```go
var ErrNotURL = errors.New("please enter a valid URL")
```

<a name="ErrRequired"></a>ErrRequired indicates that the required value is missing.

```go
var ErrRequired = errors.New("is required")
```

<a name="FailIfNoPanic"></a>
## func [FailIfNoPanic](<https://github.com/cinar/checker/blob/main/test_helper.go#L11>)

```go
func FailIfNoPanic(t *testing.T)
```

FailIfNoPanic fails if test didn't panic. Use this function with the defer.

<a name="IsASCII"></a>
## func [IsASCII](<https://github.com/cinar/checker/blob/main/ascii.go#L21>)

```go
func IsASCII(value string) error
```

IsASCII checks if the given string consists of only ASCII characters.

<details><summary>Example</summary>
<p>



```go
package main

import (
	"github.com/cinar/checker"
)

func main() {
	err := checker.IsASCII("Checker")
	if err != nil {
		// Send the errors back to the user
	}
}
```

</p>
</details>

<a name="IsAlphanumeric"></a>
## func [IsAlphanumeric](<https://github.com/cinar/checker/blob/main/alphanumeric.go#L21>)

```go
func IsAlphanumeric(value string) error
```

IsAlphanumeric checks if the given string consists of only alphanumeric characters.

<details><summary>Example</summary>
<p>



```go
package main

import (
	"github.com/cinar/checker"
)

func main() {
	err := checker.IsAlphanumeric("ABcd1234")
	if err != nil {
		// Send the errors back to the user
	}
}
```

</p>
</details>

<a name="IsAmexCreditCard"></a>
## func [IsAmexCreditCard](<https://github.com/cinar/checker/blob/main/credit_card.go#L77>)

```go
func IsAmexCreditCard(number string) error
```

IsAmexCreditCard checks if the given valie is a valid AMEX credit card.

<details><summary>Example</summary>
<p>



```go
package main

import (
	"github.com/cinar/checker"
)

func main() {
	err := checker.IsAmexCreditCard("378282246310005")

	if err != nil {
		// Send the errors back to the user
	}
}
```

</p>
</details>

<a name="IsAnyCreditCard"></a>
## func [IsAnyCreditCard](<https://github.com/cinar/checker/blob/main/credit_card.go#L72>)

```go
func IsAnyCreditCard(number string) error
```

IsAnyCreditCard checks if the given value is a valid credit card number.

<details><summary>Example</summary>
<p>



```go
package main

import (
	"github.com/cinar/checker"
)

func main() {
	err := checker.IsAnyCreditCard("6011111111111117")

	if err != nil {
		// Send the errors back to the user
	}
}
```

</p>
</details>

<a name="IsCidr"></a>
## func [IsCidr](<https://github.com/cinar/checker/blob/main/cidr.go#L21>)

```go
func IsCidr(value string) error
```

IsCidr checker checks if the value is a valid CIDR notation IP address and prefix length.

<details><summary>Example</summary>
<p>



```go
package main

import (
	"github.com/cinar/checker"
)

func main() {
	err := checker.IsCidr("2001:db8::/32")
	if err != nil {
		// Send the errors back to the user
	}
}
```

</p>
</details>

<a name="IsDigits"></a>
## func [IsDigits](<https://github.com/cinar/checker/blob/main/digits.go#L21>)

```go
func IsDigits(value string) error
```

IsDigits checks if the given string consists of only digit characters.

<details><summary>Example</summary>
<p>



```go
package main

import (
	"github.com/cinar/checker"
)

func main() {
	err := checker.IsDigits("1234")
	if err != nil {
		// Send the errors back to the user
	}
}
```

</p>
</details>

<a name="IsDinersCreditCard"></a>
## func [IsDinersCreditCard](<https://github.com/cinar/checker/blob/main/credit_card.go#L82>)

```go
func IsDinersCreditCard(number string) error
```

IsDinersCreditCard checks if the given valie is a valid Diners credit card.

<details><summary>Example</summary>
<p>



```go
package main

import (
	"github.com/cinar/checker"
)

func main() {
	err := checker.IsDinersCreditCard("36227206271667")

	if err != nil {
		// Send the errors back to the user
	}
}
```

</p>
</details>

<a name="IsDiscoverCreditCard"></a>
## func [IsDiscoverCreditCard](<https://github.com/cinar/checker/blob/main/credit_card.go#L87>)

```go
func IsDiscoverCreditCard(number string) error
```

IsDiscoverCreditCard checks if the given valie is a valid Discover credit card.

<details><summary>Example</summary>
<p>



```go
package main

import (
	"github.com/cinar/checker"
)

func main() {
	err := checker.IsDiscoverCreditCard("6011111111111117")

	if err != nil {
		// Send the errors back to the user
	}
}
```

</p>
</details>

<a name="IsEmail"></a>
## func [IsEmail](<https://github.com/cinar/checker/blob/main/email.go#L28>)

```go
func IsEmail(email string) error
```

IsEmail checks if the given string is an email address.

<details><summary>Example</summary>
<p>



```go
package main

import (
	"github.com/cinar/checker"
)

func main() {
	err := checker.IsEmail("user@zdo.com")
	if err != nil {
		// Send the errors back to the user
	}
}
```

</p>
</details>

<a name="IsFqdn"></a>
## func [IsFqdn](<https://github.com/cinar/checker/blob/main/fqdn.go#L25>)

```go
func IsFqdn(domain string) error
```

IsFqdn checks if the given string is a fully qualified domain name.

<details><summary>Example</summary>
<p>



```go
package main

import (
	"github.com/cinar/checker"
)

func main() {
	err := checker.IsFqdn("zdo.com")
	if err != nil {
		// Send the errors back to the user
	}
}
```

</p>
</details>

<a name="IsIP"></a>
## func [IsIP](<https://github.com/cinar/checker/blob/main/ip.go#L21>)

```go
func IsIP(value string) error
```

IsIP checks if the given value is an IP address.

<details><summary>Example</summary>
<p>



```go
package main

import (
	"github.com/cinar/checker"
)

func main() {
	err := checker.IsIP("2001:db8::68")
	if err != nil {
		// Send the errors back to the user
	}
}
```

</p>
</details>

<a name="IsIPV4"></a>
## func [IsIPV4](<https://github.com/cinar/checker/blob/main/ipv4.go#L21>)

```go
func IsIPV4(value string) error
```

IsIPV4 checks if the given value is an IPv4 address.

<details><summary>Example</summary>
<p>



```go
package main

import (
	"github.com/cinar/checker"
)

func main() {
	err := checker.IsIPV4("192.168.1.1")
	if err != nil {
		// Send the errors back to the user
	}
}
```

</p>
</details>

<a name="IsIPV6"></a>
## func [IsIPV6](<https://github.com/cinar/checker/blob/main/ipv6.go#L21>)

```go
func IsIPV6(value string) error
```

IsIPV6 checks if the given value is an IPv6 address.

<details><summary>Example</summary>
<p>



```go
package main

import (
	"github.com/cinar/checker"
)

func main() {
	err := checker.IsIPV6("2001:db8::68")

	if err != nil {
		// Send the errors back to the user
	}
}
```

</p>
</details>

<a name="IsISBN"></a>
## func [IsISBN](<https://github.com/cinar/checker/blob/main/isbn.go#L77>)

```go
func IsISBN(value string) error
```

IsISBN checks if the given value is a valid ISBN number.

<details><summary>Example</summary>
<p>



```go
package main

import (
	"github.com/cinar/checker"
)

func main() {
	err := checker.IsISBN("1430248270")
	if err != nil {
		// Send the errors back to the user
	}
}
```

</p>
</details>

<a name="IsISBN10"></a>
## func [IsISBN10](<https://github.com/cinar/checker/blob/main/isbn.go#L27>)

```go
func IsISBN10(value string) error
```

IsISBN10 checks if the given value is a valid ISBN\-10 number.

<details><summary>Example</summary>
<p>



```go
package main

import (
	"github.com/cinar/checker"
)

func main() {
	err := checker.IsISBN10("1430248270")
	if err != nil {
		// Send the errors back to the user
	}
}
```

</p>
</details>

<a name="IsISBN13"></a>
## func [IsISBN13](<https://github.com/cinar/checker/blob/main/isbn.go#L50>)

```go
func IsISBN13(value string) error
```

IsISBN13 checks if the given value is a valid ISBN\-13 number.

<details><summary>Example</summary>
<p>



```go
package main

import (
	"github.com/cinar/checker"
)

func main() {
	err := checker.IsISBN13("9781430248279")
	if err != nil {
		// Send the errors back to the user
	}
}
```

</p>
</details>

<a name="IsJcbCreditCard"></a>
## func [IsJcbCreditCard](<https://github.com/cinar/checker/blob/main/credit_card.go#L92>)

```go
func IsJcbCreditCard(number string) error
```

IsJcbCreditCard checks if the given valie is a valid JCB 15 credit card.

<details><summary>Example</summary>
<p>



```go
package main

import (
	"github.com/cinar/checker"
)

func main() {
	err := checker.IsJcbCreditCard("3530111333300000")

	if err != nil {
		// Send the errors back to the user
	}
}
```

</p>
</details>

<a name="IsLuhn"></a>
## func [IsLuhn](<https://github.com/cinar/checker/blob/main/luhn.go#L23>)

```go
func IsLuhn(number string) error
```

IsLuhn checks if the given number is valid based on the Luhn algorithm.

<details><summary>Example</summary>
<p>



```go
package main

import (
	"github.com/cinar/checker"
)

func main() {
	err := checker.IsLuhn("4012888888881881")
	if err != nil {
		// Send the errors back to the user
	}
}
```

</p>
</details>

<a name="IsMac"></a>
## func [IsMac](<https://github.com/cinar/checker/blob/main/mac.go#L21>)

```go
func IsMac(value string) error
```

IsMac checks if the given value is a valid an IEEE 802 MAC\-48, EUI\-48, EUI\-64, or a 20\-octet IP over InfiniBand link\-layer address.

<details><summary>Example</summary>
<p>



```go
package main

import (
	"github.com/cinar/checker"
)

func main() {
	err := checker.IsMac("00:00:5e:00:53:01")
	if err != nil {
		// Send the errors back to the user
	}
}
```

</p>
</details>

<a name="IsMasterCardCreditCard"></a>
## func [IsMasterCardCreditCard](<https://github.com/cinar/checker/blob/main/credit_card.go#L97>)

```go
func IsMasterCardCreditCard(number string) error
```

IsMasterCardCreditCard checks if the given valie is a valid MasterCard credit card.

<details><summary>Example</summary>
<p>



```go
package main

import (
	"github.com/cinar/checker"
)

func main() {
	err := checker.IsMasterCardCreditCard("5555555555554444")

	if err != nil {
		// Send the errors back to the user
	}
}
```

</p>
</details>

<a name="IsMax"></a>
## func [IsMax](<https://github.com/cinar/checker/blob/main/max.go#L18>)

```go
func IsMax(value interface{}, max float64) error
```

IsMax checks if the given value is below than the given maximum.

<details><summary>Example</summary>
<p>



```go
package main

import (
	"github.com/cinar/checker"
)

func main() {
	quantity := 5

	err := checker.IsMax(quantity, 10)
	if err != nil {
		// Send the errors back to the user
	}
}
```

</p>
</details>

<a name="IsMaxLength"></a>
## func [IsMaxLength](<https://github.com/cinar/checker/blob/main/maxlength.go#L18>)

```go
func IsMaxLength(value interface{}, maxLength int) error
```

IsMaxLength checks if the length of the given value is less than the given maximum length.

<details><summary>Example</summary>
<p>



```go
package main

import (
	"github.com/cinar/checker"
)

func main() {
	s := "1234"

	err := checker.IsMaxLength(s, 4)
	if err != nil {
		// Send the errors back to the user
	}
}
```

</p>
</details>

<a name="IsMin"></a>
## func [IsMin](<https://github.com/cinar/checker/blob/main/min.go#L18>)

```go
func IsMin(value interface{}, min float64) error
```

IsMin checks if the given value is above than the given minimum.

<details><summary>Example</summary>
<p>



```go
package main

import (
	"github.com/cinar/checker"
)

func main() {
	age := 45

	err := checker.IsMin(age, 21)
	if err != nil {
		// Send the errors back to the user
	}
}
```

</p>
</details>

<a name="IsMinLength"></a>
## func [IsMinLength](<https://github.com/cinar/checker/blob/main/minlenght.go#L18>)

```go
func IsMinLength(value interface{}, minLength int) error
```

IsMinLength checks if the length of the given value is greather than the given minimum length.

<details><summary>Example</summary>
<p>



```go
package main

import (
	"github.com/cinar/checker"
)

func main() {
	s := "1234"

	err := checker.IsMinLength(s, 4)
	if err != nil {
		// Send the errors back to the user
	}
}
```

</p>
</details>

<a name="IsRequired"></a>
## func [IsRequired](<https://github.com/cinar/checker/blob/main/required.go#L20>)

```go
func IsRequired(v interface{}) error
```

IsRequired checks if the given required value is present.

<details><summary>Example</summary>
<p>



```go
package main

import (
	"github.com/cinar/checker"
)

func main() {
	var name string

	err := checker.IsRequired(name)
	if err != nil {
		// Send the err back to the user
	}
}
```

</p>
</details>

<a name="IsURL"></a>
## func [IsURL](<https://github.com/cinar/checker/blob/main/url.go#L21>)

```go
func IsURL(value string) error
```

IsURL checks if the given value is a valid URL.

<details><summary>Example</summary>
<p>



```go
package main

import (
	"github.com/cinar/checker"
)

func main() {
	err := checker.IsURL("https://zdo.com")
	if err != nil {
		// Send the errors back to the user
	}
}
```

</p>
</details>

<a name="IsUnionPayCreditCard"></a>
## func [IsUnionPayCreditCard](<https://github.com/cinar/checker/blob/main/credit_card.go#L102>)

```go
func IsUnionPayCreditCard(number string) error
```

IsUnionPayCreditCard checks if the given valie is a valid UnionPay credit card.

<details><summary>Example</summary>
<p>



```go
package main

import (
	"github.com/cinar/checker"
)

func main() {
	err := checker.IsUnionPayCreditCard("6200000000000005")

	if err != nil {
		// Send the errors back to the user
	}
}
```

</p>
</details>

<a name="IsVisaCreditCard"></a>
## func [IsVisaCreditCard](<https://github.com/cinar/checker/blob/main/credit_card.go#L107>)

```go
func IsVisaCreditCard(number string) error
```

IsVisaCreditCard checks if the given valie is a valid Visa credit card.

<details><summary>Example</summary>
<p>



```go
package main

import (
	"github.com/cinar/checker"
)

func main() {
	err := checker.IsVisaCreditCard("4111111111111111")

	if err != nil {
		// Send the errors back to the user
	}
}
```

</p>
</details>

<a name="Register"></a>
## func [Register](<https://github.com/cinar/checker/blob/main/checker.go#L67>)

```go
func Register(name string, maker MakeFunc)
```

Register registers the given checker name and the maker function.

<a name="CheckFunc"></a>
## type [CheckFunc](<https://github.com/cinar/checker/blob/main/checker.go#L16>)

CheckFunc defines the signature for the checker functions.

```go
type CheckFunc func(value, parent reflect.Value) error
```

<a name="MakeRegexpChecker"></a>
### func [MakeRegexpChecker](<https://github.com/cinar/checker/blob/main/regexp.go#L28>)

```go
func MakeRegexpChecker(expression string, invalidError error) CheckFunc
```

MakeRegexpChecker makes a regexp checker for the given regexp expression with the given invalid result.

<a name="Errors"></a>
## type [Errors](<https://github.com/cinar/checker/blob/main/checker.go#L22>)

Errors provides a mapping of the checker errors keyed by the field names.

```go
type Errors map[string]error
```

<a name="Check"></a>
### func [Check](<https://github.com/cinar/checker/blob/main/checker.go#L72>)

```go
func Check(s interface{}) (Errors, bool)
```

Check function checks the given struct based on the checkers listed in field tag names.

<a name="MakeFunc"></a>
## type [MakeFunc](<https://github.com/cinar/checker/blob/main/checker.go#L19>)

MakeFunc defines the signature for the checker maker functions.

```go
type MakeFunc func(params string) CheckFunc
```

<a name="MakeRegexpMaker"></a>
### func [MakeRegexpMaker](<https://github.com/cinar/checker/blob/main/regexp.go#L21>)

```go
func MakeRegexpMaker(expression string, invalidError error) MakeFunc
```

MakeRegexpMaker makes a regexp checker maker for the given regexp expression with the given invalid result.

Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)


<!-- gomarkdoc:embed:end -->

# Contributing to the Project

Anyone can contribute to Checkers library. Please make sure to read our [Contributor Covenant Code of Conduct](./CODE_OF_CONDUCT.md) guide first. Follow the [How to Contribute to Checker](./CONTRIBUTING.md) to contribute.

# License

This library is free to use, modify, and distribute under the terms of the MIT license. The full license text can be found in the [LICENSE](./LICENSE) file.

The MIT license is a permissive license that allows you to do almost anything with the library, as long as you retain the copyright notice and the license text. This means that you can use the library in commercial products, modify it, and redistribute it without having to ask for permission from the authors.

The [LICENSE](./LICENSE) file is located in the root directory of the library. You can open it in a text editor to read the full license text.
