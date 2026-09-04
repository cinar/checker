# Checker for Gin

This package integrates [Checker](https://github.com/cinar/checker) with the [Gin](https://github.com/gin-gonic/gin) web framework. It binds a request body and runs Checker's struct tag validation in a single call, writing a JSON `400` response automatically when binding or validation fails.

## Install

```bash
go get github.com/cinar/checker/v2/gin
```

## Usage

```golang
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	checkergin "github.com/cinar/checker/v2/gin"
)

type Registration struct {
	Name  string `json:"name" checkers:"trim required"`
	Email string `json:"email" checkers:"required email"`
}

func main() {
	router := gin.Default()

	router.POST("/register", func(c *gin.Context) {
		var registration Registration

		if !checkergin.Bind(c, &registration) {
			// The 400 response has already been written by Bind.
			return
		}

		c.JSON(http.StatusOK, registration)
	})

	router.Run()
}
```

A `POST /register` with an invalid body, such as `{"name":"","email":""}`, gets back:

```json
{
	"name": {"code": "REQUIRED", "message": "Required value is missing."},
	"email": {"code": "REQUIRED", "message": "Required value is missing."}
}
```

If you have already bound the request body some other way, or you want to validate a struct assembled from more than one source (path params, query string, and body, for example), call `Check` directly instead:

```golang
router.POST("/register/:plan", func(c *gin.Context) {
	registration := Registration{
		Plan: c.Param("plan"),
	}

	if err := c.ShouldBindJSON(&registration); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !checkergin.Check(c, &registration) {
		return
	}

	c.JSON(http.StatusOK, registration)
})
```

## License

This library is free to use, modify, and distribute under the terms of the MIT license found in the [LICENSE](../LICENSE) file of the parent [Checker](https://github.com/cinar/checker) repository.
