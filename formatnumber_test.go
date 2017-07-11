package accounting

import (
	"math/big"
	"testing"

	"github.com/cockroachdb/apd"
)

func TestFormatNumber(t *testing.T) {
	AssertEqual(t, FormatNumber(123456789.213123, 3, ",", "."), "123,456,789.213")
	AssertEqual(t, FormatNumber(123456789.213123, 3, ".", ","), "123.456.789,213")
	AssertEqual(t, FormatNumber(-12345.123123, 5, ",", "."), "-12,345.12312")
	AssertEqual(t, FormatNumber(-1234.123123, 5, ",", "."), "-1,234.12312")
	AssertEqual(t, FormatNumber(-123.123123, 5, ",", "."), "-123.12312")
	AssertEqual(t, FormatNumber(-12.123123, 5, ",", "."), "-12.12312")
	AssertEqual(t, FormatNumber(-1.123123, 5, ",", "."), "-1.12312")
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
	AssertEqual(t, FormatNumber(uint(1000000), 3, ",", "."), "1,000,000.000")

	AssertEqual(t, FormatNumber(big.NewRat(77777777, 3), 3, ",", "."), "25,925,925.667")
	AssertEqual(t, FormatNumber(big.NewRat(-77777777, 3), 3, ",", "."), "-25,925,925.667")
	AssertEqual(t, FormatNumber(big.NewRat(-7777777, 3), 3, ",", "."), "-2,592,592.333")
	AssertEqual(t, FormatNumber(big.NewRat(-777776, 3), 3, ",", "."), "-259,258.667")

	AssertEqual(t, FormatNumber(apd.New(123456789213123, -6), 3, ",", "."), "123,456,789.213")
	AssertEqual(t, FormatNumber(apd.New(-12345123123, -6), 5, ",", "."), "-12,345.12312")
	AssertEqual(t, FormatNumber(apd.New(-1234123123, -6), 5, ",", "."), "-1,234.12312")
	AssertEqual(t, FormatNumber(apd.New(-123123123, -6), 5, ",", "."), "-123.12312")
	AssertEqual(t, FormatNumber(apd.New(-12123123, -6), 5, ",", "."), "-12.12312")
	AssertEqual(t, FormatNumber(apd.New(-1123123, -6), 5, ",", "."), "-1.12312")

	func() {
		defer func() {
			recover()
		}()
		FormatNumber(false, 3, ",", ".") // panic: Unsupported type - bool
	}()
	func() {
		defer func() {
			recover()
		}()
		FormatNumber(big.NewInt(1), 3, ",", ".") // panic: Unsupported type - *big.Int
	}()
}

func TestFormatNumberInt(t *testing.T) {
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
	AssertEqual(t, FormatNumberFloat64(123456789.213123, 3, ",", "."), "123,456,789.213")
	AssertEqual(t, FormatNumberFloat64(-12345.123123, 5, ",", "."), "-12,345.12312")
	AssertEqual(t, FormatNumberFloat64(-1234.123123, 5, ",", "."), "-1,234.12312")
	AssertEqual(t, FormatNumberFloat64(-123.123123, 5, ",", "."), "-123.12312")
	AssertEqual(t, FormatNumberFloat64(-12.123123, 5, ",", "."), "-12.12312")
	AssertEqual(t, FormatNumberFloat64(-1.123123, 5, ",", "."), "-1.12312")
}

func TestFormatNumberBigRat(t *testing.T) {
	AssertEqual(t, FormatNumberBigRat(big.NewRat(77777777, 3), 3, ",", "."), "25,925,925.667")
	AssertEqual(t, FormatNumberBigRat(big.NewRat(-77777777, 3), 3, ",", "."), "-25,925,925.667")
	AssertEqual(t, FormatNumberBigRat(big.NewRat(-7777777, 3), 3, ",", "."), "-2,592,592.333")
	AssertEqual(t, FormatNumberBigRat(big.NewRat(-777776, 3), 3, ",", "."), "-259,258.667")
}

func TestFormatNumberBigDecimal(t *testing.T) {
	AssertEqual(t, FormatNumberBigDecimal(apd.New(123456789213123, -6), 3, ",", "."), "123,456,789.213")
	AssertEqual(t, FormatNumberBigDecimal(apd.New(-12345123123, -6), 5, ",", "."), "-12,345.12312")
	AssertEqual(t, FormatNumberBigDecimal(apd.New(-1234123123, -6), 5, ",", "."), "-1,234.12312")
	AssertEqual(t, FormatNumberBigDecimal(apd.New(-123123123, -6), 5, ",", "."), "-123.12312")
	AssertEqual(t, FormatNumberBigDecimal(apd.New(-12123123, -6), 5, ",", "."), "-12.12312")
	AssertEqual(t, FormatNumberBigDecimal(apd.New(-1123123, -6), 5, ",", "."), "-1.12312")
}
