package nulls

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"gitlab.teamc.io/teamc.io/golang/structs"
	"log"
	"strings"
	"time"
)

// NullString Реализация NullString
type NullString struct {
	String string
	Valid  bool // Valid is true if String is not NULL
}

// NewNullStringf Создание NullString переменной по текстовому формату с аргументами
func NewNullStringf(format string, a ...interface{}) *NullString {
	return NewNullString(fmt.Sprintf(format, a...))
}

// NewNullString Создание NullString переменной
func NewNullString(v interface{}) *NullString {
	var n NullString
	err := n.Scan(v)
	if err != nil {
		log.Print(err)
	}
	return &n
}

// Scan implements the Scanner interface.
func (ns *NullString) Scan(value interface{}) error {
	if value == nil {
		*ns = NullString{String: "", Valid: false}
		return nil
	}

	ns.Valid = false
	switch v := value.(type) {
	case string:
		if v != "" {
			*ns = NullString{String: v, Valid: true}
		}
		return nil
	case NullString:
		*ns = v
		return nil
	case []byte:
		if v == nil {
			return structs.ErrNilPtr
		}

		if string(v) != "" {
			*ns = NullString{String: string(v), Valid: true}
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
	case NullTime:
		ns.String, ns.Valid = v.Time.Format(structs.TimeFormat()), true
		return nil
	}

	return fmt.Errorf("unsupported Scan, storing driver.Value type %T into type %T", value, ns)
}

// Value implements the driver Valuer interface.
func (ns NullString) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return ns.String, nil
}

// MarshalJSON correctly serializes a NullString to JSON
func (ns NullString) MarshalJSON() ([]byte, error) {
	if ns.Valid {
		return json.Marshal(ns.String)
	}

	return structs.NullString, nil
}

// MarshalJSON correctly serializes a NullString to JSON
func (ns *NullString) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	_ = json.Unmarshal(b, &s)
	// Ignore null, like in the main JSON package.
	if s == "null" {
		ns.String = ""
		return
	}

	*ns = NullString{String: s, Valid: true}
	return
}
