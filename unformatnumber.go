package accounting

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// ErrUnknownCurrency is returned by UnformatNumberStrict when the currency
// is not found in LocaleInfo.
var ErrUnknownCurrency = errors.New("accounting: unknown currency")

// ErrInvalidNumber is returned by UnformatNumberStrict when the input does
// not contain a valid number once the currency formatting is stripped.
var ErrInvalidNumber = errors.New("accounting: invalid number")

var nonNumericRegexp = regexp.MustCompile(`[^0-9-., ]`) // Remove anything thats not a digit, minus, space, comma, or decimal

// UnformatNumber takes a string of the number to strip currency info on
// and precision for decimals.
// It pulls the currency descripter from the LocaleInfo map and uses it to return an unformatted value
// based on thous sep and decimal sep.
// UnformatNumber panics if the currency is not found in LocaleInfo, and
// returns zero (e.g. "0.00") if the input does not contain a valid number.
// Use UnformatNumberStrict to get an error in those cases instead.
func UnformatNumber(n string, precision int, currency string) string {
	lc, ok := LocaleInfo[strings.ToUpper(currency)]
	if !ok {
		panic("No Locale Info Found")
	}

	v, _ := strconv.ParseFloat(stripFormatting(n, lc), 64)
	return setPrecision(v, precision)
}

// UnformatNumberStrict is like UnformatNumber, but it returns an error
// instead of panicking for an unknown currency (ErrUnknownCurrency), and
// instead of returning zero for input that does not contain a valid number,
// or ±Inf for a number out of the range of float64 (ErrInvalidNumber).
// On success it returns the same result as UnformatNumber.
func UnformatNumberStrict(n string, precision int, currency string) (string, error) {
	lc, ok := LocaleInfo[strings.ToUpper(currency)]
	if !ok {
		return "", fmt.Errorf("%w: %q", ErrUnknownCurrency, currency)
	}

	v, err := strconv.ParseFloat(stripFormatting(n, lc), 64)
	if err != nil {
		return "", fmt.Errorf("%w: %q", ErrInvalidNumber, n)
	}
	return setPrecision(v, precision), nil
}

// stripFormatting removes the currency formatting from n and returns the
// number with a point as the decimal separator, ready for strconv.ParseFloat.
func stripFormatting(n string, lc Locale) string {
	num := nonNumericRegexp.ReplaceAllString(n, "")

	// Strip out thousands seperator, whatever it is
	if lc.ThouSep != "" {
		num = strings.Replace(num, lc.ThouSep, "", -1)
	}

	// Replace the locale decimal separator with a point.
	// An empty separator is not a comma. CLP groups with "." and has no
	// decimal separator, so "1,234" is not the number 1.
	if lc.DecSep != "" && lc.DecSep != "." {
		num = strings.Replace(num, lc.DecSep, ".", -1)
	}

	return strings.Trim(num, " ")
}

func setPrecision(v float64, precision int) string {
	p := fmt.Sprintf("%%.%vf", precision)
	return fmt.Sprintf(p, v)
}
