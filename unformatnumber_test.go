package accounting

import (
	"errors"
	"strings"
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

func TestUnformatNumberEmptyThousandSeparator(t *testing.T) {
	// ALL and AZN have no thousand separator in LocaleInfo.
	AssertEqual(t, UnformatNumber("Lek 1234.50", 2, "ALL"), "1234.50")
	AssertEqual(t, UnformatNumber("Lek 1234.50", 2, "AZN"), "1234.50")
}

func TestUnformatNumberInvalidNumberIsZero(t *testing.T) {
	// UnformatNumber cannot report an error, so it returns zero.
	// Existing callers rely on this.
	AssertEqual(t, UnformatNumber("", 2, "USD"), "0.00")
	AssertEqual(t, UnformatNumber(" ", 2, "USD"), "0.00")
	AssertEqual(t, UnformatNumber("abc", 2, "USD"), "0.00")
	AssertEqual(t, UnformatNumber("-", 2, "USD"), "0.00")
	AssertEqual(t, UnformatNumber("1.2.3", 2, "USD"), "0.00")
	AssertEqual(t, UnformatNumber("abc", 0, "USD"), "0")
}

func TestUnformatNumberUnknownCurrencyPanics(t *testing.T) {
	defer func() {
		if r := recover(); r != "No Locale Info Found" {
			t.Error("Expected panic No Locale Info Found, got ", r)
		}
	}()
	UnformatNumber("$45,567.10", 2, "zzz")
}

func AssertErrorIs(t *testing.T, err, target error) {
	if !errors.Is(err, target) {
		t.Error("Expected error ", target, ", got ", err)
	}
}

func TestUnformatNumberStrict(t *testing.T) {
	n, err := UnformatNumberStrict("$4,500.23", 2, "USD")
	AssertEqual(t, n, "4500.23")
	AssertErrorIs(t, err, nil)

	n, err = UnformatNumberStrict("-$4,500.23", 2, "USD")
	AssertEqual(t, n, "-4500.23")
	AssertErrorIs(t, err, nil)

	n, err = UnformatNumberStrict("$45,000.50", 0, "USD")
	AssertEqual(t, n, "45000")
	AssertErrorIs(t, err, nil)

	n, err = UnformatNumberStrict("EUR 45.000,33", 2, "eur")
	AssertEqual(t, n, "45000.33")
	AssertErrorIs(t, err, nil)

	n, err = UnformatNumberStrict("EUR 12.500,3474", 3, "EUR")
	AssertEqual(t, n, "12500.347")
	AssertErrorIs(t, err, nil)

	n, err = UnformatNumberStrict("Lek 1234.50", 2, "ALL")
	AssertEqual(t, n, "1234.50")
	AssertErrorIs(t, err, nil)
}

func TestUnformatNumberStrictInvalidNumber(t *testing.T) {
	for _, s := range []string{"", " ", "abc", "$", "-", "1.2.3", "1-2", strings.Repeat("9", 400)} {
		n, err := UnformatNumberStrict(s, 2, "USD")
		AssertEqual(t, n, "")
		AssertErrorIs(t, err, ErrInvalidNumber)
	}
}

func TestUnformatNumberStrictUnknownCurrency(t *testing.T) {
	n, err := UnformatNumberStrict("$45,567.10", 2, "zzz")
	AssertEqual(t, n, "")
	AssertErrorIs(t, err, ErrUnknownCurrency)
}
