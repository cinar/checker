package locales

const (
	// FrFR is the fr_fr locale.
	FrFR = "fr-FR"
)

// FrFRMessages is the map of fr-FR messages.
var FrFRMessages = map[string]string{
	"NOT_AFTER":           "La valeur doit être postérieure à {{ .reference }}.",
	"NOT_ALPHA":           "Ne peut contenir que des lettres.",
	"NOT_ALPHANUMERIC":    "N'est pas une chaîne alphanumérique.",
	"NOT_ASCII":           "Ne peut contenir que des caractères ASCII.",
	"NOT_BEFORE":          "La valeur doit être antérieure à {{ .reference }}.",
	"NOT_CIDR":            "N'est pas une notation CIDR valide.",
	"NOT_CREDIT_CARD":     "N'est pas un numéro de carte de crédit valide.",
	"NOT_DIGITS":          "Ne peut contenir que des chiffres.",
	"NOT_EMAIL":           "N'est pas une adresse e-mail valide.",
	"EQ":                  "La valeur ne doit pas être égale à {{ .forbidden }}.",
	"NOT_EQ":              "La valeur doit être égale à {{ .expected }}.",
	"NOT_EOA":             "N'est pas une adresse détenue en externe (EOA) valide.",
	"NOT_EQ_FIELD":        "La valeur doit correspondre au champ {{ .field }}.",
	"NOT_FQDN":            "N'est pas un nom de domaine complet (FQDN).",
	"NOT_GT":              "La valeur doit être supérieure à {{ .n }}.",
	"NOT_GTE":             "La valeur ne peut pas être inférieure à {{ .n }}.",
	"NOT_HASH":            "N'est pas un hachage {{ .algorithm }} valide.",
	"NOT_HEX":             "Ne peut contenir que des caractères hexadécimaux.",
	"NOT_IP":              "N'est pas une adresse IP valide.",
	"NOT_IPV4":            "N'est pas une adresse IPv4 valide.",
	"NOT_IPV6":            "N'est pas une adresse IPv6 valide.",
	"NOT_ISBN":            "N'est pas un numéro ISBN valide.",
	"NOT_ISO31661_ALPHA2": "N'est pas un code pays ISO 3166-1 alpha-2 valide.",
	"NOT_ISO31661_ALPHA3": "N'est pas un code pays ISO 3166-1 alpha-3 valide.",
	"NOT_ISO6391":         "N'est pas un code de langue ISO 639-1 valide.",
	"NOT_LT":              "La valeur doit être inférieure à {{ .n }}.",
	"NOT_LTE":             "La valeur ne peut pas être inférieure à {{ .n }}.",
	"NOT_LUHN":            "N'est pas un numéro LUHN valide.",
	"NOT_MAC":             "N'est pas une adresse MAC valide.",
	"NOT_MAX_LEN":         "La valeur ne peut pas être supérieure à {{ .max }}.",
	"NOT_MIN_LEN":         "La valeur ne peut pas être inférieure à {{ .min }}.",
	"NOT_NUMERIC":         "N'est pas une chaîne numérique valide.",
	"NOT_ONE_OF":          "La valeur doit être l'une des suivantes : {{ .allowed }}.",
	"NOT_TIME":            "N'est pas une heure valide.",
	"REQUIRED":            "Valeur obligatoire manquante.",
	"NOT_URL":             "N'est pas une URL valide.",
	"NOT_UUID":            "N'est pas un UUID valide.",
}
