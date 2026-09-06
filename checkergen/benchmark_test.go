// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package checkergen_test

import (
	"testing"

	checker "github.com/cinar/checker/v2"
	"github.com/cinar/checker/v2/checkergen/testdata/fixture"
)

func BenchmarkCheckStruct_SignupRequest(b *testing.B) {
	req := fixture.SignupRequest{
		Email:           "  ALICE@EXAMPLE.COM  ",
		Password:        "supersecret",
		ConfirmPassword: "supersecret",
		Age:             30,
		Nickname:        "ali",
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		r := req
		_, _ = checker.CheckStruct(&r)
	}
}

func BenchmarkGenerated_SignupRequest(b *testing.B) {
	req := fixture.SignupRequest{
		Email:           "  ALICE@EXAMPLE.COM  ",
		Password:        "supersecret",
		ConfirmPassword: "supersecret",
		Age:             30,
		Nickname:        "ali",
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		r := req
		_, _ = fixture.CheckSignupRequest(&r)
	}
}

func BenchmarkCheckStruct_Coverage(b *testing.B) {
	cov := fixture.Coverage{
		Handle:    "@onur",
		Code:      "ABC",
		Sum:       "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
		Zip:       "94103",
		Status:    "active",
		Excluded:  "allowed",
		Role:      "admin",
		Weight:    10,
		Ratio:     4,
		Threshold: 0,
		Country:   "US",
		State:     "CA",
		Type:      "member",
		Email:     "onur@example.com",
		ReturnAt:  "2024-06-01",
		DepartAt:  "2024-01-01",
		BornAt:    "2000-01-01",
		DiesAt:    "2090-01-01",
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		c := cov
		_, _ = checker.CheckStruct(&c)
	}
}

func BenchmarkGenerated_Coverage(b *testing.B) {
	cov := fixture.Coverage{
		Handle:    "@onur",
		Code:      "ABC",
		Sum:       "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
		Zip:       "94103",
		Status:    "active",
		Excluded:  "allowed",
		Role:      "admin",
		Weight:    10,
		Ratio:     4,
		Threshold: 0,
		Country:   "US",
		State:     "CA",
		Type:      "member",
		Email:     "onur@example.com",
		ReturnAt:  "2024-06-01",
		DepartAt:  "2024-01-01",
		BornAt:    "2000-01-01",
		DiesAt:    "2090-01-01",
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		c := cov
		_, _ = fixture.CheckCoverage(&c)
	}
}
