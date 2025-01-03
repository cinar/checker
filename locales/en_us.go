package locales

const (
	// EnUS is the en_us locale.
	EnUS = "en-US"
)

// EnUSMessages is the map of en-US messages.
var EnUSMessages = map[string]string{
	"NOT_ALPHANUMERIC": "Not an alphanumeric string.",
	"NOT_ASCII":        "Can only contain ASCII characters.",
	"NOT_CIDR":         "Not a valid CIDR notation.",
	"NOT_CREDIT_CARD":  "Not a valid credit card number.",
	"NOT_DIGITS":       "Can only contain digits.",
	"NOT_EMAIL":        "Not a valid email address.",
	"NOT_FQDN":         "Not a fully qualified domain name (FQDN).",
	"NOT_GTE":          "Value cannot be less than {{ .n }}.",
	"NOT_HEX":          "Can only contain hexadecimal characters.",
	"NOT_IP":           "Not a valid IP address.",
	"NOT_IPV4":         "Not a valid IPv4 address.",
	"NOT_IPV6":         "Not a valid IPv6 address.",
	"NOT_ISBN":         "Not a valid ISBN number.",
	"NOT_LTE":          "Value cannot be less than {{ .n }}.",
	"NOT_LUHN":         "Not a valid LUHN number.",
	"NOT_MAC":          "Not a valid MAC address.",
	"NOT_MAX_LEN":      "Value cannot be greater than {{ .max }}.",
	"NOT_MIN_LEN":      "Value cannot be less than {{ .min }}.",
	"NOT_TIME":         "Not a valid time.",
	"REQUIRED":         "Required value is missing.",
	"NOT_URL":          "Not a valid URL.",
}
