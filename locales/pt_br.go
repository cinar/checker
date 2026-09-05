package locales

const (
	// PtBR is the pt_br locale.
	PtBR = "pt-BR"
)

// PtBRMessages is the map of pt-BR messages.
var PtBRMessages = map[string]string{
	"NOT_AFTER":           "O valor deve ser posterior a {{ .reference }}.",
	"NOT_ALPHANUMERIC":    "Não é uma string alfanumérica.",
	"NOT_ASCII":           "Só pode conter caracteres ASCII.",
	"NOT_BEFORE":          "O valor deve ser anterior a {{ .reference }}.",
	"NOT_CIDR":            "Não é uma notação CIDR válida.",
	"NOT_CREDIT_CARD":     "Não é um número de cartão de crédito válido.",
	"NOT_DIGITS":          "Só pode conter dígitos.",
	"NOT_EMAIL":           "Não é um endereço de e-mail válido.",
	"NOT_EOA":             "Não é um endereço de propriedade externa (EOA) válido.",
	"NOT_EQ_FIELD":        "O valor deve corresponder ao campo {{ .field }}.",
	"NOT_FQDN":            "Não é um nome de domínio totalmente qualificado (FQDN).",
	"NOT_GTE":             "O valor não pode ser menor que {{ .n }}.",
	"NOT_HASH":            "Não é um hash {{ .algorithm }} válido.",
	"NOT_HEX":             "Só pode conter caracteres hexadecimais.",
	"NOT_IP":              "Não é um endereço IP válido.",
	"NOT_IPV4":            "Não é um endereço IPv4 válido.",
	"NOT_IPV6":            "Não é um endereço IPv6 válido.",
	"NOT_ISBN":            "Não é um número ISBN válido.",
	"NOT_ISO31661_ALPHA2": "Não é um código de país ISO 3166-1 alfa-2 válido.",
	"NOT_ISO31661_ALPHA3": "Não é um código de país ISO 3166-1 alfa-3 válido.",
	"NOT_ISO6391":         "Não é um código de idioma ISO 639-1 válido.",
	"NOT_LTE":             "O valor não pode ser menor que {{ .n }}.",
	"NOT_LUHN":            "Não é um número LUHN válido.",
	"NOT_MAC":             "Não é um endereço MAC válido.",
	"NOT_MAX_LEN":         "O valor não pode ser maior que {{ .max }}.",
	"NOT_MIN_LEN":         "O valor não pode ser menor que {{ .min }}.",
	"NOT_TIME":            "Não é um horário válido.",
	"REQUIRED":            "Falta um valor obrigatório.",
	"NOT_URL":             "Não é uma URL válida.",
	"NOT_UUID":            "Não é um UUID válido.",
}
