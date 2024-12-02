// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

import (
	"reflect"
)

const (
	// checkerTag is the name of the field tag used for checker.
	checkerTag = "checker"
)

// checkStructJob defines a check strcut job.
type checkStructJob struct {
	Name   string
	Value  reflect.Value
	Config string
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
	return Check(value, makeChecks(config)...)
}

// CheckStruct checks the given struct based on the validation rules specified in the
// "checker" tag of each struct field. It returns a map of field names to their
// corresponding errors, and a boolean indicating if all checks passed.
func CheckStruct(st any) (map[string]error, bool) {
	errs := make(map[string]error)

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
				value := reflect.Indirect(job.Value.FieldByIndex(field.Index))

				jobs = append(jobs, &checkStructJob{
					Name:   name,
					Value:  value,
					Config: field.Tag.Get("checker"),
				})
			}
		}

		if job.Config != "" {
			newValue, err := ReflectCheckWithConfig(job.Value, job.Config)
			if err != nil {
				errs[job.Name] = err
			}

			job.Value.Set(newValue)
		}
	}

	return errs, len(errs) == 0
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
