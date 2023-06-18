# Trim Left Normalizer

The ```trim-left``` normalizer removes the whitespaces at the beginning of the given value. It can be mixed with checkers and other normalizers when defining the validation steps for user data.

```golang
type User struct {
  Username string `checkers:"trim-left"`
}

user := &User{
  Username: "      normalizer",
}

Check(user)

fmt.Println(user.Username) // normalizer
```
