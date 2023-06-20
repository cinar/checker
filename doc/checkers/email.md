# Email Checker

The ```email``` checker checks if the given string is an email address. If the given string is not a valid email address, the checker will return the ```NOT_EMAIL``` result. Here is an example:

```golang
type User struct {
  Email string `checkers:"email"`
}

user := &User{
  Email: "user@zdo.com",
}

_, valid := Check(user)
if !valid {
  // Send the mistakes back to the user
}
```

In your custom checkers, you can call the ```email``` checker function ```IsEmail``` to validate the user input. Here is an example:

```golang
result := IsEmail("user@zdo.com")

if result != ResultValid {
  // Send the mistakes back to the user
}
```
