// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

// Package cli implements the checker command-line interface: a thin
// wrapper around the core github.com/cinar/checker/v2 module that lets
// shell scripts, CI pipelines, and Git hooks run any registered
// checker/normalizer against a value without writing any Go code.
//
// It deliberately doesn't reimplement or hardcode the checker vocabulary:
// every subcommand goes through checker.CheckWithConfig and
// checker.RegisteredMakerNames/RegisteredFieldMakerNames, so it
// automatically supports every checker and normalizer the linked core
// module knows about, built-in or custom-registered, present or future.
package cli

import (
	"fmt"
	"io"
)

// Exit codes returned by Run, following the conventions most shells and CI
// systems already expect: 0 for success, 1 for a normal failure the caller
// should act on, 2 for a usage mistake (bad flags, unknown command).
const (
	// ExitOK indicates the requested check passed, or the requested
	// informational command (list, version, help) completed successfully.
	ExitOK = 0

	// ExitFailed indicates the requested check ran successfully but the
	// value failed validation.
	ExitFailed = 1

	// ExitUsage indicates the command line itself was invalid: an unknown
	// command, bad flags, or the wrong number of arguments. A malformed
	// checker configuration string (an unknown checker name, or a
	// field-relative checker used outside a struct) is reported as
	// ExitFailed instead, since it's discovered while running the check
	// rather than while parsing the command line.
	ExitUsage = 2
)

// Run parses args (as in os.Args[1:]) and executes the requested
// subcommand, reading from stdin and writing to stdout/stderr as
// appropriate. It returns the process exit code; it never calls
// os.Exit itself, so it can be exercised directly from tests.
func Run(args []string, stdin io.Reader, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		printUsage(stderr)
		return ExitUsage
	}

	switch args[0] {
	case "check":
		return runCheck(args[1:], stdin, stdout, stderr)
	case "list":
		return runList(args[1:], stdout, stderr)
	case "version":
		return runVersion(args[1:], stdout, stderr)
	case "help", "-h", "--help":
		printUsage(stdout)
		return ExitOK
	default:
		fmt.Fprintf(stderr, "checker: unknown command %q\n\n", args[0])
		printUsage(stderr)
		return ExitUsage
	}
}

// printUsage writes the top-level usage text to w.
func printUsage(w io.Writer) {
	fmt.Fprint(w, `checker is a command-line interface to the cinar/checker Go validation
library (https://github.com/cinar/checker), for running its checkers and
normalizers from shell scripts, CI pipelines, and Git hooks.

Usage:

  checker check [--locale=<tag>] [--json] <config> [value]
        Run the given checkers/normalizers config string -- the same
        syntax as a struct field's checkers:"..." tag -- against value,
        or against stdin if value is omitted. Prints the resulting
        (possibly normalized) value to stdout and exits 0 if every check
        passes; prints the error to stderr and exits 1 otherwise.

  checker list
        List every registered checker/normalizer name, and separately
        every field-relative checker name (eq-field and friends, which
        need a struct and so can't be driven from "checker check").

  checker version
        Print the version of the github.com/cinar/checker/v2 module this
        binary was built against.

Examples:

  checker check email "user@example.com"
  echo "  Test@Example.com  " | checker check "trim lower email"
  checker check --locale=de-DE required ""
  checker check --json email "not-an-email"
`)
}
