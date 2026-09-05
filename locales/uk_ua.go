package locales

const (
	// UkUA is the uk_ua locale.
	UkUA = "uk-UA"
)

// UkUAMessages is the map of uk-UA messages.
var UkUAMessages = map[string]string{
	"NOT_AFTER":           "Значення має бути пізніше {{ .reference }}.",
	"NOT_ALPHANUMERIC":    "Рядок має бути буквено-цифровим.",
	"NOT_ASCII":           "Може містити лише символи ASCII.",
	"NOT_BEFORE":          "Значення має бути раніше {{ .reference }}.",
	"NOT_CIDR":            "Недійсна нотація CIDR.",
	"NOT_CREDIT_CARD":     "Недійсний номер кредитної картки.",
	"NOT_DIGITS":          "Може містити лише цифри.",
	"NOT_EMAIL":           "Недійсна адреса електронної пошти.",
	"NOT_EOA":             "Недійсна зовнішня адреса (EOA).",
	"NOT_EQ_FIELD":        "Значення має збігатися зі значенням поля {{ .field }}.",
	"NOT_FQDN":            "Недійсне повне доменне ім'я (FQDN).",
	"NOT_GTE":             "Значення не може бути меншим за {{ .n }}.",
	"NOT_HASH":            "Недійсний хеш {{ .algorithm }}.",
	"NOT_HEX":             "Може містити лише шістнадцяткові символи.",
	"NOT_IP":              "Недійсна IP-адреса.",
	"NOT_IPV4":            "Недійсна адреса IPv4.",
	"NOT_IPV6":            "Недійсна адреса IPv6.",
	"NOT_ISBN":            "Недійсний номер ISBN.",
	"NOT_ISO31661_ALPHA2": "Недійсний код країни ISO 3166-1 alpha-2.",
	"NOT_ISO31661_ALPHA3": "Недійсний код країни ISO 3166-1 alpha-3.",
	"NOT_ISO6391":         "Недійсний код мови ISO 639-1.",
	"NOT_LTE":             "Значення не може бути меншим за {{ .n }}.",
	"NOT_LUHN":            "Недійсний номер LUHN.",
	"NOT_MAC":             "Недійсна MAC-адреса.",
	"NOT_MAX_LEN":         "Значення не може бути більшим за {{ .max }}.",
	"NOT_MIN_LEN":         "Значення не може бути меншим за {{ .min }}.",
	"NOT_TIME":            "Недійсний час.",
	"REQUIRED":            "Відсутнє обов'язкове значення.",
	"NOT_URL":             "Недійсна URL-адреса.",
}
