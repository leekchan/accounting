package accounting

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
	"AFA": Locale{"Afghani", 0, empty, empty, "060B", "&#x060B;", empty, "؋", true},
	"ALL": Locale{"Lek", 2, empty, empty, "", empty, empty, "Lek ", true},
	"AMD": Locale{"Armenian Dram", 2, ",", ".", "", empty, empty, "֏", false},
	"ANG": Locale{"Antillian Guilder", 2, ".", ",", " ", "0192", "&#x0192;", "ƒ", true},
	"AOA": Locale{"New Kwanza", 0, empty, empty, "", empty, empty, "Kz ", true},
	"ARS": Locale{"Argentine Peso", 2, ".", ",", "", "20B1", "&#x20B1;", "$", true},
	"ATS": Locale{"Schilling", 2, ".", ",", " ", empty, empty, "öS ", true},
	"AUD": Locale{"Australian Dollar", 2, " ", ".", "", "0024", "&#x0024;", "$", true},
	"AWG": Locale{"Aruban Guilder", 2, ",", ".", " ", "0192", "&#x0192;", "ƒ", true},
	"AZN": Locale{"Azerbaijanian Manat", 2, empty, empty, "", empty, empty, "₼", true},
	"BAM": Locale{"Convertible Marks", 2, ",", ".", "", empty, empty, "KM ", false},
	"BBD": Locale{"Barbados Dollar", 2, empty, empty, "", "0024", "&#x0024;", "Bds$ ", true},
	"BDT": Locale{"Taka", 2, ",", ".", " ", empty, empty, "Tk ", true},
	"BEF": Locale{"Belgian Franc", 0, ".", "", " ", "20A3", "&#x20A3;", "BEF ", true},
	"BGN": Locale{"Lev", 2, " ", ",", " ", empty, empty, "лв", false},
	"BHD": Locale{"Bahraini Dinar", 3, ",", ".", " ", empty, empty, "د.ب", true},
	"BIF": Locale{"Burundi Franc", 0, empty, empty, "", empty, empty, "FBu ", true},
	"BMD": Locale{"Bermudian Dollar", 2, ",", ".", "", "0024", "&#x0024;", "$", true},
	"BND": Locale{"Brunei Dollar", 2, ",", ".", "", "0024", "&#x0024;", "$", true},
	"BOB": Locale{"Bolivian Boliviano", 2, ",", ".", "", empty, empty, "$b", true},
	"BRL": Locale{"Brazilian Real", 2, ".", ",", " ", "0052 0024", "R$", "R$", true},
	"BSD": Locale{"Bahamian Dollar", 2, ",", ".", "", "0024", "&#x0024;", "$", true},
	"BTN": Locale{"Bhutan Ngultrum", 2, empty, empty, "", empty, empty, "BTN ", true},
	"BWP": Locale{"Pula", 2, ",", ".", "", empty, empty, "P", true},
	"BYR": Locale{"Belarussian Ruble", 0, empty, empty, "", empty, empty, "p. ", true},
	"BZD": Locale{"Belize Dollar", 2, ",", ".", "", "0024", "&#x0024;", "$", true},
	"CAD": Locale{"Canadian Dollar", 2, ",", ".", "", "0024", "&#x0024;", "CA$", true},
	"CDF": Locale{"Franc Congolais", 2, empty, empty, "", empty, empty, "FC ", true},
	"CHF": Locale{"Swiss Franc", 2, "'", ".", " ", empty, empty, "CHF ", true},
	"CLP": Locale{"Chilean Peso", 0, ".", "", "", "20B1", "&#x20B1;", "$", true},
	"CNY": Locale{"Yuan Renminbi", 2, ",", ".", "", "5713", "&#x5713;", "¥", true},
	"COP": Locale{"Colombian Peso", 2, ".", ",", "", "20B1", "&#x20B1;", "$", true},
	"CRC": Locale{"Costa Rican Colon", 2, ".", ",", " ", "20A1", "&#x20A1;", "₡", true},
	"CUP": Locale{"Cuban Peso", 2, ",", ".", " ", "20B1", "&#x20B1;", "$", true},
	"CVE": Locale{"Cape Verde Escudo", 0, empty, empty, "", empty, empty, "$", true},
	"CYP": Locale{"Cyprus Pound", 2, ".", ",", "", "00A3", "&#x00A3;", "£", true},
	"CZK": Locale{"Czech Koruna", 2, ".", ",", " ", empty, empty, "Kč", false},
	"DEM": Locale{"Deutsche Mark", 2, ".", ",", "", empty, empty, "DM ", false},
	"DJF": Locale{"Djibouti Franc", 0, empty, empty, "", empty, empty, "DJF ", true},
	"DKK": Locale{"Danish Krone", 2, ".", ",", "", empty, empty, "kr.", true},
	"DOP": Locale{"Dominican Peso", 2, ",", ".", " ", "20B1", "&#x20B1;", "$", true},
	"DZD": Locale{"Algerian Dinar", 2, empty, empty, "", empty, empty, "DA", true},
	"ECS": Locale{"Sucre", 0, empty, empty, "", empty, empty, "S.", true},
	"EEK": Locale{"Kroon", 2, " ", ",", " ", empty, empty, "kr", false},
	"EGP": Locale{"Egyptian Pound", 2, ",", ".", " ", "00A3", "&#x00A3;", "£", true},
	"ERN": Locale{"Nakfa", 0, empty, empty, "", empty, empty, "NKf ", true},
	"ESP": Locale{"Spanish Peseta", 0, ".", "", " ", "20A7", "&#x20A7;", "Ptas", false},
	"ETB": Locale{"Ethiopian Birr", 0, empty, empty, "", empty, empty, "BR", true},
	"EUR": Locale{"Euro", 2, ".", ",", "", "20AC", "&#x20AC;", "€", true},
	"FIM": Locale{"Markka", 2, " ", ",", " ", empty, empty, "mk", false},
	"FJD": Locale{"Fiji Dollar", 0, empty, empty, "", "0024", "&#x0024;", "FJ$", true},
	"FKP": Locale{"Pound", 0, empty, empty, "", "00A3", "&#x00A3;", "£", true},
	"FRF": Locale{"French Franc", 2, " ", ",", " ", "20A3", "&#x20A3;", "Fr", false},
	"GBP": Locale{"Pound Sterling", 2, ",", ".", "", "00A3", "&#x00A3;", "£", true},
	"GEL": Locale{"Lari", 0, empty, empty, "", empty, empty, "GEL ", true},
	"GHS": Locale{"Cedi", 2, ",", ".", "", "20B5", "&#x20B5;", "₵", true},
	"GIP": Locale{"Gibraltar Pound", 2, ",", ".", "", "00A3", "&#x00A3;", "£", true},
	"GMD": Locale{"Dalasi", 0, empty, empty, "", empty, empty, "GMD ", true},
	"GNF": Locale{"Guinea Franc", 0, empty, empty, empty, empty, empty, "FG", true},
	"GRD": Locale{"Drachma", 2, ".", ",", " ", "20AF", "&#x20AF;", "GRD ", false},
	"GTQ": Locale{"Quetzal", 2, ",", ".", "", empty, empty, "Q.", true},
	"GWP": Locale{"Guinea-Bissau Peso", 0, empty, empty, empty, empty, empty, "GWP ", true},
	"GYD": Locale{"Guyana Dollar", 0, empty, empty, "", "0024", "&#x0024;", "$", true},
	"HKD": Locale{"Hong Kong Dollar", 2, ",", ".", "", "0024", "&#x0024;", "HK$", true},
	"HNL": Locale{"Lempira", 2, ",", ".", " ", empty, empty, "L", true},
	"HRK": Locale{"Kuna", 2, ".", ",", " ", empty, empty, "kn", false},
	"HTG": Locale{"Gourde", 0, empty, empty, "", empty, empty, "G", true},
	"HUF": Locale{"Forint", 0, ".", "", " ", empty, empty, "Ft", false},
	"IDR": Locale{"Rupiah", 0, ".", ",", "", empty, empty, "Rp", true},
	"IEP": Locale{"Irish Pound", 2, ",", ".", "", "00A3", "&#x00A3;", "£", true},
	"ILS": Locale{"New Israeli Sheqel", 2, ",", ".", " ", "20AA", "&#x20AA;", "₪", false},
	"INR": Locale{"Indian Rupee", 2, ",", ".", "", "20A8", "&#x20A8;", "₹", true},
	"IQD": Locale{"Iraqi Dinar", 3, empty, empty, "", empty, empty, "د.ع", true},
	"IRR": Locale{"Iranian Rial", 2, ",", ".", " ", "FDFC", "&#xFDFC;", "﷼", true},
	"ISK": Locale{"Iceland Krona", 2, ".", ",", " ", empty, empty, "kr", false},
	"ITL": Locale{"Italian Lira", 0, ".", "", " ", "20A4", "&#x20A4;", "L.", true},
	"JMD": Locale{"Jamaican Dollar", 2, ",", ".", "", "0024", "&#x0024;", "$", true},
	"JOD": Locale{"Jordanian Dinar", 3, ",", ".", " ", empty, empty, "JD", true},
	"JPY": Locale{"Yen", 0, ",", "", "", "00A5", "&#x00A5;", "¥", true},
	"KES": Locale{"Kenyan Shilling", 2, ",", ".", "", empty, empty, "Ksh", true},
	"KGS": Locale{"Som", 0, empty, empty, "", empty, empty, "лв", true},
	"KHR": Locale{"Riel", 2, empty, empty, "", "17DB", "&#x17DB;", "៛", true},
	"KMF": Locale{"Comoro Franc", 0, empty, empty, "", empty, empty, "KMF ", true},
	"KPW": Locale{"North Korean Won", 0, empty, empty, "", "20A9", "&#x20A9;", "₩", true},
	"KRW": Locale{"Won", 0, ",", "", "", "20A9", "&#x20A9;", "₩", true},
	"KWD": Locale{"Kuwaiti Dinar", 3, ",", ".", " ", empty, empty, "ك", true},
	"KYD": Locale{"Cayman Islands Dollar", 2, ",", ".", "", "0024", "&#x0024;", "$", true},
	"KZT": Locale{"Tenge", 0, empty, empty, "", empty, empty, "₸", true},
	"LAK": Locale{"Kip", 0, empty, empty, "", "20AD", "&#x20AD;", "₭", true},
	"LBP": Locale{"Lebanese Pound", 0, " ", "", "", "00A3", "&#x00A3;", "ل.ل", false},
	"LKR": Locale{"Sri Lanka Rupee", 0, empty, empty, "", "0BF9", "&#x0BF9;", "₨", true},
	"LRD": Locale{"Liberian Dollar", 0, empty, empty, "", "0024", "&#x0024;", "$", true},
	"LSL": Locale{"Lesotho Maloti", 0, empty, empty, "", empty, empty, "LSL ", true},
	"LTL": Locale{"Lithuanian Litas", 2, " ", ",", " ", empty, empty, "Lt", false},
	"LUF": Locale{"Luxembourg Franc", 0, "'", "", " ", "20A3", "&#x20A3;", "F", false},
	"LVL": Locale{"Latvian Lats", 2, ",", ".", " ", empty, empty, "Ls", true},
	"LYD": Locale{"Libyan Dinar", 0, empty, empty, "", empty, empty, "LD ", true},
	"MAD": Locale{"Moroccan Dirham", 0, empty, empty, "", empty, empty, "MAD ", true},
	"MDL": Locale{"Moldovan Leu", 0, empty, empty, "", empty, empty, "MDL ", true},
	"MGF": Locale{"Malagasy Franc", 0, empty, empty, "", empty, empty, "MF", true},
	"MKD": Locale{"Denar", 2, ",", ".", " ", empty, empty, "ден", false},
	"MMK": Locale{"Kyat", 0, empty, empty, "", empty, empty, "K", true},
	"MNT": Locale{"Tugrik", 0, empty, empty, "", "20AE", "&#x20AE;", "₮", true},
	"MOP": Locale{"Pataca", 0, empty, empty, "", empty, empty, "MOP$", true},
	"MRO": Locale{"Ouguiya", 0, empty, empty, "", empty, empty, "MRO ", true},
	"MTL": Locale{"Maltese Lira", 2, ",", ".", "", "20A4", "&#x20A4;", "Lm", true},
	"MUR": Locale{"Mauritius Rupee", 0, ",", "", "", "20A8", "&#x20A8;", "Rs", true},
	"MVR": Locale{"Rufiyaa", 0, empty, empty, "", empty, empty, "MVR ", true},
	"MWK": Locale{"Kwacha", 2, ",", ".", "", empty, empty, "MK ", true},
	"MXN": Locale{"Mexican Peso", 2, ",", ".", " ", "0024", "&#x0024;", "$", true},
	"MYR": Locale{"Malaysian Ringgit", 2, ",", ".", "", empty, empty, "RM", true},
	"MZN": Locale{"Metical", 2, ".", ",", " ", empty, empty, "Mt", false},
	"NAD": Locale{"Namibian Dollar", 0, empty, empty, "", "0024", "&#x0024;", "$", true},
	"NGN": Locale{"Naira", 2, ",", ".", ".", "20A6", "&#x20A6;", "₦", true},
	"NIO": Locale{"Cordoba Oro", 0, empty, empty, "", empty, empty, "C$", true},
	"NLG": Locale{"Netherlands Guilder", 2, ".", ",", " ", "0192", "&#x0192;", "ƒ", true},
	"NOK": Locale{"Norwegian Krone", 2, ".", ",", " ", "kr", "kr", "kr", true},
	"NPR": Locale{"Nepalese Rupee", 2, ",", ".", " ", "20A8", "&#x20A8;", "Rs.", true},
	"NZD": Locale{"New Zealand Dollar", 2, ",", ".", "", "0024", "&#x0024;", "$", true},
	"OMR": Locale{"Rial Omani", 3, ",", ".", " ", "FDFC", "&#xFDFC;", "RO", true},
	"PAB": Locale{"Balboa", 0, empty, empty, "", empty, empty, "B/.", true},
	"PEN": Locale{"Nuevo Sol", 2, ",", ".", " ", "S/.", "S/.", "S/.", true},
	"PGK": Locale{"Kina", 0, empty, empty, "", empty, empty, "K", true},
	"PHP": Locale{"Philippine Peso", 2, ",", ".", "", "20B1", "&#x20B1;", "₱", true},
	"PKR": Locale{"Pakistan Rupee", 2, ",", ".", "", "20A8", "&#x20A8;", "Rs", true},
	"PLN": Locale{"Zloty", 2, " ", ",", " ", empty, empty, "zł", false},
	"PTE": Locale{"Portuguese Escudo", 0, ".", "", " ", empty, empty, " $", false},
	"PYG": Locale{"Guarani", 0, empty, empty, "", "20B2", "&#x20B2;", "Gs", true},
	"QAR": Locale{"Qatari Rial", 0, empty, empty, "", "FDFC", "&#xFDFC;", "﷼ ", true},
	"RON": Locale{"Leu", 2, ".", ",", " ", empty, empty, "lei", false},
	"RSD": Locale{"Serbian Dinar", 2, empty, empty, empty, empty, empty, "РСД", false},
	"RUB": Locale{"Russian Ruble", 2, ".", ",", empty, "0440 0443 0431", "&#x0440;&#x0443;&#x0431;", "₽", true},
	"RWF": Locale{"Rwanda Franc", 0, empty, empty, "", empty, empty, "RWF ", true},
	"SAC": Locale{"S. African Rand Commerc.", 0, empty, empty, "", empty, empty, "SAC ", true},
	"SAR": Locale{"Saudi Riyal", 2, ",", ".", " ", "FDFC", "&#xFDFC;", "﷼", true},
	"SBD": Locale{"Solomon Islands Dollar", 0, empty, empty, "", "0024", "&#x0024;", "$", true},
	"SCR": Locale{"Seychelles Rupee", 0, empty, empty, "", "20A8", "&#x20A8;", "₨", true},
	"SDG": Locale{"Sudanese Dinar", 0, empty, empty, empty, empty, empty, "LSd", true},
	"SDP": Locale{"Sudanese Pound", 0, empty, empty, "", empty, empty, "SDP ", true},
	"SEK": Locale{"Swedish Krona", 2, " ", ",", " ", empty, empty, "kr", false},
	"SGD": Locale{"Singapore Dollar", 2, ",", ".", "", "0024", "&#x0024;", "$", true},
	"SHP": Locale{"St Helena Pound", 0, empty, empty, "", "00A3", "&#x00A3;", "£", true},
	"SIT": Locale{"Tolar", 2, ".", ",", " ", empty, empty, "SIT ", false},
	"SKK": Locale{"Slovak Koruna", 2, " ", ",", " ", empty, empty, "Sk", false},
	"SLL": Locale{"Leone", 0, empty, empty, "", empty, empty, "Le ", true},
	"SOS": Locale{"Somali Shilling", 0, empty, empty, "", empty, empty, "S", true},
	"SRG": Locale{"Surinam Guilder", 0, empty, empty, "", empty, empty, "SRG ", true},
	"STD": Locale{"Dobra", 0, empty, empty, "", empty, empty, "DB", true},
	"SVC": Locale{"El Salvador Colon", 2, ",", ".", "", "20A1", "&#x20A1;", "¢", true},
	"SYP": Locale{"Syrian Pound", 0, empty, empty, "", "00A3", "&#x00A3;", "£", true},
	"SZL": Locale{"Lilangeni", 2, "", ".", "", empty, empty, "E", true},
	"THB": Locale{"Baht", 2, ",", ".", " ", "0E3F", "&#x0E3F;", "Bt", false},
	"TJR": Locale{"Tajik Ruble", 0, empty, empty, "", empty, empty, "TJR ", true},
	"TJS": Locale{"Somoni", 0, empty, empty, empty, empty, empty, "TJS ", true},
	"TMM": Locale{"Manat", 0, empty, empty, "", empty, empty, "T", true},
	"TND": Locale{"Tunisian Dinar", 3, empty, empty, "", empty, empty, "TND ", true},
	"TOP": Locale{"Pa'anga", 2, ",", ".", " ", empty, empty, "$", true},
	"TPE": Locale{"Timor Escudo", 0, empty, empty, empty, empty, empty, "TPE ", true},
	"TRY": Locale{"Turkish Lira", 0, ",", "", "", "20A4", "&#x20A4;", "TL", false},
	"TTD": Locale{"Trinidad and Tobago Dollar", 0, empty, empty, "", "0024", "&#x0024;", "TT$", true},
	"TWD": Locale{"New Taiwan Dollar", 0, empty, empty, "", "0024", "&#x0024;", "NT$", true},
	"TZS": Locale{"Tanzanian Shilling", 2, ",", ".", " ", empty, empty, "TZs", false},
	"UAH": Locale{"Hryvnia", 2, " ", ",", "", "20B4", "&#x20B4", "UAH ", false},
	"UGX": Locale{"Uganda Shilling", 0, empty, empty, "", empty, empty, "UGX", true},
	"USD": Locale{"US Dollar", 2, ",", ".", "", "0024", "&#x0024;", "$", true},
	"UYU": Locale{"Peso Uruguayo", 2, ".", ",", "", "20B1", "&#x20B1;", "$", true},
	"UZS": Locale{"Uzbekistan Sum", 0, empty, empty, "", empty, empty, "лв", true},
	"VEF": Locale{"Bolivar", 2, ".", ",", " ", empty, empty, "Bs.", true},
	"VND": Locale{"Dong", 2, ".", ",", " ", "20AB", "&#x20AB;", "₫", true},
	"VUV": Locale{"Vatu", 0, ",", "", "", empty, empty, "VT", false},
	"WST": Locale{"Tala", 0, empty, empty, "", empty, empty, "WST ", true},
	"XAF": Locale{"CFA Franc BEAC", 0, empty, empty, "", empty, empty, "$", true},
	"XCD": Locale{"East Caribbean Dollar", 2, ",", ".", "", "0024", "&#x0024;", "$", true},
	"XOF": Locale{"CFA Franc BCEAO", 0, empty, empty, empty, empty, empty, "XOF ", true},
	"XPF": Locale{"CFP Franc", 0, empty, empty, "", empty, empty, "XPF ", true},
	"YER": Locale{"Yemeni Rial", 0, empty, empty, "", "FDFC", "&#xFDFC;", "﷼ ", true},
	"YUN": Locale{"New Dinar", 0, empty, empty, "", empty, empty, "YUN ", true},
	"ZAR": Locale{"Rand", 2, " ", ".", " ", "0052", "&#x0052;", "R", true},
	"ZMK": Locale{"Kwacha", 0, empty, empty, "", empty, empty, "ZMK ", true},
	"ZRN": Locale{"New Zaire", 0, empty, empty, empty, empty, empty, "ZRN ", true},
	"ZWD": Locale{"Zimbabwe Dollar ", 2, " ", ".", "", "0024", "&#x0024;", "Z$", true},
}
