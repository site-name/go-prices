package goprices

import (
	"errors"

	"github.com/site-name/decimal"
)

var (
	ErrNotSameCurrency = errors.New("not same currency")     // ErrNotSameCurrency is used when perform operations between money with different currencies
	ErrUnknownType     = errors.New("unknown given type")    // ErrUnknownType is returned when a type is invalid
	ErrUnknownCurrency = errors.New("unknown currency unit") // ErrUnknownCurrency is returned when given currency unit is invalid
	ErrNillValue       = errors.New("argument must not be nil")
	ErrDivisorNotZero  = errors.New("divisor must not be zero")
	ErrInvalidRounding = errors.New("invalid rounding")
)

// Currencyable
type Currencyable interface {
	MyCurrency() string
}

// RoundFunc
type RoundFunc func(places int32) decimal.Decimal

// Rounding up/down money
type Rounding uint8

// Rounding variants
const (
	Up Rounding = iota
	Down
	Ceil
	Floor
)

var (
	CurrenciesMap map[string]string // map with keys are currency units, values are full currency names
)

// most well-known money units
const (
	AED = "AED"
	AFN = "AFN"
	ALL = "ALL"
	AMD = "AMD"
	ANG = "ANG"
	AOA = "AOA"
	ARS = "ARS"
	AUD = "AUD"
	AWG = "AWG"
	AZN = "AZN"
	BAM = "BAM"
	BBD = "BBD"
	BDT = "BDT"
	BGN = "BGN"
	BHD = "BHD"
	BIF = "BIF"
	BMD = "BMD"
	BND = "BND"
	BOB = "BOB"
	BRL = "BRL"
	BSD = "BSD"
	BTN = "BTN"
	BWP = "BWP"
	BYN = "BYN"
	BYR = "BYR"
	BZD = "BZD"
	CAD = "CAD"
	CDF = "CDF"
	CHF = "CHF"
	CLF = "CLF"
	CLP = "CLP"
	CNY = "CNY"
	COP = "COP"
	CRC = "CRC"
	CUC = "CUC"
	CUP = "CUP"
	CVE = "CVE"
	CZK = "CZK"
	DJF = "DJF"
	DKK = "DKK"
	DOP = "DOP"
	DZD = "DZD"
	EEK = "EEK"
	EGP = "EGP"
	ERN = "ERN"
	ETB = "ETB"
	EUR = "EUR"
	FJD = "FJD"
	FKP = "FKP"
	GBP = "GBP"
	GEL = "GEL"
	GGP = "GGP"
	GHC = "GHC"
	GHS = "GHS"
	GIP = "GIP"
	GMD = "GMD"
	GNF = "GNF"
	GTQ = "GTQ"
	GYD = "GYD"
	HKD = "HKD"
	HNL = "HNL"
	HRK = "HRK"
	HTG = "HTG"
	HUF = "HUF"
	IDR = "IDR"
	ILS = "ILS"
	IMP = "IMP"
	INR = "INR"
	IQD = "IQD"
	IRR = "IRR"
	ISK = "ISK"
	JEP = "JEP"
	JMD = "JMD"
	JOD = "JOD"
	JPY = "JPY"
	KES = "KES"
	KGS = "KGS"
	KHR = "KHR"
	KMF = "KMF"
	KPW = "KPW"
	KRW = "KRW"
	KWD = "KWD"
	KYD = "KYD"
	KZT = "KZT"
	LAK = "LAK"
	LBP = "LBP"
	LKR = "LKR"
	LRD = "LRD"
	LSL = "LSL"
	LTL = "LTL"
	LVL = "LVL"
	LYD = "LYD"
	MAD = "MAD"
	MDL = "MDL"
	MKD = "MKD"
	MMK = "MMK"
	MNT = "MNT"
	MOP = "MOP"
	MUR = "MUR"
	MVR = "MVR"
	MWK = "MWK"
	MXN = "MXN"
	MYR = "MYR"
	MZN = "MZN"
	NAD = "NAD"
	NGN = "NGN"
	NIO = "NIO"
	NOK = "NOK"
	NPR = "NPR"
	NZD = "NZD"
	OMR = "OMR"
	PAB = "PAB"
	PEN = "PEN"
	PGK = "PGK"
	PHP = "PHP"
	PKR = "PKR"
	PLN = "PLN"
	PYG = "PYG"
	QAR = "QAR"
	RON = "RON"
	RSD = "RSD"
	RUB = "RUB"
	RUR = "RUR"
	RWF = "RWF"
	SAR = "SAR"
	SBD = "SBD"
	SCR = "SCR"
	SDG = "SDG"
	SEK = "SEK"
	SGD = "SGD"
	SHP = "SHP"
	SKK = "SKK"
	SLL = "SLL"
	SOS = "SOS"
	SRD = "SRD"
	SSP = "SSP"
	STD = "STD"
	SVC = "SVC"
	SYP = "SYP"
	SZL = "SZL"
	THB = "THB"
	TJS = "TJS"
	TMT = "TMT"
	TND = "TND"
	TOP = "TOP"
	TRL = "TRL"
	TRY = "TRY"
	TTD = "TTD"
	TWD = "TWD"
	TZS = "TZS"
	UAH = "UAH"
	UGX = "UGX"
	USD = "USD"
	UYU = "UYU"
	UZS = "UZS"
	VEF = "VEF"
	VND = "VND"
	VUV = "VUV"
	WST = "WST"
	XAF = "XAF"
	XAG = "XAG"
	XAU = "XAU"
	XCD = "XCD"
	XDR = "XDR"
	XPF = "XPF"
	YER = "YER"
	ZAR = "ZAR"
	ZMW = "ZMW"
	ZWD = "ZWD"
)

func init() {
	CurrenciesMap = map[string]string{
		AED: "United Arab Emirates Dirham",
		AFN: "Afghan Afghani",
		ALL: "Albanian Lek",
		AMD: "Armenian Dram",
		ANG: "Netherlands Antillean Guilder",
		AOA: "Angolan Kwanza",
		ARS: "Argentine Peso",
		AUD: "Australian Dollar",
		AWG: "Aruban Florin",
		AZN: "Azerbaijani Manat",
		BAM: "Bosnia-Herzegovina Convertible Mark",
		BBD: "Barbadian Dollar",
		BDT: "Bangladeshi Taka",
		BGN: "Bulgarian Lev",
		BHD: "Bahraini Dinar",
		BIF: "Burundian Franc",
		BMD: "Bermudan Dollar",
		BND: "Brunei Dollar",
		BOB: "Bolivian Boliviano",
		BRL: "Brazilian Real",
		BSD: "Bahamian Dollar",
		BTN: "Bhutanese Ngultrum",
		BWP: "Botswanan Pula",
		BYN: "Belarusian Ruble",
		BZD: "Belize Dollar",
		CAD: "Canadian Dollar",
		CDF: "Congolese Franc",
		CHF: "Swiss Franc",
		CLF: "Chilean Unit of Account (UF)",
		CLP: "Chilean Peso",
		CNY: "Chinese Yuan",
		COP: "Colombian Peso",
		CRC: "Costa Rican Colón",
		CUC: "Cuban Convertible Peso",
		CUP: "Cuban Peso",
		CVE: "Cape Verdean Escudo",
		CZK: "Czech Republic Koruna",
		DJF: "Djiboutian Franc",
		DKK: "Danish Krone",
		DOP: "Dominican Peso",
		DZD: "Algerian Dinar",
		EGP: "Egyptian Pound",
		ERN: "Eritrean Nakfa",
		ETB: "Ethiopian Birr",
		EUR: "Euro",
		FJD: "Fijian Dollar",
		FKP: "Falkland Islands Pound",
		GBP: "British Pound Sterling",
		GEL: "Georgian Lari",
		GGP: "Guernsey Pound",
		GHS: "Ghanaian Cedi",
		GIP: "Gibraltar Pound",
		GMD: "Gambian Dalasi",
		GNF: "Guinean Franc",
		GTQ: "Guatemalan Quetzal",
		GYD: "Guyanaese Dollar",
		HKD: "Hong Kong Dollar",
		HNL: "Honduran Lempira",
		HRK: "Croatian Kuna",
		HTG: "Haitian Gourde",
		HUF: "Hungarian Forint",
		IDR: "Indonesian Rupiah",
		ILS: "Israeli New Sheqel",
		IMP: "Manx pound",
		INR: "Indian Rupee",
		IQD: "Iraqi Dinar",
		IRR: "Iranian Rial",
		ISK: "Icelandic Króna",
		JEP: "Jersey Pound",
		JMD: "Jamaican Dollar",
		JOD: "Jordanian Dinar",
		JPY: "Japanese Yen",
		KES: "Kenyan Shilling",
		KGS: "Kyrgystani Som",
		KHR: "Cambodian Riel",
		KMF: "Comorian Franc",
		KPW: "North Korean Won",
		KRW: "South Korean Won",
		KWD: "Kuwaiti Dinar",
		KYD: "Cayman Islands Dollar",
		KZT: "Kazakhstani Tenge",
		LAK: "Laotian Kip",
		LBP: "Lebanese Pound",
		LKR: "Sri Lankan Rupee",
		LRD: "Liberian Dollar",
		LSL: "Lesotho Loti",
		LYD: "Libyan Dinar",
		MAD: "Moroccan Dirham",
		MDL: "Moldovan Leu",
		MKD: "Macedonian Denar",
		MMK: "Myanma Kyat",
		MNT: "Mongolian Tugrik",
		MOP: "Macanese Pataca",
		MUR: "Mauritian Rupee",
		MVR: "Maldivian Rufiyaa",
		MWK: "Malawian Kwacha",
		MXN: "Mexican Peso",
		MYR: "Malaysian Ringgit",
		MZN: "Mozambican Metical",
		NAD: "Namibian Dollar",
		NGN: "Nigerian Naira",
		NIO: "Nicaraguan Córdoba",
		NOK: "Norwegian Krone",
		NPR: "Nepalese Rupee",
		NZD: "New Zealand Dollar",
		OMR: "Omani Rial",
		PAB: "Panamanian Balboa",
		PEN: "Peruvian Nuevo Sol",
		PGK: "Papua New Guinean Kina",
		PHP: "Philippine Peso",
		PKR: "Pakistani Rupee",
		PLN: "Polish Zloty",
		PYG: "Paraguayan Guarani",
		QAR: "Qatari Rial",
		RON: "Romanian Leu",
		RSD: "Serbian Dinar",
		RUB: "Russian Ruble",
		RWF: "Rwandan Franc",
		SAR: "Saudi Riyal",
		SBD: "Solomon Islands Dollar",
		SCR: "Seychellois Rupee",
		SDG: "Sudanese Pound",
		SEK: "Swedish Krona",
		SGD: "Singapore Dollar",
		SHP: "Saint Helena Pound",
		SLL: "Sierra Leonean Leone",
		SOS: "Somali Shilling",
		SRD: "Surinamese Dollar",
		SSP: "South Sudanese Pound",
		STD: "São Tomé and Príncipe Dobra (pre-2018)",
		SVC: "Salvadoran Colón",
		SYP: "Syrian Pound",
		SZL: "Swazi Lilangeni",
		THB: "Thai Baht",
		TJS: "Tajikistani Somoni",
		TMT: "Turkmenistani Manat",
		TND: "Tunisian Dinar",
		TOP: "Tongan Pa'anga",
		TRY: "Turkish Lira",
		TTD: "Trinidad and Tobago Dollar",
		TWD: "New Taiwan Dollar",
		TZS: "Tanzanian Shilling",
		UAH: "Ukrainian Hryvnia",
		UGX: "Ugandan Shilling",
		USD: "United States Dollar",
		UYU: "Uruguayan Peso",
		UZS: "Uzbekistan Som",
		VEF: "Venezuelan Bolívar Fuerte (Old)",
		XPF: "CFP Franc",
		YER: "Yemeni Rial",
		ZAR: "South African Rand",
		ZMW: "Zambian Kwacha",
		VND: "Vietnamese Dong",
		VUV: "Vanuatu Vatu",
		WST: "Samoan Tala",
		XAF: "CFA Franc BEAC",
		XAG: "Silver Ounce",
		XAU: "Gold Ounce",
		XCD: "East Caribbean Dollar",
		XDR: "Special Drawing Rights",
		// XOF: "CFA Franc BCEAO",
		// XPD: "Palladium Ounce",
		// VES: "Venezuelan Bolívar Soberano",
		// STN: "São Tomé and Príncipe Dobra",
		// MRO: "Mauritanian Ouguiya (pre-2018)",
		// MRU: "Mauritanian Ouguiya",
		// MGA: "Malagasy Ariary",
		// CNH: "Chinese Yuan (Offshore)",
		// BTC: "Bitcoin",
		// XPT: "Platinum Ounce",
		// ZWL: "Zimbabwean Dollar",
	}
}
