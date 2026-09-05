package locales

const (
	// ItIT is the it_it locale.
	ItIT = "it-IT"
)

// ItITMessages is the map of it-IT messages.
var ItITMessages = map[string]string{
	"NOT_AFTER":           "Il valore deve essere successivo a {{ .reference }}.",
	"NOT_ALPHANUMERIC":    "Non è una stringa alfanumerica.",
	"NOT_ASCII":           "Può contenere solo caratteri ASCII.",
	"NOT_BEFORE":          "Il valore deve essere precedente a {{ .reference }}.",
	"NOT_CIDR":            "Non è una notazione CIDR valida.",
	"NOT_CREDIT_CARD":     "Non è un numero di carta di credito valido.",
	"NOT_DIGITS":          "Può contenere solo cifre.",
	"NOT_EMAIL":           "Non è un indirizzo email valido.",
	"NOT_EOA":             "Non è un indirizzo a proprietà esterna (EOA) valido.",
	"NOT_EQ_FIELD":        "Il valore deve corrispondere al campo {{ .field }}.",
	"NOT_FQDN":            "Non è un nome di dominio completamente qualificato (FQDN).",
	"NOT_GTE":             "Il valore non può essere inferiore a {{ .n }}.",
	"NOT_HASH":            "Non è un hash {{ .algorithm }} valido.",
	"NOT_HEX":             "Può contenere solo caratteri esadecimali.",
	"NOT_IP":              "Non è un indirizzo IP valido.",
	"NOT_IPV4":            "Non è un indirizzo IPv4 valido.",
	"NOT_IPV6":            "Non è un indirizzo IPv6 valido.",
	"NOT_ISBN":            "Non è un numero ISBN valido.",
	"NOT_ISO31661_ALPHA2": "Non è un codice paese ISO 3166-1 alpha-2 valido.",
	"NOT_ISO31661_ALPHA3": "Non è un codice paese ISO 3166-1 alpha-3 valido.",
	"NOT_ISO6391":         "Non è un codice lingua ISO 639-1 valido.",
	"NOT_LTE":             "Il valore non può essere inferiore a {{ .n }}.",
	"NOT_LUHN":            "Non è un numero LUHN valido.",
	"NOT_MAC":             "Non è un indirizzo MAC valido.",
	"NOT_MAX_LEN":         "Il valore non può essere superiore a {{ .max }}.",
	"NOT_MIN_LEN":         "Il valore non può essere inferiore a {{ .min }}.",
	"NOT_TIME":            "Non è un orario valido.",
	"REQUIRED":            "Manca un valore obbligatorio.",
	"NOT_URL":             "Non è un URL valido.",
	"NOT_UUID":            "Non è un UUID valido.",
}
