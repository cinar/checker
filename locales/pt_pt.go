package locales

const (
	// PtPT is the pt_pt locale.
	PtPT = "pt-PT"
)

// PtPTMessages is the map of pt-PT messages.
var PtPTMessages = map[string]string{
	"NOT_AFTER":           "O valor deve ser posterior a {{ .reference }}.",
	"NOT_ALPHANUMERIC":    "Não é uma cadeia alfanumérica.",
	"NOT_ASCII":           "Só pode conter caracteres ASCII.",
	"NOT_BEFORE":          "O valor deve ser anterior a {{ .reference }}.",
	"NOT_CIDR":            "Não é uma notação CIDR válida.",
	"NOT_CREDIT_CARD":     "Não é um número de cartão de crédito válido.",
	"NOT_DIGITS":          "Só pode conter dígitos.",
	"NOT_EMAIL":           "Não é um endereço de email válido.",
	"EQ":                  "O valor não deve ser igual a {{ .forbidden }}.",
	"NOT_EQ":              "O valor deve ser igual a {{ .expected }}.",
	"NOT_EOA":             "Não é um endereço de propriedade externa (EOA) válido.",
	"NOT_EQ_FIELD":        "O valor deve corresponder ao campo {{ .field }}.",
	"NOT_FQDN":            "Não é um nome de domínio totalmente qualificado (FQDN).",
	"NOT_GT":              "O valor deve ser superior a {{ .n }}.",
	"NOT_GTE":             "O valor não pode ser inferior a {{ .n }}.",
	"NOT_HASH":            "Não é um hash {{ .algorithm }} válido.",
	"NOT_HEX":             "Só pode conter caracteres hexadecimais.",
	"NOT_IP":              "Não é um endereço IP válido.",
	"NOT_IPV4":            "Não é um endereço IPv4 válido.",
	"NOT_IPV6":            "Não é um endereço IPv6 válido.",
	"NOT_ISBN":            "Não é um número ISBN válido.",
	"NOT_ISO31661_ALPHA2": "Não é um código de país ISO 3166-1 alfa-2 válido.",
	"NOT_ISO31661_ALPHA3": "Não é um código de país ISO 3166-1 alfa-3 válido.",
	"NOT_ISO6391":         "Não é um código de idioma ISO 639-1 válido.",
	"NOT_LT":              "O valor deve ser inferior a {{ .n }}.",
	"NOT_LTE":             "O valor não pode ser inferior a {{ .n }}.",
	"NOT_LUHN":            "Não é um número LUHN válido.",
	"NOT_MAC":             "Não é um endereço MAC válido.",
	"NOT_MAX_LEN":         "O valor não pode ser superior a {{ .max }}.",
	"NOT_MIN_LEN":         "O valor não pode ser inferior a {{ .min }}.",
	"NOT_ONE_OF":          "O valor deve ser um dos seguintes: {{ .allowed }}.",
	"NOT_TIME":            "Não é uma hora válida.",
	"REQUIRED":            "Falta um valor obrigatório.",
	"NOT_URL":             "Não é um URL válido.",
	"NOT_UUID":            "Não é um UUID válido.",
}
