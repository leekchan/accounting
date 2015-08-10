package accounting

import (
	"math"
	"testing"
)

func TestFormatNumber(t *testing.T) {
	AssertEqual(t, FormatNumber(123456789.213123, 3, ",", "."), "123,456,789.213")
	AssertEqual(t, FormatNumber(123456789.213123, 3, ".", ","), "123.456.789,213")
	AssertEqual(t, FormatNumber(-12345.123123, 5, ",", "."), "-12,345.12312")
	AssertEqual(t, FormatNumber(-1234.123123, 5, ",", "."), "-1,234.12312")
	AssertEqual(t, FormatNumber(-123.123123, 5, ",", "."), "-123.12312")
	AssertEqual(t, FormatNumber(-12.123123, 5, ",", "."), "-12.12312")
	AssertEqual(t, FormatNumber(-1.123123, 5, ",", "."), "-1.12312")
	AssertEqual(t, FormatNumber(math.MaxInt64, 10, ",", "."), "9,223,372,036,854,775,807.0000000000")
	AssertEqual(t, FormatNumber(math.MinInt64, 10, ",", "."), "-9,223,372,036,854,775,808.0000000000")
	AssertEqual(t, FormatNumber(-1, 3, ",", "."), "-1.000")
	AssertEqual(t, FormatNumber(-10, 3, ",", "."), "-10.000")
	AssertEqual(t, FormatNumber(-100, 3, ",", "."), "-100.000")
	AssertEqual(t, FormatNumber(-1000, 3, ",", "."), "-1,000.000")
	AssertEqual(t, FormatNumber(-10000, 3, ",", "."), "-10,000.000")
	AssertEqual(t, FormatNumber(-100000, 3, ",", "."), "-100,000.000")
	AssertEqual(t, FormatNumber(-1000000, 3, ",", "."), "-1,000,000.000")
	AssertEqual(t, FormatNumber(1, 3, ",", "."), "1.000")
	AssertEqual(t, FormatNumber(10, 3, ",", "."), "10.000")
	AssertEqual(t, FormatNumber(100, 3, ",", "."), "100.000")
	AssertEqual(t, FormatNumber(1000, 3, ",", "."), "1,000.000")
	AssertEqual(t, FormatNumber(10000, 3, ",", "."), "10,000.000")
	AssertEqual(t, FormatNumber(100000, 3, ",", "."), "100,000.000")
	AssertEqual(t, FormatNumber(1000000, 3, ",", "."), "1,000,000.000")
	AssertEqual(t, FormatNumber(1000000, 10, " ", "."), "1 000 000.0000000000")
	AssertEqual(t, FormatNumber(1000000, 10, "   ", "."), "1   000   000.0000000000")
}

func TestFormatNumberInt(t *testing.T) {
	AssertEqual(t, FormatNumberInt(math.MaxInt64, 10, ",", "."), "9,223,372,036,854,775,807.0000000000")
	AssertEqual(t, FormatNumberInt(math.MinInt64+1, 10, ",", "."), "-9,223,372,036,854,775,807.0000000000")
	AssertEqual(t, FormatNumberInt(math.MinInt64, 10, ",", "."), "-9,223,372,036,854,775,808.0000000000")
	AssertEqual(t, FormatNumberInt(-1, 3, ",", "."), "-1.000")
	AssertEqual(t, FormatNumberInt(-10, 3, ",", "."), "-10.000")
	AssertEqual(t, FormatNumberInt(-100, 3, ",", "."), "-100.000")
	AssertEqual(t, FormatNumberInt(-1000, 3, ",", "."), "-1,000.000")
	AssertEqual(t, FormatNumberInt(-10000, 3, ",", "."), "-10,000.000")
	AssertEqual(t, FormatNumberInt(-100000, 3, ",", "."), "-100,000.000")
	AssertEqual(t, FormatNumberInt(-1000000, 3, ",", "."), "-1,000,000.000")
	AssertEqual(t, FormatNumberInt(1, 3, ",", "."), "1.000")
	AssertEqual(t, FormatNumberInt(10, 3, ",", "."), "10.000")
	AssertEqual(t, FormatNumberInt(100, 3, ",", "."), "100.000")
	AssertEqual(t, FormatNumberInt(1000, 3, ",", "."), "1,000.000")
	AssertEqual(t, FormatNumberInt(10000, 3, ",", "."), "10,000.000")
	AssertEqual(t, FormatNumberInt(100000, 3, ",", "."), "100,000.000")
	AssertEqual(t, FormatNumberInt(1000000, 3, ",", "."), "1,000,000.000")
	AssertEqual(t, FormatNumberInt(1000000, 10, " ", "."), "1 000 000.0000000000")
	AssertEqual(t, FormatNumberInt(1000000, 10, "   ", "."), "1   000   000.0000000000")
}

func TestFormatNumberFloat64(t *testing.T) {
	AssertEqual(t, FormatNumber(123456789.213123, 3, ",", "."), "123,456,789.213")
	AssertEqual(t, FormatNumber(-12345.123123, 5, ",", "."), "-12,345.12312")
	AssertEqual(t, FormatNumber(-1234.123123, 5, ",", "."), "-1,234.12312")
	AssertEqual(t, FormatNumber(-123.123123, 5, ",", "."), "-123.12312")
	AssertEqual(t, FormatNumber(-12.123123, 5, ",", "."), "-12.12312")
	AssertEqual(t, FormatNumber(-1.123123, 5, ",", "."), "-1.12312")
}
