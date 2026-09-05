package locales

const (
	// LvLV is the lv_lv locale.
	LvLV = "lv-LV"
)

// LvLVMessages is the map of lv-LV messages.
var LvLVMessages = map[string]string{
	"NOT_AFTER":           "Vērtībai jābūt pēc {{ .reference }}.",
	"NOT_ALPHANUMERIC":    "Nav burtciparu virkne.",
	"NOT_ASCII":           "Var saturēt tikai ASCII rakstzīmes.",
	"NOT_BEFORE":          "Vērtībai jābūt pirms {{ .reference }}.",
	"NOT_CIDR":            "Nav derīga CIDR pieraksta.",
	"NOT_CREDIT_CARD":     "Nav derīga kredītkartes numura.",
	"NOT_DIGITS":          "Var saturēt tikai ciparus.",
	"NOT_EMAIL":           "Nav derīgas e-pasta adreses.",
	"NOT_EOA":             "Nav derīgas ārēji piederošas adreses (EOA).",
	"NOT_EQ_FIELD":        "Vērtībai jāsakrīt ar lauku {{ .field }}.",
	"NOT_FQDN":            "Nav pilnībā kvalificēta domēna nosaukuma (FQDN).",
	"NOT_GT":              "Vērtībai jābūt lielākai par {{ .n }}.",
	"NOT_GTE":             "Vērtība nedrīkst būt mazāka par {{ .n }}.",
	"NOT_HASH":            "Nav derīga {{ .algorithm }} jaucējvērtība.",
	"NOT_HEX":             "Var saturēt tikai heksadecimālas rakstzīmes.",
	"NOT_IP":              "Nav derīgas IP adreses.",
	"NOT_IPV4":            "Nav derīgas IPv4 adreses.",
	"NOT_IPV6":            "Nav derīgas IPv6 adreses.",
	"NOT_ISBN":            "Nav derīga ISBN numura.",
	"NOT_ISO31661_ALPHA2": "Nav derīga ISO 3166-1 alpha-2 valsts koda.",
	"NOT_ISO31661_ALPHA3": "Nav derīga ISO 3166-1 alpha-3 valsts koda.",
	"NOT_ISO6391":         "Nav derīga ISO 639-1 valodas koda.",
	"NOT_LT":              "Vērtībai jābūt mazākai par {{ .n }}.",
	"NOT_LTE":             "Vērtība nedrīkst būt mazāka par {{ .n }}.",
	"NOT_LUHN":            "Nav derīga LUHN numura.",
	"NOT_MAC":             "Nav derīgas MAC adreses.",
	"NOT_MAX_LEN":         "Vērtība nedrīkst būt lielāka par {{ .max }}.",
	"NOT_MIN_LEN":         "Vērtība nedrīkst būt mazāka par {{ .min }}.",
	"NOT_ONE_OF":          "Vērtībai jābūt vienai no {{ .allowed }}.",
	"NOT_TIME":            "Nav derīga laika.",
	"REQUIRED":            "Trūkst obligātās vērtības.",
	"NOT_URL":             "Nav derīga URL.",
	"NOT_UUID":            "Nav derīga UUID.",
}
