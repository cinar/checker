// Copyright (c) 2024-2026 Onur Cinar. All Rights Reserved.
// The source code is provided under MIT License.
//
// https://github.com/cinar/checker
//
// Try this on Go Playground: https://go.dev/play/p/U04Du4M6spX

package main

import (
	"encoding/json"
	"fmt"

	checker "github.com/cinar/checker/v2"
)

type Product struct {
	SKU     string   `json:"sku" checkers:"required alphanumeric min-len:6 max-len:12"`
	Name    string   `json:"name" checkers:"trim required"`
	Price   float64  `json:"price" checkers:"gte:0"`
	Website string   `json:"website,omitempty" checkers:"url"`
	Tags    []string `json:"tags" checkers:"@max-len:5 trim alphanumeric"`
}

func main() {
	// Generate Draft 2020-12 JSON Schema directly from struct tags:
	schema := checker.JSONSchema(&Product{})

	data, err := json.MarshalIndent(schema, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println("Generated JSON Schema for Product:")
	fmt.Println(string(data))
}
