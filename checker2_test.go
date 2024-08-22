package checker_test

import (
	"reflect"
	"testing"

	"github.com/cinar/checker"
)

func TestRegisterRulesFromConfig(t *testing.T) {
	type Person struct {
		Name  string
		Email string
	}

	err := checker.RegisterRulesFromConfig(reflect.TypeFor[Person](), map[string]string{
		"Name":  "trim required",
		"Email": "required email",
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestRegisterRulesFromTag(t *testing.T) {
	type Address struct {
		Street string `checkers:"required"`
	}

	type Person struct {
		Name  string `checkers:"trim required"`
		Email string `checkers:"required email"`
		Home  Address
		Work  Address
	}

	err := checker.RegisterRulesFromTag(reflect.TypeFor[Person]())
	if err != nil {
		t.Fatal(err)
	}
}
