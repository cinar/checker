# Checker Examples

This directory contains standalone, runnable examples demonstrating the capabilities of [Checker](https://github.com/cinar/checker).

## Examples Overview

| Example | Description | Run Locally | In-Browser Playground |
| :--- | :--- | :--- | :---: |
| [**Basic**](./basic/main.go) | Normalization in-place (`trim`, `lower`), cross-field validation (`eq-field`), slice rules, and JSON error output. | `go run ./basic` | [**Try on Go Playground**](https://go.dev/play/p/c9ohvThGz1D) |
| [**JSON Schema**](./jsonschema/main.go) | Generating Draft 2020-12 JSON Schema documents directly from struct validation tags. | `go run ./jsonschema` | [**Try on Go Playground**](https://go.dev/play/p/U04Du4M6spX) |
| [**Locales (i18n)**](./locales/main.go) | Opt-in localization supporting 23 languages without adding bloat to the core module. | `go run ./locales` | [**Try on Go Playground**](https://go.dev/play/p/GQEENYcQPgD) |
| [**HTTP Server**](./http/main.go) | Standard library `net/http` handler validating requests and returning structured JSON errors with zero dependencies. | `go run ./http` | [**Try on Go Playground**](https://go.dev/play/p/M_tKEKwL38G) |
| [**Gin Integration**](./gin/main.go) | Binding and validating requests in one call with the `checker/v2/gin` adapter. | `go run ./gin` | — |
| [**Echo Integration**](./echo/main.go) | Binding and validating requests in one call with the `checker/v2/echo` adapter. | `go run ./echo` | — |

## Running Locally

To run any of the examples locally:

```bash
cd examples

# Run the basic example:
go run ./basic

# Run the JSON Schema example:
go run ./jsonschema

# Run the Locales example:
go run ./locales

# Run the net/http server example:
go run ./http

# Run the Gin server example:
go run ./gin

# Run the Echo server example:
go run ./echo
```
