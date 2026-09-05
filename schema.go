// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

import (
	"reflect"
	"sort"
	"strings"
)

// jsonSchemaDialect is the JSON Schema dialect declared by the root Schema
// returned from JSONSchema.
const jsonSchemaDialect = "https://json-schema.org/draft/2020-12/schema"

// Schema is a JSON Schema document, or subschema, generated from a struct's
// checker tags. It covers the common validation keywords that map cleanly
// from checker tags, not the full JSON Schema specification.
type Schema struct {
	// Dialect identifies the JSON Schema dialect. Only set on the root Schema
	// returned from JSONSchema, not on nested subschemas.
	Dialect string `json:"$schema,omitempty"`

	// Title is the name of the Go type the schema was generated from.
	Title string `json:"title,omitempty"`

	// Type is the JSON Schema type: "object", "array", "string", "integer",
	// "number", or "boolean".
	Type string `json:"type,omitempty"`

	// Properties maps a struct field's JSON name to its Schema. Only set when
	// Type is "object" and generated from a struct.
	Properties map[string]*Schema `json:"properties,omitempty"`

	// Required lists the property names, sorted, whose checkers tag includes
	// "required". Only set when Type is "object" and generated from a struct.
	Required []string `json:"required,omitempty"`

	// AdditionalProperties is the Schema every value in a map must satisfy.
	// Only set when Type is "object" and generated from a map.
	AdditionalProperties *Schema `json:"additionalProperties,omitempty"`

	// Items is the Schema every element in an array must satisfy. Only set
	// when Type is "array".
	Items *Schema `json:"items,omitempty"`

	// Format is a JSON Schema format hint, such as "email" or "uri".
	Format string `json:"format,omitempty"`

	// Pattern is a regular expression a string value must match.
	Pattern string `json:"pattern,omitempty"`

	// Enum lists the only values a value may hold. Not set by any built-in
	// checker, since checkers like iso639-1 back it with a code list too
	// large to be worth inlining into every generated schema; register a
	// SchemaMakeFunc via RegisterSchemaMaker to populate it for a custom
	// checker instead.
	Enum []string `json:"enum,omitempty"`

	// MinLength is the minimum length of a string value.
	MinLength *int `json:"minLength,omitempty"`

	// MaxLength is the maximum length of a string value.
	MaxLength *int `json:"maxLength,omitempty"`

	// MinItems is the minimum number of elements in an array value.
	MinItems *int `json:"minItems,omitempty"`

	// MaxItems is the maximum number of elements in an array value.
	MaxItems *int `json:"maxItems,omitempty"`

	// MinProperties is the minimum number of entries in a map value.
	MinProperties *int `json:"minProperties,omitempty"`

	// MaxProperties is the maximum number of entries in a map value.
	MaxProperties *int `json:"maxProperties,omitempty"`

	// Minimum is the smallest value a numeric value can hold.
	Minimum *float64 `json:"minimum,omitempty"`

	// Maximum is the largest value a numeric value can hold.
	Maximum *float64 `json:"maximum,omitempty"`

	// ExclusiveMinimum is the smallest value a numeric value can hold,
	// exclusive: the value itself doesn't satisfy the bound.
	ExclusiveMinimum *float64 `json:"exclusiveMinimum,omitempty"`

	// ExclusiveMaximum is the largest value a numeric value can hold,
	// exclusive: the value itself doesn't satisfy the bound.
	ExclusiveMaximum *float64 `json:"exclusiveMaximum,omitempty"`

	// XChecker lists the raw checker tag tokens that have no JSON Schema
	// equivalent, such as "eq-field:Password" or "luhn", so the constraint
	// isn't silently lost even though it isn't translated.
	XChecker []string `json:"x-checker,omitempty"`
}

// schemaField carries a field's Schema along with whether its checkers tag
// included "required", since "required" contributes to the parent object's
// Required list rather than the field's own Schema.
type schemaField struct {
	schema   *Schema
	required bool
}

// JSONSchema generates a JSON Schema document describing the shape and
// validation rules declared in st's "checkers" struct tags. st must be a
// struct, or a pointer to one. Panics if a checker's parameter cannot be
// parsed, same as CheckStruct, since that indicates a struct tag typo rather
// than a data problem.
//
// JSONSchema is a static, type-level operation: it never inspects st's
// field values, only its type and tags, so a zero value works just as well
// as a populated one.
func JSONSchema(st any) *Schema {
	t := reflect.TypeOf(st)

	for t.Kind() == reflect.Pointer {
		t = t.Elem()
	}

	if t.Kind() != reflect.Struct {
		panic("JSONSchema requires a struct or a pointer to a struct")
	}

	schema := structSchema(t)
	schema.Dialect = jsonSchemaDialect

	return schema
}

// typeSchema generates the Schema for the given type and checkers config,
// dereferencing pointers first.
func typeSchema(t reflect.Type, config string) schemaField {
	for t.Kind() == reflect.Pointer {
		t = t.Elem()
	}

	switch t.Kind() {
	case reflect.Struct:
		schema := structSchema(t)

		field := schemaField{schema: schema}
		applyConfig(schema, config, &field)

		return field

	case reflect.Slice, reflect.Array:
		return sliceSchema(t, config)

	case reflect.Map:
		return mapSchema(t, config)

	default:
		return scalarSchema(t, config)
	}
}

// structSchema generates an "object" Schema for the given struct type, with
// a property for each exported field.
func structSchema(t reflect.Type) *Schema {
	schema := &Schema{
		Title:      t.Name(),
		Type:       "object",
		Properties: make(map[string]*Schema),
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		if field.PkgPath != "" {
			// Unexported field, never part of the JSON representation.
			continue
		}

		name, ok := jsonPropertyName(field)
		if !ok {
			continue
		}

		property := typeSchema(field.Type, field.Tag.Get(checkerTag))

		schema.Properties[name] = property.schema

		if property.required {
			schema.Required = append(schema.Required, name)
		}
	}

	sort.Strings(schema.Required)

	return schema
}

// sliceSchema generates an "array" Schema for the given slice or array type.
// An "@"-prefixed checker in config applies to the array itself; the rest
// apply to each element.
func sliceSchema(t reflect.Type, config string) schemaField {
	arrayConfig, itemConfig := splitSliceConfig(config)

	item := typeSchema(t.Elem(), itemConfig)

	schema := &Schema{
		Type:  "array",
		Items: item.schema,
	}

	field := schemaField{schema: schema}
	applyConfig(schema, arrayConfig, &field)

	return field
}

// mapSchema generates an "object" Schema for the given map type, using
// AdditionalProperties to describe the shape every value must satisfy. An
// "@"-prefixed checker in config applies to the map itself; the rest apply
// to each value.
func mapSchema(t reflect.Type, config string) schemaField {
	mapConfig, itemConfig := splitSliceConfig(config)

	item := typeSchema(t.Elem(), itemConfig)

	schema := &Schema{
		Type:                 "object",
		AdditionalProperties: item.schema,
	}

	field := schemaField{schema: schema}
	applyConfig(schema, mapConfig, &field)

	return field
}

// scalarSchema generates a Schema for a leaf (non-struct, non-slice,
// non-map) type, applying every checker in config to it.
func scalarSchema(t reflect.Type, config string) schemaField {
	schema := &Schema{Type: jsonTypeForKind(t.Kind())}

	field := schemaField{schema: schema}
	applyConfig(schema, config, &field)

	return field
}

// jsonTypeForKind returns the JSON Schema type name for the given Go kind,
// or an empty string if there isn't a direct equivalent.
func jsonTypeForKind(kind reflect.Kind) string {
	switch {
	case kind == reflect.String:
		return "string"
	case kind == reflect.Bool:
		return "boolean"
	case kind >= reflect.Int && kind <= reflect.Uint64:
		return "integer"
	case kind == reflect.Float32 || kind == reflect.Float64:
		return "number"
	default:
		return ""
	}
}

// applyConfig applies every checker in config to schema, tracking whether
// "required" was present via field. Checkers with a registered SchemaMakeFunc
// refine the schema; normalizers are ignored, since they don't constrain the
// shape of the data; anything else is recorded in schema.XChecker instead of
// being silently dropped.
func applyConfig(schema *Schema, config string, field *schemaField) {
	for _, token := range strings.Fields(config) {
		name, params, _ := strings.Cut(token, ":")

		switch {
		case name == nameRequired:
			field.required = true

		case ignoredForSchema[name]:
			// Normalizers transform data; they don't constrain its shape.

		default:
			schemaMakersMu.RLock()
			maker, ok := schemaMakers[name]
			schemaMakersMu.RUnlock()

			if ok {
				maker(schema, params)
			} else {
				schema.XChecker = append(schema.XChecker, token)
			}
		}
	}
}

// jsonPropertyName returns the JSON property name for the given struct
// field, and whether it should be included at all (a field tagged
// `json:"-"` is excluded, matching encoding/json).
func jsonPropertyName(field reflect.StructField) (string, bool) {
	tag, ok := field.Tag.Lookup("json")
	if !ok || tag == "" {
		return field.Name, true
	}

	if tag == "-" {
		return "", false
	}

	name, _, _ := strings.Cut(tag, ",")
	if name == "" {
		name = field.Name
	}

	return name, true
}
