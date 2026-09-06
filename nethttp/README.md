# Checker for net/http

This package integrates [Checker](https://github.com/cinar/checker) with the standard library's `net/http`. It decodes a JSON request body and runs Checker's struct tag validation in a single call, writing a JSON `400` response automatically when decoding or validation fails. It has no external dependencies beyond the core `checker` module: only `encoding/json` and `net/http`, both standard library.

## Install

```bash
go get github.com/cinar/checker/v2/nethttp
```

## Usage

```golang
package main

import (
	"encoding/json"
	"net/http"

	checkernethttp "github.com/cinar/checker/v2/nethttp"
)

type Registration struct {
	Name  string `json:"name" checkers:"trim required"`
	Email string `json:"email" checkers:"required email"`
}

func main() {
	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		var registration Registration

		if !checkernethttp.Bind(w, r, &registration) {
			// The 400 response has already been written by Bind.
			return
		}

		json.NewEncoder(w).Encode(registration)
	})

	http.ListenAndServe(":8080", nil)
}
```

A `POST /register` with an invalid body, such as `{"name":"","email":""}`, gets back:

```json
{
	"name": {"code": "REQUIRED", "message": "Required value is missing."},
	"email": {"code": "REQUIRED", "message": "Required value is missing."}
}
```

If you have already decoded the request body some other way, or you want to validate a struct assembled from more than one source (path params, query string, and body, for example), call `Check` directly instead:

```golang
http.HandleFunc("/register/", func(w http.ResponseWriter, r *http.Request) {
	registration := Registration{
		Plan: r.PathValue("plan"),
	}

	if err := json.NewDecoder(r.Body).Decode(&registration); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !checkernethttp.Check(w, &registration) {
		return
	}

	json.NewEncoder(w).Encode(registration)
})
```

## License

This library is free to use, modify, and distribute under the terms of the MIT license found in the [LICENSE](../LICENSE) file of the parent [Checker](https://github.com/cinar/checker) repository.
