// Copyright (c) 2023-2026 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package cli

import (
	checker "github.com/cinar/checker/v2"
	"github.com/cinar/checker/v2/locales"
)

// registerAllLocales registers every translated locale the core module
// ships with, so "checker check --locale=<tag>" works for all 23 out of
// the box. The core module leaves this opt-in (see checker.RegisterLocale)
// so importing it never pulls in translations a library caller doesn't
// use; a standalone CLI has no such caller to defer to, so it registers
// all of them itself, once, at startup.
func registerAllLocales() {
	for _, locale := range []struct {
		name     string
		messages map[string]string
	}{
		{locales.ArSA, locales.ArSAMessages},
		{locales.DeDE, locales.DeDEMessages},
		{locales.EnUS, locales.EnUSMessages},
		{locales.EsES, locales.EsESMessages},
		{locales.FaIR, locales.FaIRMessages},
		{locales.FrFR, locales.FrFRMessages},
		{locales.HyAM, locales.HyAMMessages},
		{locales.IDID, locales.IDIDMessages},
		{locales.ItIT, locales.ItITMessages},
		{locales.JaJP, locales.JaJPMessages},
		{locales.KoKR, locales.KoKRMessages},
		{locales.LvLV, locales.LvLVMessages},
		{locales.NlNL, locales.NlNLMessages},
		{locales.PlPL, locales.PlPLMessages},
		{locales.PtBR, locales.PtBRMessages},
		{locales.PtPT, locales.PtPTMessages},
		{locales.RuRU, locales.RuRUMessages},
		{locales.ThTH, locales.ThTHMessages},
		{locales.TrTR, locales.TrTRMessages},
		{locales.UkUA, locales.UkUAMessages},
		{locales.ViVN, locales.ViVNMessages},
		{locales.ZhCN, locales.ZhCNMessages},
		{locales.ZhTW, locales.ZhTWMessages},
	} {
		checker.RegisterLocale(locale.name, locale.messages)
	}
}

// init eagerly registers every shipped locale. A CLI process is
// short-lived and has no other caller to defer this cost to, unlike a
// library import, so there's no reason to make it lazy or opt-in here.
func init() {
	registerAllLocales()
}
