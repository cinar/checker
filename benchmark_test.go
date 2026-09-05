// Copyright (c) 2024-2026 Onur Cinar. All Rights Reserved.
// The source code is provided under MIT License.
//
// https://github.com/cinar/checker

package v2_test

import (
	"testing"

	checker "github.com/cinar/checker/v2"
)

type BenchmarkSimple struct {
	Name  string `checkers:"trim required"`
	Email string `checkers:"trim required email"`
	Age   int    `checkers:"gte:18"`
}

type BenchmarkAddress struct {
	Street  string `checkers:"trim required"`
	City    string `checkers:"trim required"`
	ZipCode string `checkers:"required digits"`
}

type BenchmarkComplex struct {
	ID              string            `checkers:"required hex"`
	Email           string            `checkers:"trim lower required email"`
	Password        string            `checkers:"required min-len:8"`
	ConfirmPassword string            `checkers:"required eq-field:Password"`
	Website         string            `checkers:"url"`
	Country         string            `checkers:"required iso3166-1-alpha-2"`
	State           string            `checkers:"required-if:Country:US"`
	Tags            []string          `checkers:"@max-len:5 trim alphanumeric"`
	Metadata        map[string]string `checkers:"@max-len:10 trim"`
	Address         BenchmarkAddress
}

func BenchmarkCheckStruct_Simple(b *testing.B) {
	user := BenchmarkSimple{
		Name:  "  Alice Smith  ",
		Email: "alice@example.com",
		Age:   25,
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		u := user
		_, _ = checker.CheckStruct(&u)
	}
}

func BenchmarkCheckStruct_Complex(b *testing.B) {
	user := BenchmarkComplex{
		ID:              "a1b2c3d4e5f6",
		Email:           "  ALICE@EXAMPLE.COM  ",
		Password:        "supersecret123",
		ConfirmPassword: "supersecret123",
		Website:         "https://example.com",
		Country:         "US",
		State:           "CA",
		Tags:            []string{"golang", "backend", "api"},
		Metadata:        map[string]string{"env": "prod"},
		Address: BenchmarkAddress{
			Street:  "123 Main St",
			City:    "San Jose",
			ZipCode: "95112",
		},
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		u := user
		_, _ = checker.CheckStruct(&u)
	}
}

func BenchmarkJSONSchema(b *testing.B) {
	st := &BenchmarkComplex{}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = checker.JSONSchema(st)
	}
}
