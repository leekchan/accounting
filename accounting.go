package accounting

import (
	"strings"
)

type Accounting struct {
	Symbol         string
	Precision      int
	Thousand       string
	Decimal        string
	Format         string
	FormatNegative string
	FormatZero     string
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
		accounting.FormatNegative = strings.Replace(strings.Replace(accounting.Format, "-", "", -1), "%v", "-%v", -1)
	}

	if accounting.FormatZero == "" {
		accounting.FormatZero = "0"
	}
}

func (accounting *Accounting) formatMoneyString(formattedNumber string) string {
	var format string

	if formattedNumber[0] == '-' {
		format = accounting.FormatNegative
		formattedNumber = formattedNumber[1:]
	} else if formattedNumber == "0" {
		format = accounting.FormatZero
	} else {
		format = accounting.Format
	}

	result := strings.Replace(format, "%s", accounting.Symbol, -1)
	result = strings.Replace(result, "%v", formattedNumber, -1)

	return result
}

func (accounting *Accounting) FormatMoney(value interface{}) string {
	accounting.init()
	formattedNumber := FormatNumber(value, accounting.Precision, accounting.Thousand, accounting.Decimal)
	return accounting.formatMoneyString(formattedNumber)
}

func (accounting *Accounting) FormatMoneyInt(value int) string {
	accounting.init()
	formattedNumber := FormatNumberInt(value, accounting.Precision, accounting.Thousand, accounting.Decimal)
	return accounting.formatMoneyString(formattedNumber)
}

func (accounting *Accounting) FormatMoneyFloat64(value float64) string {
	accounting.init()
	formattedNumber := FormatNumberFloat64(value, accounting.Precision, accounting.Thousand, accounting.Decimal)
	return accounting.formatMoneyString(formattedNumber)
}
