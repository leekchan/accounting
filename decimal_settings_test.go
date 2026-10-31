package accounting

import (
	"github.com/shopspring/decimal"
	"testing"
)

func TestFormatMoneyDecimalPreservesExplicitSettings(t *testing.T) {
	for _, tc := range []struct{ name, value, want string }{
		{"positive", "1234.5", "$123450"},
		{"negative", "-1234.5", "($123450)"},
		{"zero", "0", "zero"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			ac := NewAccounting("$", 2, "", "", "%s%v", "(%s%v)", "zero")
			value, err := decimal.NewFromString(tc.value)
			if err != nil {
				t.Fatal(err)
			}
			if got := ac.FormatMoneyDecimal(value); got != tc.want {
				t.Errorf("got %q, want %q", got, tc.want)
			}
			if ac.Thousand != "" || ac.Decimal != "" {
				t.Errorf("formatting changed explicit separators: %#v", ac)
			}
		})
	}
	t.Run("literal defaults", func(t *testing.T) {
		ac := Accounting{Symbol: "$", Precision: 2}
		if got := ac.FormatMoneyDecimal(decimal.NewFromFloat(1234.5)); got != "$1,234.50" {
			t.Errorf("got %q", got)
		}
	})
}
