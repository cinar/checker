// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package checkergen_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/cinar/checker/v2/checkergen"
)

// copyToScratch copies every non-generated file from srcDir into a fresh
// directory under testdata/, and returns its path. The scratch directory
// has to live inside this module's own tree, not under t.TempDir() (which
// lands outside it, under the system temp directory): go/packages needs a
// directory it can resolve within a module. Tests that call
// checkergen.Generate with anything other than a plain, unfiltered rerun
// (e.g. a -type filter) must target a scratch copy like this one, never a
// real testdata package directly -- Generate overwrites that package's
// committed, shared checkergen_generated.go file as a side effect.
func copyToScratch(t *testing.T, srcDir string) string {
	t.Helper()

	scratch, err := os.MkdirTemp("testdata", "scratch-")
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		_ = os.RemoveAll(scratch)
	})

	entries, err := os.ReadDir(srcDir)
	if err != nil {
		t.Fatal(err)
	}

	for _, entry := range entries {
		if entry.IsDir() || entry.Name() == checkergen.GeneratedFileName {
			continue
		}

		data, err := os.ReadFile(filepath.Join(srcDir, entry.Name()))
		if err != nil {
			t.Fatal(err)
		}

		if err := os.WriteFile(filepath.Join(scratch, entry.Name()), data, 0o644); err != nil {
			t.Fatal(err)
		}
	}

	return scratch
}

// TestGenerateSkipsIneligibleStructs confirms a struct with an unmapped
// checker or an ineligible field type is skipped with a clear reason,
// without preventing other structs in the same package from being
// generated.
func TestGenerateSkipsIneligibleStructs(t *testing.T) {
	result, err := checkergen.Generate("testdata/ineligible", nil)
	if err != nil {
		t.Fatal(err)
	}

	wantGenerated := map[string]bool{"Eligible": true, "WithEmbedded": true}

	if len(result.Generated) != len(wantGenerated) {
		t.Fatalf("expected %v to be generated, got %v", wantGenerated, result.Generated)
	}

	for _, name := range result.Generated {
		if !wantGenerated[name] {
			t.Fatalf("unexpected struct %s in generated list %v", name, result.Generated)
		}
	}

	wantSkipped := map[string]string{
		"UnknownChecker":   "no checkergen mapping",
		"NamedType":        "not eligible for generation",
		"SliceField":       "not eligible for generation",
		"BadIntParam":      "not an integer",
		"BadNumberParam":   "not a number",
		"MalformedTwoPart": "two ':'-separated parameters",
		"WrongNumericType": "requires a float64 field",
		"MissingSibling":   "field Password not found",
	}

	if len(result.Skipped) != len(wantSkipped) {
		t.Fatalf("actual %v expected %v", result.Skipped, wantSkipped)
	}

	for name, substr := range wantSkipped {
		reason, ok := result.Skipped[name]
		if !ok {
			t.Fatalf("expected %s to be skipped, got %v", name, result.Skipped)
		}

		if !strings.Contains(reason, substr) {
			t.Fatalf("%s: actual reason %q does not contain %q", name, reason, substr)
		}
	}
}

// TestGenerateTypeFilter confirms -type restricts generation to the named
// struct(s), leaving other eligible structs in the package alone. Runs
// against a scratch copy of testdata/fixture, since a filtered Generate
// call would otherwise overwrite that package's committed, shared
// checkergen_generated.go with a partial result.
func TestGenerateTypeFilter(t *testing.T) {
	scratch := copyToScratch(t, "testdata/fixture")

	result, err := checkergen.Generate(scratch, []string{"SignupRequest"})
	if err != nil {
		t.Fatal(err)
	}

	if len(result.Generated) != 1 || result.Generated[0] != "SignupRequest" {
		t.Fatalf("expected only SignupRequest to be generated, got %v", result.Generated)
	}

	data, err := os.ReadFile(filepath.Join(scratch, checkergen.GeneratedFileName))
	if err != nil {
		t.Fatal(err)
	}

	if strings.Contains(string(data), "CheckCoverage") {
		t.Fatal("expected the type filter to exclude Coverage")
	}
}

// TestGenerateNoEligibleStructs confirms Generate is a no-op (no output
// file, no error) for a package with no checkers/validate tag anywhere.
func TestGenerateNoEligibleStructs(t *testing.T) {
	result, err := checkergen.Generate("testdata/empty", nil)
	if err != nil {
		t.Fatal(err)
	}

	if len(result.Generated) != 0 || len(result.Skipped) != 0 {
		t.Fatalf("expected no generated or skipped structs, got %+v", result)
	}

	if _, err := os.Stat(filepath.Join("testdata/empty", checkergen.GeneratedFileName)); err == nil {
		t.Fatal("expected no output file to be written")
	}
}

// TestGenerateBadDirectory confirms Generate surfaces a package load
// failure as a returned error instead of panicking.
func TestGenerateBadDirectory(t *testing.T) {
	if _, err := checkergen.Generate("testdata/does-not-exist", nil); err == nil {
		t.Fatal("expected an error for a nonexistent directory")
	}
}

// TestGenerateWriteFailure confirms Generate surfaces an output-write
// failure as a returned error: dir itself, passed as the output file's
// parent directory, can't be created as a regular file.
func TestGenerateWriteFailure(t *testing.T) {
	scratch := copyToScratch(t, "testdata/fixture")

	blocker := filepath.Join(scratch, checkergen.GeneratedFileName)
	if err := os.Mkdir(blocker, 0o755); err != nil {
		t.Fatal(err)
	}

	if _, err := checkergen.Generate(scratch, nil); err == nil {
		t.Fatal("expected an error writing over a directory")
	}
}

// TestGenerateMatchesCommittedFixture regenerates testdata/fixture's
// checkergen_generated.go into a scratch copy of the fixture package and
// byte-compares the result against the version committed to the repo.
// Go's build model compiles a test binary before running it, so a test
// can't regenerate a file and then exercise the newly generated code in
// the same run (see TestGeneratedMatchesCheckStruct, which exercises the
// committed copy directly) -- this test instead exists to catch the
// committed fixture drifting out of sync with the generator itself, e.g.
// after a callSpecs change that wasn't followed by re-running go generate
// in testdata/fixture.
func TestGenerateMatchesCommittedFixture(t *testing.T) {
	const fixtureDir = "testdata/fixture"

	committed, err := os.ReadFile(filepath.Join(fixtureDir, checkergen.GeneratedFileName))
	if err != nil {
		t.Fatal(err)
	}

	scratch := copyToScratch(t, fixtureDir)

	result, err := checkergen.Generate(scratch, nil)
	if err != nil {
		t.Fatal(err)
	}

	if len(result.Skipped) != 0 {
		t.Fatalf("unexpected skips regenerating the fixture: %v", result.Skipped)
	}

	regenerated, err := os.ReadFile(filepath.Join(scratch, checkergen.GeneratedFileName))
	if err != nil {
		t.Fatal(err)
	}

	if string(regenerated) != string(committed) {
		t.Fatalf(
			"testdata/fixture/%s is out of date with the current generator; "+
				"run `go run ./cmd/checkergen %s` and commit the result.\n\ngot:\n%s\n\nwant:\n%s",
			checkergen.GeneratedFileName, fixtureDir, regenerated, committed,
		)
	}
}
