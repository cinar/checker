package locales

const (
	// TrTR is the tr_tr locale.
	TrTR = "tr-TR"
)

// TrTRMessages is the map of tr-TR messages.
var TrTRMessages = map[string]string{
	"NOT_AFTER":           "Değer {{ .reference }} tarihinden sonra olmalıdır.",
	"NOT_ALPHANUMERIC":    "Alfanümerik bir dize değil.",
	"NOT_ASCII":           "Yalnızca ASCII karakterleri içerebilir.",
	"NOT_BEFORE":          "Değer {{ .reference }} tarihinden önce olmalıdır.",
	"NOT_CIDR":            "Geçerli bir CIDR gösterimi değil.",
	"NOT_CREDIT_CARD":     "Geçerli bir kredi kartı numarası değil.",
	"NOT_DIGITS":          "Yalnızca rakam içerebilir.",
	"NOT_EMAIL":           "Geçerli bir e-posta adresi değil.",
	"NOT_EOA":             "Geçerli bir harici sahipli adres (EOA) değil.",
	"NOT_EQ_FIELD":        "Değer, {{ .field }} alanıyla eşleşmelidir.",
	"NOT_FQDN":            "Tam nitelikli bir alan adı (FQDN) değil.",
	"NOT_GTE":             "Değer {{ .n }} değerinden küçük olamaz.",
	"NOT_HASH":            "Geçerli bir {{ .algorithm }} özeti değil.",
	"NOT_HEX":             "Yalnızca onaltılık karakterler içerebilir.",
	"NOT_IP":              "Geçerli bir IP adresi değil.",
	"NOT_IPV4":            "Geçerli bir IPv4 adresi değil.",
	"NOT_IPV6":            "Geçerli bir IPv6 adresi değil.",
	"NOT_ISBN":            "Geçerli bir ISBN numarası değil.",
	"NOT_ISO31661_ALPHA2": "Geçerli bir ISO 3166-1 alpha-2 ülke kodu değil.",
	"NOT_ISO31661_ALPHA3": "Geçerli bir ISO 3166-1 alpha-3 ülke kodu değil.",
	"NOT_ISO6391":         "Geçerli bir ISO 639-1 dil kodu değil.",
	"NOT_LTE":             "Değer {{ .n }} değerinden küçük olamaz.",
	"NOT_LUHN":            "Geçerli bir LUHN numarası değil.",
	"NOT_MAC":             "Geçerli bir MAC adresi değil.",
	"NOT_MAX_LEN":         "Değer {{ .max }} değerinden büyük olamaz.",
	"NOT_MIN_LEN":         "Değer {{ .min }} değerinden küçük olamaz.",
	"NOT_TIME":            "Geçerli bir saat değil.",
	"REQUIRED":            "Gerekli değer eksik.",
	"NOT_URL":             "Geçerli bir URL değil.",
	"NOT_UUID":            "Geçerli bir UUID değil.",
}
