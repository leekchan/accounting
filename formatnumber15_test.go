// +build go1.5

package accounting

import (
	"math/big"
	"testing"
)

func TestFormatNumberBigFloat(t *testing.T) {
	AssertEqual(t, FormatNumberBigFloat(big.NewFloat(123456789.213123), 3, ",", "."), "123,456,789.213")
	AssertEqual(t, FormatNumberBigFloat(big.NewFloat(-12345.123123), 5, ",", "."), "-12,345.12312")
	AssertEqual(t, FormatNumberBigFloat(big.NewFloat(-1234.123123), 5, ",", "."), "-1,234.12312")
	AssertEqual(t, FormatNumberBigFloat(big.NewFloat(-123.123123), 5, ",", "."), "-123.12312")
	AssertEqual(t, FormatNumberBigFloat(big.NewFloat(-12.123123), 5, ",", "."), "-12.12312")
	AssertEqual(t, FormatNumberBigFloat(big.NewFloat(-1.123123), 5, ",", "."), "-1.12312")
}
