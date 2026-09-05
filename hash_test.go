// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2_test

import (
	"crypto/md5"  //nolint:gosec // used only to produce a test fixture, not for security
	"crypto/sha1" //nolint:gosec // used only to produce a test fixture, not for security
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"testing"

	v2 "github.com/cinar/checker/v2"
)

func TestIsHashSuccess(t *testing.T) {
	input := []byte("Onur Cinar")

	md5Sum := md5.Sum(input)   //nolint:gosec // test fixture, not a security use
	sha1Sum := sha1.Sum(input) //nolint:gosec // test fixture, not a security use
	sha256Sum := sha256.Sum256(input)
	sha384Sum := sha512.Sum384(input)
	sha512Sum := sha512.Sum512(input)

	tests := []struct {
		algorithm string
		value     string
	}{
		{"md5", hex.EncodeToString(md5Sum[:])},
		{"sha1", hex.EncodeToString(sha1Sum[:])},
		{"sha256", hex.EncodeToString(sha256Sum[:])},
		{"sha384", hex.EncodeToString(sha384Sum[:])},
		{"sha512", hex.EncodeToString(sha512Sum[:])},
	}

	for _, test := range tests {
		value, err := v2.IsHash(test.algorithm, test.value)
		if err != nil {
			t.Fatalf("%s: unexpected error %v", test.algorithm, err)
		}

		if value != test.value {
			t.Fatalf("%s: actual %s expected %s", test.algorithm, value, test.value)
		}
	}
}

func TestIsHashWrongLength(t *testing.T) {
	_, err := v2.IsHash("md5", "not-the-right-length")

	if !errors.Is(err, v2.ErrNotHash) {
		t.Fatalf("got unexpected error %v", err)
	}
}

func TestIsHashNotHex(t *testing.T) {
	_, err := v2.IsHash("md5", "zzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz")

	if !errors.Is(err, v2.ErrNotHash) {
		t.Fatalf("got unexpected error %v", err)
	}
}

func TestIsHashUnknownAlgorithm(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	v2.IsHash("unknown", "abcd")
}

func TestCheckStructHash(t *testing.T) {
	type File struct {
		Checksum string `checkers:"hash:sha256"`
	}

	sum := sha256.Sum256([]byte("Onur Cinar"))

	file := &File{
		Checksum: "not-a-valid-checksum",
	}

	errs, ok := v2.CheckStruct(file)
	if ok {
		t.Fatal("expected errors")
	}

	if !errors.Is(errs["Checksum"], v2.ErrNotHash) {
		t.Fatalf("expected checksum not hash %v", errs)
	}

	file.Checksum = hex.EncodeToString(sum[:])

	errs, ok = v2.CheckStruct(file)
	if !ok {
		t.Fatalf("got unexpected errors %v", errs)
	}
}
