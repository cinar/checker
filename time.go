// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

import (
	"reflect"
	"time"
)

const (
	// time is the name of the time format check.
	nameTime = "time"
)

var (
	// ErrTime indicates that the value is not a valid based on the given time format.
	ErrTime = NewCheckError("NOT_TIME")
)

// timeLayouts is a map of time layouts that can be used in the time check.
var timeLayouts = map[string]string{
	"Layout":      time.Layout,
	"ANSIC":       time.ANSIC,
	"UnixDate":    time.UnixDate,
	"RubyDate":    time.RubyDate,
	"RFC822":      time.RFC822,
	"RFC822Z":     time.RFC822Z,
	"RFC850":      time.RFC850,
	"RFC1123":     time.RFC1123,
	"RFC1123Z":    time.RFC1123Z,
	"RFC3339":     time.RFC3339,
	"RFC3339Nano": time.RFC3339Nano,
	"Kitchen":     time.Kitchen,
	"Stamp":       time.Stamp,
	"StampMilli":  time.StampMilli,
	"StampMicro":  time.StampMicro,
	"StampNano":   time.StampNano,
	"DateTime":    time.DateTime,
	"DateOnly":    time.DateOnly,
	"TimeOnly":    time.TimeOnly,
}

// IsTime checks if the given value is a valid time based on the given layout.
func IsTime(params, value string) (string, error) {
	layout, ok := timeLayouts[params]
	if !ok {
		layout = params
	}

	_, err := time.Parse(layout, value)
	if err != nil {
		return value, ErrTime
	}

	return value, nil
}

// makeTime makes a time format check based on the given time layout.
func makeTime(params string) CheckFunc[reflect.Value] {
	return func(value reflect.Value) (reflect.Value, error) {
		_, err := IsTime(params, value.String())
		return value, err
	}
}
