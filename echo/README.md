# Checker for Echo

This package integrates [Checker](https://github.com/cinar/checker) with the [Echo](https://github.com/labstack/echo) web framework. It binds a request and runs Checker's struct tag validation in a single call, writing a JSON `400` response automatically when binding or validation fails.

## Install

```bash
go get github.com/cinar/checker/v2/echo
```

## Usage

```golang
package main

import (
	"net/http"

	"github.com/labstack/echo/v4"

	checkerecho "github.com/cinar/checker/v2/echo"
)

type Registration struct {
	Name  string `json:"name" checkers:"trim required"`
	Email string `json:"email" checkers:"required email"`
}

func main() {
	e := echo.New()

	e.POST("/register", func(c echo.Context) error {
		var registration Registration

		if !checkerecho.Bind(c, &registration) {
			// The 400 response has already been written by Bind.
			return nil
		}

		return c.JSON(http.StatusOK, registration)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
```

A `POST /register` with an invalid body, such as `{"name":"","email":""}`, gets back:

```json
{
	"name": {"code": "REQUIRED", "message": "Required value is missing."},
	"email": {"code": "REQUIRED", "message": "Required value is missing."}
}
```

If you have already bound the request some other way, or you want to validate a struct assembled from more than one source (path params and body, for example), call `Check` directly instead:

```golang
e.POST("/register/:plan", func(c echo.Context) error {
	registration := Registration{
		Plan: c.Param("plan"),
	}

	if err := c.Bind(&registration); err != nil {
		return err
	}

	if !checkerecho.Check(c, &registration) {
		return nil
	}

	return c.JSON(http.StatusOK, registration)
})
```

## License

This library is free to use, modify, and distribute under the terms of the MIT license found in the [LICENSE](../LICENSE) file of the parent [Checker](https://github.com/cinar/checker) repository.
