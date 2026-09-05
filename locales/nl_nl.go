package locales

const (
	// NlNL is the nl_nl locale.
	NlNL = "nl-NL"
)

// NlNLMessages is the map of nl-NL messages.
var NlNLMessages = map[string]string{
	"NOT_AFTER":           "De waarde moet na {{ .reference }} liggen.",
	"NOT_ALPHA":           "Mag alleen letters bevatten.",
	"NOT_ALPHANUMERIC":    "Geen alfanumerieke tekenreeks.",
	"NOT_ASCII":           "Mag alleen ASCII-tekens bevatten.",
	"NOT_BEFORE":          "De waarde moet vóór {{ .reference }} liggen.",
	"NOT_CIDR":            "Geen geldige CIDR-notatie.",
	"NOT_CONTAINS":        "De waarde moet {{ .substr }} bevatten.",
	"NOT_CREDIT_CARD":     "Geen geldig creditcardnummer.",
	"NOT_DIGITS":          "Mag alleen cijfers bevatten.",
	"NOT_EMAIL":           "Geen geldig e-mailadres.",
	"EQ":                  "De waarde mag niet gelijk zijn aan {{ .forbidden }}.",
	"NOT_EQ":              "De waarde moet gelijk zijn aan {{ .expected }}.",
	"NOT_ENDS_WITH":       "De waarde moet eindigen met {{ .suffix }}.",
	"NOT_EOA":             "Geen geldig extern beheerd adres (EOA).",
	"NOT_EQ_FIELD":        "De waarde moet overeenkomen met het veld {{ .field }}.",
	"NOT_FQDN":            "Geen volledig gekwalificeerde domeinnaam (FQDN).",
	"NOT_GT":              "De waarde moet groter zijn dan {{ .n }}.",
	"NOT_GTE":             "De waarde mag niet kleiner zijn dan {{ .n }}.",
	"NOT_HASH":            "Geen geldige {{ .algorithm }}-hash.",
	"NOT_HEX":             "Mag alleen hexadecimale tekens bevatten.",
	"NOT_IP":              "Geen geldig IP-adres.",
	"NOT_IPV4":            "Geen geldig IPv4-adres.",
	"NOT_IPV6":            "Geen geldig IPv6-adres.",
	"NOT_ISBN":            "Geen geldig ISBN-nummer.",
	"NOT_ISO31661_ALPHA2": "Geen geldige ISO 3166-1 alpha-2-landcode.",
	"NOT_ISO31661_ALPHA3": "Geen geldige ISO 3166-1 alpha-3-landcode.",
	"NOT_ISO6391":         "Geen geldige ISO 639-1-taalcode.",
	"NOT_LT":              "De waarde moet kleiner zijn dan {{ .n }}.",
	"NOT_LEN":             "De waarde moet een lengte van {{ .len }} hebben.",
	"NOT_LTE":             "De waarde mag niet kleiner zijn dan {{ .n }}.",
	"NOT_LUHN":            "Geen geldig LUHN-nummer.",
	"NOT_MAC":             "Geen geldig MAC-adres.",
	"NOT_MAX_LEN":         "De waarde mag niet groter zijn dan {{ .max }}.",
	"NOT_MIN_LEN":         "De waarde mag niet kleiner zijn dan {{ .min }}.",
	"NOT_NUMERIC":         "Geen geldige numerieke tekenreeks.",
	"NOT_ONE_OF":          "De waarde moet een van {{ .allowed }} zijn.",
	"NOT_STARTS_WITH":     "De waarde moet beginnen met {{ .prefix }}.",
	"NOT_TIME":            "Geen geldige tijd.",
	"REQUIRED":            "Verplichte waarde ontbreekt.",
	"NOT_URL":             "Geen geldige URL.",
	"NOT_UUID":            "Geen geldige UUID.",
}
