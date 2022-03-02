package accounting

import (
	"fmt"
	"math/big"
	"runtime"
	"testing"

	"github.com/cockroachdb/apd"
	"github.com/shopspring/decimal"
)

func AssertEqual(t *testing.T, x, y string) {
	if x != y {
		t.Error("Expected ", y, ", got ", x)
	}
}

func init() {
	fmt.Printf("version: %s", runtime.Version())
}

func TestAccounting_SetFormat(t *testing.T) {
	accounting := DefaultAccounting("$", 2)
	accounting.SetFormat("%v %s")
	AssertEqual(t, accounting.FormatMoney(123456789.213123), "123,456,789.21 $")
}

func TestAccounting_SetFormatZero(t *testing.T) {
	accounting := DefaultAccounting("$", 2)
	accounting.SetFormatZero("%s --")
	AssertEqual(t, accounting.FormatMoney(0), "$ --")
}

func TestAccounting_SetFormatNegative(t *testing.T) {
	accounting := DefaultAccounting("$", 2)
	accounting.SetFormatNegative("%s(%v)")
	AssertEqual(t, accounting.FormatMoney(-123456789.213123), "$(123,456,789.21)")
}

func TestAccounting_SetThousandSeparator(t *testing.T) {
	accounting := DefaultAccounting("$", 2)
	accounting.SetThousandSeparator("'")
	AssertEqual(t, accounting.FormatMoney(123456789.213123), "$123'456'789.21")
}

func TestAccounting_SetDecimalSeparator(t *testing.T) {
	accounting := DefaultAccounting("$", 2)
	accounting.SetDecimalSeparator("'")
	AssertEqual(t, accounting.FormatMoney(123456789.213123), "$123,456,789'21")
}

func TestFormatMoney(t *testing.T) {
	accounting := DefaultAccounting("$", 2)
	AssertEqual(t, accounting.FormatMoney(123456789.213123), "$123,456,789.21")
	AssertEqual(t, accounting.FormatMoney(12345678), "$12,345,678.00")
	AssertEqual(t, accounting.FormatMoney(-12345678), "-$12,345,678.00")
	AssertEqual(t, accounting.FormatMoney(0), "$0.00")
	AssertEqual(t, accounting.FormatMoney(big.NewRat(77777777, 3)), "$25,925,925.67")
	AssertEqual(t, accounting.FormatMoney(big.NewRat(-77777777, 3)), "-$25,925,925.67")
	AssertEqual(t, accounting.FormatMoney(apd.New(499999, -2)), "$4,999.99")
	AssertEqual(t, accounting.FormatMoney(apd.New(500000, 0)), "$500,000.00")
	AssertEqual(t, accounting.FormatMoney(decimal.New(499999, -2)), "$4,999.99")
	AssertEqual(t, accounting.FormatMoney(decimal.New(500000, 0)), "$500,000.00")

	accounting = NewAccounting("$", 0, ",", ".", "%s %v", "-%s %v", "%s %v")
	AssertEqual(t, accounting.FormatMoney(123456789.213123), "$ 123,456,789")
	AssertEqual(t, accounting.FormatMoney(12345678), "$ 12,345,678")
	AssertEqual(t, accounting.FormatMoney(-12345678), "-$ 12,345,678")
	AssertEqual(t, accounting.FormatMoney(0), "$ 0")

	accounting2 := Accounting{Symbol: "€", Precision: 2, Thousand: ".", Decimal: ","}
	AssertEqual(t, accounting2.FormatMoney(4999.99), "€4.999,99")

	accounting2 = Accounting{Symbol: "£ ", Precision: 0}
	AssertEqual(t, accounting2.FormatMoney(500000), "£ 500,000")

	accounting2 = Accounting{Symbol: "GBP", Precision: 0,
		Format: "%s %v", FormatNegative: "%s (%v)", FormatZero: "%s --"}
	AssertEqual(t, accounting2.FormatMoney(1000000), "GBP 1,000,000")
	AssertEqual(t, accounting2.FormatMoney(-5000), "GBP (5,000)")
	AssertEqual(t, accounting2.FormatMoney(0), "GBP --")

	accounting2 = Accounting{Symbol: "GBP", Precision: 2,
		Format: "%s %v", FormatNegative: "%s (%v)", FormatZero: "%s --"}
	AssertEqual(t, accounting2.FormatMoney(1000000), "GBP 1,000,000.00")
	AssertEqual(t, accounting2.FormatMoney(-5000), "GBP (5,000.00)")
	AssertEqual(t, accounting2.FormatMoney(0), "GBP --")

	accounting2 = Accounting{Symbol: "€", Precision: 2,
		Decimal: ",", FormatZero: "0.-"}
	AssertEqual(t, accounting2.FormatMoney(0), "0.-")
}

func TestFormatMoneyInt(t *testing.T) {
	accounting := Accounting{Symbol: "$", Precision: 2}
	AssertEqual(t, accounting.FormatMoneyInt(12345678), "$12,345,678.00")

	accounting = Accounting{Symbol: "€", Precision: 2, Thousand: ".", Decimal: ","}
	AssertEqual(t, accounting.FormatMoneyInt(4999), "€4.999,00")

	accounting = Accounting{Symbol: "£ ", Precision: 0}
	AssertEqual(t, accounting.FormatMoneyInt(500000), "£ 500,000")

	accounting = Accounting{Symbol: "GBP", Precision: 0,
		Format: "%s %v", FormatNegative: "%s (%v)", FormatZero: "%s --"}
	AssertEqual(t, accounting.FormatMoneyInt(1000000), "GBP 1,000,000")
	AssertEqual(t, accounting.FormatMoneyInt(-5000), "GBP (5,000)")
	AssertEqual(t, accounting.FormatMoneyInt(0), "GBP --")
}

func TestFormatMoneyFloat64(t *testing.T) {
	accounting := Accounting{Symbol: "$", Precision: 2}
	AssertEqual(t, accounting.FormatMoneyFloat64(123456789.213123), "$123,456,789.21")

	accounting = Accounting{Symbol: "€", Precision: 2, Thousand: ".", Decimal: ","}
	AssertEqual(t, accounting.FormatMoneyFloat64(4999.99), "€4.999,99")

	accounting = Accounting{Symbol: "£ ", Precision: 0}
	AssertEqual(t, accounting.FormatMoneyFloat64(500000.0), "£ 500,000")

	accounting = Accounting{Symbol: "GBP", Precision: 0,
		Format: "%s %v", FormatNegative: "%s (%v)", FormatZero: "%s --"}
	AssertEqual(t, accounting.FormatMoneyFloat64(1000000.0), "GBP 1,000,000")
	AssertEqual(t, accounting.FormatMoneyFloat64(-5000.0), "GBP (5,000)")
	AssertEqual(t, accounting.FormatMoneyFloat64(0.0), "GBP --")
}

func TestFormatMoneyBigRat(t *testing.T) {
	accounting := Accounting{Symbol: "$", Precision: 2}
	AssertEqual(t, accounting.FormatMoneyBigRat(big.NewRat(77777777, 3)), "$25,925,925.67")
	AssertEqual(t, accounting.FormatMoneyBigRat(big.NewRat(-77777777, 3)), "-$25,925,925.67")

	accounting = Accounting{Symbol: "€", Precision: 2, Thousand: ".", Decimal: ","}
	AssertEqual(t, accounting.FormatMoneyBigRat(big.NewRat(77777777, 3)), "€25.925.925,67")

	accounting = Accounting{Symbol: "£ ", Precision: 0}
	AssertEqual(t, accounting.FormatMoneyBigRat(big.NewRat(77777777, 3)), "£ 25,925,926")

	accounting = Accounting{Symbol: "GBP", Precision: 0,
		Format: "%s %v", FormatNegative: "%s (%v)", FormatZero: "%s --"}
	AssertEqual(t, accounting.FormatMoneyBigRat(big.NewRat(77777777, 3)), "GBP 25,925,926")
	AssertEqual(t, accounting.FormatMoneyBigRat(big.NewRat(-77777777, 3)), "GBP (25,925,926)")
	AssertEqual(t, accounting.FormatMoneyBigRat(big.NewRat(0, 3)), "GBP --")
}

func TestFormatMoneyBigDecimal(t *testing.T) {
	accounting := Accounting{Symbol: "$", Precision: 2}
	AssertEqual(t, accounting.FormatMoneyBigDecimal(apd.New(123456789213123, -6)), "$123,456,789.21")

	accounting = Accounting{Symbol: "€", Precision: 2, Thousand: ".", Decimal: ","}
	AssertEqual(t, accounting.FormatMoneyBigDecimal(apd.New(499999, -2)), "€4.999,99")

	accounting = Accounting{Symbol: "£ ", Precision: 0}
	AssertEqual(t, accounting.FormatMoneyBigDecimal(apd.New(500000, 0)), "£ 500,000")

	accounting = Accounting{Symbol: "GBP", Precision: 0,
		Format: "%s %v", FormatNegative: "%s (%v)", FormatZero: "%s --"}
	AssertEqual(t, accounting.FormatMoneyBigDecimal(apd.New(1000000, 0)), "GBP 1,000,000")
	AssertEqual(t, accounting.FormatMoneyBigDecimal(apd.New(-5000, 0)), "GBP (5,000)")
	AssertEqual(t, accounting.FormatMoneyBigDecimal(apd.New(0, 0)), "GBP --")
}

func TestFormatMoneyDecimal(t *testing.T) {
	accounting := Accounting{Symbol: "$", Precision: 2}
	AssertEqual(t, accounting.FormatMoneyDecimal(decimal.New(123456789213123, -6)), "$123,456,789.21")

	accounting = Accounting{Symbol: "€", Precision: 2, Thousand: ".", Decimal: ","}
	AssertEqual(t, accounting.FormatMoneyDecimal(decimal.New(499999, -2)), "€4.999,99")

	accounting = Accounting{Symbol: "£ ", Precision: 0}
	AssertEqual(t, accounting.FormatMoneyDecimal(decimal.New(500000, 0)), "£ 500,000")

	accounting = Accounting{Symbol: "GBP", Precision: 0,
		Format: "%s %v", FormatNegative: "%s (%v)", FormatZero: "%s --"}
	AssertEqual(t, accounting.FormatMoneyDecimal(decimal.New(1000000, 0)), "GBP 1,000,000")
	AssertEqual(t, accounting.FormatMoneyDecimal(decimal.New(-5000, 0)), "GBP (5,000)")
	AssertEqual(t, accounting.FormatMoneyDecimal(decimal.New(0, 0)), "GBP --")
}
