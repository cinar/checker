package locales

const (
	// ZhTW is the zh_tw locale.
	ZhTW = "zh-TW"
)

// ZhTWMessages is the map of zh-TW messages.
var ZhTWMessages = map[string]string{
	"NOT_AFTER":           "值必須在{{ .reference }}之後。",
	"NOT_ALPHANUMERIC":    "不是字母數字字串。",
	"NOT_ASCII":           "只能包含ASCII字元。",
	"NOT_BEFORE":          "值必須在{{ .reference }}之前。",
	"NOT_CIDR":            "不是有效的CIDR表示法。",
	"NOT_CREDIT_CARD":     "不是有效的信用卡號。",
	"NOT_DIGITS":          "只能包含數字。",
	"NOT_EMAIL":           "不是有效的電子郵件地址。",
	"NOT_EOA":             "不是有效的外部擁有帳戶地址(EOA)。",
	"NOT_EQ_FIELD":        "值必須與欄位{{ .field }}相符。",
	"NOT_FQDN":            "不是有效的完全限定網域名稱(FQDN)。",
	"NOT_GTE":             "值不能小於{{ .n }}。",
	"NOT_HASH":            "不是有效的{{ .algorithm }}雜湊值。",
	"NOT_HEX":             "只能包含十六進位字元。",
	"NOT_IP":              "不是有效的IP位址。",
	"NOT_IPV4":            "不是有效的IPv4位址。",
	"NOT_IPV6":            "不是有效的IPv6位址。",
	"NOT_ISBN":            "不是有效的ISBN號碼。",
	"NOT_ISO31661_ALPHA2": "不是有效的ISO 3166-1 alpha-2國家代碼。",
	"NOT_ISO31661_ALPHA3": "不是有效的ISO 3166-1 alpha-3國家代碼。",
	"NOT_ISO6391":         "不是有效的ISO 639-1語言代碼。",
	"NOT_LTE":             "值不能小於{{ .n }}。",
	"NOT_LUHN":            "不是有效的LUHN號碼。",
	"NOT_MAC":             "不是有效的MAC位址。",
	"NOT_MAX_LEN":         "值不能大於{{ .max }}。",
	"NOT_MIN_LEN":         "值不能小於{{ .min }}。",
	"NOT_TIME":            "不是有效的時間。",
	"REQUIRED":            "缺少必填值。",
	"NOT_URL":             "不是有效的URL。",
	"NOT_UUID":            "不是有效的UUID。",
}
