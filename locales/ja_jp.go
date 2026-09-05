package locales

const (
	// JaJP is the ja_jp locale.
	JaJP = "ja-JP"
)

// JaJPMessages is the map of ja-JP messages.
var JaJPMessages = map[string]string{
	"NOT_AFTER":           "値は{{ .reference }}より後である必要があります。",
	"NOT_ALPHA":           "文字のみを使用できます。",
	"NOT_ALPHANUMERIC":    "英数字の文字列ではありません。",
	"NOT_ASCII":           "ASCII文字のみを使用できます。",
	"NOT_BEFORE":          "値は{{ .reference }}より前である必要があります。",
	"NOT_CIDR":            "有効なCIDR表記ではありません。",
	"NOT_CONTAINS":        "値には{{ .substr }}を含める必要があります。",
	"NOT_CREDIT_CARD":     "有効なクレジットカード番号ではありません。",
	"NOT_DIGITS":          "数字のみを使用できます。",
	"NOT_EMAIL":           "有効なメールアドレスではありません。",
	"EQ":                  "値は{{ .forbidden }}と等しくてはなりません。",
	"NOT_EQ":              "値は{{ .expected }}と等しくなければなりません。",
	"NOT_ENDS_WITH":       "値は{{ .suffix }}で終わる必要があります。",
	"NOT_EOA":             "有効な外部所有アドレス(EOA)ではありません。",
	"NOT_EQ_FIELD":        "値はフィールド{{ .field }}と一致する必要があります。",
	"NOT_FQDN":            "完全修飾ドメイン名(FQDN)ではありません。",
	"NOT_GT":              "値は{{ .n }}より大きくなければなりません。",
	"NOT_GTE":             "値は{{ .n }}未満にはできません。",
	"NOT_HASH":            "有効な{{ .algorithm }}ハッシュではありません。",
	"NOT_HEX":             "16進数の文字のみを使用できます。",
	"NOT_IP":              "有効なIPアドレスではありません。",
	"NOT_IPV4":            "有効なIPv4アドレスではありません。",
	"NOT_IPV6":            "有効なIPv6アドレスではありません。",
	"NOT_ISBN":            "有効なISBN番号ではありません。",
	"NOT_ISO31661_ALPHA2": "有効なISO 3166-1 alpha-2国コードではありません。",
	"NOT_ISO31661_ALPHA3": "有効なISO 3166-1 alpha-3国コードではありません。",
	"NOT_ISO6391":         "有効なISO 639-1言語コードではありません。",
	"NOT_LT":              "値は{{ .n }}より小さくなければなりません。",
	"NOT_LEN":             "値の長さは{{ .len }}でなければなりません。",
	"NOT_LTE":             "値は{{ .n }}未満にはできません。",
	"NOT_LUHN":            "有効なLUHN番号ではありません。",
	"NOT_MAC":             "有効なMACアドレスではありません。",
	"NOT_MAX_LEN":         "値は{{ .max }}を超えることはできません。",
	"NOT_MIN_LEN":         "値は{{ .min }}未満にはできません。",
	"NOT_NUMERIC":         "有効な数値文字列ではありません。",
	"NOT_ONE_OF":          "値は次のいずれかである必要があります: {{ .allowed }}。",
	"NOT_STARTS_WITH":     "値は{{ .prefix }}で始まる必要があります。",
	"NOT_TIME":            "有効な時刻ではありません。",
	"REQUIRED":            "必須の値がありません。",
	"NOT_URL":             "有効なURLではありません。",
	"NOT_UUID":            "有効なUUIDではありません。",
}
