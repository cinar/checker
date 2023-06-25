package checker_test

import (
	"testing"

	"github.com/cinar/checker"
)

func ExampleIsISBN10() {
	result := checker.IsISBN10("1430248270")
	if result != checker.ResultValid {
		// Send the mistakes back to the user
	}
}

func TestIsISBN10Valid(t *testing.T) {
	result := checker.IsISBN10("1430248270")
	if result != checker.ResultValid {
		t.Fail()
	}
}

func TestIsISBN10ValidX(t *testing.T) {
	result := checker.IsISBN10("007462542X")
	if result != checker.ResultValid {
		t.Fail()
	}
}

func TestIsISBN10ValidWithDashes(t *testing.T) {
	result := checker.IsISBN10("1-4302-4827-0")
	if result != checker.ResultValid {
		t.Fail()
	}
}

func TestIsISBN10InvalidLength(t *testing.T) {
	result := checker.IsISBN10("143024827")
	if result != checker.ResultNotISBN {
		t.Fail()
	}
}

func TestIsISBN10InvalidCheck(t *testing.T) {
	result := checker.IsISBN10("1430248272")
	if result != checker.ResultNotISBN {
		t.Fail()
	}
}

func ExampleIsISBN13() {
	result := checker.IsISBN13("9781430248279")
	if result != checker.ResultValid {
		// Send the mistakes back to the user
	}
}

func TestIsISBN13Valid(t *testing.T) {
	result := checker.IsISBN13("9781430248279")
	if result != checker.ResultValid {
		t.Fail()
	}
}

func TestIsISBN13ValidWithDashes(t *testing.T) {
	result := checker.IsISBN13("978-1-4302-4827-9")
	if result != checker.ResultValid {
		t.Fail()
	}
}

func TestIsISBN13InvalidLength(t *testing.T) {
	result := checker.IsISBN13("978143024827")
	if result != checker.ResultNotISBN {
		t.Fail()
	}
}

func TestIsISBN13InvalidCheck(t *testing.T) {
	result := checker.IsISBN13("9781430248272")
	if result != checker.ResultNotISBN {
		t.Fail()
	}
}

func ExampleIsISBN() {
	result := checker.IsISBN("1430248270")
	if result != checker.ResultValid {
		// Send the mistakes back to the user
	}
}

func TestIsISBNValid10(t *testing.T) {
	result := checker.IsISBN("1430248270")
	if result != checker.ResultValid {
		t.Fail()
	}
}

func TestIsISBNValid13(t *testing.T) {
	result := checker.IsISBN("9781430248279")
	if result != checker.ResultValid {
		t.Fail()
	}
}

func TestIsISBNInvalidLenght(t *testing.T) {
	result := checker.IsISBN("978143024827")
	if result != checker.ResultNotISBN {
		t.Fail()
	}
}

func TestCheckISBNNonString(t *testing.T) {
	defer checker.FailIfNoPanic(t)

	type Book struct {
		ISBN int `checkers:"isbn"`
	}

	book := &Book{}

	checker.Check(book)
}

func TestCheckISBNValid(t *testing.T) {
	type Book struct {
		ISBN string `checkers:"isbn"`
	}

	book := &Book{
		ISBN: "1430248270",
	}

	_, valid := checker.Check(book)
	if !valid {
		t.Fail()
	}
}

func TestCheckISBNInvalid(t *testing.T) {
	defer checker.FailIfNoPanic(t)

	type Book struct {
		ISBN string `checkers:"isbn:20"`
	}

	book := &Book{
		ISBN: "1430248270",
	}

	checker.Check(book)
}

func TestCheckISBNValid10(t *testing.T) {
	type Book struct {
		ISBN string `checkers:"isbn:10"`
	}

	book := &Book{
		ISBN: "1430248270",
	}

	_, valid := checker.Check(book)
	if !valid {
		t.Fail()
	}
}

func TestCheckISBNValid13(t *testing.T) {
	type Book struct {
		ISBN string `checkers:"isbn:13"`
	}

	book := &Book{
		ISBN: "9781430248279",
	}

	_, valid := checker.Check(book)
	if !valid {
		t.Fail()
	}
}
