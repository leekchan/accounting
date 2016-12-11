# accounting - money and currency formatting for golang
[![Build Status](https://travis-ci.org/leekchan/accounting.svg?branch=master)](https://travis-ci.org/leekchan/accounting)
[![Coverage Status](https://coveralls.io/repos/leekchan/accounting/badge.svg?branch=master&service=github)](https://coveralls.io/github/leekchan/accounting?branch=master)
[![GoDoc](https://godoc.org/github.com/leekchan/accounting?status.svg)](https://godoc.org/github.com/leekchan/accounting)

accounting is a library for money and currency formatting. (inspired by [accounting.js](https://github.com/openexchangerates/accounting.js))


## Quick Start

```
go get github.com/leekchan/accounting
```

example.go

```Go
package main

import (
    "fmt"
    "math/big"

    "github.com/leekchan/accounting"
)

func main() {
    ac := accounting.Accounting{Symbol: "$", Precision: 2}
    fmt.Println(ac.FormatMoney(123456789.213123))                       // "$123,456,789.21"
    fmt.Println(ac.FormatMoney(12345678))                               // "$12,345,678.00"
    fmt.Println(ac.FormatMoney(big.NewRat(77777777, 3)))                // "$25,925,925.67"
    fmt.Println(ac.FormatMoney(big.NewRat(-77777777, 3)))               // "-$25,925,925.67"
    fmt.Println(ac.FormatMoneyBigFloat(big.NewFloat(123456789.213123))) // "$123,456,789.21"

    ac = accounting.Accounting{Symbol: "€", Precision: 2, Thousand: ".", Decimal: ","}
    fmt.Println(ac.FormatMoney(4999.99))  // "€4.999,99"

    ac = accounting.Accounting{Symbol: "£ ", Precision: 0}
    fmt.Println(ac.FormatMoney(500000)) // "£ 500,000"

    ac = accounting.Accounting{Symbol: "GBP", Precision: 0,
        Format: "%s %v", FormatNegative: "%s (%v)", FormatZero: "%s --"}
    fmt.Println(ac.FormatMoney(1000000)) // "GBP 1,000,000"
    fmt.Println(ac.FormatMoney(-5000))   // "GBP (5,000)"
    fmt.Println(ac.FormatMoney(0))       // "GBP --"
}
```

## Caution

Please do not use float64 to count money. Floats can have errors when you perform operations on them. Using [big.Rat](https://golang.org/pkg/math/big/#Rat) (< Go 1.5) or [big.Float](https://golang.org/pkg/math/big/#Float) (>= Go 1.5) is highly recommended.
(accounting supports float64, but it is just for convenience.)

* [FormatMoneyBigFloat(value *big.Float) string](#formatmoneybigfloatvalue-bigfloat-string)
* [FormatMoneyBigRat(value *big.Rat) string](#formatmoneybigratvalue-bigrat-string)

## Initialization

### Accounting struct

```Go
type Accounting struct {
    Symbol         string // currency symbol (required)
    Precision      int    // currency precision (decimal places) (optional / default: 0)
    Thousand       string // thousand separator (optional / default: ,)
    Decimal        string // decimal separator (optional / default: .)
    Format         string // simple format string allows control of symbol position (%v = value, %s = symbol) (default: %s%v)
    FormatNegative string // format string for negative values (optional / default: strings.Replace(strings.Replace(accounting.Format, "-", "", -1), "%v", "-%v", -1))
    FormatZero     string // format string for zero values (optional / default: Format)
}
```

Field | Type | Description | Default | Example
-------------| ------------- | ------------- | ------------- | -------------
Symbol         | string | currency symbol | no default value | $
Precision      | int    | currency precision (decimal places) | 0 | 2
Thousand       | string | thousand separator | , | .
Decimal        | string | decimal separator | . | ,
Format         | string | simple format string allows control of symbol position (%v = value, %s = symbol) | %s%v | %s %v
FormatNegative | string | format string for negative values | strings.Replace(strings.Replace(accounting.Format, "-", "", -1), "%v", "-%v", -1)) | %s (%v)
FormatZero     | string | format string for zero values | Format | %s --

**Examples:**

```Go
ac := accounting.Accounting{Symbol: "$", Precision: 2}

ac = accounting.Accounting{Symbol: "€", Precision: 2, Thousand: ".", Decimal: ","}

ac = accounting.Accounting{Symbol: "GBP", Precision: 0,
        Format: "%s %v", FormatNegative: "%s (%v)", FormatZero: "%s --"}
```

## FormatMoney(value interface{}) string

FormatMoney is a function for formatting numbers as money values, with customisable currency symbol, precision (decimal places), and thousand/decimal separators.

FormatMoney supports various types of value by runtime reflection. If you don't need runtime type evaluation, please refer to FormatMoneyInt, FormatMoneyBigRat, FormatMoneyBigRat, or FormatMoneyFloat64.

* supported value types : int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, *big.Rat

**Examples:**

```Go
ac := accounting.Accounting{Symbol: "$", Precision: 2}
fmt.Println(ac.FormatMoney(123456789.213123))         // "$123,456,789.21"
fmt.Println(ac.FormatMoney(12345678))                 // "$12,345,678.00"
fmt.Println(ac.FormatMoney(big.NewRat(77777777, 3)))  // "$25,925,925.67"
fmt.Println(ac.FormatMoney(big.NewRat(-77777777, 3))) // "-$25,925,925.67"

ac = accounting.Accounting{Symbol: "€", Precision: 2, Thousand: ".", Decimal: ","}
fmt.Println(ac.FormatMoney(4999.99))  // "€4.999,99"

ac = accounting.Accounting{Symbol: "£ ", Precision: 0}
fmt.Println(ac.FormatMoney(500000)) // "£ 500,000"

ac = accounting.Accounting{Symbol: "GBP", Precision: 0,
    Format: "%s %v", FormatNegative: "%s (%v)", FormatZero: "%s --"}
fmt.Println(ac.FormatMoney(1000000)) // "GBP 1,000,000"
fmt.Println(ac.FormatMoney(-5000))   // "GBP (5,000)"
fmt.Println(ac.FormatMoney(0))       // "GBP --"
```

## FormatMoneyBigFloat(value *big.Float) string

**(>= Go 1.5)**

FormatMoneyBigFloat only supports [*big.Float](https://golang.org/pkg/math/big/#Float) value. It is faster than FormatMoney, because it does not do any runtime type evaluation.

**Examples:**

```Go
ac = accounting.Accounting{Symbol: "$", Precision: 2}
fmt.Println(ac.FormatMoneyBigFloat(big.NewFloat(123456789.213123))) // "$123,456,789.21"
fmt.Println(ac.FormatMoneyBigFloat(big.NewFloat(12345678)))         // "$12,345,678.00"

ac = accounting.Accounting{Symbol: "€", Precision: 2, Thousand: ".", Decimal: ","}
fmt.Println(ac.FormatMoneyBigFloat(big.NewFloat(4999.99)))  // "€4.999,99"

ac = accounting.Accounting{Symbol: "£ ", Precision: 0}
fmt.Println(ac.FormatMoneyBigFloat(big.NewFloat(500000))) // "£ 500,000"

ac = accounting.Accounting{Symbol: "GBP", Precision: 0,
    Format: "%s %v", FormatNegative: "%s (%v)", FormatZero: "%s --"}
fmt.Println(ac.FormatMoneyBigFloat(big.NewFloat(1000000))) // "GBP 1,000,000"
fmt.Println(ac.FormatMoneyBigFloat(big.NewFloat(-5000)))   // "GBP (5,000)"
fmt.Println(ac.FormatMoneyBigFloat(big.NewFloat(0)))       // "GBP --"
```

## FormatMoneyInt(value int) string

FormatMoneyInt only supports int value. It is faster than FormatMoney, because it does not do any runtime type evaluation.

**Examples:**

```Go
ac = accounting.Accounting{Symbol: "$", Precision: 2}
fmt.Println(ac.FormatMoneyInt(12345678)) // "$12,345,678.00"

ac = accounting.Accounting{Symbol: "€", Precision: 2, Thousand: ".", Decimal: ","}
fmt.Println(ac.FormatMoneyInt(4999))  // "€4.999,00"

ac = accounting.Accounting{Symbol: "£ ", Precision: 0}
fmt.Println(ac.FormatMoneyInt(500000)) // "£ 500,000"

ac = accounting.Accounting{Symbol: "GBP", Precision: 0,
    Format: "%s %v", FormatNegative: "%s (%v)", FormatZero: "%s --"}
fmt.Println(ac.FormatMoneyInt(1000000)) // "GBP 1,000,000"
fmt.Println(ac.FormatMoneyInt(-5000))   // "GBP (5,000)"
fmt.Println(ac.FormatMoneyInt(0))       // "GBP --"
```

## FormatMoneyBigRat(value *big.Rat) string

FormatMoneyBigRat only supports [*big.Rat](https://golang.org/pkg/math/big/#Rat) value. It is faster than FormatMoney, because it does not do any runtime type evaluation.

**Examples:**

```Go
ac = accounting.Accounting{Symbol: "$", Precision: 2}
fmt.Println(ac.FormatMoneyBigRat(big.NewRat(77777777, 3)))  // "$25,925,925.67"
fmt.Println(ac.FormatMoneyBigRat(big.NewRat(-77777777, 3))) // "-$25,925,925.67"

ac = accounting.Accounting{Symbol: "€", Precision: 2, Thousand: ".", Decimal: ","}
fmt.Println(ac.FormatMoneyBigRat(big.NewRat(77777777, 3)))  // "€25.925.925,67"

ac = accounting.Accounting{Symbol: "£ ", Precision: 0}
fmt.Println(ac.FormatMoneyBigRat(big.NewRat(77777777, 3)))  // "£ 25,925,926"

ac = accounting.Accounting{Symbol: "GBP", Precision: 0,
    Format: "%s %v", FormatNegative: "%s (%v)", FormatZero: "%s --"}
fmt.Println(ac.FormatMoneyBigRat(big.NewRat(77777777, 3)))  // "GBP 25,925,926"
fmt.Println(ac.FormatMoneyBigRat(big.NewRat(-77777777, 3))) // "GBP (25,925,926)"
fmt.Println(ac.FormatMoneyBigRat(big.NewRat(0, 3)))         // "GBP --"
```

## FormatMoneyFloat64(value float64) string

FormatMoneyFloat64 only supports float64 value. It is faster than FormatMoney, because it does not do any runtime type evaluation.

**Examples:**

```Go
ac = accounting.Accounting{Symbol: "$", Precision: 2}
fmt.Println(ac.FormatMoneyFloat64(123456789.213123)) // "$123,456,789.21"
fmt.Println(ac.FormatMoneyFloat64(12345678))         // "$12,345,678.00"

ac = accounting.Accounting{Symbol: "€", Precision: 2, Thousand: ".", Decimal: ","}
fmt.Println(ac.FormatMoneyFloat64(4999.99))  // "€4.999,99"

ac = accounting.Accounting{Symbol: "£ ", Precision: 0}
fmt.Println(ac.FormatMoneyFloat64(500000)) // "£ 500,000"

ac = accounting.Accounting{Symbol: "GBP", Precision: 0,
    Format: "%s %v", FormatNegative: "%s (%v)", FormatZero: "%s --"}
fmt.Println(ac.FormatMoneyFloat64(1000000)) // "GBP 1,000,000"
fmt.Println(ac.FormatMoneyFloat64(-5000))   // "GBP (5,000)"
fmt.Println(ac.FormatMoneyFloat64(0))       // "GBP --"
```

## FormatNumber(value interface{}, precision int, thousand string, decimal string) string

FormatNumber is a base function of the library which formats a number with custom precision and separators. 

FormatNumber supports various types of value by runtime reflection. If you don't need runtime type evaluation, please refer to FormatNumberInt, FormatNumberBigRat, or FormatNumberFloat64.

* supported value types : int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, *big.Rat

**Examples:**

```Go
fmt.Println(accounting.FormatNumber(123456789.213123, 3, ",", ".")) // "123,456,789.213"
fmt.Println(accounting.FormatNumber(1000000, 3, " ", ","))          // "1 000 000,000"
```

## FormatNumberBigFloat(value *big.Float, precision int, thousand string, decimal string) string

**(>= Go 1.5)**

FormatNumberBigFloat only supports [*big.Float](https://golang.org/pkg/math/big/#Float) value. It is faster than FormatNumber, because it does not do any runtime type evaluation.

**Examples:**

```Go
fmt.Println(accounting.FormatNumberBigFloat(big.NewFloat(123456789.213123), 3, ",", ".")) // "123,456,789.213"
```

## FormatNumberInt(value int, precision int, thousand string, decimal string) string

FormatNumberInt only supports int value. It is faster than FormatNumber, because it does not do any runtime type evaluation.

**Examples:**

```Go
fmt.Println(accounting.FormatNumberInt(123456789, 3, ",", ".")) // "123,456,789.000"
```

## FormatNumberBigRat(value *big.Rat, precision int, thousand string, decimal string) string

FormatNumberBigRat only supports [*big.Rat](https://golang.org/pkg/math/big/#Rat) value. It is faster than FormatNumber, because it does not do any runtime type evaluation.

**Examples:**

```Go
fmt.Println(accounting.FormatNumberBigRat(big.NewRat(77777777, 3), 3, ",", ".")) // "25,925,925.667"
```

## FormatNumberFloat64(value float64, precision int, thousand string, decimal string) string

FormatNumberFloat64 only supports float64 value. It is faster than FormatNumber, because it does not do any runtime type evaluation.

**Examples:**

```Go
fmt.Println(accounting.FormatNumberFloat64(123456789.213123, 3, ",", ".")) // "123,456,789.213"
```