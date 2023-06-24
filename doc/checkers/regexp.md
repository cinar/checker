# Regexp Checker

The `regexp` checker checks if the given string matches the given regexp. If the given string does not match, the checker will return the `NOT_MATCH` result. Here is an example:

```golang
type User struct {
  Username string `checkers:"regexp:^[A-Za-z]+$"`
}

user := &User{
  Username: "abcd",
}

_, valid := checker.Check(user)
if !valid {
  // Send the mistakes back to the user
}
```

The `regexp` checker can be used to build other checkers for other regexp patterns. In order to do that, you can use the [`MakeRegexpChecker`](https://pkg.go.dev/github.com/cinar/checker#MakeRegexpChecker) function. The function takes an expression and a result to return when the the given string is not a match. Here is an example:

```golang
checkHex := checker.MakeRegexpChecker("^[A-Fa-f0-9]+$", "NOT_HEX")

result := checkHex(reflect.ValueOf("f0f0f0"), reflect.ValueOf(nil))
if result != checker.ResultValid {
  // Send the mistakes back to the user
}
```

To register the new regexp checker to validate user input in struct, [`Register`](https://pkg.go.dev/github.com/cinar/checker#Register) function can be used. Here is an example:

```golang
checker.Register("hex", checker.MakeRegexpMaker("^[A-Fa-f0-9]+$", "NOT_HEX"))

type Theme struct {
  Color string `checkers:hex`
}

theme := &Theme{
  Color: "f0f0f0",
}

_, valid := checker.Check(theme)
if !valid {
  // Send the mistakes back to the user
}
```
