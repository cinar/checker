// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

import (
	"reflect"
)

type checkStructJob struct {
	Name  string
	Value reflect.Value
}

// Check applies one or more check functions to a value. It returns the
// final value and the first encountered error, if any.
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

// CheckWithConfig applies one or more check functions specified by the config to the given value.
// It returns the first encountered error, if any.
func CheckWithConfig[T any](value T, config string) (T, error) {
	newValue, err := ReflectCheckWithConfig(reflect.Indirect(reflect.ValueOf(value)), config)
	return newValue.Interface().(T), err
}

// ReflectCheckWithConfig applies one or more check functions specified by the
// config to the given value. It returns the first encountered error, if any.
func ReflectCheckWithConfig(value reflect.Value, config string) (reflect.Value, error) {
	return Check(value, makeChecks(config)...)
}

// CheckStruct checks the given struct based on the checks specified through
// the struct field tags. It returns a map of field name to error.
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

		for i := range job.Value.NumField() {
			field := job.Value.Type().Field(i)

			name := fieldName(job.Name, field)
			value := reflect.Indirect(job.Value.FieldByIndex(field.Index))

			if value.Kind() == reflect.Struct {
				jobs = append(jobs, &checkStructJob{
					Name:  name,
					Value: value,
				})

				continue
			}

			newValue, err := ReflectCheckWithConfig(value, field.Tag.Get("checker"))
			if err != nil {
				errs[name] = err
			}

			value.Set(newValue)
		}
	}

	return errs, len(errs) == 0
}

// fieldName returns the JSON name for the field if defined, otherwise it returns
// the field name. If a prefix is provided, it prepends the prefix with a dot.
func fieldName(prefix string, field reflect.StructField) string {
	name, ok := field.Tag.Lookup("json")
	if !ok {
		name = field.Name
	}

	if prefix != "" {
		name = prefix + "." + name
	}

	return name
}
