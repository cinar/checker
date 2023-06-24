# Email Checker

The `email` checker checks if the given string is an email address. If the given string is not a valid email address, the checker will return the `NOT_EMAIL` result. Here is an example:

```golang
type User struct {
  Email string `checkers:"email"`
}

user := &User{
  Email: "user@zdo.com",
}

_, valid := checker.Check(user)
if !valid {
  // Send the mistakes back to the user
}
```

In your custom checkers, you can call the `email` checker function [`IsEmail`](https://pkg.go.dev/github.com/cinar/checker#IsEmail) to validate the user input. Here is an example:

```golang
result := checker.IsEmail("user@zdo.com")

if result != checker.ResultValid {
  // Send the mistakes back to the user
}
```
