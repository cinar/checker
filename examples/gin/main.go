// Copyright (c) 2024-2026 Onur Cinar. All Rights Reserved.
// The source code is provided under MIT License.
//
// https://github.com/cinar/checker

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	checkergin "github.com/cinar/checker/v2/gin"
)

type Registration struct {
	Name            string `json:"name" checkers:"trim required"`
	Email           string `json:"email" checkers:"trim lower required email"`
	Password        string `json:"password" checkers:"required min-len:8"`
	ConfirmPassword string `json:"confirm_password" checkers:"required eq-field:Password"`
}

func main() {
	router := gin.Default()

	router.POST("/register", func(c *gin.Context) {
		var registration Registration

		// checkergin.Bind binds the JSON request body and runs CheckStruct.
		// If binding or validation fails, it automatically responds with HTTP 400 Bad Request
		// containing structured JSON errors.
		if !checkergin.Bind(c, &registration) {
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "User registered successfully",
			"user":    registration,
		})
	})

	router.Run(":8080")
}
