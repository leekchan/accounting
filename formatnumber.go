package accounting

import (
	"bytes"
	"fmt"
	"math"
	"reflect"
	"strings"
)

func formatNumberString(x string, precision int, thousand string, decimal string) string {
	lastIndex := strings.Index(x, ".") - 1

	if lastIndex < 0 {
		lastIndex = len(x) - 1
	}

	var buffer []byte
	var strBuffer bytes.Buffer

	j := 0
	for i := lastIndex; i >= 0; i-- {
		j++
		buffer = append(buffer, x[i])

		if j == 3 && i > 0 && !(i == 1 && x[0] == '-') {
			buffer = append(buffer, ',')
			j = 0
		}
	}

	for i := len(buffer) - 1; i >= 0; i-- {
		strBuffer.WriteByte(buffer[i])
	}
	result := strBuffer.String()

	if thousand != "," {
		result = strings.Replace(result, ",", thousand, -1)
	}

	extra := x[lastIndex+1:]
	if decimal != "." {
		extra = strings.Replace(extra, ".", decimal, 1)
	}

	return result + extra
}

func FormatNumber(value interface{}, precision int, thousand string, decimal string) string {
	v := reflect.ValueOf(value)

	var x string
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		x = fmt.Sprintf("%d", v.Int())
		if precision > 0 {
			x += "." + strings.Repeat("0", precision)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		x = fmt.Sprintf("%d", v.Uint())
		if precision > 0 {
			x += "." + strings.Repeat("0", precision)
		}
	case reflect.Float32, reflect.Float64:
		x = fmt.Sprintf(fmt.Sprintf("%%.%df", precision), v.Float())
	default:
		return ""
	}

	return formatNumberString(x, precision, thousand, decimal)
}

func FormatNumberInt(x int, precision int, thousand string, decimal string) string {
	var result string
	var minus bool

	if x == math.MinInt64 {
		return FormatNumber(x, precision, thousand, decimal)
	}

	if x < 0 {
		minus = true
		x *= -1
	}

	for x >= 1000 {
		result = fmt.Sprintf("%s%03d%s", thousand, x%1000, result)
		x /= 1000
	}
	result = fmt.Sprintf("%d%s", x, result)

	if minus {
		result = "-" + result
	}

	if precision > 0 {
		result += decimal + strings.Repeat("0", precision)
	}

	return result
}

func FormatNumberFloat64(x float64, precision int, thousand string, decimal string) string {
	return formatNumberString(fmt.Sprintf(fmt.Sprintf("%%.%df", precision), x), precision, thousand, decimal)
}
