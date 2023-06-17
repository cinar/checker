# Trim Right Normalizer

The ```trim``` normalizer removes the whitespaces at the beginning and at the end of the given value. It can be mixed with checkers and other normalizers when defining the validation steps for user data.

```golang
type User struct {
  Username string `checkers:"trim"`
}

user := &User{
  Username: "normalizer      ",
}

Check(user)
```
