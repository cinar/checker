// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package cli

import (
	"flag"
	"fmt"
	"io"
	"runtime/debug"
)

// checkerModulePath is the module path debug.ReadBuildInfo reports the
// core checker dependency under, used to look up its resolved version.
const checkerModulePath = "github.com/cinar/checker/v2"

// unknownVersion is reported when the core module's version can't be
// determined from the binary's build info.
const unknownVersion = "(unknown)"

// readBuildInfo is a seam over debug.ReadBuildInfo so tests can exercise
// coreModuleVersion's fallback paths without needing to actually produce a
// binary built without embedded module info.
var readBuildInfo = debug.ReadBuildInfo

// runVersion implements "checker version".
func runVersion(args []string, stdout, stderr io.Writer) int {
	fs := flag.NewFlagSet("version", flag.ContinueOnError)
	fs.SetOutput(stderr)

	fs.Usage = func() {
		fmt.Fprintln(stderr, "usage: checker version")
	}

	if err := fs.Parse(args); err != nil {
		return ExitUsage
	}

	if fs.NArg() != 0 {
		fs.Usage()
		return ExitUsage
	}

	fmt.Fprintf(stdout, "checker %s\n", coreModuleVersion())

	return ExitOK
}

// coreModuleVersion returns the resolved version of the core
// github.com/cinar/checker/v2 module this binary was built against, read
// from the build info embedded by "go build"/"go install". It falls back
// to "(unknown)" when build info isn't available (e.g. "go run") or
// doesn't list the dependency (e.g. a local replace directive without a
// pseudo-version, as in this repository's own development build).
func coreModuleVersion() string {
	info, ok := readBuildInfo()
	if !ok {
		return unknownVersion
	}

	return moduleVersion(info)
}

// moduleVersion extracts the resolved version of checkerModulePath from
// info, applying the same replace/pseudo-version fallbacks as
// coreModuleVersion's doc comment describes.
func moduleVersion(info *debug.BuildInfo) string {
	for _, dep := range info.Deps {
		if dep.Path != checkerModulePath {
			continue
		}

		if dep.Replace != nil {
			if dep.Replace.Version != "" {
				return dep.Replace.Version
			}

			return "(devel, replaced with " + dep.Replace.Path + ")"
		}

		return dep.Version
	}

	return unknownVersion
}
