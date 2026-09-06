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

func TestVersionOK(t *testing.T) {
	var stdout, stderr bytes.Buffer

	code := cli.Run([]string{"version"}, strings.NewReader(""), &stdout, &stderr)

	if code != cli.ExitOK {
		t.Fatalf("expected ExitOK, got %d, stderr=%q", code, stderr.String())
	}

	if !strings.HasPrefix(stdout.String(), "checker ") {
		t.Fatalf("unexpected version output: %q", stdout.String())
	}
}

func TestVersionTooManyArgs(t *testing.T) {
	var stdout, stderr bytes.Buffer

	code := cli.Run([]string{"version", "extra"}, strings.NewReader(""), &stdout, &stderr)

	if code != cli.ExitUsage {
		t.Fatalf("expected ExitUsage, got %d", code)
	}

	if !strings.Contains(stderr.String(), "usage:") {
		t.Fatalf("expected usage message, got %q", stderr.String())
	}
}

func TestVersionBadFlag(t *testing.T) {
	var stdout, stderr bytes.Buffer

	code := cli.Run([]string{"version", "--bogus-flag"}, strings.NewReader(""), &stdout, &stderr)

	if code != cli.ExitUsage {
		t.Fatalf("expected ExitUsage, got %d", code)
	}

	if stderr.Len() == 0 {
		t.Fatal("expected flag error on stderr")
	}
}
