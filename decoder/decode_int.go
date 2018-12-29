package decoder

import (
	"fmt"
	"math"
)

var digits []int8

const maxInt64toMultiply = math.MaxInt64 / 10
const maxInt32toMultiply = math.MaxInt32 / 10
const maxInt16toMultiply = math.MaxInt16 / 10
const maxInt8toMultiply = math.MaxInt8 / 10
const maxUint8toMultiply = math.MaxUint8 / 10
const maxUint16toMultiply = math.MaxUint16 / 10
const maxUint32toMultiply = math.MaxUint32 / 10
const maxUint64toMultiply = math.MaxUint64 / 10
const maxUint32Length = 10
const maxUint64Length = 20
const maxUint16Length = 5
const maxUint8Length = 3
const maxInt32Length = 10
const maxInt64Length = 19
const maxInt16Length = 5
const maxInt8Length = 3
const invalidNumber = int8(-1)

var pow10uint64 = [21]uint64{
	0,
	1,
	10,
	100,
	1000,
	10000,
	100000,
	1000000,
	10000000,
	100000000,
	1000000000,
	10000000000,
	100000000000,
	1000000000000,
	10000000000000,
	100000000000000,
	1000000000000000,
	10000000000000000,
	100000000000000000,
	1000000000000000000,
	10000000000000000000,
}

func (dec *Decoder) DecodeInt(v *int64) error {
	for ; dec.cursor < dec.Length || dec.read(); dec.cursor++ {
		switch c := dec.Data[dec.cursor]; c {
		case ' ', '\n', '\t', '\r', ',':
			continue
		// we don't look for 0 as leading zeros are invalid per RFC
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			val, err := dec.getInt64()
			if err != nil {
				return err
			}
			*v = int64(val)
			return nil
		case '-':
			dec.cursor = dec.cursor + 1
			val, err := dec.getInt64Negative()
			if err != nil {
				return err
			}
			*v = -int64(val)
			return nil
		case 'n':
			dec.cursor++
			err := dec.assertNull()
			if err != nil {
				return err
			}
			return nil
		default:
			dec.Err = InvalidUnmarshalError(
				fmt.Sprintf(
					"Cannot unmarshall to int, wrong char '%s' found at pos %d",
					string(dec.Data[dec.cursor]),
					dec.cursor,
				),
			)
			err := dec.skipData()
			if err != nil {
				return err
			}
			return nil
		}
	}
	return dec.raiseInvalidJSONErr(dec.cursor)
}
func (dec *Decoder) getInt64() (int64, error) {
	var end = dec.cursor
	var start = dec.cursor
	// look for following numbers
	for j := dec.cursor + 1; j < dec.Length || dec.read(); j++ {
		switch dec.Data[j] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			end = j
			continue
		case ' ', '\t', '\n', ',', '}', ']':
			dec.cursor = j
			return dec.atoi64(start, end), nil
		case '.':
			// if dot is found
			// look for exponent (e,E) as exponent can change the
			// way number should be parsed to int.
			// if no exponent found, just unmarshal the number before decimal point
			j++
			startDecimal := j
			endDecimal := j - 1
			for ; j < dec.Length || dec.read(); j++ {
				switch dec.Data[j] {
				case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
					endDecimal = j
					continue
				case 'e', 'E':
					// if eg 1.E
					if startDecimal > endDecimal {
						return 0, dec.raiseInvalidJSONErr(dec.cursor)
					}
					dec.cursor = j + 1
					// can try unmarshalling to int as Exponent might change decimal number to non decimal
					// let's get the float value first
					// we get part before decimal as integer
					beforeDecimal := dec.atoi64(start, end)
					// get number after the decimal point
					// multiple the before decimal point portion by 10 using bitwise
					for i := startDecimal; i <= endDecimal; i++ {
						beforeDecimal = (beforeDecimal << 3) + (beforeDecimal << 1)
					}
					// then we add both integers
					// then we divide the number by the power found
					afterDecimal := dec.atoi64(startDecimal, endDecimal)
					expI := endDecimal - startDecimal + 2
					if expI >= len(pow10uint64) || expI < 0 {
						return 0, dec.raiseInvalidJSONErr(dec.cursor)
					}
					pow := pow10uint64[expI]
					floatVal := float64(beforeDecimal+afterDecimal) / float64(pow)
					// we have the floating value, now multiply by the exponent
					exp, err := dec.getExponent()
					if err != nil {
						return 0, err
					}
					pExp := (exp + (exp >> 31)) ^ (exp >> 31) + 1 // abs
					if pExp >= int64(len(pow10uint64)) || pExp < 0 {
						return 0, dec.raiseInvalidJSONErr(dec.cursor)
					}
					val := floatVal * float64(pow10uint64[pExp])
					return int64(val), nil
				case ' ', '\t', '\n', ',', ']', '}':
					dec.cursor = j
					return dec.atoi64(start, end), nil
				default:
					dec.cursor = j
					return 0, dec.raiseInvalidJSONErr(dec.cursor)
				}
			}
			return dec.atoi64(start, end), nil
		case 'e', 'E':
			// get init n
			dec.cursor = j + 1
			return dec.getInt64WithExp(dec.atoi64(start, end))
		}
		// invalid json we expect numbers, dot (single one), comma, or spaces
		return 0, dec.raiseInvalidJSONErr(dec.cursor)
	}
	return dec.atoi64(start, end), nil
}
func (dec *Decoder) getInt64Negative() (int64, error) {
	// look for following numbers
	for ; dec.cursor < dec.Length || dec.read(); dec.cursor++ {
		switch dec.Data[dec.cursor] {
		case '1', '2', '3', '4', '5', '6', '7', '8', '9':
			return dec.getInt64()
		default:
			return 0, dec.raiseInvalidJSONErr(dec.cursor)
		}
	}
	return 0, dec.raiseInvalidJSONErr(dec.cursor)
}
func (dec *Decoder) atoi64(start, end int) int64 {
	var ll = end + 1 - start
	var val = int64(digits[dec.Data[start]])
	end = end + 1
	if ll < maxInt64Length {
		for i := start + 1; i < end; i++ {
			intv := int64(digits[dec.Data[i]])
			val = (val << 3) + (val << 1) + intv
		}
		return val
	} else if ll == maxInt64Length {
		for i := start + 1; i < end; i++ {
			intv := int64(digits[dec.Data[i]])
			if val > maxInt64toMultiply {
				dec.Err = dec.makeInvalidUnmarshalErr(val)
				return 0
			}
			val = (val << 3) + (val << 1)
			if math.MaxInt64-val < intv {
				dec.Err = dec.makeInvalidUnmarshalErr(val)
				return 0
			}
			val += intv
		}
	} else {
		dec.Err = dec.makeInvalidUnmarshalErr(val)
		return 0
	}
	return val
}
func (dec *Decoder) getExponent() (int64, error) {
	start := dec.cursor
	end := dec.cursor
	for ; dec.cursor < dec.Length || dec.read(); dec.cursor++ {
		switch dec.Data[dec.cursor] { // is positive
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			end = dec.cursor + 1
		case '-':
			dec.cursor++
			exp, err := dec.getExponent()
			return -exp, err
		case '+':
			dec.cursor++
			return dec.getExponent()
		default:
			// if nothing return 0
			// could raise error
			if start == end {
				return 0, dec.raiseInvalidJSONErr(dec.cursor)
			}
			return dec.atoi64(start, end-1), nil
		}
	}
	if start == end {

		return 0, dec.raiseInvalidJSONErr(dec.cursor)
	}
	return dec.atoi64(start, end-1), nil
}
func (dec *Decoder) getInt64WithExp(init int64) (int64, error) {
	var exp uint64
	var sign = int64(1)
	for ; dec.cursor < dec.Length || dec.read(); dec.cursor++ {
		switch dec.Data[dec.cursor] {
		case '+':
			continue
		case '-':
			sign = -1
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			uintv := uint64(digits[dec.Data[dec.cursor]])
			exp = (exp << 3) + (exp << 1) + uintv
			dec.cursor++
			for ; dec.cursor < dec.Length || dec.read(); dec.cursor++ {
				switch dec.Data[dec.cursor] {
				case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
					uintv := uint64(digits[dec.Data[dec.cursor]])
					exp = (exp << 3) + (exp << 1) + uintv
				case ' ', '\t', '\n', '}', ',', ']':
					if exp+1 >= uint64(len(pow10uint64)) {
						return 0, dec.raiseInvalidJSONErr(dec.cursor)
					}
					if sign == -1 {
						return init * (1 / int64(pow10uint64[exp+1])), nil
					}
					return init * int64(pow10uint64[exp+1]), nil
				default:
					return 0, dec.raiseInvalidJSONErr(dec.cursor)
				}
			}
			if exp+1 >= uint64(len(pow10uint64)) {
				return 0, dec.raiseInvalidJSONErr(dec.cursor)
			}
			if sign == -1 {
				return init * (1 / int64(pow10uint64[exp+1])), nil
			}
			return init * int64(pow10uint64[exp+1]), nil
		default:
			return 0, dec.raiseInvalidJSONErr(dec.cursor)
		}
	}
	return 0, dec.raiseInvalidJSONErr(dec.cursor)
}
