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
)

const (
	// checkerTag is the name of the field tag used for checker.
	checkerTag = "checkers"

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
	return Check(value, makeChecks(config, parent)...)
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
			for i := 0; i < job.Value.NumField(); i++ {
				field := job.Value.Type().Field(i)

				name := fieldName(job.Name, field)
				value := indirectOrNilPointer(job.Value.FieldByIndex(field.Index))

				jobs = append(jobs, &checkStructJob{
					Name:    name,
					Value:   value,
					Parent:  job.Value,
					Config:  field.Tag.Get(checkerTag),
					SetFunc: safeSetFunc(value),
				})
			}

		case reflect.Slice:
			sliceConfig, itemConfig := splitSliceConfig(job.Config)
			job.Config = sliceConfig

			for i := 0; i < job.Value.Len(); i++ {
				name := fmt.Sprintf("%s[%d]", job.Name, i)
				value := indirectOrNilPointer(job.Value.Index(i))

				jobs = append(jobs, &checkStructJob{
					Name:    name,
					Value:   value,
					Config:  itemConfig,
					SetFunc: safeSetFunc(value),
				})
			}

		case reflect.Map:
			mapConfig, itemConfig := splitSliceConfig(job.Config)
			job.Config = mapConfig

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

// fieldName returns the field name. If a "json" tag is present, it uses the
// tag value instead. It also prepends the parent struct's name (if any) to
// create a fully qualified field name.
func fieldName(prefix string, field reflect.StructField) string {
	// Default to field name
	name := field.Name

	// Use json tag if present
	if jsonTag, ok := field.Tag.Lookup("json"); ok {
		name = jsonTag
	}

	// Prepend parent name
	if prefix != "" {
		name = prefix + "." + name
	}

	return name
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
