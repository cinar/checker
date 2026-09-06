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

func TestListOK(t *testing.T) {
	var stdout, stderr bytes.Buffer

	code := cli.Run([]string{"list"}, strings.NewReader(""), &stdout, &stderr)

	if code != cli.ExitOK {
		t.Fatalf("expected ExitOK, got %d, stderr=%q", code, stderr.String())
	}

	out := stdout.String()

	if !strings.Contains(out, "  email\n") {
		t.Fatalf("expected email checker in output, got %q", out)
	}

	if !strings.Contains(out, "  eq-field\n") {
		t.Fatalf("expected eq-field checker in output, got %q", out)
	}

	if !strings.Contains(out, "Field-relative checkers") {
		t.Fatalf("expected field-relative section header, got %q", out)
	}
}

func TestListTooManyArgs(t *testing.T) {
	var stdout, stderr bytes.Buffer

	code := cli.Run([]string{"list", "extra"}, strings.NewReader(""), &stdout, &stderr)

	if code != cli.ExitUsage {
		t.Fatalf("expected ExitUsage, got %d", code)
	}

	if !strings.Contains(stderr.String(), "usage:") {
		t.Fatalf("expected usage message, got %q", stderr.String())
	}
}

func TestListBadFlag(t *testing.T) {
	var stdout, stderr bytes.Buffer

	code := cli.Run([]string{"list", "--bogus-flag"}, strings.NewReader(""), &stdout, &stderr)

	if code != cli.ExitUsage {
		t.Fatalf("expected ExitUsage, got %d", code)
	}

	if stderr.Len() == 0 {
		t.Fatal("expected flag error on stderr")
	}
}
