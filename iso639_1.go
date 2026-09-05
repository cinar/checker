// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

import "reflect"

const (
	// nameISO6391 is the name of the ISO 639-1 language code check.
	nameISO6391 = "iso639-1"
)

var (
	// ErrNotISO6391 indicates that the given value is not a valid ISO 639-1 language code.
	ErrNotISO6391 = NewCheckError("NOT_ISO6391")

	// iso6391Codes is the set of two-letter ISO 639-1 language codes.
	iso6391Codes = map[string]struct{}{
		"aa": {}, "ab": {}, "ae": {}, "af": {}, "ak": {}, "am": {}, "an": {}, "ar": {}, "as": {}, "av": {},
		"ay": {}, "az": {}, "ba": {}, "be": {}, "bg": {}, "bi": {}, "bm": {}, "bn": {}, "bo": {}, "br": {},
		"bs": {}, "ca": {}, "ce": {}, "ch": {}, "co": {}, "cr": {}, "cs": {}, "cu": {}, "cv": {}, "cy": {},
		"da": {}, "de": {}, "dv": {}, "dz": {}, "ee": {}, "el": {}, "en": {}, "eo": {}, "es": {}, "et": {},
		"eu": {}, "fa": {}, "ff": {}, "fi": {}, "fj": {}, "fo": {}, "fr": {}, "fy": {}, "ga": {}, "gd": {},
		"gl": {}, "gn": {}, "gu": {}, "gv": {}, "ha": {}, "he": {}, "hi": {}, "ho": {}, "hr": {}, "ht": {},
		"hu": {}, "hy": {}, "hz": {}, "ia": {}, "id": {}, "ie": {}, "ig": {}, "ii": {}, "ik": {}, "io": {},
		"is": {}, "it": {}, "iu": {}, "ja": {}, "jv": {}, "ka": {}, "kg": {}, "ki": {}, "kj": {}, "kk": {},
		"kl": {}, "km": {}, "kn": {}, "ko": {}, "kr": {}, "ks": {}, "ku": {}, "kv": {}, "kw": {}, "ky": {},
		"la": {}, "lb": {}, "lg": {}, "li": {}, "ln": {}, "lo": {}, "lt": {}, "lu": {}, "lv": {}, "mg": {},
		"mh": {}, "mi": {}, "mk": {}, "ml": {}, "mn": {}, "mr": {}, "ms": {}, "mt": {}, "my": {}, "na": {},
		"nb": {}, "nd": {}, "ne": {}, "ng": {}, "nl": {}, "nn": {}, "no": {}, "nr": {}, "nv": {}, "ny": {},
		"oc": {}, "oj": {}, "om": {}, "or": {}, "os": {}, "pa": {}, "pi": {}, "pl": {}, "ps": {}, "pt": {},
		"qu": {}, "rm": {}, "rn": {}, "ro": {}, "ru": {}, "rw": {}, "sa": {}, "sc": {}, "sd": {}, "se": {},
		"sg": {}, "si": {}, "sk": {}, "sl": {}, "sm": {}, "sn": {}, "so": {}, "sq": {}, "sr": {}, "ss": {},
		"st": {}, "su": {}, "sv": {}, "sw": {}, "ta": {}, "te": {}, "tg": {}, "th": {}, "ti": {}, "tk": {},
		"tl": {}, "tn": {}, "to": {}, "tr": {}, "ts": {}, "tt": {}, "tw": {}, "ty": {}, "ug": {}, "uk": {},
		"ur": {}, "uz": {}, "ve": {}, "vi": {}, "vo": {}, "wa": {}, "wo": {}, "xh": {}, "yi": {}, "yo": {},
		"za": {}, "zh": {}, "zu": {},
	}
)

// IsISO6391 checks if the value is a valid two-letter ISO 639-1 language code,
// such as "en" or "tr". The check is case-sensitive; combine it with the lower
// normalizer if the input's case is not already guaranteed.
func IsISO6391(value string) (string, error) {
	if _, ok := iso6391Codes[value]; !ok {
		return value, ErrNotISO6391
	}

	return value, nil
}

// checkISO6391 checks if the value is a valid ISO 639-1 language code.
func checkISO6391(value reflect.Value) (reflect.Value, error) {
	_, err := IsISO6391(reflectString(value))
	return value, err
}

// makeISO6391 makes a checker function for the ISO 639-1 checker.
func makeISO6391(_ string) CheckFunc[reflect.Value] {
	return checkISO6391
}
