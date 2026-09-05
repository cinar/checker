package locales

const (
	// ZhCN is the zh_cn locale.
	ZhCN = "zh-CN"
)

// ZhCNMessages is the map of zh-CN messages.
var ZhCNMessages = map[string]string{
	"NOT_AFTER":           "值必须在{{ .reference }}之后。",
	"NOT_ALPHA":           "只能包含字母。",
	"NOT_ALPHANUMERIC":    "不是字母数字字符串。",
	"NOT_ASCII":           "只能包含ASCII字符。",
	"NOT_BEFORE":          "值必须在{{ .reference }}之前。",
	"NOT_CIDR":            "不是有效的CIDR表示法。",
	"NOT_CREDIT_CARD":     "不是有效的信用卡号。",
	"NOT_DIGITS":          "只能包含数字。",
	"NOT_EMAIL":           "不是有效的电子邮件地址。",
	"EQ":                  "值不能等于{{ .forbidden }}。",
	"NOT_EQ":              "值必须等于{{ .expected }}。",
	"NOT_EOA":             "不是有效的外部拥有账户地址(EOA)。",
	"NOT_EQ_FIELD":        "值必须与字段{{ .field }}匹配。",
	"NOT_FQDN":            "不是有效的完全限定域名(FQDN)。",
	"NOT_GT":              "值必须大于{{ .n }}。",
	"NOT_GTE":             "值不能小于{{ .n }}。",
	"NOT_HASH":            "不是有效的{{ .algorithm }}哈希值。",
	"NOT_HEX":             "只能包含十六进制字符。",
	"NOT_IP":              "不是有效的IP地址。",
	"NOT_IPV4":            "不是有效的IPv4地址。",
	"NOT_IPV6":            "不是有效的IPv6地址。",
	"NOT_ISBN":            "不是有效的ISBN号。",
	"NOT_ISO31661_ALPHA2": "不是有效的ISO 3166-1 alpha-2国家代码。",
	"NOT_ISO31661_ALPHA3": "不是有效的ISO 3166-1 alpha-3国家代码。",
	"NOT_ISO6391":         "不是有效的ISO 639-1语言代码。",
	"NOT_LT":              "值必须小于{{ .n }}。",
	"NOT_LTE":             "值不能小于{{ .n }}。",
	"NOT_LUHN":            "不是有效的LUHN号。",
	"NOT_MAC":             "不是有效的MAC地址。",
	"NOT_MAX_LEN":         "值不能大于{{ .max }}。",
	"NOT_MIN_LEN":         "值不能小于{{ .min }}。",
	"NOT_NUMERIC":         "不是有效的数字字符串。",
	"NOT_ONE_OF":          "值必须是{{ .allowed }}之一。",
	"NOT_TIME":            "不是有效的时间。",
	"REQUIRED":            "缺少必填值。",
	"NOT_URL":             "不是有效的URL。",
	"NOT_UUID":            "不是有效的UUID。",
}
