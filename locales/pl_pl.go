package locales

const (
	// PlPL is the pl_pl locale.
	PlPL = "pl-PL"
)

// PlPLMessages is the map of pl-PL messages.
var PlPLMessages = map[string]string{
	"NOT_AFTER":           "Wartość musi być późniejsza niż {{ .reference }}.",
	"NOT_ALPHA":           "Może zawierać tylko litery.",
	"NOT_ALPHANUMERIC":    "Nie jest ciągiem alfanumerycznym.",
	"NOT_ASCII":           "Może zawierać tylko znaki ASCII.",
	"NOT_BEFORE":          "Wartość musi być wcześniejsza niż {{ .reference }}.",
	"NOT_CIDR":            "Nieprawidłowa notacja CIDR.",
	"NOT_CONTAINS":        "Wartość musi zawierać {{ .substr }}.",
	"NOT_CREDIT_CARD":     "Nieprawidłowy numer karty kredytowej.",
	"NOT_DIGITS":          "Może zawierać tylko cyfry.",
	"NOT_EMAIL":           "Nieprawidłowy adres e-mail.",
	"EQ":                  "Wartość nie może być równa {{ .forbidden }}.",
	"NOT_EQ":              "Wartość musi być równa {{ .expected }}.",
	"NOT_ENDS_WITH":       "Wartość musi kończyć się na {{ .suffix }}.",
	"NOT_EOA":             "Nieprawidłowy adres zewnętrzny (EOA).",
	"NOT_EQ_FIELD":        "Wartość musi być zgodna z polem {{ .field }}.",
	"NOT_FQDN":            "Nieprawidłowa w pełni kwalifikowana nazwa domeny (FQDN).",
	"NOT_GT":              "Wartość musi być większa niż {{ .n }}.",
	"NOT_GTE":             "Wartość nie może być mniejsza niż {{ .n }}.",
	"NOT_HASH":            "Nieprawidłowy hash {{ .algorithm }}.",
	"NOT_HEX":             "Może zawierać tylko znaki szesnastkowe.",
	"NOT_IP":              "Nieprawidłowy adres IP.",
	"NOT_IPV4":            "Nieprawidłowy adres IPv4.",
	"NOT_IPV6":            "Nieprawidłowy adres IPv6.",
	"NOT_ISBN":            "Nieprawidłowy numer ISBN.",
	"NOT_ISO31661_ALPHA2": "Nieprawidłowy kod kraju ISO 3166-1 alfa-2.",
	"NOT_ISO31661_ALPHA3": "Nieprawidłowy kod kraju ISO 3166-1 alfa-3.",
	"NOT_ISO6391":         "Nieprawidłowy kod języka ISO 639-1.",
	"NOT_LT":              "Wartość musi być mniejsza niż {{ .n }}.",
	"NOT_LEN":             "Wartość musi mieć długość {{ .len }}.",
	"NOT_LTE":             "Wartość nie może być mniejsza niż {{ .n }}.",
	"NOT_LUHN":            "Nieprawidłowy numer LUHN.",
	"NOT_MAC":             "Nieprawidłowy adres MAC.",
	"NOT_MAX_LEN":         "Wartość nie może być większa niż {{ .max }}.",
	"NOT_MIN_LEN":         "Wartość nie może być mniejsza niż {{ .min }}.",
	"NOT_NUMERIC":         "Nieprawidłowy ciąg liczbowy.",
	"NOT_ONE_OF":          "Wartość musi być jedną z {{ .allowed }}.",
	"NOT_STARTS_WITH":     "Wartość musi zaczynać się od {{ .prefix }}.",
	"NOT_TIME":            "Nieprawidłowy czas.",
	"REQUIRED":            "Brak wymaganej wartości.",
	"NOT_URL":             "Nieprawidłowy adres URL.",
	"NOT_UUID":            "Nieprawidłowy UUID.",
}
