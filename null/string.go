package null

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/smgladkovskiy/structs"
	"log"
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
	if value == nil {
		*ns = String{String: "", Valid: false}
		return nil
	}

	ns.Valid = false
	switch v := value.(type) {
	case string:
		if v != "" {
			*ns = String{String: v, Valid: true}
		}
		return nil
	case String:
		*ns = v
		return nil
	case []byte:
		if v == nil {
			return structs.ErrNilPtr
		}

		if string(v) != "" {
			*ns = String{String: string(v), Valid: true}
		}
		return nil
	case structs.RawBytes:
		if v == nil {
			return structs.ErrNilPtr
		}
		var d structs.RawBytes
		ns.String = string(append((d)[:0], v...))
		if ns.String != "" {
			ns.Valid = true
		}
		return nil
	case time.Time:
		ns.String, ns.Valid = v.Format(structs.TimeFormat()), true
		return nil
	case Time:
		ns.String, ns.Valid = v.Time.Format(structs.TimeFormat()), true
		return nil
	}

	return fmt.Errorf("unsupported Scan, storing driver.va type %T into type %T", value, ns)
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
	if ns.Valid {
		return json.Marshal(ns.String)
	}

	return structs.NullString, nil
}

// MarshalJSON correctly serializes a String to JSON
func (ns *String) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	_ = json.Unmarshal(b, &s)
	// Ignore null, like in the main JSON package.
	if s == "null" {
		ns.String = ""
		return
	}

	*ns = String{String: s, Valid: true}
	return
}
