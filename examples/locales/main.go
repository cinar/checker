// Copyright (c) 2024-2026 Onur Cinar. All Rights Reserved.
// The source code is provided under MIT License.
//
// https://github.com/cinar/checker
//
// Try this on Go Playground: https://go.dev/play/p/GQEENYcQPgD

package main

import (
	"fmt"

	checker "github.com/cinar/checker/v2"
	"github.com/cinar/checker/v2/locales"
)

type User struct {
	Email string `checkers:"required email"`
	Age   int    `checkers:"required gte:18"`
}

func main() {
	// Register opt-in locales (no unused translation bloat in core):
	checker.RegisterLocale(locales.DeDE, locales.DeDEMessages)
	checker.RegisterLocale(locales.EsES, locales.EsESMessages)
	checker.RegisterLocale(locales.FrFR, locales.FrFRMessages)
	checker.RegisterLocale(locales.JaJP, locales.JaJPMessages)

	invalidUser := &User{
		Email: "not-an-email",
		Age:   15,
	}

	errs, ok := checker.CheckStruct(invalidUser)
	if ok {
		return
	}

	fmt.Println("=== Default Locale (en-US) ===")
	jsonEn, _ := errs.JSON()
	fmt.Println(string(jsonEn))

	fmt.Println("\n=== German (de-DE) ===")
	jsonDe, _ := errs.JSONWithLocale(locales.DeDE)
	fmt.Println(string(jsonDe))

	fmt.Println("\n=== Spanish (es-ES) ===")
	jsonEs, _ := errs.JSONWithLocale(locales.EsES)
	fmt.Println(string(jsonEs))

	fmt.Println("\n=== Japanese (ja-JP) ===")
	jsonJa, _ := errs.JSONWithLocale(locales.JaJP)
	fmt.Println(string(jsonJa))
}
