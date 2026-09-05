package locales

const (
	// KoKR is the ko_kr locale.
	KoKR = "ko-KR"
)

// KoKRMessages is the map of ko-KR messages.
var KoKRMessages = map[string]string{
	"NOT_AFTER":           "값은 {{ .reference }} 이후여야 합니다.",
	"NOT_ALPHANUMERIC":    "영숫자 문자열이 아닙니다.",
	"NOT_ASCII":           "ASCII 문자만 포함할 수 있습니다.",
	"NOT_BEFORE":          "값은 {{ .reference }} 이전이어야 합니다.",
	"NOT_CIDR":            "유효한 CIDR 표기법이 아닙니다.",
	"NOT_CREDIT_CARD":     "유효한 신용카드 번호가 아닙니다.",
	"NOT_DIGITS":          "숫자만 포함할 수 있습니다.",
	"NOT_EMAIL":           "유효한 이메일 주소가 아닙니다.",
	"NOT_EOA":             "유효한 외부 소유 주소(EOA)가 아닙니다.",
	"NOT_EQ_FIELD":        "값은 {{ .field }} 필드와 일치해야 합니다.",
	"NOT_FQDN":            "정규화된 전체 도메인 이름(FQDN)이 아닙니다.",
	"NOT_GTE":             "값은 {{ .n }}보다 작을 수 없습니다.",
	"NOT_HASH":            "유효한 {{ .algorithm }} 해시가 아닙니다.",
	"NOT_HEX":             "16진수 문자만 포함할 수 있습니다.",
	"NOT_IP":              "유효한 IP 주소가 아닙니다.",
	"NOT_IPV4":            "유효한 IPv4 주소가 아닙니다.",
	"NOT_IPV6":            "유효한 IPv6 주소가 아닙니다.",
	"NOT_ISBN":            "유효한 ISBN 번호가 아닙니다.",
	"NOT_ISO31661_ALPHA2": "유효한 ISO 3166-1 alpha-2 국가 코드가 아닙니다.",
	"NOT_ISO31661_ALPHA3": "유효한 ISO 3166-1 alpha-3 국가 코드가 아닙니다.",
	"NOT_ISO6391":         "유효한 ISO 639-1 언어 코드가 아닙니다.",
	"NOT_LTE":             "값은 {{ .n }}보다 작을 수 없습니다.",
	"NOT_LUHN":            "유효한 LUHN 번호가 아닙니다.",
	"NOT_MAC":             "유효한 MAC 주소가 아닙니다.",
	"NOT_MAX_LEN":         "값은 {{ .max }}보다 클 수 없습니다.",
	"NOT_MIN_LEN":         "값은 {{ .min }}보다 작을 수 없습니다.",
	"NOT_TIME":            "유효한 시간이 아닙니다.",
	"REQUIRED":            "필수 값이 없습니다.",
	"NOT_URL":             "유효한 URL이 아닙니다.",
	"NOT_UUID":            "유효한 UUID가 아닙니다.",
}
