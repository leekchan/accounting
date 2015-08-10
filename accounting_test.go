package accounting

import (
	"testing"
)

func AssertEqual(t *testing.T, x, y string) {
	if x != y {
		t.Error("Expected ", y, ", got ", x)
	}
}

func TestFormatMoney(t *testing.T) {
	accounting := Accounting{Symbol: "$", Precision: 2}
	AssertEqual(t, accounting.FormatMoney(123456789.213123), "$123,456,789.21")
	AssertEqual(t, accounting.FormatMoney(12345678), "$12,345,678.00")

	accounting = Accounting{Symbol: "€", Precision: 2, Thousand: ".", Decimal: ","}
	AssertEqual(t, accounting.FormatMoney(4999.99), "€4.999,99")
	AssertEqual(t, accounting.FormatMoney(-4999.99), "€-4.999,99")

	accounting = Accounting{Symbol: "£ ", Precision: 0}
	AssertEqual(t, accounting.FormatMoney(-500000), "£ -500,000")

	accounting = Accounting{Symbol: "GBP", Precision: 0,
		Format: "%s %v", FormatNegative: "%s (%v)", FormatZero: "%s --"}
	AssertEqual(t, accounting.FormatMoney(1000000), "GBP 1,000,000")
	AssertEqual(t, accounting.FormatMoney(-5000), "GBP (5,000)")
	AssertEqual(t, accounting.FormatMoney(0), "GBP --")
}

func TestFormatMoneyInt(t *testing.T) {
	accounting := Accounting{Symbol: "$", Precision: 2}
	AssertEqual(t, accounting.FormatMoneyInt(12345678), "$12,345,678.00")

	accounting = Accounting{Symbol: "€", Precision: 2, Thousand: ".", Decimal: ","}
	AssertEqual(t, accounting.FormatMoneyInt(4999), "€4.999,00")
	AssertEqual(t, accounting.FormatMoneyInt(-4999), "€-4.999,00")

	accounting = Accounting{Symbol: "£ ", Precision: 0}
	AssertEqual(t, accounting.FormatMoneyInt(-500000), "£ -500,000")

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
	AssertEqual(t, accounting.FormatMoneyFloat64(-4999.99), "€-4.999,99")

	accounting = Accounting{Symbol: "£ ", Precision: 0}
	AssertEqual(t, accounting.FormatMoneyFloat64(-500000.0), "£ -500,000")

	accounting = Accounting{Symbol: "GBP", Precision: 0,
		Format: "%s %v", FormatNegative: "%s (%v)", FormatZero: "%s --"}
	AssertEqual(t, accounting.FormatMoneyFloat64(1000000.0), "GBP 1,000,000")
	AssertEqual(t, accounting.FormatMoneyFloat64(-5000.0), "GBP (5,000)")
	AssertEqual(t, accounting.FormatMoneyFloat64(0.0), "GBP --")
}
