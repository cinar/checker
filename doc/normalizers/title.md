# Title Case Normalizer

The ```title``` normalizer maps the first letter of each word to their upper case. It can be mixed with checkers and other normalizers when defining the validation steps for user data.

```golang
type Book struct {
  Chapter string `checkers:"title"`
}

book := &Book{
  Chapter: "THE Checker",
}

Check(book)

fmt.Println(book.Chapter) // The Checker
```
