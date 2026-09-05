// Copyright (c) 2024-2026 Onur Cinar. All Rights Reserved.
// The source code is provided under MIT License.
//
// https://github.com/cinar/checker

package main

import (
	"net/http"

	"github.com/labstack/echo/v4"

	checkerecho "github.com/cinar/checker/v2/echo"
)

type Registration struct {
	Name            string `json:"name" checkers:"trim required"`
	Email           string `json:"email" checkers:"trim lower required email"`
	Password        string `json:"password" checkers:"required min-len:8"`
	ConfirmPassword string `json:"confirm_password" checkers:"required eq-field:Password"`
}

func main() {
	e := echo.New()

	e.POST("/register", func(c echo.Context) error {
		var registration Registration

		// checkerecho.Bind binds the request body and runs CheckStruct.
		// If binding or validation fails, it automatically writes an HTTP 400 response
		// with structured JSON errors.
		if !checkerecho.Bind(c, &registration) {
			return nil
		}

		return c.JSON(http.StatusOK, map[string]any{
			"message": "User registered successfully",
			"user":    registration,
		})
	})

	e.Start(":8080")
}
