// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package cli

import (
	"flag"
	"fmt"
	"io"
	"sort"

	checker "github.com/cinar/checker/v2"
)

// runList implements "checker list".
func runList(args []string, stdout, stderr io.Writer) int {
	fs := flag.NewFlagSet("list", flag.ContinueOnError)
	fs.SetOutput(stderr)

	fs.Usage = func() {
		fmt.Fprintln(stderr, "usage: checker list")
	}

	if err := fs.Parse(args); err != nil {
		return ExitUsage
	}

	if fs.NArg() != 0 {
		fs.Usage()
		return ExitUsage
	}

	names := checker.RegisteredMakerNames()
	sort.Strings(names)

	fmt.Fprintln(stdout, "Checkers and normalizers (usable with \"checker check\"):")
	for _, name := range names {
		fmt.Fprintf(stdout, "  %s\n", name)
	}

	fieldNames := checker.RegisteredFieldMakerNames()
	sort.Strings(fieldNames)

	fmt.Fprintln(stdout)
	fmt.Fprintln(stdout, "Field-relative checkers (require a struct; not usable with \"checker check\"):")
	for _, name := range fieldNames {
		fmt.Fprintf(stdout, "  %s\n", name)
	}

	return ExitOK
}
