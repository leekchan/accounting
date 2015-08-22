// +build go1.5

package accounting

import (
	"math/big"
)

// FormatNumberBigFloat only supports *big.Float value.
// It is faster than FormatNumber, because it does not do any runtime type evaluation.
func FormatNumberBigFloat(x *big.Float, precision int, thousand string, decimal string) string {
	return formatNumberString(x.Text('f', precision), precision, thousand, decimal)
}
