// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

// Command checker is the standalone command-line interface to the
// github.com/cinar/checker/v2 validation library. See "checker help" or
// the cli package documentation for usage.
package main

import (
	"os"

	"github.com/cinar/checker/v2/cli"
)

func main() {
	os.Exit(cli.Run(os.Args[1:], os.Stdin, os.Stdout, os.Stderr))
}
