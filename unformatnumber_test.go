package accounting

import (
	"testing"
)

func TestUnformatNumberCommaDecimal(t *testing.T) {
	AssertEqual(t, UnformatNumber("$4,500.23", 2, "USD"), "4500.23")
}

func TestUnformatNumberDecimalComma(t *testing.T) {
	AssertEqual(t, UnformatNumber("EUR 45.000,33", 2, "eur"), "45000.33")

	func() {
		defer func() {
			recover()
		}()
		UnformatNumber("$45,567.10", 2, "zzz")
	}()
}
