// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package cli

import (
	"runtime/debug"
	"testing"
)

// withBuildInfo temporarily overrides readBuildInfo for the duration of a
// test, restoring it afterward.
func withBuildInfo(t *testing.T, info *debug.BuildInfo, ok bool) {
	t.Helper()

	original := readBuildInfo
	readBuildInfo = func() (*debug.BuildInfo, bool) {
		return info, ok
	}

	t.Cleanup(func() {
		readBuildInfo = original
	})
}

func TestCoreModuleVersionNoBuildInfo(t *testing.T) {
	withBuildInfo(t, nil, false)

	if got := coreModuleVersion(); got != unknownVersion {
		t.Fatalf("expected %q, got %q", unknownVersion, got)
	}
}

func TestCoreModuleVersionDependencyMissing(t *testing.T) {
	withBuildInfo(t, &debug.BuildInfo{}, true)

	if got := coreModuleVersion(); got != unknownVersion {
		t.Fatalf("expected %q, got %q", unknownVersion, got)
	}
}

func TestCoreModuleVersionTagged(t *testing.T) {
	info := &debug.BuildInfo{
		Deps: []*debug.Module{
			{Path: "example.com/unrelated", Version: "v1.0.0"},
			{Path: checkerModulePath, Version: "v2.1.0"},
		},
	}
	withBuildInfo(t, info, true)

	if got := coreModuleVersion(); got != "v2.1.0" {
		t.Fatalf("expected v2.1.0, got %q", got)
	}
}

func TestCoreModuleVersionReplacedWithVersion(t *testing.T) {
	info := &debug.BuildInfo{
		Deps: []*debug.Module{
			{
				Path:    checkerModulePath,
				Version: "v2.0.1",
				Replace: &debug.Module{Path: "example.com/fork", Version: "v0.0.1-fork"},
			},
		},
	}
	withBuildInfo(t, info, true)

	if got := coreModuleVersion(); got != "v0.0.1-fork" {
		t.Fatalf("expected v0.0.1-fork, got %q", got)
	}
}

func TestCoreModuleVersionReplacedWithLocalPath(t *testing.T) {
	info := &debug.BuildInfo{
		Deps: []*debug.Module{
			{
				Path:    checkerModulePath,
				Version: "v2.0.1",
				Replace: &debug.Module{Path: "../"},
			},
		},
	}
	withBuildInfo(t, info, true)

	want := "(devel, replaced with ../)"
	if got := coreModuleVersion(); got != want {
		t.Fatalf("expected %q, got %q", want, got)
	}
}
