package accounting

import (
	"math/big"
	"strings"

	"github.com/cockroachdb/apd"
	"github.com/shopspring/decimal"
)

type Accounting struct {
	Symbol         string // currency symbol (required)
	Precision      int    // currency precision (decimal places) (optional / default: 0)
	Thousand       string // thousand separator (optional / default: ,)
	Decimal        string // decimal separator (optional / default: .)
	Format         string // simple format string allows control of symbol position (%v = value, %s = symbol) (default: %s%v)
	FormatNegative string // format string for negative values (optional / default: strings.Replace(strings.Replace(accounting.Format, "-", "", -1), "%v", "-%v", -1))
	FormatZero     string // format string for zero values (optional / default: Format)
	isInitialized bool // is set to true if used via DefaultAccounting or NewAccounting
}


// DefaultAccounting returns the Accounting with default settings
func DefaultAccounting(symbol string, precision int) *Accounting {
	ac := &Accounting{Symbol:symbol,Precision:precision}
	ac.init()
	ac.isInitialized = true
	return ac
}


// NewAccounting returns the Accounting with default settings
func NewAccounting(symbol string, precision int, thousand, decimal, format, formatNegative, formatZero string) *Accounting {
	ac := &Accounting{
		Symbol: symbol,
		Precision: precision,
		Thousand: thousand,
		Decimal: decimal,
		Format: format,
		FormatNegative: formatNegative,
		FormatZero: formatZero,
	}
	ac.isInitialized = true
	return ac
}

// SetThousandSeparator sets the separator for the thousands separation
func (accounting *Accounting) SetThousandSeparator(str string) {
	accounting.Thousand = str
}

// SetDecimalSeparator sets the separator for the decimal separation
func (accounting *Accounting) SetDecimalSeparator(str string) {
	accounting.Decimal = str
}

// SetFormat sets the Format default: %s%v (%s=Symbol;%v=Value)
func (accounting *Accounting) SetFormat(str string) {
	accounting.Format = str
}

// SetFormatNegative sets the Format for negative values default: -%s%v (%s=Symbol;%v=Value)
func (accounting *Accounting) SetFormatNegative(str string) {
	accounting.FormatNegative = str
}

// SetFormatZero sets the Format for zero values default: %s%v (%s=Symbol;%v=Value)
func (accounting *Accounting) SetFormatZero(str string) {
	accounting.FormatZero = str
}

func (accounting *Accounting) init() {
	if accounting.Thousand == "" {
		accounting.Thousand = ","
	}

	if accounting.Decimal == "" {
		accounting.Decimal = "."
	}

	if accounting.Format == "" {
		accounting.Format = "%s%v"
	}

	if accounting.FormatNegative == "" {
		accounting.FormatNegative = "-" + accounting.Format
	}

	if accounting.FormatZero == "" {
		accounting.FormatZero = accounting.Format
	}
}

func (accounting *Accounting) formatMoneyString(formattedNumber string) string {
	var format string

	formattedZero := FormatNumber(0, accounting.Precision, accounting.Thousand, accounting.Decimal)

	if formattedNumber[0] == '-' {
		format = accounting.FormatNegative
		formattedNumber = formattedNumber[1:]
	} else if formattedNumber == formattedZero {
		format = accounting.FormatZero
	} else {
		format = accounting.Format
	}

	result := strings.Replace(format, "%s", accounting.Symbol, -1)
	result = strings.Replace(result, "%v", formattedNumber, -1)

	return result
}

// FormatMoney is a function for formatting numbers as money values,
// with customisable currency symbol, precision (decimal places), and thousand/decimal separators.
// FormatMoney supports various types of value by runtime reflection.
// If you don't need runtime type evaluation, please refer to FormatMoneyInt or FormatMoneyFloat64.
// (supported value types : int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, *big.Rat)
func (accounting *Accounting) FormatMoney(value interface{}) string {
	if !accounting.isInitialized {
		accounting.init()
	}
	formattedNumber := FormatNumber(value, accounting.Precision, accounting.Thousand, accounting.Decimal)
	return accounting.formatMoneyString(formattedNumber)
}

// FormatMoneyInt only supports int type value. It is faster than FormatMoney,
// because it does not do any runtime type evaluation.
func (accounting *Accounting) FormatMoneyInt(value int) string {
	if !accounting.isInitialized {
		accounting.init()
	}
	formattedNumber := FormatNumberInt(value, accounting.Precision, accounting.Thousand, accounting.Decimal)
	return accounting.formatMoneyString(formattedNumber)
}

// FormatMoneyFloat64 only supports float64 value. It is faster than FormatMoney,
// because it does not do any runtime type evaluation.
// (Caution: Please do not use float64 to count money.
// Floats can have errors when you perform operations on them.
// Using big.Rat is highly recommended.)
func (accounting *Accounting) FormatMoneyFloat64(value float64) string {
	if !accounting.isInitialized {
		accounting.init()
	}
	formattedNumber := FormatNumberFloat64(value, accounting.Precision, accounting.Thousand, accounting.Decimal)
	return accounting.formatMoneyString(formattedNumber)
}

// FormatMoneyBigRat only supports *big.Rat value. It is faster than FormatMoney,
// because it does not do any runtime type evaluation.
func (accounting *Accounting) FormatMoneyBigRat(value *big.Rat) string {
	if !accounting.isInitialized {
		accounting.init()
	}
	formattedNumber := FormatNumberBigRat(value, accounting.Precision, accounting.Thousand, accounting.Decimal)
	return accounting.formatMoneyString(formattedNumber)
}

// FormatMoneyBigDecimal only supports *apd.Decimal value. It is faster than FormatMoney,
// because it does not do any runtime type evaluation.
func (accounting *Accounting) FormatMoneyBigDecimal(value *apd.Decimal) string {
	if !accounting.isInitialized {
		accounting.init()
	}
	formattedNumber := FormatNumberBigDecimal(value, accounting.Precision, accounting.Thousand, accounting.Decimal)
	return accounting.formatMoneyString(formattedNumber)
}

// FormatMoneyDecimal only supports decimal.Decimal value. It is faster than FormatMoney,
// because it does not do any runtime type evaluation.
func (accounting *Accounting) FormatMoneyDecimal(value decimal.Decimal) string {
	accounting.init()
	formattedNumber := FormatNumberDecimal(value, accounting.Precision, accounting.Thousand, accounting.Decimal)
	return accounting.formatMoneyString(formattedNumber)
}
