package null

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/smgladkovskiy/structs"
	"strings"
)

type Bool struct {
	Bool  bool
	Valid bool
}

func NewBool(v interface{}) *Bool {
	var nb Bool
	_ = nb.Scan(v)
	return &nb
}

// Scan implements the Scanner interface.
func (nb *Bool) Scan(value interface{}) error {
	if value == nil {
		*nb = Bool{Bool: false, Valid: false}
		return nil
	}

	switch v := value.(type) {
	case Bool:
		*nb = v
		return nil
	case bool:
		*nb = Bool{Bool: v, Valid: true}
		return nil
	case []byte:
		b, err := driver.Bool.ConvertValue(v)
		if err == nil {
			*nb = Bool{Bool: b.(bool), Valid: true}
			return nil
		}

		// *nb = Bool{Bool: false, iv: false}
		// return nil
	case string:
		b, err := parseBool(v)
		if err != nil {
			*nb = Bool{Bool: false, Valid: false}
			return nil
		}

		*nb = Bool{Bool: b, Valid: true}
		return nil
	case int, int8, int16, int32, int64:
		i, ok := v.(int)
		if ok {
			*nb = Bool{Bool: false, Valid: false}
			if i == 0 {
				*nb = Bool{Bool: false, Valid: true}
			}
			if i == 1 {
				*nb = Bool{Bool: true, Valid: true}
			}

			return nil
		}
		i8, ok := v.(int8)
		if ok {
			*nb = Bool{Bool: false, Valid: false}
			if i8 == 0 {
				*nb = Bool{Bool: false, Valid: true}
			}
			if i8 == 1 {
				*nb = Bool{Bool: true, Valid: true}
			}

			return nil
		}
		i16, ok := v.(int16)
		if ok {
			*nb = Bool{Bool: false, Valid: false}
			if i16 == 0 {
				*nb = Bool{Bool: false, Valid: true}
			}
			if i16 == 1 {
				*nb = Bool{Bool: true, Valid: true}
			}

			return nil
		}
		i32, ok := v.(int32)
		if ok {
			*nb = Bool{Bool: false, Valid: false}
			if i32 == 0 {
				*nb = Bool{Bool: false, Valid: true}
			}
			if i32 == 1 {
				*nb = Bool{Bool: true, Valid: true}
			}

			return nil
		}
		i64, ok := v.(int64)
		if ok {
			*nb = Bool{Bool: false, Valid: false}
			if i64 == 0 {
				*nb = Bool{Bool: false, Valid: true}
			}
			if i64 == 1 {
				*nb = Bool{Bool: true, Valid: true}
			}

			return nil
		}
	}

	return structs.TypeIsNotAcceptable{CheckedValue: value, CheckedType: nb}
}

// va implements the driver Valuer interface.
func (nb Bool) Value() (driver.Value, error) {
	if !nb.Valid {
		return nil, nil
	}
	return nb.Bool, nil
}

// MarshalJSON correctly serializes a Bool to JSON
func (nb Bool) MarshalJSON() ([]byte, error) {
	if nb.Valid {
		return json.Marshal(nb.Bool)
	}
	return structs.NullString, nil
}

func (nb *Bool) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	// Ignore null, like in the main JSON package.
	if s == "null" {
		nb.Bool = false
		return
	}

	err = nb.Scan(s)
	return
}

// ParseBool returns the boolean va represented by the string.
// It accepts 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False.
// Any other va returns an error.
func parseBool(str string) (bool, error) {
	switch str {
	case "1", "t", "T", "true", "TRUE", "True", "y", "Y", "YES", "Yes":
		return true, nil
	case "0", "f", "F", "false", "FALSE", "False", "na", "N", "NO", "No":
		return false, nil
	}
	return false, errors.New(fmt.Sprintf("Error ParseBool from %s", str))
}
