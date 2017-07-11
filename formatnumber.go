package accounting

import (
	"bytes"
	"fmt"
	"math/big"
	"reflect"
	"strings"

	"github.com/cockroachdb/apd"
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

// FormatNumber is a base function of the library which formats a number with custom precision and separators.
// FormatNumber supports various types of value by runtime reflection.
// If you don't need runtime type evaluation, please refer to FormatNumberInt or FormatNumberFloat64.
// (supported value types : int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, *big.Rat)
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
	case reflect.Ptr:
		switch v.Type().String() {
		case "*big.Rat":
			x = value.(*big.Rat).FloatString(precision)
		case "*apd.Decimal":
			v := value.(*apd.Decimal)
			d := apd.New(0, 0)
			apd.BaseContext.WithPrecision(uint32(v.NumDigits())+uint32(precision)).Quantize(d, v, int32(-precision))
			x = d.Text('f')
		default:
			panic("Unsupported type - " + v.Type().String())
		}
	default:
		panic("Unsupported type - " + v.Kind().String())
	}

	return formatNumberString(x, precision, thousand, decimal)
}

// FormatNumberInt only supports int value. It is faster than FormatNumber,
// because it does not do any runtime type evaluation.
func FormatNumberInt(x int, precision int, thousand string, decimal string) string {
	var result string
	var minus bool

	if x < 0 {
		if x*-1 < 0 {
			return FormatNumber(x, precision, thousand, decimal)
		}

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

// FormatNumberFloat64 only supports float64 value.
// It is faster than FormatNumber, because it does not do any runtime type evaluation.
func FormatNumberFloat64(x float64, precision int, thousand string, decimal string) string {
	return formatNumberString(fmt.Sprintf(fmt.Sprintf("%%.%df", precision), x), precision, thousand, decimal)
}

// FormatNumberBigRat only supports *big.Rat value.
// It is faster than FormatNumber, because it does not do any runtime type evaluation.
func FormatNumberBigRat(x *big.Rat, precision int, thousand string, decimal string) string {
	return formatNumberString(x.FloatString(precision), precision, thousand, decimal)
}

// FormatNumberBigDecimal only supports *apd.Decimal value.
// It is faster than FormatNumber, because it does not do any runtime type evaluation.
func FormatNumberBigDecimal(x *apd.Decimal, precision int, thousand string, decimal string) string {
	d := apd.New(0, 0)
	apd.BaseContext.WithPrecision(uint32(x.NumDigits())+uint32(precision)).Quantize(d, x, int32(-precision))
	return formatNumberString(d.Text('f'), precision, thousand, decimal)
}
