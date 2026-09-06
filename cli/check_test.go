// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package cli_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/cinar/checker/v2/cli"
)

func runCheckArgs(t *testing.T, stdin string, args ...string) (stdout, stderr string, code int) {
	t.Helper()

	var outBuf, errBuf bytes.Buffer

	fullArgs := append([]string{"check"}, args...)
	code = cli.Run(fullArgs, strings.NewReader(stdin), &outBuf, &errBuf)

	return outBuf.String(), errBuf.String(), code
}

func TestCheckValidValueArg(t *testing.T) {
	stdout, _, code := runCheckArgs(t, "", "email", "user@example.com")

	if code != cli.ExitOK {
		t.Fatalf("expected ExitOK, got %d", code)
	}

	if strings.TrimSpace(stdout) != "user@example.com" {
		t.Fatalf("unexpected stdout: %q", stdout)
	}
}

func TestCheckInvalidValueArg(t *testing.T) {
	stdout, stderr, code := runCheckArgs(t, "", "email", "not-an-email")

	if code != cli.ExitFailed {
		t.Fatalf("expected ExitFailed, got %d", code)
	}

	if stdout != "" {
		t.Fatalf("expected no stdout on failure, got %q", stdout)
	}

	if !strings.Contains(stderr, "Not a valid email address.") {
		t.Fatalf("unexpected stderr: %q", stderr)
	}
}

func TestCheckNormalizes(t *testing.T) {
	stdout, _, code := runCheckArgs(t, "", "trim lower email", "  Test@Example.com  ")

	if code != cli.ExitOK {
		t.Fatalf("expected ExitOK, got %d", code)
	}

	if strings.TrimSpace(stdout) != "test@example.com" {
		t.Fatalf("unexpected stdout: %q", stdout)
	}
}

func TestCheckValueFromStdin(t *testing.T) {
	stdout, _, code := runCheckArgs(t, "  Test@Example.com  \n", "trim lower email")

	if code != cli.ExitOK {
		t.Fatalf("expected ExitOK, got %d", code)
	}

	if strings.TrimSpace(stdout) != "test@example.com" {
		t.Fatalf("unexpected stdout: %q", stdout)
	}
}

func TestCheckValueFromStdinCRLF(t *testing.T) {
	stdout, _, code := runCheckArgs(t, "user@example.com\r\n", "email")

	if code != cli.ExitOK {
		t.Fatalf("expected ExitOK, got %d", code)
	}

	if strings.TrimSpace(stdout) != "user@example.com" {
		t.Fatalf("unexpected stdout: %q", stdout)
	}
}

func TestCheckJSONValid(t *testing.T) {
	stdout, _, code := runCheckArgs(t, "", "--json", "email", "user@example.com")

	if code != cli.ExitOK {
		t.Fatalf("expected ExitOK, got %d", code)
	}

	want := `{"valid":true,"value":"user@example.com"}`
	if strings.TrimSpace(stdout) != want {
		t.Fatalf("unexpected stdout: %q", stdout)
	}
}

func TestCheckJSONInvalid(t *testing.T) {
	stdout, _, code := runCheckArgs(t, "", "--json", "email", "not-an-email")

	if code != cli.ExitFailed {
		t.Fatalf("expected ExitFailed, got %d", code)
	}

	want := `{"valid":false,"error":{"code":"NOT_EMAIL","message":"Not a valid email address."}}`
	if strings.TrimSpace(stdout) != want {
		t.Fatalf("unexpected stdout: %q", stdout)
	}
}

func TestCheckLocale(t *testing.T) {
	_, stderr, code := runCheckArgs(t, "", "--locale=de-DE", "required", "")

	if code != cli.ExitFailed {
		t.Fatalf("expected ExitFailed, got %d", code)
	}

	if !strings.Contains(stderr, "Erforderlicher Wert fehlt.") {
		t.Fatalf("unexpected stderr: %q", stderr)
	}
}

func TestCheckJSONLocale(t *testing.T) {
	stdout, _, code := runCheckArgs(t, "", "--json", "--locale=de-DE", "required", "")

	if code != cli.ExitFailed {
		t.Fatalf("expected ExitFailed, got %d", code)
	}

	if !strings.Contains(stdout, "Erforderlicher Wert fehlt.") {
		t.Fatalf("unexpected stdout: %q", stdout)
	}
}

func TestCheckUnknownCheckerNameRecovered(t *testing.T) {
	_, stderr, code := runCheckArgs(t, "", "bogus-checker-name", "x")

	if code != cli.ExitFailed {
		t.Fatalf("expected ExitFailed, got %d", code)
	}

	if !strings.Contains(stderr, "invalid check configuration") {
		t.Fatalf("expected recovered panic message, got %q", stderr)
	}
}

func TestCheckUnknownCheckerNameRecoveredJSON(t *testing.T) {
	stdout, _, code := runCheckArgs(t, "", "--json", "bogus-checker-name", "x")

	if code != cli.ExitFailed {
		t.Fatalf("expected ExitFailed, got %d", code)
	}

	if !strings.Contains(stdout, `"code":"ERROR"`) {
		t.Fatalf("expected ERROR code in JSON output, got %q", stdout)
	}
}

func TestCheckFieldRelativeCheckerRecovered(t *testing.T) {
	_, stderr, code := runCheckArgs(t, "", "eq-field:Foo", "x")

	if code != cli.ExitFailed {
		t.Fatalf("expected ExitFailed, got %d", code)
	}

	if !strings.Contains(stderr, "invalid check configuration") {
		t.Fatalf("expected recovered panic message, got %q", stderr)
	}
}

func TestCheckNoArgs(t *testing.T) {
	_, stderr, code := runCheckArgs(t, "")

	if code != cli.ExitUsage {
		t.Fatalf("expected ExitUsage, got %d", code)
	}

	if !strings.Contains(stderr, "usage:") {
		t.Fatalf("expected usage message, got %q", stderr)
	}
}

func TestCheckTooManyArgs(t *testing.T) {
	_, stderr, code := runCheckArgs(t, "", "email", "a@b.com", "extra")

	if code != cli.ExitUsage {
		t.Fatalf("expected ExitUsage, got %d", code)
	}

	if !strings.Contains(stderr, "usage:") {
		t.Fatalf("expected usage message, got %q", stderr)
	}
}

func TestCheckBadFlag(t *testing.T) {
	_, stderr, code := runCheckArgs(t, "", "--bogus-flag", "email", "a@b.com")

	if code != cli.ExitUsage {
		t.Fatalf("expected ExitUsage, got %d", code)
	}

	if stderr == "" {
		t.Fatal("expected flag error on stderr")
	}
}
