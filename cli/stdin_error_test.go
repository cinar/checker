// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package cli_test

import (
	"bytes"
	"errors"
	"testing"

	"github.com/cinar/checker/v2/cli"
)

// errReader is an io.Reader whose Read always fails, used to exercise the
// stdin-read-failure path of "checker check" without a value argument.
type errReader struct{}

func (errReader) Read([]byte) (int, error) {
	return 0, errors.New("boom")
}

func TestCheckStdinReadError(t *testing.T) {
	var stdout, stderr bytes.Buffer

	code := cli.Run([]string{"check", "email"}, errReader{}, &stdout, &stderr)

	if code != cli.ExitUsage {
		t.Fatalf("expected ExitUsage, got %d", code)
	}

	if stderr.Len() == 0 {
		t.Fatal("expected error message on stderr")
	}
}
