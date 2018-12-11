package null

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/smgladkovskiy/structs"
	"github.com/smgladkovskiy/structs/decoder"
	"strings"
)

type Bool struct {
	Bool  bool
	Valid bool
}

func NewBool(v interface{}) (*Bool, error) {
	var nb Bool
	err := nb.Scan(v)
	return &nb, err
}

// Scan implements the Scanner interface.
func (nb *Bool) Scan(value interface{}) error {
	switch v := value.(type) {
	case nil:
		return nil
	case Bool:
		nb.Bool, nb.Valid = v.Bool, v.Valid
		return nil
	case *Bool:
		nb.Bool, nb.Valid = v.Bool, v.Valid
		return nil
	case bool:
		nb.Bool, nb.Valid = v, true
		return nil
	case []byte:
		var b bool
		dec := &decoder.Decoder{}
		dec.Length = len(v)
		dec.Data = v
		err := dec.DecodeBool(&b)
		if err != nil {
			return err
		}
		if dec.Err != nil {
			return dec.Err
		}

		nb.Bool, nb.Valid = b, true
		return nil
	case string:
		b, err := parseBool(v)
		if err != nil {
			nb.Bool, nb.Valid = false, false
			return nil
		}

		nb.Bool, nb.Valid = b, true
		return nil
	case int, uint, int8, uint8, int16, uint16, int32, uint32, int64, uint64:
		switch v {
		case int(0), uint(0), int8(0), uint8(0), int16(0), uint16(0), int32(0), uint32(0), int64(0), uint64(0):
			nb.Bool, nb.Valid = false, true
		case int(1), uint(1), int8(1), uint8(1), int16(1), uint16(1), int32(1), uint32(1), int64(1), uint64(1):
			nb.Bool, nb.Valid = true, true
		default:
			nb.Bool, nb.Valid = false, false
		}

		return nil
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
