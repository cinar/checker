// Package checker is a Go library for validating user input through struct tags.
//
// Copyright (c) 2023 Onur Cinar. All Rights Reserved.
// The source code is provided under MIT License.
//
// https://github.com/cinar/checker
package checker_test

import (
	"testing"

	"github.com/cinar/checker"
)

func TestCheckEmailNonString(t *testing.T) {
	defer checker.FailIfNoPanic(t)

	type User struct {
		Email int `checkers:"email"`
	}

	user := &User{}

	checker.Check(user)
}

func TestCheckEmailValid(t *testing.T) {
	type User struct {
		Email string `checkers:"email"`
	}

	user := &User{
		Email: "user@zdo.com",
	}

	_, valid := checker.Check(user)
	if !valid {
		t.Fail()
	}
}

func TestIsEmailValid(t *testing.T) {
	validEmails := []string{
		"simple@example.com",
		"very.common@example.com",
		"disposable.style.email.with+symbol@example.com",
		"other.email-with-hyphen@and.subdomains.example.com",
		"fully-qualified-domain@example.com",
		"user.name+tag+sorting@example.com",
		"x@example.com",
		"example-indeed@strange-example.com",
		"test/test@test.com",
		"example@s.example",
		"\" \"@example.org",
		"\"john..doe\"@example.org",
		"mailhost!username@example.org",
		"\"very.(),:;<>[]\\\".VERY.\\\"very@\\\\ \\\"very\\\".unusual\"@strange.example.com",
		"user%example.com@example.org",
		"user-@example.org",
		"postmaster@[123.123.123.123]",
		"postmaster@[IPv6:2001:0db8:85a3:0000:0000:8a2e:0370:7334]",
	}

	for _, email := range validEmails {
		if checker.IsEmail(email) != checker.ResultValid {
			t.Fatal(email)
		}
	}
}

func TestIsEmailInvalid(t *testing.T) {
	validEmails := []string{
		"Abc.example.com",
		"A@b@c@example.com",
		"a\"b(c)d,e:f;g<h>i[j\\k]l@example.com",
		"just\"not\"right@example.com",
		"this is\"not\\allowed@example.com",
		"this\\ still\\\"not\\\\allowed@example.com",
		"1234567890123456789012345678901234567890123456789012345678901234+x@example.com",
		"i_like_underscore@but_its_not_allowed_in_this_part.example.com",
		"QA[icon]CHOCOLATE[icon]@test.com",
		".cannot.start.with.dot@example.com",
		"cannot.end.with.dot.@example.com",
		"cannot.have..double.dots@example.com",
		"user@domaincannotbemorethan255charactersdomaincannotbemorethan255charactersdomaincannotbemorethan255charactersdomaincannotbemorethan255charactersdomaincannotbemorethan255charactersdomaincannotbemorethan255charactersdomaincannotbemorethan255charactersdomaincannotbemorethan255characters.com",
	}

	for _, email := range validEmails {
		if checker.IsEmail(email) == checker.ResultValid {
			t.Fatal(email)
		}
	}
}
