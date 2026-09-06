// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2_test

import (
	"io/fs"
	"sync"
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

// TestCheckErrorWithDataConcurrentSameTemplate exercises the parsed-template
// cache from many goroutines rendering the same templated message
// concurrently. Run with `go test -race`.
func TestCheckErrorWithDataConcurrentSameTemplate(t *testing.T) {
	code := "TEST_CONCURRENT"
	message := "Test message {{.Name}}"

	locales.EnUSMessages[code] = message

	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			err := v2.NewCheckErrorWithData(code, map[string]interface{}{
				"Name": "Onur",
			})

			if err.Error() != "Test message Onur" {
				t.Error(err)
			}
		}()
	}

	wg.Wait()
}

func BenchmarkErrorWithLocaleStatic(b *testing.B) {
	err := v2.NewCheckError("REQUIRED")

	for i := 0; i < b.N; i++ {
		_ = err.Error()
	}
}

func BenchmarkErrorWithLocaleTemplated(b *testing.B) {
	err := v2.NewCheckErrorWithData("NOT_MIN_LEN", map[string]interface{}{
		"min": 8,
	})

	for i := 0; i < b.N; i++ {
		_ = err.Error()
	}
}

func TestCheckErrorWithDataDoesNotHTMLEscape(t *testing.T) {
	code := "TEST_NO_ESCAPE"
	message := "Got: {{.Value}}"

	locales.EnUSMessages[code] = message

	err := v2.NewCheckErrorWithData(code, map[string]interface{}{
		"Value": `<script>&"'</script>`,
	})

	expected := `Got: <script>&"'</script>`

	if err.Error() != expected {
		t.Fatalf("actual %q expected %q", err.Error(), expected)
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

func TestCheckErrorWithMessageOverridesLocale(t *testing.T) {
	code := "TEST"
	locales.EnUSMessages[code] = "Locale message"

	err := v2.NewCheckError(code).WithMessage("Custom message")

	if err.Error() != "Custom message" {
		t.Fatalf("actual %s expected %s", err.Error(), "Custom message")
	}
}

func TestCheckErrorWithMessageDoesNotMutateOriginal(t *testing.T) {
	code := "TEST"
	locales.EnUSMessages[code] = "Locale message"

	original := v2.NewCheckError(code)
	_ = original.WithMessage("Custom message")

	if original.Error() != "Locale message" {
		t.Fatalf("actual %s expected %s", original.Error(), "Locale message")
	}
}

func TestCheckErrorWithMessageTemplatePlaceholder(t *testing.T) {
	code := "TEST"

	err := v2.NewCheckErrorWithData(code, map[string]interface{}{
		"min": 8,
	}).WithMessage("Need at least {{ .min }}")

	expected := "Need at least 8"

	if err.Error() != expected {
		t.Fatalf("actual %s expected %s", err.Error(), expected)
	}
}

func TestCheckErrorWithMessageInvalidTemplateFallsBackToCode(t *testing.T) {
	code := "TEST"

	err := v2.NewCheckError(code).WithMessage("Bad template {{}")

	if err.Error() != code {
		t.Fatalf("actual %s expected %s", err.Error(), code)
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
