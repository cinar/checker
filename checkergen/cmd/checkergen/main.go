// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

// Command checkergen generates reflection-free Go validation code from
// Checker's "checkers" (and "validate" fallback) struct tags. Add a
// generate directive in the package containing the struct(s) to generate
// for:
//
//	//go:generate go run github.com/cinar/checker/v2/checkergen/cmd/checkergen
//
// See the checkergen package doc for what makes a struct field eligible
// for generation.
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/cinar/checker/v2/checkergen"
)

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}

// run is checkergen's whole CLI logic, exercised directly from tests (no
// os.Exit inside it) and wrapped by main.
func run(args []string, stdout, stderr *os.File) int {
	flags := flag.NewFlagSet("checkergen", flag.ContinueOnError)
	flags.SetOutput(stderr)

	typeFlag := flags.String("type", "", "comma-separated struct names to generate (default: every eligible struct)")

	if err := flags.Parse(args); err != nil {
		return 2
	}

	dir := "."
	if flags.NArg() > 0 {
		dir = flags.Arg(0)
	}

	var typeFilter []string
	if *typeFlag != "" {
		typeFilter = strings.Split(*typeFlag, ",")
	}

	result, err := checkergen.Generate(dir, typeFilter)
	if err != nil {
		fmt.Fprintf(stderr, "checkergen: %v\n", err)
		return 1
	}

	for _, name := range result.Generated {
		fmt.Fprintf(stdout, "checkergen: generated Check%s\n", name)
	}

	for name, reason := range result.Skipped {
		fmt.Fprintf(stderr, "checkergen: skipped %s: %s\n", name, reason)
	}

	if len(result.Generated) == 0 && len(result.Skipped) == 0 {
		fmt.Fprintln(stderr, "checkergen: no struct with a checkers/validate tag found")
	}

	return 0
}
