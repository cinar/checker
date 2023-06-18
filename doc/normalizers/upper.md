# Upper Case Normalizer

The ```upper``` normalizer maps all Unicode letters in the given value to their upper case. It can be mixed with checkers and other normalizers when defining the validation steps for user data.

```golang
type User struct {
  Username string `checkers:"upper"`
}

user := &User{
  Username: "chECker",
}

Check(user)
```
