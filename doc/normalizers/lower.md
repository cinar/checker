# Lower Case Normalizer

The `lower` normalizer maps all Unicode letters in the given value to their lower case. It can be mixed with checkers and other normalizers when defining the validation steps for user data.

```golang
type User struct {
  Username string `checkers:"lower"`
}

user := &User{
  Username: "chECker",
}

checker.Check(user)

fmt.Println(user.Username) // checker
```
