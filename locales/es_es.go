package locales

const (
	// EsES is the es_es locale.
	EsES = "es-ES"
)

// EsESMessages is the map of es-ES messages.
var EsESMessages = map[string]string{
	"NOT_AFTER":           "El valor debe ser posterior a {{ .reference }}.",
	"NOT_ALPHA":           "Solo puede contener letras.",
	"NOT_ALPHANUMERIC":    "No es una cadena alfanumérica.",
	"NOT_ASCII":           "Solo puede contener caracteres ASCII.",
	"NOT_BEFORE":          "El valor debe ser anterior a {{ .reference }}.",
	"NOT_CIDR":            "No es una notación CIDR válida.",
	"NOT_CONTAINS":        "El valor debe contener {{ .substr }}.",
	"NOT_CREDIT_CARD":     "No es un número de tarjeta de crédito válido.",
	"NOT_DIGITS":          "Solo puede contener dígitos.",
	"NOT_EMAIL":           "No es una dirección de correo electrónico válida.",
	"EQ":                  "El valor no debe ser igual a {{ .forbidden }}.",
	"NOT_EQ":              "El valor debe ser igual a {{ .expected }}.",
	"NOT_ENDS_WITH":       "El valor debe terminar con {{ .suffix }}.",
	"NOT_EOA":             "No es una dirección de cuenta externa (EOA) válida.",
	"NOT_EQ_FIELD":        "El valor debe coincidir con el campo {{ .field }}.",
	"NOT_FQDN":            "No es un nombre de dominio completamente calificado (FQDN).",
	"NOT_GT":              "El valor debe ser mayor que {{ .n }}.",
	"NOT_GTE":             "El valor no puede ser menor que {{ .n }}.",
	"NOT_HASH":            "No es un hash {{ .algorithm }} válido.",
	"NOT_HEX":             "Solo puede contener caracteres hexadecimales.",
	"NOT_IP":              "No es una dirección IP válida.",
	"NOT_IPV4":            "No es una dirección IPv4 válida.",
	"NOT_IPV6":            "No es una dirección IPv6 válida.",
	"NOT_ISBN":            "No es un número ISBN válido.",
	"NOT_ISO31661_ALPHA2": "No es un código de país ISO 3166-1 alfa-2 válido.",
	"NOT_ISO31661_ALPHA3": "No es un código de país ISO 3166-1 alfa-3 válido.",
	"NOT_ISO6391":         "No es un código de idioma ISO 639-1 válido.",
	"NOT_LT":              "El valor debe ser menor que {{ .n }}.",
	"NOT_LTE":             "El valor no puede ser menor que {{ .n }}.",
	"NOT_LUHN":            "No es un número LUHN válido.",
	"NOT_MAC":             "No es una dirección MAC válida.",
	"NOT_MAX_LEN":         "El valor no puede ser mayor que {{ .max }}.",
	"NOT_MIN_LEN":         "El valor no puede ser menor que {{ .min }}.",
	"NOT_NUMERIC":         "No es una cadena numérica válida.",
	"NOT_ONE_OF":          "El valor debe ser uno de {{ .allowed }}.",
	"NOT_STARTS_WITH":     "El valor debe comenzar con {{ .prefix }}.",
	"NOT_TIME":            "No es una hora válida.",
	"REQUIRED":            "Falta un valor obligatorio.",
	"NOT_URL":             "No es una URL válida.",
	"NOT_UUID":            "No es un UUID válido.",
}
