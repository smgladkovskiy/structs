package null

import (
	"database/sql/driver"
	"fmt"
	"github.com/smgladkovskiy/structs"
	"github.com/smgladkovskiy/structs/decoder"
	"github.com/smgladkovskiy/structs/encoder"
	"github.com/smgladkovskiy/structs/zero"
	"log"
	"strconv"
	"strings"
	"time"
)

// String Реализация String
type String struct {
	String string
	Valid  bool // iv is true if String is not NULL
}

// NewStringf Создание String переменной по текстовому формату с аргументами
func NewStringf(format string, a ...interface{}) *String {
	return NewString(fmt.Sprintf(format, a...))
}

// NewString Создание String переменной
func NewString(v interface{}) *String {
	var n String
	err := n.Scan(v)
	if err != nil {
		log.Print(err)
	}
	return &n
}

// Scan implements the Scanner interface.
func (ns *String) Scan(value interface{}) error {
	switch v := value.(type) {
	case nil:
		ns.String = ""
		return nil
	case string:
		if v != "" {
			ns.String, ns.Valid = v, true
		}
		return nil
	case String:
		*ns = v
		return nil
	case []byte:
		if string(v) == "false" || string(v) == "true" {
			break
		}
		if v != nil && string(v) != "null" && string(v) != "" && string(v) != "\"\"" {
			ns.String, ns.Valid = strings.Trim(string(v), "\""), true
		}
		return nil
	case structs.RawBytes:
		if v == nil {
			return nil
		}
		var d structs.RawBytes
		return ns.Scan(append((d)[:0], v...))
	case time.Time:
		ns.String, ns.Valid = v.Format(structs.TimeFormat()), true
		return nil
	case Time:
		ns.String, ns.Valid = v.Time.Format(structs.TimeFormat()), v.Valid
		return nil
	case zero.Time:
		ns.String, ns.Valid = v.Time.Format(structs.TimeFormat()), true
		return nil
	case int, uint, int8, uint8, int16, uint16, int32, uint32, int64, uint64:
		i, ok := v.(int)
		if ok {
			ns.String, ns.Valid = strconv.FormatInt(int64(i), 10), true
			return nil
		}
		ui, ok := v.(uint)
		if ok {
			ns.String, ns.Valid = strconv.FormatInt(int64(ui), 10), true
			return nil
		}
		i8, ok := v.(int8)
		if ok {
			ns.String, ns.Valid = strconv.FormatInt(int64(i8), 10), true
			return nil
		}
		ui8, ok := v.(uint8)
		if ok {
			ns.String, ns.Valid = strconv.FormatInt(int64(ui8), 10), true
			return nil
		}
		i16, ok := v.(int16)
		if ok {
			ns.String, ns.Valid = strconv.FormatInt(int64(i16), 10), true
			return nil
		}
		ui16, ok := v.(uint16)
		if ok {
			ns.String, ns.Valid = strconv.FormatInt(int64(ui16), 10), true
			return nil
		}
		i32, ok := v.(int32)
		if ok {
			ns.String, ns.Valid = strconv.FormatInt(int64(i32), 10), true
			return nil
		}
		ui32, ok := v.(uint32)
		if ok {
			ns.String, ns.Valid = strconv.FormatInt(int64(ui32), 10), true
			return nil
		}
		i64, ok := v.(int64)
		if ok {
			ns.String, ns.Valid = strconv.FormatInt(i64, 10), true
			return nil
		}
		ui64, ok := v.(uint64)
		if ok {
			ns.String, ns.Valid = strconv.FormatInt(int64(ui64), 10), true
			return nil
		}
	}

	return structs.TypeIsNotAcceptable{CheckedValue: value, CheckedType: ns}
}

// va implements the driver Valuer interface.
func (ns String) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return ns.String, nil
}

// MarshalJSON correctly serializes a String to JSON
func (ns String) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return structs.NullString, nil
	}

	bytes := encoder.StringToBytes(ns.String)

	return bytes, nil
}

// MarshalJSON correctly serializes a String to JSON
func (ns *String) UnmarshalJSON(b []byte) (err error) {
	// s := strings.Trim(string(b), "\"")

	var str string
	dec := &decoder.Decoder{}
	dec.Length = len(b)
	dec.Data = b
	err = dec.DecodeString(&str)
	if err != nil {
		return err
	}

	// // Ignore null, like in the main JSON package.
	if &str == nil {
		ns.String = ""
		return
	}

	ns.String, ns.Valid = str, err == nil
	return
}
