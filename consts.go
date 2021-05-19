package goprices

import "errors"

var (
	ErrNotSameCurrency = errors.New("not same currency")     // ErrNotSameCurrency used when manipulate not-same-type money amounts
	ErrUnknownType     = errors.New("unknown given type")    // ErrUnknownType is returned when given interface{} value is invalid
	ErrUnknownCurrency = errors.New("unknown currency unit") // ErrUnknownCurrency is returned when given currency unit is invalid
)

// money codes for all countries in the world
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
