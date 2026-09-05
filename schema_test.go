// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2_test

import (
	"encoding/json"
	"fmt"
	"testing"

	v2 "github.com/cinar/checker/v2"
)

func TestJSONSchemaStructFields(t *testing.T) {
	type Person struct {
		Name    string `json:"name" checkers:"trim required min-len:1 max-len:64"`
		Email   string `checkers:"required email"`
		Age     int    `checkers:"gte:0 lte:150"`
		Website string `checkers:"url"`
		Zip     string `checkers:"regexp:^[0-9]{5}$"`
		Ignored string `json:"-"`
		Comma   string `json:"comma,omitempty" checkers:"required"`
		Empty   string `json:""`
		NoName  string `json:",omitempty"`
		private string
	}

	person := &Person{private: "unexported"}

	schema := v2.JSONSchema(person)

	if person.private != "unexported" {
		t.Fatal("expected the unexported field to be left untouched")
	}

	if schema.Dialect == "" {
		t.Fatal("expected root schema to declare a dialect")
	}

	if schema.Title != "Person" {
		t.Fatalf("actual %s expected Person", schema.Title)
	}

	if schema.Type != "object" {
		t.Fatalf("actual %s expected object", schema.Type)
	}

	if _, ok := schema.Properties["Ignored"]; ok {
		t.Fatal("expected json:\"-\" field to be excluded")
	}

	if _, ok := schema.Properties["private"]; ok {
		t.Fatal("expected unexported field to be excluded")
	}

	name, ok := schema.Properties["name"]
	if !ok {
		t.Fatal("expected name property")
	}

	if name.Type != "string" || *name.MinLength != 1 || *name.MaxLength != 64 {
		t.Fatalf("unexpected name schema %+v", name)
	}

	email, ok := schema.Properties["Email"]
	if !ok || email.Format != "email" {
		t.Fatalf("unexpected email schema %+v", email)
	}

	age, ok := schema.Properties["Age"]
	if !ok || age.Type != "integer" || *age.Minimum != 0 || *age.Maximum != 150 {
		t.Fatalf("unexpected age schema %+v", age)
	}

	website, ok := schema.Properties["Website"]
	if !ok || website.Format != "uri" {
		t.Fatalf("unexpected website schema %+v", website)
	}

	zip, ok := schema.Properties["Zip"]
	if !ok || zip.Pattern != "^[0-9]{5}$" {
		t.Fatalf("unexpected zip schema %+v", zip)
	}

	if _, ok := schema.Properties["comma"]; !ok {
		t.Fatal("expected json tag with options to still resolve a name")
	}

	if _, ok := schema.Properties["Empty"]; !ok {
		t.Fatal("expected an empty json tag to fall back to the field name")
	}

	if _, ok := schema.Properties["NoName"]; !ok {
		t.Fatal("expected a json tag with no name before the comma to fall back to the field name")
	}

	expectedRequired := []string{"Email", "comma", "name"}

	if len(schema.Required) != len(expectedRequired) {
		t.Fatalf("actual %v expected %v", schema.Required, expectedRequired)
	}

	for i, name := range expectedRequired {
		if schema.Required[i] != name {
			t.Fatalf("actual %v expected %v", schema.Required, expectedRequired)
		}
	}
}

func TestJSONSchemaBoolAndFloat(t *testing.T) {
	type Settings struct {
		Enabled bool
		Ratio   float64
	}

	schema := v2.JSONSchema(Settings{})

	if schema.Properties["Enabled"].Type != "boolean" {
		t.Fatalf("actual %s expected boolean", schema.Properties["Enabled"].Type)
	}

	if schema.Properties["Ratio"].Type != "number" {
		t.Fatalf("actual %s expected number", schema.Properties["Ratio"].Type)
	}
}

func TestJSONSchemaUnknownKind(t *testing.T) {
	type Weird struct {
		Callback func()
	}

	schema := v2.JSONSchema(Weird{})

	if schema.Properties["Callback"].Type != "" {
		t.Fatalf("expected no type for an unmappable kind, got %s", schema.Properties["Callback"].Type)
	}
}

func TestJSONSchemaNestedStruct(t *testing.T) {
	type Address struct {
		Street string `checkers:"required"`
	}

	type Person struct {
		Name    string `checkers:"required"`
		Address *Address
	}

	schema := v2.JSONSchema(&Person{})

	address, ok := schema.Properties["Address"]
	if !ok || address.Type != "object" {
		t.Fatalf("unexpected address schema %+v", address)
	}

	street, ok := address.Properties["Street"]
	if !ok || len(address.Required) != 1 || address.Required[0] != "Street" {
		t.Fatalf("unexpected address required %+v", address)
	}

	if street.Type != "string" {
		t.Fatalf("unexpected street schema %+v", street)
	}
}

func TestJSONSchemaSlice(t *testing.T) {
	type Person struct {
		Emails []string `checkers:"@min-len:1 @max-len:5 email"`
	}

	schema := v2.JSONSchema(&Person{})

	emails := schema.Properties["Emails"]

	if emails.Type != "array" {
		t.Fatalf("actual %s expected array", emails.Type)
	}

	if *emails.MinItems != 1 || *emails.MaxItems != 5 {
		t.Fatalf("unexpected emails schema %+v", emails)
	}

	if emails.Items.Type != "string" || emails.Items.Format != "email" {
		t.Fatalf("unexpected emails item schema %+v", emails.Items)
	}
}

func TestJSONSchemaSliceOfStruct(t *testing.T) {
	type Item struct {
		Name string `checkers:"required"`
	}

	type Order struct {
		Items []Item
	}

	schema := v2.JSONSchema(&Order{})

	items := schema.Properties["Items"]

	if items.Type != "array" || items.Items.Type != "object" {
		t.Fatalf("unexpected items schema %+v", items)
	}

	if len(items.Items.Required) != 1 || items.Items.Required[0] != "Name" {
		t.Fatalf("unexpected items.Items required %+v", items.Items)
	}
}

func TestJSONSchemaMap(t *testing.T) {
	type Person struct {
		Tags map[string]string `checkers:"@min-len:2 @max-len:10 max-len:32"`
	}

	schema := v2.JSONSchema(&Person{})

	tags := schema.Properties["Tags"]

	if tags.Type != "object" {
		t.Fatalf("actual %s expected object", tags.Type)
	}

	if *tags.MinProperties != 2 || *tags.MaxProperties != 10 {
		t.Fatalf("unexpected tags schema %+v", tags)
	}

	if tags.AdditionalProperties.Type != "string" || *tags.AdditionalProperties.MaxLength != 32 {
		t.Fatalf("unexpected tags value schema %+v", tags.AdditionalProperties)
	}
}

func TestJSONSchemaXChecker(t *testing.T) {
	type Registration struct {
		Password        string `checkers:"required"`
		ConfirmPassword string `checkers:"eq-field:Password"`
	}

	schema := v2.JSONSchema(&Registration{})

	confirm := schema.Properties["ConfirmPassword"]

	if len(confirm.XChecker) != 1 || confirm.XChecker[0] != "eq-field:Password" {
		t.Fatalf("unexpected x-checker %+v", confirm.XChecker)
	}
}

func TestJSONSchemaUUIDFormat(t *testing.T) {
	type Resource struct {
		ID string `checkers:"uuid"`
	}

	schema := v2.JSONSchema(&Resource{})

	id, ok := schema.Properties["ID"]
	if !ok || id.Format != "uuid" {
		t.Fatalf("unexpected id schema %+v", id)
	}
}

func TestJSONSchemaOneOfEnum(t *testing.T) {
	type Role struct {
		Name string `checkers:"oneof:admin,user,guest"`
	}

	schema := v2.JSONSchema(&Role{})

	name, ok := schema.Properties["Name"]
	if !ok {
		t.Fatal("expected Name property")
	}

	expected := []string{"admin", "user", "guest"}

	if len(name.Enum) != len(expected) {
		t.Fatalf("actual %v expected %v", name.Enum, expected)
	}

	for i, v := range expected {
		if name.Enum[i] != v {
			t.Fatalf("actual %v expected %v", name.Enum, expected)
		}
	}
}

func TestJSONSchemaIgnoresNormalizers(t *testing.T) {
	type Person struct {
		Name string `checkers:"trim lower upper title trim-left trim-right html-escape html-unescape url-escape url-unescape omitempty required"`
	}

	schema := v2.JSONSchema(&Person{})

	name := schema.Properties["Name"]

	if len(name.XChecker) != 0 {
		t.Fatalf("expected normalizers to be ignored, got %v", name.XChecker)
	}
}

func TestJSONSchemaNotAStruct(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	v2.JSONSchema("not-a-struct")
}

func TestJSONSchemaBadMinLen(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Person struct {
		Name string `checkers:"min-len:abc"`
	}

	v2.JSONSchema(&Person{})
}

func TestJSONSchemaBadMaxLen(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Person struct {
		Name string `checkers:"max-len:abc"`
	}

	v2.JSONSchema(&Person{})
}

func TestJSONSchemaBadMinimum(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Person struct {
		Age int `checkers:"gte:abc"`
	}

	v2.JSONSchema(&Person{})
}

func TestJSONSchemaBadMaximum(t *testing.T) {
	defer FailIfNoPanic(t, "expected panic")

	type Person struct {
		Age int `checkers:"lte:abc"`
	}

	v2.JSONSchema(&Person{})
}

func TestJSONSchemaMarshalsToJSON(t *testing.T) {
	type Person struct {
		Name string `checkers:"required"`
	}

	schema := v2.JSONSchema(&Person{})

	data, err := json.Marshal(schema)
	if err != nil {
		t.Fatal(err)
	}

	var decoded map[string]interface{}

	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatal(err)
	}

	if decoded["type"] != "object" {
		t.Fatalf("unexpected marshaled schema %s", data)
	}
}

func ExampleJSONSchema() {
	type Person struct {
		Name  string `json:"name" checkers:"trim required"`
		Email string `json:"email" checkers:"required email"`
	}

	schema := v2.JSONSchema(&Person{})

	data, err := json.MarshalIndent(schema, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
	// Output:
	// {
	//   "$schema": "https://json-schema.org/draft/2020-12/schema",
	//   "title": "Person",
	//   "type": "object",
	//   "properties": {
	//     "email": {
	//       "type": "string",
	//       "format": "email"
	//     },
	//     "name": {
	//       "type": "string"
	//     }
	//   },
	//   "required": [
	//     "email",
	//     "name"
	//   ]
	// }
}

func TestRegisterSchemaMaker(t *testing.T) {
	v2.RegisterSchemaMaker("is-fruit", func(schema *v2.Schema, _ string) {
		schema.Enum = []string{"apple", "banana"}
	})

	type Item struct {
		Name string `checkers:"is-fruit"`
	}

	schema := v2.JSONSchema(&Item{})

	name := schema.Properties["Name"]

	if len(name.Enum) != 2 || name.Enum[0] != "apple" || name.Enum[1] != "banana" {
		t.Fatalf("expected custom schema maker to run, got %+v", name)
	}
}
