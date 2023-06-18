package checker

import "testing"

func TestNormalizeTitleNonString(t *testing.T) {
	defer FailIfNoPanic(t)

	type Book struct {
		Chapter int `checkers:"title"`
	}

	book := &Book{}

	Check(book)
}

func TestNormalizeTitleResultValid(t *testing.T) {
	type Book struct {
		Chapter string `checkers:"title"`
	}

	book := &Book{
		Chapter: "THE Checker",
	}

	_, valid := Check(book)
	if !valid {
		t.Fail()
	}
}

func TestNormalizeTitle(t *testing.T) {
	type Book struct {
		Chapter string `checkers:"title"`
	}

	book := &Book{
		Chapter: "THE Checker",
	}

	Check(book)

	if book.Chapter != "The Checker" {
		t.Fail()
	}
}
