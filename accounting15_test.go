// +build go1.5

package accounting

import (
	"math/big"
	"testing"
)

func TestFormatMoneyBigFloat(t *testing.T) {
	accounting := Accounting{Symbol: "$", Precision: 2}
	AssertEqual(t, accounting.FormatMoneyBigFloat(big.NewFloat(123456789.213123)), "$123,456,789.21")

	accounting = Accounting{Symbol: "€", Precision: 2, Thousand: ".", Decimal: ","}
	AssertEqual(t, accounting.FormatMoneyBigFloat(big.NewFloat(4999.99)), "€4.999,99")

	accounting = Accounting{Symbol: "£ ", Precision: 0}
	AssertEqual(t, accounting.FormatMoneyBigFloat(big.NewFloat(500000.0)), "£ 500,000")

	accounting = Accounting{Symbol: "GBP", Precision: 0,
		Format: "%s %v", FormatNegative: "%s (%v)", FormatZero: "%s --"}
	AssertEqual(t, accounting.FormatMoneyBigFloat(big.NewFloat(1000000.0)), "GBP 1,000,000")
	AssertEqual(t, accounting.FormatMoneyBigFloat(big.NewFloat(-5000.0)), "GBP (5,000)")
	AssertEqual(t, accounting.FormatMoneyBigFloat(big.NewFloat(0.0)), "GBP --")
}
