package null

import (
	"database/sql/driver"
	"strings"

	"github.com/smgladkovskiy/structs"
	"github.com/smgladkovskiy/structs/decoder"
	"github.com/smgladkovskiy/structs/encoder"
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
	case *bool:
		nb.Bool, nb.Valid = *v, true
		return nil
	case []byte:
		return nb.UnmarshalJSON(v)
	case string:
		var err error
		nb.Bool, err = parseBool(v)
		nb.Valid = err == nil
		return err
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
	if !nb.Valid {
		return structs.NullString, nil
	}

	if nb.Bool {
		return encoder.StringToBytes("true"), nil
	}

	return encoder.StringToBytes("false"), nil
}

func (nb *Bool) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		nb.Valid = false
		return nil
	}

	var bo bool
	dec := &decoder.Decoder{}
	dec.Length = len(b)
	dec.Data = b
	err := dec.DecodeBool(&bo)
	if err != nil {
		return err
	}
	if dec.Err != nil {
		return dec.Err
	}

	nb.Bool, nb.Valid = bo, true
	return nil
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
	return false, structs.ValueIsNotAcceptable{CheckedValue: str, CheckedType: Bool{}}
}
