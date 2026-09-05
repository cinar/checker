package locales

const (
	// DeDE is the de_de locale.
	DeDE = "de-DE"
)

// DeDEMessages is the map of de-DE messages.
var DeDEMessages = map[string]string{
	"NOT_AFTER":           "Der Wert muss nach {{ .reference }} liegen.",
	"NOT_ALPHA":           "Darf nur Buchstaben enthalten.",
	"NOT_ALPHANUMERIC":    "Keine alphanumerische Zeichenfolge.",
	"NOT_ASCII":           "Darf nur ASCII-Zeichen enthalten.",
	"NOT_BEFORE":          "Der Wert muss vor {{ .reference }} liegen.",
	"NOT_CIDR":            "Keine gültige CIDR-Notation.",
	"NOT_CONTAINS":        "Der Wert muss {{ .substr }} enthalten.",
	"NOT_CREDIT_CARD":     "Keine gültige Kreditkartennummer.",
	"NOT_DIGITS":          "Darf nur Ziffern enthalten.",
	"NOT_EMAIL":           "Keine gültige E-Mail-Adresse.",
	"EQ":                  "Der Wert darf nicht {{ .forbidden }} entsprechen.",
	"NOT_EQ":              "Der Wert muss {{ .expected }} entsprechen.",
	"NOT_ENDS_WITH":       "Der Wert muss mit {{ .suffix }} enden.",
	"NOT_EOA":             "Keine gültige extern kontrollierte Adresse (EOA).",
	"NOT_EQ_FIELD":        "Der Wert muss mit dem Feld {{ .field }} übereinstimmen.",
	"NOT_FQDN":            "Kein vollqualifizierter Domainname (FQDN).",
	"NOT_GT":              "Der Wert muss größer als {{ .n }} sein.",
	"NOT_GTE":             "Der Wert darf nicht kleiner als {{ .n }} sein.",
	"NOT_HASH":            "Kein gültiger {{ .algorithm }}-Hash.",
	"NOT_HEX":             "Darf nur hexadezimale Zeichen enthalten.",
	"NOT_IP":              "Keine gültige IP-Adresse.",
	"NOT_IPV4":            "Keine gültige IPv4-Adresse.",
	"NOT_IPV6":            "Keine gültige IPv6-Adresse.",
	"NOT_ISBN":            "Keine gültige ISBN-Nummer.",
	"NOT_ISO31661_ALPHA2": "Kein gültiger ISO-3166-1-Alpha-2-Ländercode.",
	"NOT_ISO31661_ALPHA3": "Kein gültiger ISO-3166-1-Alpha-3-Ländercode.",
	"NOT_ISO6391":         "Kein gültiger ISO-639-1-Sprachcode.",
	"NOT_LT":              "Der Wert muss kleiner als {{ .n }} sein.",
	"NOT_LEN":             "Der Wert muss eine Länge von {{ .len }} haben.",
	"NOT_LTE":             "Der Wert darf nicht kleiner als {{ .n }} sein.",
	"NOT_LUHN":            "Keine gültige LUHN-Nummer.",
	"NOT_MAC":             "Keine gültige MAC-Adresse.",
	"NOT_MAX_LEN":         "Der Wert darf nicht größer als {{ .max }} sein.",
	"NOT_MIN_LEN":         "Der Wert darf nicht kleiner als {{ .min }} sein.",
	"NOT_NUMERIC":         "Keine gültige numerische Zeichenfolge.",
	"NOT_ONE_OF":          "Der Wert muss einer der folgenden sein: {{ .allowed }}.",
	"NOT_STARTS_WITH":     "Der Wert muss mit {{ .prefix }} beginnen.",
	"NOT_TIME":            "Keine gültige Uhrzeit.",
	"REQUIRED":            "Erforderlicher Wert fehlt.",
	"NOT_URL":             "Keine gültige URL.",
	"NOT_UUID":            "Keine gültige UUID.",
}
