package locales

const (
	// ArSA is the ar_sa locale.
	ArSA = "ar-SA"
)

// ArSAMessages is the map of ar-SA messages.
var ArSAMessages = map[string]string{
	"NOT_AFTER":           "يجب أن تكون القيمة بعد {{ .reference }}.",
	"NOT_ALPHA":           "يمكن أن تحتوي على أحرف فقط.",
	"NOT_ALPHANUMERIC":    "ليست سلسلة أبجدية رقمية.",
	"NOT_ASCII":           "يمكن أن تحتوي على أحرف ASCII فقط.",
	"NOT_BEFORE":          "يجب أن تكون القيمة قبل {{ .reference }}.",
	"NOT_CIDR":            "ليست ترميز CIDR صالحًا.",
	"NOT_CREDIT_CARD":     "ليس رقم بطاقة ائتمان صالحًا.",
	"NOT_DIGITS":          "يمكن أن تحتوي على أرقام فقط.",
	"NOT_EMAIL":           "ليس عنوان بريد إلكتروني صالحًا.",
	"EQ":                  "يجب ألا تكون القيمة مساوية لـ {{ .forbidden }}.",
	"NOT_EQ":              "يجب أن تكون القيمة مساوية لـ {{ .expected }}.",
	"NOT_EOA":             "ليس عنوانًا مملوكًا خارجيًا (EOA) صالحًا.",
	"NOT_EQ_FIELD":        "يجب أن تطابق القيمة الحقل {{ .field }}.",
	"NOT_FQDN":            "ليس اسم نطاق مؤهلاً بالكامل (FQDN).",
	"NOT_GT":              "يجب أن تكون القيمة أكبر من {{ .n }}.",
	"NOT_GTE":             "لا يمكن أن تكون القيمة أقل من {{ .n }}.",
	"NOT_HASH":            "ليس تجزئة {{ .algorithm }} صالحة.",
	"NOT_HEX":             "يمكن أن تحتوي على أحرف سداسية عشرية فقط.",
	"NOT_IP":              "ليس عنوان IP صالحًا.",
	"NOT_IPV4":            "ليس عنوان IPv4 صالحًا.",
	"NOT_IPV6":            "ليس عنوان IPv6 صالحًا.",
	"NOT_ISBN":            "ليس رقم ISBN صالحًا.",
	"NOT_ISO31661_ALPHA2": "ليس رمز دولة ISO 3166-1 alpha-2 صالحًا.",
	"NOT_ISO31661_ALPHA3": "ليس رمز دولة ISO 3166-1 alpha-3 صالحًا.",
	"NOT_ISO6391":         "ليس رمز لغة ISO 639-1 صالحًا.",
	"NOT_LT":              "يجب أن تكون القيمة أقل من {{ .n }}.",
	"NOT_LTE":             "لا يمكن أن تكون القيمة أقل من {{ .n }}.",
	"NOT_LUHN":            "ليس رقم LUHN صالحًا.",
	"NOT_MAC":             "ليس عنوان MAC صالحًا.",
	"NOT_MAX_LEN":         "لا يمكن أن تكون القيمة أكبر من {{ .max }}.",
	"NOT_MIN_LEN":         "لا يمكن أن تكون القيمة أقل من {{ .min }}.",
	"NOT_NUMERIC":         "ليست سلسلة رقمية صالحة.",
	"NOT_ONE_OF":          "يجب أن تكون القيمة واحدة من {{ .allowed }}.",
	"NOT_TIME":            "ليس وقتًا صالحًا.",
	"REQUIRED":            "القيمة المطلوبة مفقودة.",
	"NOT_URL":             "ليس رابط URL صالحًا.",
	"NOT_UUID":            "ليس UUID صالحًا.",
}
