package checker_test

import (
	"testing"

	"github.com/cinar/checker"
)

func ExampleIsURL() {
	result := checker.IsURL("https://zdo.com")
	if result != checker.ResultValid {
		// Send the mistakes back to the user
	}
}

func TestIsURLValid(t *testing.T) {
	result := checker.IsURL("https://zdo.com")
	if result != checker.ResultValid {
		t.Fail()
	}
}

func TestIsURLInvalid(t *testing.T) {
	result := checker.IsURL("https:://index.html")
	if result == checker.ResultValid {
		t.Fail()
	}
}

func TestCheckURLNonString(t *testing.T) {
	defer checker.FailIfNoPanic(t)

	type Bookmark struct {
		URL int `checkers:"url"`
	}

	bookmark := &Bookmark{}

	checker.Check(bookmark)
}

func TestCheckURLValid(t *testing.T) {
	type Bookmark struct {
		URL string `checkers:"url"`
	}

	bookmark := &Bookmark{
		URL: "https://zdo.com",
	}

	_, valid := checker.Check(bookmark)
	if !valid {
		t.Fail()
	}
}

func TestCheckURLInvalid(t *testing.T) {
	type Bookmark struct {
		URL string `checkers:"url"`
	}

	bookmark := &Bookmark{
		URL: "zdo.com/index.html",
	}

	_, valid := checker.Check(bookmark)
	if valid {
		t.Fail()
	}
}

func TestCheckURLWithoutSchema(t *testing.T) {
	type Bookmark struct {
		URL string `checkers:"url"`
	}

	bookmark := &Bookmark{
		URL: "//zdo.com/index.html",
	}

	_, valid := checker.Check(bookmark)
	if valid {
		t.Fail()
	}
}

func TestCheckURLWithoutHost(t *testing.T) {
	type Bookmark struct {
		URL string `checkers:"url"`
	}

	bookmark := &Bookmark{
		URL: "https:://index.html",
	}

	_, valid := checker.Check(bookmark)
	if valid {
		t.Fail()
	}
}
