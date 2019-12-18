package accounting

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// UnformatNumber takes a string of the number to strip currency info on
// and precision for decimals.
// It pulls the currency descripter from the LocaleInfo map and uses it to return an unformatted value
// based on thous sep and decimal sep
func UnformatNumber(n string, precision int, currency string) string {
	var lc Locale
	currency = strings.ToUpper(currency)

	if val, ok := LocaleInfo[currency]; ok {
		lc = val
	} else {
		panic("No Locale Info Found")
	}

	r := regexp.MustCompile(`[^0-9-., ]`) // Remove anything thats not a space, comma, or decimal
	num := r.ReplaceAllString(n, "${1}")

	r = regexp.MustCompile(fmt.Sprintf("\\%v", lc.ThouSep)) // Strip out thousands seperator, whatever it is
	num = r.ReplaceAllString(num, "${1}")

	// Replace decimal seperator with a decimal at specified precision
	if lc.DecSep != "." {
		r = regexp.MustCompile(`\,`)
		num = r.ReplaceAllString(num, ".")
	}

	num = setPrecision(num, precision)
	return num
}

func setPrecision(num string, precision int) string {
	p := fmt.Sprintf("%%.%vf", precision)
	num = strings.Trim(num, " ")
	v, _ := strconv.ParseFloat(num, 64)
	return fmt.Sprintf(p, v)
}
