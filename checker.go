// Package checker is a Go library for validating user input through struct tags.
//
// https://github.com/cinar/checker
//
// Copyright 2023 Onur Cinar. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package checker

import (
	"fmt"
	"reflect"
	"strings"
)

// Result is a unique textual identifier for the mistake.
type Result string

// CheckFunc defines the checker function.
type CheckFunc func(value, parent reflect.Value) Result

// MakeFunc defines the maker function.
type MakeFunc func(params string) CheckFunc

// Mistakes provides mapping to checker result for the invalid fields.
type Mistakes map[string]Result

type checkerJob struct {
	Parent reflect.Value
	Name   string
	Value  reflect.Value
	Config string
}

// ResultValid result indicates that the user input is valid.
const ResultValid Result = "VALID"

// makers provides mapping to maker function for the checkers.
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

// Check checks the given struct based on the checkers listed in each field's strcut tag named checkers.
func Check(s interface{}) (Mistakes, bool) {
	root := reflect.Indirect(reflect.ValueOf(s))
	if root.Kind() != reflect.Struct {
		panic("expecting struct")
	}

	mistakes := Mistakes{}

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
				if result := checker(job.Value, job.Parent); result != ResultValid {
					mistakes[job.Name] = result
					break
				}
			}
		}
	}

	return mistakes, len(mistakes) == 0
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
