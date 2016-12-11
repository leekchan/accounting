// +build 386

package accounting

import (
	"math"
	"testing"
)

func TestFormatNumber32Bit(t *testing.T) {
	AssertEqual(t, FormatNumber(math.MaxInt32, 10, ",", "."), "2,147,483,647.0000000000")
	AssertEqual(t, FormatNumber(math.MinInt32, 10, ",", "."), "-2,147,483,648.0000000000")
}
