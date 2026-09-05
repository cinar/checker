package locales

const (
	// EnUS is the en_us locale.
	EnUS = "en-US"
)

// EnUSMessages is the map of en-US messages.
var EnUSMessages = map[string]string{
	"NOT_AFTER":           "Value must be after {{ .reference }}.",
	"NOT_ALPHA":           "Can only contain letters.",
	"NOT_ALPHANUMERIC":    "Not an alphanumeric string.",
	"NOT_ASCII":           "Can only contain ASCII characters.",
	"NOT_BEFORE":          "Value must be before {{ .reference }}.",
	"NOT_CIDR":            "Not a valid CIDR notation.",
	"NOT_CONTAINS":        "Value must contain {{ .substr }}.",
	"NOT_CREDIT_CARD":     "Not a valid credit card number.",
	"NOT_DIGITS":          "Can only contain digits.",
	"NOT_EMAIL":           "Not a valid email address.",
	"EQ":                  "Value must not equal {{ .forbidden }}.",
	"NOT_EQ":              "Value must equal {{ .expected }}.",
	"NOT_ENDS_WITH":       "Value must end with {{ .suffix }}.",
	"NOT_EOA":             "Not a valid externally owned address (EOA).",
	"NOT_EQ_FIELD":        "Value must match the {{ .field }} field.",
	"NOT_FQDN":            "Not a fully qualified domain name (FQDN).",
	"NOT_GT":              "Value must be greater than {{ .n }}.",
	"NOT_GTE":             "Value cannot be less than {{ .n }}.",
	"NOT_HASH":            "Not a valid {{ .algorithm }} hash.",
	"NOT_HEX":             "Can only contain hexadecimal characters.",
	"NOT_IP":              "Not a valid IP address.",
	"NOT_IPV4":            "Not a valid IPv4 address.",
	"NOT_IPV6":            "Not a valid IPv6 address.",
	"NOT_ISBN":            "Not a valid ISBN number.",
	"NOT_ISO31661_ALPHA2": "Not a valid ISO 3166-1 alpha-2 country code.",
	"NOT_ISO31661_ALPHA3": "Not a valid ISO 3166-1 alpha-3 country code.",
	"NOT_ISO6391":         "Not a valid ISO 639-1 language code.",
	"NOT_LT":              "Value must be less than {{ .n }}.",
	"NOT_LEN":             "Value must have a length of {{ .len }}.",
	"NOT_LTE":             "Value cannot be less than {{ .n }}.",
	"NOT_LUHN":            "Not a valid LUHN number.",
	"NOT_MAC":             "Not a valid MAC address.",
	"NOT_MAX_LEN":         "Value cannot be greater than {{ .max }}.",
	"NOT_MIN_LEN":         "Value cannot be less than {{ .min }}.",
	"NOT_NUMERIC":         "Not a valid numeric string.",
	"NOT_ONE_OF":          "Value must be one of {{ .allowed }}.",
	"NOT_STARTS_WITH":     "Value must start with {{ .prefix }}.",
	"NOT_TIME":            "Not a valid time.",
	"REQUIRED":            "Required value is missing.",
	"NOT_URL":             "Not a valid URL.",
	"NOT_UUID":            "Not a valid UUID.",
}
