package accounting

import (
	"math/big"
	"strings"

	"github.com/cockroachdb/apd"
)

type Accounting struct {
	Symbol         string // currency symbol (required)
	Precision      int    // currency precision (decimal places) (optional / default: 0)
	Thousand       string // thousand separator (optional / default: ,)
	Decimal        string // decimal separator (optional / default: .)
	Format         string // simple format string allows control of symbol position (%v = value, %s = symbol) (default: %s%v)
	FormatNegative string // format string for negative values (optional / default: strings.Replace(strings.Replace(accounting.Format, "-", "", -1), "%v", "-%v", -1))
	FormatZero     string // format string for zero values (optional / default: Format)
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

	zero := "0"
	if accounting.Precision > 0 {
		zero += "." + strings.Repeat("0", accounting.Precision)
	}

	if formattedNumber[0] == '-' {
		format = accounting.FormatNegative
		formattedNumber = formattedNumber[1:]
	} else if formattedNumber == zero {
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
	accounting.init()
	formattedNumber := FormatNumber(value, accounting.Precision, accounting.Thousand, accounting.Decimal)
	return accounting.formatMoneyString(formattedNumber)
}

// FormatMoneyInt only supports int type value. It is faster than FormatMoney,
// because it does not do any runtime type evaluation.
func (accounting *Accounting) FormatMoneyInt(value int) string {
	accounting.init()
	formattedNumber := FormatNumberInt(value, accounting.Precision, accounting.Thousand, accounting.Decimal)
	return accounting.formatMoneyString(formattedNumber)
}

// FormatMoneyFloat64 only supports float64 value. It is faster than FormatMoney,
// because it does not do any runtime type evaluation.
// (Caution: Please do not use float64 to count money.
// Floats can have errors when you perform operations on them.
// Using big.Rat is highly recommended.)
func (accounting *Accounting) FormatMoneyFloat64(value float64) string {
	accounting.init()
	formattedNumber := FormatNumberFloat64(value, accounting.Precision, accounting.Thousand, accounting.Decimal)
	return accounting.formatMoneyString(formattedNumber)
}

// FormatMoneyBigRat only supports *big.Rat value. It is faster than FormatMoney,
// because it does not do any runtime type evaluation.
func (accounting *Accounting) FormatMoneyBigRat(value *big.Rat) string {
	accounting.init()
	formattedNumber := FormatNumberBigRat(value, accounting.Precision, accounting.Thousand, accounting.Decimal)
	return accounting.formatMoneyString(formattedNumber)
}

// FormatMoneyBigDecimal only supports *apd.Decimal value. It is faster than FormatMoney,
// because it does not do any runtime type evaluation.
func (accounting *Accounting) FormatMoneyBigDecimal(value *apd.Decimal) string {
	accounting.init()
	formattedNumber := FormatNumberBigDecimal(value, accounting.Precision, accounting.Thousand, accounting.Decimal)
	return accounting.formatMoneyString(formattedNumber)
}
