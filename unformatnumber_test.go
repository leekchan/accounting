package accounting

import (
	"testing"
)

func TestUnformatNumberCommaDecimal(t *testing.T) {
	n := UnformatNumber("$4,500.23", 2, "USD")
	AssertEqual(t, "4500.23", n)
}

func TestUnformatNumberDecimalComma(t *testing.T) {
	n := UnformatNumber("EUR 45.000,33", 2, "eur")
	AssertEqual(t, "45000.33", n)
}
