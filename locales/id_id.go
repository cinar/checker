package locales

const (
	// IDID is the id_id locale.
	IDID = "id-ID"
)

// IDIDMessages is the map of id-ID messages.
var IDIDMessages = map[string]string{
	"NOT_AFTER":           "Nilai harus setelah {{ .reference }}.",
	"NOT_ALPHA":           "Hanya boleh berisi huruf.",
	"NOT_ALPHANUMERIC":    "Bukan string alfanumerik.",
	"NOT_ASCII":           "Hanya boleh berisi karakter ASCII.",
	"NOT_BEFORE":          "Nilai harus sebelum {{ .reference }}.",
	"NOT_CIDR":            "Bukan notasi CIDR yang valid.",
	"NOT_CONTAINS":        "Nilai harus mengandung {{ .substr }}.",
	"NOT_CREDIT_CARD":     "Bukan nomor kartu kredit yang valid.",
	"NOT_DIGITS":          "Hanya boleh berisi angka.",
	"NOT_EMAIL":           "Bukan alamat email yang valid.",
	"EQ":                  "Nilai tidak boleh sama dengan {{ .forbidden }}.",
	"NOT_EQ":              "Nilai harus sama dengan {{ .expected }}.",
	"NOT_ENDS_WITH":       "Nilai harus diakhiri dengan {{ .suffix }}.",
	"NOT_EOA":             "Bukan alamat yang dimiliki secara eksternal (EOA) yang valid.",
	"NOT_EQ_FIELD":        "Nilai harus sama dengan bidang {{ .field }}.",
	"NOT_FQDN":            "Bukan nama domain yang memenuhi syarat sepenuhnya (FQDN).",
	"NOT_GT":              "Nilai harus lebih besar dari {{ .n }}.",
	"NOT_GTE":             "Nilai tidak boleh kurang dari {{ .n }}.",
	"NOT_HASH":            "Bukan hash {{ .algorithm }} yang valid.",
	"NOT_HEX":             "Hanya boleh berisi karakter heksadesimal.",
	"NOT_IP":              "Bukan alamat IP yang valid.",
	"NOT_IPV4":            "Bukan alamat IPv4 yang valid.",
	"NOT_IPV6":            "Bukan alamat IPv6 yang valid.",
	"NOT_ISBN":            "Bukan nomor ISBN yang valid.",
	"NOT_ISO31661_ALPHA2": "Bukan kode negara ISO 3166-1 alpha-2 yang valid.",
	"NOT_ISO31661_ALPHA3": "Bukan kode negara ISO 3166-1 alpha-3 yang valid.",
	"NOT_ISO6391":         "Bukan kode bahasa ISO 639-1 yang valid.",
	"NOT_LT":              "Nilai harus kurang dari {{ .n }}.",
	"NOT_LTE":             "Nilai tidak boleh kurang dari {{ .n }}.",
	"NOT_LUHN":            "Bukan nomor LUHN yang valid.",
	"NOT_MAC":             "Bukan alamat MAC yang valid.",
	"NOT_MAX_LEN":         "Nilai tidak boleh lebih besar dari {{ .max }}.",
	"NOT_MIN_LEN":         "Nilai tidak boleh kurang dari {{ .min }}.",
	"NOT_NUMERIC":         "Bukan string numerik yang valid.",
	"NOT_ONE_OF":          "Nilai harus salah satu dari {{ .allowed }}.",
	"NOT_STARTS_WITH":     "Nilai harus dimulai dengan {{ .prefix }}.",
	"NOT_TIME":            "Bukan waktu yang valid.",
	"REQUIRED":            "Nilai yang wajib diisi tidak ada.",
	"NOT_URL":             "Bukan URL yang valid.",
	"NOT_UUID":            "Bukan UUID yang valid.",
}
