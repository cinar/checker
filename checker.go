// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

// Package v2 Checker is a Go library for validating user input through checker rules provided in struct tags.
package v2

import (
	"fmt"
	"reflect"
	"strings"
	"sync"
)

const (
	// checkerTag is the name of the field tag used for checker.
	checkerTag = "checkers"

	// validateTag is a fallback field tag name, checked when checkerTag is
	// absent. It doesn't give this package any understanding of
	// go-playground/validator's own tag syntax -- only the tag *name* is a
	// fallback, not its contents -- but it lowers the switching cost for a
	// codebase that already writes checker-compatible rules under the
	// conventional `validate:"..."` tag name.
	validateTag = "validate"

	// sliceConfigPrefix is the prefix used to distinguish slice or map-level checks from item-level checks.
	sliceConfigPrefix = "@"
)

// checkStructJob defines a check strcut job.
type checkStructJob struct {
	// Name is the fully qualified name of the value being checked.
	Name string

	// Value is the value being checked.
	Value reflect.Value

	// Parent is the enclosing struct's reflect.Value, used by field-relative checks
	// such as eq-field. It is only set for direct struct fields.
	Parent reflect.Value

	// Config is the checker config string for the value.
	Config string

	// SetFunc writes the checked value back to its place, which may not always be
	// as simple as Value.Set, such as for values held in a map.
	SetFunc func(reflect.Value)
}

// Check applies the given check functions to a value sequentially.
// It returns the final value and the first encountered error, if any.
func Check[T any](value T, checks ...CheckFunc[T]) (T, error) {
	var err error

	for _, check := range checks {
		value, err = check(value)
		if err != nil {
			break
		}
	}

	return value, err
}

// CheckWithConfig applies the check functions specified by the config string to the given value.
// It returns the modified value and the first encountered error, if any.
func CheckWithConfig[T any](value T, config string) (T, error) {
	newValue, err := ReflectCheckWithConfig(reflect.Indirect(reflect.ValueOf(value)), config)
	return newValue.Interface().(T), err
}

// ReflectCheckWithConfig applies the check functions specified by the config string
// to the given reflect.Value. It returns the modified reflect.Value and the first
// encountered error, if any.
func ReflectCheckWithConfig(value reflect.Value, config string) (reflect.Value, error) {
	return reflectCheckFieldWithConfig(value, reflect.Value{}, config)
}

// reflectCheckFieldWithConfig applies the check functions specified by the config string
// to the given reflect.Value, making the parent struct's reflect.Value available to any
// field-relative checks, such as eq-field.
func reflectCheckFieldWithConfig(value, parent reflect.Value, config string) (reflect.Value, error) {
	compiled := getCompiledConfig(config)

	if compiled.omitEmpty && value.IsValid() && value.IsZero() {
		return value, nil
	}

	return runCompiledChecks(value, parent, compiled.checks)
}

// extractOmitEmpty removes the "omitempty" token from config, if present, and
// reports whether it was found. It follows the same space-separated token
// convention as the rest of the checkers tag DSL.
func extractOmitEmpty(config string) (string, bool) {
	fields := strings.Fields(config)
	kept := make([]string, 0, len(fields))
	found := false

	for _, field := range fields {
		if field == nameOmitEmpty {
			found = true
			continue
		}

		kept = append(kept, field)
	}

	return strings.Join(kept, " "), found
}

// CheckStruct checks the given struct based on the validation rules specified in the
// "checker" tag of each struct field. It returns CheckErrors, a map of field names
// to their corresponding errors, and a boolean indicating if all checks passed.
func CheckStruct(st any) (CheckErrors, bool) {
	errs := make(CheckErrors)

	jobs := []*checkStructJob{
		{
			Name:  "",
			Value: reflect.Indirect(reflect.ValueOf(st)),
		},
	}

	for len(jobs) > 0 {
		job := jobs[0]
		jobs = jobs[1:]

		switch job.Value.Kind() {
		case reflect.Struct:
			meta := getStructFieldMeta(job.Value.Type())

			for _, fm := range meta {
				name := joinFieldName(job.Name, fm.localName)
				value := indirectOrNilPointer(job.Value.FieldByIndex(fm.index))

				jobs = append(jobs, &checkStructJob{
					Name:    name,
					Value:   value,
					Parent:  job.Value,
					Config:  fm.rawConfig,
					SetFunc: safeSetFunc(value),
				})
			}

		case reflect.Slice:
			split := getSliceSplit(job.Config)
			job.Config = split.sliceConfig

			for i := 0; i < job.Value.Len(); i++ {
				name := fmt.Sprintf("%s[%d]", job.Name, i)
				value := indirectOrNilPointer(job.Value.Index(i))

				jobs = append(jobs, &checkStructJob{
					Name:    name,
					Value:   value,
					Config:  split.itemConfig,
					SetFunc: safeSetFunc(value),
				})
			}

		case reflect.Map:
			split := getSliceSplit(job.Config)
			job.Config = split.sliceConfig
			itemConfig := split.itemConfig

			mapValue := job.Value

			for _, key := range mapValue.MapKeys() {
				name := fmt.Sprintf("%s[%v]", job.Name, key.Interface())
				value := indirectOrNilPointer(mapValue.MapIndex(key))

				jobs = append(jobs, &checkStructJob{
					Name:   name,
					Value:  value,
					Config: itemConfig,
					SetFunc: func(newValue reflect.Value) {
						// A pointer map value is addressable through its indirected
						// value, so it can be mutated in place. A non-pointer map
						// value is a copy, so it must be written back through the map.
						if value.CanSet() {
							value.Set(newValue)
						} else {
							mapValue.SetMapIndex(key, newValue)
						}
					},
				})
			}
		}

		if job.Config != "" {
			newValue, err := reflectCheckFieldWithConfig(job.Value, job.Parent, job.Config)
			if err != nil {
				errs[job.Name] = err
			}

			job.SetFunc(newValue)
		}
	}

	return errs, len(errs) == 0
}

// indirectOrNilPointer returns the pointed-to value, like reflect.Indirect,
// except a nil pointer is returned as-is instead of an invalid zero Value.
// This lets a nil pointer field still be checked (e.g. required, which
// reports IsZero() on the pointer itself), while a non-nil pointer is
// dereferenced as before so its pointee can be checked or descended into.
func indirectOrNilPointer(value reflect.Value) reflect.Value {
	if value.Kind() == reflect.Pointer && value.IsNil() {
		return value
	}

	return reflect.Indirect(value)
}

// safeSetFunc returns a SetFunc that writes back through value.Set, but
// only if value is valid and addressable. A nil pointer field indirected by
// indirectOrNilPointer stays valid and addressable, so this only actually
// skips the write when CheckStruct was called with a struct passed by
// value instead of by pointer, in which case fields aren't addressable and
// there's nothing to normalize back into.
func safeSetFunc(value reflect.Value) func(reflect.Value) {
	return func(newValue reflect.Value) {
		if value.IsValid() && value.CanSet() {
			value.Set(newValue)
		}
	}
}

// fieldConfig returns the checkers tag config for the given field, falling
// back to the validate tag if checkerTag is absent, so a struct already
// tagged for go-playground/validator-style migration doesn't need every tag
// renamed just to pick up Checker.
func fieldConfig(field reflect.StructField) string {
	if config, ok := field.Tag.Lookup(checkerTag); ok {
		return config
	}

	return field.Tag.Get(validateTag)
}

// localFieldName returns the field's own name segment, without any parent
// prefix. If a "json" tag is present, it uses the tag value instead.
func localFieldName(field reflect.StructField) string {
	// Use the json tag's property name if present, stripping any
	// comma-separated options (e.g. ",omitempty"); fields tagged json:"-"
	// still need an error key, so fall back to the Go field name for them,
	// same as fields with no json tag at all.
	name, ok := jsonPropertyName(field)
	if !ok {
		name = field.Name
	}

	return name
}

// joinFieldName prepends the parent struct's fully qualified name (if any)
// to a field's local name segment, to build the fully qualified field name.
func joinFieldName(prefix, name string) string {
	if prefix == "" {
		return name
	}

	return prefix + "." + name
}

// splitSliceConfig splits config string into slice/map and item-level configurations.
func splitSliceConfig(config string) (string, string) {
	sliceFileds := make([]string, 0)
	itemFields := make([]string, 0)

	for _, configField := range strings.Fields(config) {
		if strings.HasPrefix(configField, sliceConfigPrefix) {
			sliceFileds = append(sliceFileds, strings.TrimPrefix(configField, sliceConfigPrefix))
		} else {
			itemFields = append(itemFields, configField)
		}
	}

	return strings.Join(sliceFileds, " "), strings.Join(itemFields, " ")
}

// structFieldMeta is the pre-resolved, tag-string-parsed metadata for a
// single struct field: everything CheckStruct needs to queue that field
// for checking, without re-walking its reflect.StructField on every call.
type structFieldMeta struct {
	index     []int
	localName string
	rawConfig string
}

// structMetaCache caches []structFieldMeta by struct reflect.Type, so a
// given struct type's fields, JSON names, and checkers/validate tags are
// only ever parsed once, no matter how many times a value of that type is
// checked. Unlike configCache, this never needs invalidating: a type's own
// field tags can't change at runtime.
var structMetaCache sync.Map

// getStructFieldMeta returns t's field metadata, computing and caching it
// on first use.
func getStructFieldMeta(t reflect.Type) []structFieldMeta {
	if cached, ok := structMetaCache.Load(t); ok {
		return cached.([]structFieldMeta)
	}

	meta := make([]structFieldMeta, t.NumField())

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		meta[i] = structFieldMeta{
			index:     field.Index,
			localName: localFieldName(field),
			rawConfig: fieldConfig(field),
		}
	}

	actual, _ := structMetaCache.LoadOrStore(t, meta)

	return actual.([]structFieldMeta)
}

// sliceSplit is the parsed container/item split of a slice or map field's
// checkers/validate tag config.
type sliceSplit struct {
	sliceConfig string
	itemConfig  string
}

// sliceSplitCache caches sliceSplit by its exact source config string, for
// the same reason configCache does: a given tag value only needs its
// "@"-prefixed tokens separated out once, no matter how many slice/map
// values (across instances, or even across unrelated fields) share it.
var sliceSplitCache sync.Map

// getSliceSplit returns the sliceSplit for config, computing and caching
// it on first use.
func getSliceSplit(config string) sliceSplit {
	if cached, ok := sliceSplitCache.Load(config); ok {
		return cached.(sliceSplit)
	}

	sliceConfig, itemConfig := splitSliceConfig(config)
	split := sliceSplit{sliceConfig: sliceConfig, itemConfig: itemConfig}

	actual, _ := sliceSplitCache.LoadOrStore(config, split)

	return actual.(sliceSplit)
}
