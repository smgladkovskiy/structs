package nulls

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/smgladkovskiy/go-structs"
	"strings"
)

type NullBool struct {
	Bool  bool
	Valid bool
}

func NewNullBool(v interface{}) *NullBool {
	var nb NullBool
	nb.Scan(v)
	return &nb
}

// Scan implements the Scanner interface.
func (nb *NullBool) Scan(value interface{}) error {
	if value == nil {
		*nb = NullBool{Bool: false, Valid: false}
		return nil
	}

	switch v := value.(type) {
	case NullBool:
		*nb = v
		return nil
	case bool:
		*nb = NullBool{Bool: v, Valid: true}
		return nil
	case []byte:
		b, err := driver.Bool.ConvertValue(v)
		if err == nil {
			*nb = NullBool{Bool: b.(bool), Valid: true}
			return nil
		}

		// *nb = NullBool{Bool: false, Valid: false}
		// return nil
	case string:
		b, err := parseBool(v)
		if err != nil {
			*nb = NullBool{Bool: false, Valid: false}
			return nil
		}

		*nb = NullBool{Bool: b, Valid: true}
		return nil
	case int, int8, int16, int32, int64:
		i, ok := v.(int)
		if ok {
			*nb = NullBool{Bool: false, Valid: false}
			if i == 0 {
				*nb = NullBool{Bool: false, Valid: true}
			}
			if i == 1 {
				*nb = NullBool{Bool: true, Valid: true}
			}

			return nil
		}
		i8, ok := v.(int8)
		if ok {
			*nb = NullBool{Bool: false, Valid: false}
			if i8 == 0 {
				*nb = NullBool{Bool: false, Valid: true}
			}
			if i8 == 1 {
				*nb = NullBool{Bool: true, Valid: true}
			}

			return nil
		}
		i16, ok := v.(int16)
		if ok {
			*nb = NullBool{Bool: false, Valid: false}
			if i16 == 0 {
				*nb = NullBool{Bool: false, Valid: true}
			}
			if i16 == 1 {
				*nb = NullBool{Bool: true, Valid: true}
			}

			return nil
		}
		i32, ok := v.(int32)
		if ok {
			*nb = NullBool{Bool: false, Valid: false}
			if i32 == 0 {
				*nb = NullBool{Bool: false, Valid: true}
			}
			if i32 == 1 {
				*nb = NullBool{Bool: true, Valid: true}
			}

			return nil
		}
		i64, ok := v.(int64)
		if ok {
			*nb = NullBool{Bool: false, Valid: false}
			if i64 == 0 {
				*nb = NullBool{Bool: false, Valid: true}
			}
			if i64 == 1 {
				*nb = NullBool{Bool: true, Valid: true}
			}

			return nil
		}
	}

	return fmt.Errorf("unsupported Scan, storing driver.Value type %T into type %T", value, nb)
}

// Value implements the driver Valuer interface.
func (nb NullBool) Value() (driver.Value, error) {
	if !nb.Valid {
		return nil, nil
	}
	return nb.Bool, nil
}

// MarshalJSON correctly serializes a NullBool to JSON
func (ni NullBool) MarshalJSON() ([]byte, error) {
	if ni.Valid {
		return json.Marshal(ni.Bool)
	}
	return structs.NullString, nil
}

func (ni *NullBool) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	// Ignore null, like in the main JSON package.
	if s == "null" {
		ni.Bool = false
		return
	}

	err = ni.Scan(s)
	return
}

// ParseBool returns the boolean value represented by the string.
// It accepts 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False.
// Any other value returns an error.
func parseBool(str string) (bool, error) {
	switch str {
	case "1", "t", "T", "true", "TRUE", "True", "y", "Y", "YES", "Yes":
		return true, nil
	case "0", "f", "F", "false", "FALSE", "False", "n", "N", "NO", "No":
		return false, nil
	}
	return false, errors.New(fmt.Sprintf("Error ParseBool from %s", str))
}
