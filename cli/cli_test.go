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

func TestRunNoArgs(t *testing.T) {
	var stdout, stderr bytes.Buffer

	code := cli.Run(nil, strings.NewReader(""), &stdout, &stderr)

	if code != cli.ExitUsage {
		t.Fatalf("expected ExitUsage, got %d", code)
	}

	if stderr.Len() == 0 {
		t.Fatal("expected usage on stderr")
	}
}

func TestRunUnknownCommand(t *testing.T) {
	var stdout, stderr bytes.Buffer

	code := cli.Run([]string{"bogus"}, strings.NewReader(""), &stdout, &stderr)

	if code != cli.ExitUsage {
		t.Fatalf("expected ExitUsage, got %d", code)
	}

	if !strings.Contains(stderr.String(), "unknown command") {
		t.Fatalf("expected unknown command message, got %q", stderr.String())
	}
}

func TestRunHelp(t *testing.T) {
	for _, arg := range []string{"help", "-h", "--help"} {
		t.Run(arg, func(t *testing.T) {
			var stdout, stderr bytes.Buffer

			code := cli.Run([]string{arg}, strings.NewReader(""), &stdout, &stderr)

			if code != cli.ExitOK {
				t.Fatalf("expected ExitOK, got %d", code)
			}

			if !strings.Contains(stdout.String(), "checker check") {
				t.Fatalf("expected usage text on stdout, got %q", stdout.String())
			}
		})
	}
}

func TestRunDispatchesCheck(t *testing.T) {
	var stdout, stderr bytes.Buffer

	code := cli.Run([]string{"check", "email", "user@example.com"}, strings.NewReader(""), &stdout, &stderr)

	if code != cli.ExitOK {
		t.Fatalf("expected ExitOK, got %d, stderr=%q", code, stderr.String())
	}

	if strings.TrimSpace(stdout.String()) != "user@example.com" {
		t.Fatalf("unexpected stdout: %q", stdout.String())
	}
}

func TestRunDispatchesList(t *testing.T) {
	var stdout, stderr bytes.Buffer

	code := cli.Run([]string{"list"}, strings.NewReader(""), &stdout, &stderr)

	if code != cli.ExitOK {
		t.Fatalf("expected ExitOK, got %d, stderr=%q", code, stderr.String())
	}

	if !strings.Contains(stdout.String(), "email") {
		t.Fatalf("expected email in list output, got %q", stdout.String())
	}
}

func TestRunDispatchesVersion(t *testing.T) {
	var stdout, stderr bytes.Buffer

	code := cli.Run([]string{"version"}, strings.NewReader(""), &stdout, &stderr)

	if code != cli.ExitOK {
		t.Fatalf("expected ExitOK, got %d, stderr=%q", code, stderr.String())
	}

	if !strings.HasPrefix(stdout.String(), "checker ") {
		t.Fatalf("unexpected version output: %q", stdout.String())
	}
}
