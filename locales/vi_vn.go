package locales

const (
	// ViVN is the vi_vn locale.
	ViVN = "vi-VN"
)

// ViVNMessages is the map of vi-VN messages.
var ViVNMessages = map[string]string{
	"NOT_AFTER":           "Giá trị phải sau {{ .reference }}.",
	"NOT_ALPHA":           "Chỉ được chứa chữ cái.",
	"NOT_ALPHANUMERIC":    "Không phải là chuỗi chữ và số.",
	"NOT_ASCII":           "Chỉ được chứa ký tự ASCII.",
	"NOT_BEFORE":          "Giá trị phải trước {{ .reference }}.",
	"NOT_CIDR":            "Không phải là ký hiệu CIDR hợp lệ.",
	"NOT_CREDIT_CARD":     "Không phải là số thẻ tín dụng hợp lệ.",
	"NOT_DIGITS":          "Chỉ được chứa chữ số.",
	"NOT_EMAIL":           "Không phải là địa chỉ email hợp lệ.",
	"EQ":                  "Giá trị không được bằng {{ .forbidden }}.",
	"NOT_EQ":              "Giá trị phải bằng {{ .expected }}.",
	"NOT_EOA":             "Không phải là địa chỉ sở hữu bên ngoài (EOA) hợp lệ.",
	"NOT_EQ_FIELD":        "Giá trị phải khớp với trường {{ .field }}.",
	"NOT_FQDN":            "Không phải là tên miền đầy đủ điều kiện (FQDN) hợp lệ.",
	"NOT_GT":              "Giá trị phải lớn hơn {{ .n }}.",
	"NOT_GTE":             "Giá trị không được nhỏ hơn {{ .n }}.",
	"NOT_HASH":            "Không phải là hàm băm {{ .algorithm }} hợp lệ.",
	"NOT_HEX":             "Chỉ được chứa ký tự thập lục phân.",
	"NOT_IP":              "Không phải là địa chỉ IP hợp lệ.",
	"NOT_IPV4":            "Không phải là địa chỉ IPv4 hợp lệ.",
	"NOT_IPV6":            "Không phải là địa chỉ IPv6 hợp lệ.",
	"NOT_ISBN":            "Không phải là số ISBN hợp lệ.",
	"NOT_ISO31661_ALPHA2": "Không phải là mã quốc gia ISO 3166-1 alpha-2 hợp lệ.",
	"NOT_ISO31661_ALPHA3": "Không phải là mã quốc gia ISO 3166-1 alpha-3 hợp lệ.",
	"NOT_ISO6391":         "Không phải là mã ngôn ngữ ISO 639-1 hợp lệ.",
	"NOT_LT":              "Giá trị phải nhỏ hơn {{ .n }}.",
	"NOT_LTE":             "Giá trị không được nhỏ hơn {{ .n }}.",
	"NOT_LUHN":            "Không phải là số LUHN hợp lệ.",
	"NOT_MAC":             "Không phải là địa chỉ MAC hợp lệ.",
	"NOT_MAX_LEN":         "Giá trị không được lớn hơn {{ .max }}.",
	"NOT_MIN_LEN":         "Giá trị không được nhỏ hơn {{ .min }}.",
	"NOT_NUMERIC":         "Không phải là chuỗi số hợp lệ.",
	"NOT_ONE_OF":          "Giá trị phải là một trong {{ .allowed }}.",
	"NOT_TIME":            "Không phải là thời gian hợp lệ.",
	"REQUIRED":            "Thiếu giá trị bắt buộc.",
	"NOT_URL":             "Không phải là URL hợp lệ.",
	"NOT_UUID":            "Không phải là UUID hợp lệ.",
}
