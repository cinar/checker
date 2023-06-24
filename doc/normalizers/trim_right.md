# Trim Right Normalizer

The `trim-right` normalizer removes the whitespaces at the end of the given value. It can be mixed with checkers and other normalizers when defining the validation steps for user data.

```golang
type User struct {
  Username string `checkers:"trim-right"`
}

user := &User{
  Username: "normalizer      ",
}

checker.Check(user)

fmt.Println(user.Username) // CHECKER
```
