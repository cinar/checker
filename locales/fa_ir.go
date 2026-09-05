package locales

const (
	// FaIR is the fa_ir locale.
	FaIR = "fa-IR"
)

// FaIRMessages is the map of fa-IR messages.
var FaIRMessages = map[string]string{
	"NOT_AFTER":           "مقدار باید بعد از {{ .reference }} باشد.",
	"NOT_ALPHANUMERIC":    "یک رشته الفبایی\u200cعددی نیست.",
	"NOT_ASCII":           "فقط می\u200cتواند شامل کاراکترهای ASCII باشد.",
	"NOT_BEFORE":          "مقدار باید قبل از {{ .reference }} باشد.",
	"NOT_CIDR":            "نماد CIDR معتبر نیست.",
	"NOT_CREDIT_CARD":     "شماره کارت اعتباری معتبر نیست.",
	"NOT_DIGITS":          "فقط می\u200cتواند شامل ارقام باشد.",
	"NOT_EMAIL":           "آدرس ایمیل معتبر نیست.",
	"NOT_EOA":             "آدرس دارای مالکیت خارجی (EOA) معتبر نیست.",
	"NOT_EQ_FIELD":        "مقدار باید با فیلد {{ .field }} مطابقت داشته باشد.",
	"NOT_FQDN":            "نام دامنه کاملاً واجد شرایط (FQDN) معتبر نیست.",
	"NOT_GTE":             "مقدار نمی\u200cتواند کمتر از {{ .n }} باشد.",
	"NOT_HASH":            "هش {{ .algorithm }} معتبر نیست.",
	"NOT_HEX":             "فقط می\u200cتواند شامل کاراکترهای هگزادسیمال باشد.",
	"NOT_IP":              "آدرس IP معتبر نیست.",
	"NOT_IPV4":            "آدرس IPv4 معتبر نیست.",
	"NOT_IPV6":            "آدرس IPv6 معتبر نیست.",
	"NOT_ISBN":            "شماره ISBN معتبر نیست.",
	"NOT_ISO31661_ALPHA2": "کد کشور ISO 3166-1 alpha-2 معتبر نیست.",
	"NOT_ISO31661_ALPHA3": "کد کشور ISO 3166-1 alpha-3 معتبر نیست.",
	"NOT_ISO6391":         "کد زبان ISO 639-1 معتبر نیست.",
	"NOT_LTE":             "مقدار نمی\u200cتواند کمتر از {{ .n }} باشد.",
	"NOT_LUHN":            "شماره LUHN معتبر نیست.",
	"NOT_MAC":             "آدرس MAC معتبر نیست.",
	"NOT_MAX_LEN":         "مقدار نمی\u200cتواند بیشتر از {{ .max }} باشد.",
	"NOT_MIN_LEN":         "مقدار نمی\u200cتواند کمتر از {{ .min }} باشد.",
	"NOT_TIME":            "زمان معتبر نیست.",
	"REQUIRED":            "مقدار الزامی وجود ندارد.",
	"NOT_URL":             "URL معتبر نیست.",
}
