// Package checker is a Go library for validating user input through struct tags.
//
// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker
package checker

import (
	"fmt"
	"reflect"
	"strings"
)

// CheckFunc defines the signature for the checker functions.
type CheckFunc func(value, parent reflect.Value) error

// MakeFunc defines the signature for the checker maker functions.
type MakeFunc func(params string) CheckFunc

// Errors provides a mapping of the checker errors keyed by the field names.
type Errors map[string]error

type checkerJob struct {
	Parent reflect.Value
	Name   string
	Value  reflect.Value
	Config string
}

// makers provides a mapping of the maker functions keyed by the respective checker names.
var makers = map[string]MakeFunc{
	CheckerAlphanumeric:    makeAlphanumeric,
	CheckerASCII:           makeASCII,
	CheckerCreditCard:      makeCreditCard,
	CheckerCidr:            makeCidr,
	CheckerDigits:          makeDigits,
	CheckerEmail:           makeEmail,
	CheckerFqdn:            makeFqdn,
	CheckerIP:              makeIP,
	CheckerIPV4:            makeIPV4,
	CheckerIPV6:            makeIPV6,
	CheckerISBN:            makeISBN,
	CheckerLuhn:            makeLuhn,
	CheckerMac:             makeMac,
	CheckerMax:             makeMax,
	CheckerMaxLength:       makeMaxLength,
	CheckerMin:             makeMin,
	CheckerMinLength:       makeMinLength,
	CheckerRegexp:          makeRegexp,
	CheckerRequired:        makeRequired,
	CheckerSame:            makeSame,
	CheckerURL:             makeURL,
	NormalizerHTMLEscape:   makeHTMLEscape,
	NormalizerHTMLUnescape: makeHTMLUnescape,
	NormalizerLower:        makeLower,
	NormalizerUpper:        makeUpper,
	NormalizerTitle:        makeTitle,
	NormalizerTrim:         makeTrim,
	NormalizerTrimLeft:     makeTrimLeft,
	NormalizerTrimRight:    makeTrimRight,
	NormalizerURLEscape:    makeURLEscape,
	NormalizerURLUnescape:  makeURLUnescape,
}

// Register registers the given checker name and the maker function.
func Register(name string, maker MakeFunc) {
	makers[name] = maker
}

// Check function checks the given struct based on the checkers listed in field tag names.
func Check(s interface{}) (Errors, bool) {
	root := reflect.Indirect(reflect.ValueOf(s))
	if root.Kind() != reflect.Struct {
		panic("expecting struct")
	}

	errors := Errors{}

	jobs := []checkerJob{
		{
			Parent: reflect.ValueOf(nil),
			Name:   "",
			Value:  root,
			Config: "",
		},
	}

	for len(jobs) > 0 {
		job := jobs[0]
		jobs = jobs[1:]

		if job.Value.Kind() == reflect.Struct {
			for i := 0; i < job.Value.NumField(); i++ {
				field := job.Value.Type().Field(i)
				addJob := field.Type.Kind() == reflect.Struct
				config := ""

				if !addJob {
					config = field.Tag.Get("checkers")
					addJob = config != ""
				}

				if addJob {
					name := field.Name
					if job.Name != "" {
						name = job.Name + "." + name
					}

					jobs = append(jobs, checkerJob{
						Parent: job.Value,
						Name:   name,
						Value:  reflect.Indirect(job.Value.FieldByIndex(field.Index)),
						Config: config,
					})
				}
			}
		} else {
			for _, checker := range initCheckers(job.Config) {
				if err := checker(job.Value, job.Parent); err != nil {
					errors[job.Name] = err
					break
				}
			}
		}
	}

	return errors, len(errors) == 0
}

// initCheckers initializes the checkers provided in the config.
func initCheckers(config string) []CheckFunc {
	fields := strings.Fields(config)
	checkers := make([]CheckFunc, len(fields))

	for i, field := range fields {
		name, params, _ := strings.Cut(field, ":")

		maker, ok := makers[name]
		if !ok {
			panic(fmt.Sprintf("checker %s is unknown", name))
		}

		checkers[i] = maker(params)
	}

	return checkers
}

// numberOf gives value's numerical value given that it is either an int or a float.
func numberOf(value reflect.Value) float64 {
	switch value.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(value.Int())

	case reflect.Float32, reflect.Float64:
		return value.Float()

	default:
		panic("expecting int or float")
	}
}
