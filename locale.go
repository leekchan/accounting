package accounting

import ()

type Locale struct {
	Name           string // currency name
	FractionLength int    // default decimal length
	ThouSep        string // thousands seperator
	DecSep         string // decimal seperator
	SpaceSep       string // space seperator
	UTFSymbol      string // UTF symbol
	HTMLSymbol     string // HTML symbol
	ComSymbol      string // Common symbol
	Pre            bool   // symbol before or after currency
}

const empty string = ""

var LocaleInfo map[string]Locale = map[string]Locale{
	"AED": Locale{"UAE Dirham", 2, ",", ".", " ", empty, empty, "Dhs.", true},
	"AFA": Locale{"Afghani", 0, empty, empty, "060B", "&#x060B;", empty, true},
	"ALL": Locale{"Lek", 2, empty, empty, "", empty, empty, empty, true},
	"AMD": Locale{"Armenian Dram", 2, ",", ".", "", empty, empty, "AMD", false},
	"ANG": Locale{"Antillian Guilder", 2, ".", ",", " ", "0192", "&#x0192;", "NAf.", true},
	"AON": Locale{"New Kwanza", 0, empty, empty, "", empty, empty, empty, true},
	"ARS": Locale{"Argentine Peso", 2, ".", ",", "", "20B1", "&#x20B1;", "$", true},
	"ATS": Locale{"Schilling", 2, ".", ",", " ", empty, empty, "öS", true},
	"AUD": Locale{"Australian Dollar", 2, " ", ".", "", "0024", "&#x0024;", "$", true},
	"AWG": Locale{"Aruban Guilder", 2, ",", ".", " ", "0192", "&#x0192;", "AWG", true},
	"AZN": Locale{"Azerbaijanian Manat", 2, empty, empty, "", empty, empty, "m", true},
	"BAM": Locale{"Convertible Marks", 2, ",", ".", "", empty, empty, "AZM", false},
	"BBD": Locale{"Barbados Dollar", 2, empty, empty, "", "0024", "&#x0024;", empty, true},
	"BDT": Locale{"Taka", 2, ",", ".", " ", empty, empty, "Bt.", true},
	"BEF": Locale{"Belgian Franc", 0, ".", "", " ", "20A3", "&#x20A3;", "BEF", true},
	"BGN": Locale{"Lev", 2, " ", ",", " ", empty, empty, "lv", false},
	"BHD": Locale{"Bahraini Dinar", 3, ",", ".", " ", empty, empty, "BD", true},
	"BIF": Locale{"Burundi Franc", 0, empty, empty, "", empty, empty, empty, true},
	"BMD": Locale{"Bermudian Dollar", 2, ",", ".", "", "0024", "&#x0024;", "$", true},
	"BND": Locale{"Brunei Dollar", 2, ",", ".", "", "0024", "&#x0024;", "$", true},
	"BOB": Locale{"Bolivian Boliviano", 2, ",", ".", "", empty, empty, "Bs", true},
	"BRL": Locale{"Brazilian Real", 2, ".", ",", " ", "0052 0024", "R$", "R$", true},
	"BSD": Locale{"Bahamian Dollar", 2, ",", ".", "", "0024", "&#x0024;", "$", true},
	"BTN": Locale{"Bhutan Ngultrum", 2, empty, empty, "", empty, empty, empty, true},
	"BWP": Locale{"Pula", 2, ",", ".", "", empty, empty, "P", true},
	"BYR": Locale{"Belarussian Ruble", 0, empty, empty, "", empty, empty, empty, true},
	"BZD": Locale{"Belize Dollar", 2, ",", ".", "", "0024", "&#x0024;", "$", true},
	"CAD": Locale{"Canadian Dollar", 2, ",", ".", "", "0024", "&#x0024;", "$", true},
	"CDF": Locale{"Franc Congolais", 2, empty, empty, "", empty, empty, empty, true},
	"CHF": Locale{"Swiss Franc", 2, "'", ".", " ", empty, empty, "SFr.", true},
	"CLP": Locale{"Chilean Peso", 0, ".", "", "", "20B1", "&#x20B1;", "$", true},
	"CNY": Locale{"Yuan Renminbi", 2, ",", ".", "", "5713", "&#x5713;", "Y", true},
	"COP": Locale{"Colombian Peso", 2, ".", ",", "", "20B1", "&#x20B1;", "$", true},
	"CRC": Locale{"Costa Rican Colon", 2, ".", ",", " ", "20A1", "&#x20A1;", "₡", true},
	"CUP": Locale{"Cuban Peso", 2, ",", ".", " ", "20B1", "&#x20B1;", "$", true},
	"CVE": Locale{"Cape Verde Escudo", 0, empty, empty, "", empty, empty, empty, true},
	"CYP": Locale{"Cyprus Pound", 2, ".", ",", "", "00A3", "&#x00A3;", "£", true},
	"CZK": Locale{"Czech Koruna", 2, ".", ",", " ", empty, empty, "Kc", false},
	"DEM": Locale{"Deutsche Mark", 2, ".", ",", "", empty, empty, "DM", false},
	"DJF": Locale{"Djibouti Franc", 0, empty, empty, "", empty, empty, empty, true},
	"DKK": Locale{"Danish Krone", 2, ".", ",", "", empty, empty, "kr.", true},
	"DOP": Locale{"Dominican Peso", 2, ",", ".", " ", "20B1", "&#x20B1;", "$", true},
	"DZD": Locale{"Algerian Dinar", 2, empty, empty, "", empty, empty, empty, true},
	"ECS": Locale{"Sucre", 0, empty, empty, "", empty, empty, empty, true},
	"EEK": Locale{"Kroon", 2, " ", ",", " ", empty, empty, "EEK", false},
	"EGP": Locale{"Egyptian Pound", 2, ",", ".", " ", "00A3", "&#x00A3;", "L.E.", true},
	"ERN": Locale{"Nakfa", 0, empty, empty, "", empty, empty, empty, true},
	"ESP": Locale{"Spanish Peseta", 0, ".", "", " ", "20A7", "&#x20A7;", "Ptas", false},
	"ETB": Locale{"Ethiopian Birr", 0, empty, empty, "", empty, empty, empty, true},
	"EUR": Locale{"Euro", 2, ".", ",", "", "20AC", "&#x20AC;", "EUR", true},
	"FIM": Locale{"Markka", 2, " ", ",", " ", empty, empty, "mk", false},
	"FJD": Locale{"Fiji Dollar", 0, empty, empty, "", "0024", "&#x0024;", empty, true},
	"FKP": Locale{"Pound", 0, empty, empty, "", "00A3", "&#x00A3;", empty, true},
	"FRF": Locale{"French Franc", 2, " ", ",", " ", "20A3", "&#x20A3;", "FRF", false},
	"GBP": Locale{"Pound Sterling", 2, ",", ".", "", "00A3", "&#x00A3;", "£", true},
	"GEL": Locale{"Lari", 0, empty, empty, "", empty, empty, empty, true},
	"GHS": Locale{"Cedi", 2, ",", ".", "", "20B5", "&#x20B5;", "₵", true},
	"GIP": Locale{"Gibraltar Pound", 2, ",", ".", "", "00A3", "&#x00A3;", "£", true},
	"GMD": Locale{"Dalasi", 0, empty, empty, "", empty, empty, empty, true},
	"GNF": Locale{"Guinea Franc", 0, empty, empty, empty, empty, empty, empty, true},
	"GRD": Locale{"Drachma", 2, ".", ",", " ", "20AF", "&#x20AF;", "GRD", false},
	"GTQ": Locale{"Quetzal", 2, ",", ".", "", empty, empty, "Q.", true},
	"GWP": Locale{"Guinea-Bissau Peso", 0, empty, empty, empty, empty, empty, empty, true},
	"GYD": Locale{"Guyana Dollar", 0, empty, empty, "", "0024", "&#x0024;", empty, true},
	"HKD": Locale{"Hong Kong Dollar", 2, ",", ".", "", "0024", "&#x0024;", "HK$", true},
	"HNL": Locale{"Lempira", 2, ",", ".", " ", empty, empty, "L", true},
	"HRK": Locale{"Kuna", 2, ".", ",", " ", empty, empty, "kn", false},
	"HTG": Locale{"Gourde", 0, empty, empty, "", empty, empty, empty, true},
	"HUF": Locale{"Forint", 0, ".", "", " ", empty, empty, "Ft", false},
	"IDR": Locale{"Rupiah", 0, ".", ",", "", empty, empty, "Rp.", true},
	"IEP": Locale{"Irish Pound", 2, ",", ".", "", "00A3", "&#x00A3;", "£", true},
	"ILS": Locale{"New Israeli Sheqel", 2, ",", ".", " ", "20AA", "&#x20AA;", "NIS", false},
	"INR": Locale{"Indian Rupee", 2, ",", ".", "", "20A8", "&#x20A8;", "Rs.", true},
	"IQD": Locale{"Iraqi Dinar", 3, empty, empty, "", empty, empty, empty, true},
	"IRR": Locale{"Iranian Rial", 2, ",", ".", " ", "FDFC", "&#xFDFC;", "Rls", true},
	"ISK": Locale{"Iceland Krona", 2, ".", ",", " ", empty, empty, "kr", false},
	"ITL": Locale{"Italian Lira", 0, ".", "", " ", "20A4", "&#x20A4;", "L.", true},
	"JMD": Locale{"Jamaican Dollar", 2, ",", ".", "", "0024", "&#x0024;", "$", true},
	"JOD": Locale{"Jordanian Dinar", 3, ",", ".", " ", empty, empty, "JD", true},
	"JPY": Locale{"Yen", 0, ",", "", "", "00A5", "&#x00A5;", "¥", true},
	"KES": Locale{"Kenyan Shilling", 2, ",", ".", "", empty, empty, "Kshs.", true},
	"KGS": Locale{"Som", 0, empty, empty, "", empty, empty, empty, true},
	"KHR": Locale{"Riel", 2, empty, empty, "", "17DB", "&#x17DB;", empty, true},
	"KMF": Locale{"Comoro Franc", 0, empty, empty, "", empty, empty, empty, true},
	"KPW": Locale{"North Korean Won", 0, empty, empty, "", "20A9", "&#x20A9;", empty, true},
	"KRW": Locale{"Won", 0, ",", "", "", "20A9", "&#x20A9;", "\\", true},
	"KWD": Locale{"Kuwaiti Dinar", 3, ",", ".", " ", empty, empty, "KD", true},
	"KYD": Locale{"Cayman Islands Dollar", 2, ",", ".", "", "0024", "&#x0024;", "$", true},
	"KZT": Locale{"Tenge", 0, empty, empty, "", empty, empty, empty, true},
	"LAK": Locale{"Kip", 0, empty, empty, "", "20AD", "&#x20AD;", empty, true},
	"LBP": Locale{"Lebanese Pound", 0, " ", "", "", "00A3", "&#x00A3;", "L.L.", false},
	"LKR": Locale{"Sri Lanka Rupee", 0, empty, empty, "", "0BF9", "&#x0BF9;", empty, true},
	"LRD": Locale{"Liberian Dollar", 0, empty, empty, "", "0024", "&#x0024;", empty, true},
	"LSL": Locale{"Lesotho Maloti", 0, empty, empty, "", empty, empty, empty, true},
	"LTL": Locale{"Lithuanian Litas", 2, " ", ",", " ", empty, empty, "Lt", false},
	"LUF": Locale{"Luxembourg Franc", 0, "'", "", " ", "20A3", "&#x20A3;", "F", false},
	"LVL": Locale{"Latvian Lats", 2, ",", ".", " ", empty, empty, "Ls", true},
	"LYD": Locale{"Libyan Dinar", 0, empty, empty, "", empty, empty, empty, true},
	"MAD": Locale{"Moroccan Dirham", 0, empty, empty, "", empty, empty, empty, true},
	"MDL": Locale{"Moldovan Leu", 0, empty, empty, "", empty, empty, empty, true},
	"MGF": Locale{"Malagasy Franc", 0, empty, empty, "", empty, empty, empty, true},
	"MKD": Locale{"Denar", 2, ",", ".", " ", empty, empty, "MKD", false},
	"MMK": Locale{"Kyat", 0, empty, empty, "", empty, empty, empty, true},
	"MNT": Locale{"Tugrik", 0, empty, empty, "", "20AE", "&#x20AE;", empty, true},
	"MOP": Locale{"Pataca", 0, empty, empty, "", empty, empty, empty, true},
	"MRO": Locale{"Ouguiya", 0, empty, empty, "", empty, empty, empty, true},
	"MTL": Locale{"Maltese Lira", 2, ",", ".", "", "20A4", "&#x20A4;", "Lm", true},
	"MUR": Locale{"Mauritius Rupee", 0, ",", "", "", "20A8", "&#x20A8;", "Rs", true},
	"MVR": Locale{"Rufiyaa", 0, empty, empty, "", empty, empty, empty, true},
	"MWK": Locale{"Kwacha", 2, ",", ".", "", empty, empty, empty, true},
	"MXN": Locale{"Mexican Peso", 2, ",", ".", " ", "0024", "&#x0024;", "$", true},
	"MYR": Locale{"Malaysian Ringgit", 2, ",", ".", "", empty, empty, "RM", true},
	"MZN": Locale{"Metical", 2, ".", ",", " ", empty, empty, "Mt", false},
	"NAD": Locale{"Namibian Dollar", 0, empty, empty, "", "0024", "&#x0024;", empty, true},
	"NGN": Locale{"Naira", 2, ",", ".", ".", "20A6", "&#x20A6;", empty, true},
	"NIO": Locale{"Cordoba Oro", 0, empty, empty, "", empty, empty, empty, true},
	"NLG": Locale{"Netherlands Guilder", 2, ".", ",", " ", "0192", "&#x0192;", "f", true},
	"NOK": Locale{"Norwegian Krone", 2, ".", ",", " ", "kr", "kr", "kr", true},
	"NPR": Locale{"Nepalese Rupee", 2, ",", ".", " ", "20A8", "&#x20A8;", "Rs.", true},
	"NZD": Locale{"New Zealand Dollar", 2, ",", ".", "", "0024", "&#x0024;", "$", true},
	"OMR": Locale{"Rial Omani", 3, ",", ".", " ", "FDFC", "&#xFDFC;", "RO", true},
	"PAB": Locale{"Balboa", 0, empty, empty, "", empty, empty, empty, true},
	"PEN": Locale{"Nuevo Sol", 2, ",", ".", " ", "S/.", "S/.", "S/.", true},
	"PGK": Locale{"Kina", 0, empty, empty, "", empty, empty, empty, true},
	"PHP": Locale{"Philippine Peso", 2, ",", ".", "", "20B1", "&#x20B1;", "PHP", true},
	"PKR": Locale{"Pakistan Rupee", 2, ",", ".", "", "20A8", "&#x20A8;", "Rs.", true},
	"PLN": Locale{"Zloty", 2, " ", ",", " ", empty, empty, "zl", false},
	"PTE": Locale{"Portuguese Escudo", 0, ".", "", " ", empty, empty, "Esc", false},
	"PYG": Locale{"Guarani", 0, empty, empty, "", "20B2", "&#x20B2;", "Gs.", true},
	"QAR": Locale{"Qatari Rial", 0, empty, empty, "", "FDFC", "&#xFDFC;", empty, true},
	"RON": Locale{"Leu", 2, ".", ",", " ", empty, empty, "lei", false},
	"RSD": Locale{"Serbian Dinar", 2, empty, empty, empty, empty, empty, "din", false},
	"RUB": Locale{"Russian Ruble", 2, ".", ",", empty, "0440 0443 0431", "&#x0440;&#x0443;&#x0431;", "RUB", true},
	"RWF": Locale{"Rwanda Franc", 0, empty, empty, "", empty, empty, empty, true},
	"SAC": Locale{"S. African Rand Commerc.", 0, empty, empty, "", empty, empty, empty, true},
	"SAR": Locale{"Saudi Riyal", 2, ",", ".", " ", "FDFC", "&#xFDFC;", "SR", true},
	"SBD": Locale{"Solomon Islands Dollar", 0, empty, empty, "", "0024", "&#x0024;", empty, true},
	"SCR": Locale{"Seychelles Rupee", 0, empty, empty, "", "20A8", "&#x20A8;", empty, true},
	"SDG": Locale{"Sudanese Dinar", empty, empty, empty, empty, empty, empty, "LSd", true},
	"SDP": Locale{"Sudanese Pound", 0, empty, empty, "", empty, empty, empty, true},
	"SEK": Locale{"Swedish Krona", 2, " ", ",", " ", empty, empty, "kr", false},
	"SGD": Locale{"Singapore Dollar", 2, ",", ".", "", "0024", "&#x0024;", "$", true},
	"SHP": Locale{"St Helena Pound", 0, empty, empty, "", "00A3", "&#x00A3;", empty, true},
	"SIT": Locale{"Tolar", 2, ".", ",", " ", empty, empty, "SIT", false},
	"SKK": Locale{"Slovak Koruna", 2, " ", ",", " ", empty, empty, "Sk", false},
	"SLL": Locale{"Leone", 0, empty, empty, "", empty, empty, empty, true},
	"SOS": Locale{"Somali Shilling", 0, empty, empty, "", empty, empty, empty, true},
	"SRG": Locale{"Surinam Guilder", 0, empty, empty, "", empty, empty, empty, true},
	"STD": Locale{"Dobra", 0, empty, empty, "", empty, empty, empty, true},
	"SVC": Locale{"El Salvador Colon", 2, ",", ".", "", "20A1", "&#x20A1;", "¢", true},
	"SYP": Locale{"Syrian Pound", 0, empty, empty, "", "00A3", "&#x00A3;", empty, true},
	"SZL": Locale{"Lilangeni", 2, "", ".", "", empty, empty, "E", true},
	"THB": Locale{"Baht", 2, ",", ".", " ", "0E3F", "&#x0E3F;", "Bt", false},
	"TJR": Locale{"Tajik Ruble", 0, empty, empty, "", empty, empty, empty, true},
	"TJS": Locale{"Somoni", empty, empty, empty, empty, empty, empty, empty, true},
	"TMM": Locale{"Manat", 0, empty, empty, "", empty, empty, empty, true},
	"TND": Locale{"Tunisian Dinar", 3, empty, empty, "", empty, empty, empty, true},
	"TOP": Locale{"Pa'anga", 2, ",", ".", " ", empty, empty, "$", true},
	"TPE": Locale{"Timor Escudo", empty, empty, empty, empty, empty, empty, empty, true},
	"TRY": Locale{"Turkish Lira", 0, ",", "", "", "20A4", "&#x20A4;", "TL", false},
	"TTD": Locale{"Trinidad and Tobago Dollar", 0, empty, empty, "", "0024", "&#x0024;", empty, true},
	"TWD": Locale{"New Taiwan Dollar", 0, empty, empty, "", "0024", "&#x0024;", empty, true},
	"TZS": Locale{"Tanzanian Shilling", 2, ",", ".", " ", empty, empty, "TZs", false},
	"UAH": Locale{"Hryvnia", 2, " ", ",", "", "20B4", "&#x20B4", empty, false},
	"UGX": Locale{"Uganda Shilling", 0, empty, empty, "", empty, empty, empty, true},
	"USD": Locale{"US Dollar", 2, ",", ".", "", "0024", "&#x0024;", "$", true},
	"UYU": Locale{"Peso Uruguayo", 2, ".", ",", "", "20B1", "&#x20B1;", "$", true},
	"UZS": Locale{"Uzbekistan Sum", 0, empty, empty, "", empty, empty, empty, true},
	"VEF": Locale{"Bolivar", 2, ".", ",", " ", empty, empty, "Bs.F", true},
	"VND": Locale{"Dong", 2, ".", ",", " ", "20AB", "&#x20AB;", "Dong", false},
	"VUV": Locale{"Vatu", 0, ",", "", "", empty, empty, "VT", false},
	"WST": Locale{"Tala", 0, empty, empty, "", empty, empty, empty, true},
	"XAF": Locale{"CFA Franc BEAC", 0, empty, empty, "", empty, empty, empty, true},
	"XCD": Locale{"East Caribbean Dollar", 2, ",", ".", "", "0024", "&#x0024;", "$", true},
	"XOF": Locale{"CFA Franc BCEAO", empty, empty, empty, empty, empty, empty, empty, true},
	"XPF": Locale{"CFP Franc", 0, empty, empty, "", empty, empty, empty, true},
	"YER": Locale{"Yemeni Rial", 0, empty, empty, "", "FDFC", "&#xFDFC;", empty, true},
	"YUN": Locale{"New Dinar", 0, empty, empty, "", empty, empty, empty, true},
	"ZAR": Locale{"Rand", 2, " ", ".", " ", "0052", "&#x0052;", "R", true},
	"ZMK": Locale{"Kwacha", 0, empty, empty, "", empty, empty, empty, true},
	"ZRN": Locale{"New Zaire", empty, empty, empty, empty, empty, empty, empty, true},
	"ZWD": Locale{"Zimbabwe Dollar ", 2, " ", ".", "", "0024", "&#x0024;", "Z$", true},
}
