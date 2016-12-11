// +build amd64

package accounting

import (
	"math"
	"testing"
)

func TestFormatNumber64Bit(t *testing.T) {
	AssertEqual(t, FormatNumber(math.MaxInt64, 10, ",", "."), "9,223,372,036,854,775,807.0000000000")
	AssertEqual(t, FormatNumber(math.MinInt64, 10, ",", "."), "-9,223,372,036,854,775,808.0000000000")
}

func TestFormatNumberInt64Bit(t *testing.T) {
	AssertEqual(t, FormatNumberInt(math.MaxInt64, 10, ",", "."), "9,223,372,036,854,775,807.0000000000")
	AssertEqual(t, FormatNumberInt(math.MinInt64+1, 10, ",", "."), "-9,223,372,036,854,775,807.0000000000")
	AssertEqual(t, FormatNumberInt(math.MinInt64, 10, ",", "."), "-9,223,372,036,854,775,808.0000000000")
}
