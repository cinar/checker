package locales_test

import (
	"regexp"
	"sort"
	"testing"

	"github.com/cinar/checker/v2/locales"
)

// allMessages lists every locale's message map. Add new locales here so
// TestLocaleMessagesAreComplete can verify they stay in sync with en-US.
var allMessages = map[string]map[string]string{
	locales.ArSA: locales.ArSAMessages,
	locales.DeDE: locales.DeDEMessages,
	locales.EnUS: locales.EnUSMessages,
	locales.EsES: locales.EsESMessages,
	locales.FaIR: locales.FaIRMessages,
	locales.FrFR: locales.FrFRMessages,
	locales.HyAM: locales.HyAMMessages,
	locales.IDID: locales.IDIDMessages,
	locales.ItIT: locales.ItITMessages,
	locales.JaJP: locales.JaJPMessages,
	locales.KoKR: locales.KoKRMessages,
	locales.LvLV: locales.LvLVMessages,
	locales.NlNL: locales.NlNLMessages,
	locales.PlPL: locales.PlPLMessages,
	locales.PtBR: locales.PtBRMessages,
	locales.PtPT: locales.PtPTMessages,
	locales.RuRU: locales.RuRUMessages,
	locales.ThTH: locales.ThTHMessages,
	locales.TrTR: locales.TrTRMessages,
	locales.UkUA: locales.UkUAMessages,
	locales.ViVN: locales.ViVNMessages,
	locales.ZhCN: locales.ZhCNMessages,
	locales.ZhTW: locales.ZhTWMessages,
}

// placeholderRegexp matches a Go template action such as "{{ .reference }}".
var placeholderRegexp = regexp.MustCompile(`\{\{\s*(\.\w+)\s*\}\}`)

// TestLocaleMessagesAreComplete verifies that every registered locale defines
// exactly the same set of message codes as en-US, the default locale, with
// the same set of template placeholders for each code. This is meant to
// catch a locale falling out of sync when a new checker adds a message code,
// or a translated message losing/renaming a {{ .placeholder }} by mistake.
func TestLocaleMessagesAreComplete(t *testing.T) {
	reference := locales.EnUSMessages

	for locale, messages := range allMessages {
		if locale == locales.EnUS {
			continue
		}

		for code, want := range reference {
			got, ok := messages[code]
			if !ok {
				t.Errorf("%s: missing message for code %s", locale, code)
				continue
			}

			if got == "" {
				t.Errorf("%s: empty message for code %s", locale, code)
			}

			wantPlaceholders := sortedPlaceholders(want)
			gotPlaceholders := sortedPlaceholders(got)

			if !equal(wantPlaceholders, gotPlaceholders) {
				t.Errorf("%s: code %s has placeholders %v, want %v", locale, code, gotPlaceholders, wantPlaceholders)
			}
		}

		for code := range messages {
			if _, ok := reference[code]; !ok {
				t.Errorf("%s: code %s is not defined in en-US", locale, code)
			}
		}
	}
}

// sortedPlaceholders returns the sorted, deduplicated set of {{ .name }}
// placeholders used in the given message.
func sortedPlaceholders(message string) []string {
	matches := placeholderRegexp.FindAllStringSubmatch(message, -1)

	set := make(map[string]struct{}, len(matches))
	for _, match := range matches {
		set[match[1]] = struct{}{}
	}

	placeholders := make([]string, 0, len(set))
	for placeholder := range set {
		placeholders = append(placeholders, placeholder)
	}

	sort.Strings(placeholders)

	return placeholders
}

// equal reports whether the two sorted string slices contain the same elements.
func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}
