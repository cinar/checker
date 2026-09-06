# Checker for Fiber

This package integrates [Checker](https://github.com/cinar/checker) with the [Fiber](https://github.com/gofiber/fiber) web framework. It binds a request body and runs Checker's struct tag validation in a single call, writing a JSON `400` response automatically when binding or validation fails.

## Install

```bash
go get github.com/cinar/checker/v2/fiber
```

## Usage

```golang
package main

import (
	"github.com/gofiber/fiber/v3"

	checkerfiber "github.com/cinar/checker/v2/fiber"
)

type Registration struct {
	Name  string `json:"name" checkers:"trim required"`
	Email string `json:"email" checkers:"required email"`
}

func main() {
	app := fiber.New()

	app.Post("/register", func(c fiber.Ctx) error {
		var registration Registration

		if !checkerfiber.Bind(c, &registration) {
			// The 400 response has already been written by Bind.
			return nil
		}

		return c.JSON(registration)
	})

	app.Listen(":8080")
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
app.Post("/register/:plan", func(c fiber.Ctx) error {
	registration := Registration{
		Plan: c.Params("plan"),
	}

	if err := c.Bind().Body(&registration); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if !checkerfiber.Check(c, &registration) {
		return nil
	}

	return c.JSON(registration)
})
```

## License

This library is free to use, modify, and distribute under the terms of the MIT license found in the [LICENSE](../LICENSE) file of the parent [Checker](https://github.com/cinar/checker) repository.
