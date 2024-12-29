// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2_test

import (
	"io/fs"
	"testing"

	v2 "github.com/cinar/checker/v2"
	"github.com/cinar/checker/v2/locales"
)

func TestCheckErrorWithNotLocalizedCode(t *testing.T) {
	code := "TEST"

	err := v2.NewCheckError(code)

	if err.Error() != code {
		t.Fatalf("actual %s expected %s", err.Error(), code)
	}
}

func TestCheckErrorWithLocalizedCode(t *testing.T) {
	code := "TEST"
	message := "Test message"

	locales.EnUSMessages[code] = message

	err := v2.NewCheckError(code)

	if err.ErrorWithLocale("fr-FR") != message {
		t.Fatalf("actual %s expected %s", err.Error(), message)
	}
}

func TestCheckErrorWithDefaultLocalizedCode(t *testing.T) {
	code := "TEST"
	message := "Test message"

	locales.EnUSMessages[code] = message

	err := v2.NewCheckError(code)

	if err.Error() != message {
		t.Fatalf("actual %s expected %s", err.Error(), message)
	}
}

func TestCheckErrorWithDataAndLocalizedCode(t *testing.T) {
	code := "TEST"
	message := "Test message {{.Name}}"

	locales.EnUSMessages[code] = message

	err := v2.NewCheckErrorWithData(code, map[string]interface{}{
		"Name": "Onur",
	})

	expected := "Test message Onur"

	if err.Error() != expected {
		t.Fatalf("actual %s expected %s", err.Error(), expected)
	}
}

func TestCheckErrorWithLocalizedCodeInvalidTemplate(t *testing.T) {
	code := "TEST"
	message := "Test message {{}"

	locales.EnUSMessages[code] = message

	err := v2.NewCheckError(code)

	if err.Error() != code {
		t.Fatalf("actual %s expected %s", err.Error(), code)
	}
}

func TestCheckErrorWithLocalizedCodeInvalidExecute(t *testing.T) {
	code := "TEST"
	message := "{{ len .Name}}"

	locales.EnUSMessages[code] = message

	err := v2.NewCheckError(code)

	if err.Error() != code {
		t.Fatalf("actual %s expected %s", err.Error(), code)
	}
}

func TestCheckErrorIsSuccess(t *testing.T) {
	code := "TEST"

	err1 := v2.NewCheckError(code)
	err2 := v2.NewCheckError(code)

	if !err1.Is(err2) {
		t.Fatalf("actual %t expected %t", err1.Is(err2), true)
	}
}

func TestCheckErrorIsFailure(t *testing.T) {
	code1 := "TEST1"
	code2 := "TEST2"

	err1 := v2.NewCheckError(code1)
	err2 := v2.NewCheckError(code2)

	if err1.Is(err2) {
		t.Fatalf("actual %t expected %t", err1.Is(err2), false)
	}
}

func TestCheckErrorIsFailureWithDifferentType(t *testing.T) {
	code := "TEST"

	err1 := v2.NewCheckError(code)
	err2 := fs.ErrExist

	if err1.Is(err2) {
		t.Fatalf("actual %t expected %t", err1.Is(err2), false)
	}
}

func TestRegisterLocale(t *testing.T) {
	locale := "de-DE"
	code := "TEST"
	message := "Testmeldung"

	v2.RegisterLocale(locale, map[string]string{
		code: message,
	})

	err := v2.NewCheckError(code)

	if err.ErrorWithLocale("de-DE") != message {
		t.Fatalf("actual %s expected %s", err.Error(), message)
	}
}
