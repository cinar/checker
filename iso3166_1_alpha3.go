// Copyright (c) 2023-2024 Onur Cinar.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// https://github.com/cinar/checker

package v2

import "reflect"

const (
	// nameISO31661Alpha3 is the name of the ISO 3166-1 alpha-3 country code check.
	nameISO31661Alpha3 = "iso3166-1-alpha-3"
)

var (
	// ErrNotISO31661Alpha3 indicates that the given value is not a valid ISO 3166-1 alpha-3 country code.
	ErrNotISO31661Alpha3 = NewCheckError("NOT_ISO31661_ALPHA3")

	// iso31661Alpha3Codes is the set of three-letter ISO 3166-1 alpha-3 country codes.
	iso31661Alpha3Codes = map[string]struct{}{
		"ABW": {}, "AFG": {}, "AGO": {}, "AIA": {}, "ALA": {}, "ALB": {}, "AND": {}, "ARE": {}, "ARG": {}, "ARM": {},
		"ASM": {}, "ATA": {}, "ATF": {}, "ATG": {}, "AUS": {}, "AUT": {}, "AZE": {}, "BDI": {}, "BEL": {}, "BEN": {},
		"BES": {}, "BFA": {}, "BGD": {}, "BGR": {}, "BHR": {}, "BHS": {}, "BIH": {}, "BLM": {}, "BLR": {}, "BLZ": {},
		"BMU": {}, "BOL": {}, "BRA": {}, "BRB": {}, "BRN": {}, "BTN": {}, "BVT": {}, "BWA": {}, "CAF": {}, "CAN": {},
		"CCK": {}, "CHE": {}, "CHL": {}, "CHN": {}, "CIV": {}, "CMR": {}, "COD": {}, "COG": {}, "COK": {}, "COL": {},
		"COM": {}, "CPV": {}, "CRI": {}, "CUB": {}, "CUW": {}, "CXR": {}, "CYM": {}, "CYP": {}, "CZE": {}, "DEU": {},
		"DJI": {}, "DMA": {}, "DNK": {}, "DOM": {}, "DZA": {}, "ECU": {}, "EGY": {}, "ERI": {}, "ESH": {}, "ESP": {},
		"EST": {}, "ETH": {}, "FIN": {}, "FJI": {}, "FLK": {}, "FRA": {}, "FRO": {}, "FSM": {}, "GAB": {}, "GBR": {},
		"GEO": {}, "GGY": {}, "GHA": {}, "GIB": {}, "GIN": {}, "GLP": {}, "GMB": {}, "GNB": {}, "GNQ": {}, "GRC": {},
		"GRD": {}, "GRL": {}, "GTM": {}, "GUF": {}, "GUM": {}, "GUY": {}, "HKG": {}, "HMD": {}, "HND": {}, "HRV": {},
		"HTI": {}, "HUN": {}, "IDN": {}, "IMN": {}, "IND": {}, "IOT": {}, "IRL": {}, "IRN": {}, "IRQ": {}, "ISL": {},
		"ISR": {}, "ITA": {}, "JAM": {}, "JEY": {}, "JOR": {}, "JPN": {}, "KAZ": {}, "KEN": {}, "KGZ": {}, "KHM": {},
		"KIR": {}, "KNA": {}, "KOR": {}, "KWT": {}, "LAO": {}, "LBN": {}, "LBR": {}, "LBY": {}, "LCA": {}, "LIE": {},
		"LKA": {}, "LSO": {}, "LTU": {}, "LUX": {}, "LVA": {}, "MAC": {}, "MAF": {}, "MAR": {}, "MCO": {}, "MDA": {},
		"MDG": {}, "MDV": {}, "MEX": {}, "MHL": {}, "MKD": {}, "MLI": {}, "MLT": {}, "MMR": {}, "MNE": {}, "MNG": {},
		"MNP": {}, "MOZ": {}, "MRT": {}, "MSR": {}, "MTQ": {}, "MUS": {}, "MWI": {}, "MYS": {}, "MYT": {}, "NAM": {},
		"NCL": {}, "NER": {}, "NFK": {}, "NGA": {}, "NIC": {}, "NIU": {}, "NLD": {}, "NOR": {}, "NPL": {}, "NRU": {},
		"NZL": {}, "OMN": {}, "PAK": {}, "PAN": {}, "PCN": {}, "PER": {}, "PHL": {}, "PLW": {}, "PNG": {}, "POL": {},
		"PRI": {}, "PRK": {}, "PRT": {}, "PRY": {}, "PSE": {}, "PYF": {}, "QAT": {}, "REU": {}, "ROU": {}, "RUS": {},
		"RWA": {}, "SAU": {}, "SDN": {}, "SEN": {}, "SGP": {}, "SGS": {}, "SHN": {}, "SJM": {}, "SLB": {}, "SLE": {},
		"SLV": {}, "SMR": {}, "SOM": {}, "SPM": {}, "SRB": {}, "SSD": {}, "STP": {}, "SUR": {}, "SVK": {}, "SVN": {},
		"SWE": {}, "SWZ": {}, "SXM": {}, "SYC": {}, "SYR": {}, "TCA": {}, "TCD": {}, "TGO": {}, "THA": {}, "TJK": {},
		"TKL": {}, "TKM": {}, "TLS": {}, "TON": {}, "TTO": {}, "TUN": {}, "TUR": {}, "TUV": {}, "TWN": {}, "TZA": {},
		"UGA": {}, "UKR": {}, "UMI": {}, "URY": {}, "USA": {}, "UZB": {}, "VAT": {}, "VCT": {}, "VEN": {}, "VGB": {},
		"VIR": {}, "VNM": {}, "VUT": {}, "WLF": {}, "WSM": {}, "YEM": {}, "ZAF": {}, "ZMB": {}, "ZWE": {},
	}
)

// IsISO31661Alpha3 checks if the value is a valid three-letter ISO 3166-1
// alpha-3 country code, such as "USA" or "TUR". The check is case-sensitive;
// combine it with the upper normalizer if the input's case is not already
// guaranteed.
func IsISO31661Alpha3(value string) (string, error) {
	if _, ok := iso31661Alpha3Codes[value]; !ok {
		return value, ErrNotISO31661Alpha3
	}

	return value, nil
}

// checkISO31661Alpha3 checks if the value is a valid ISO 3166-1 alpha-3 country code.
func checkISO31661Alpha3(value reflect.Value) (reflect.Value, error) {
	_, err := IsISO31661Alpha3(value.Interface().(string))
	return value, err
}

// makeISO31661Alpha3 makes a checker function for the ISO 3166-1 alpha-3 checker.
func makeISO31661Alpha3(_ string) CheckFunc[reflect.Value] {
	return checkISO31661Alpha3
}
