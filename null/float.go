package null

import (
	"database/sql/driver"
	"strconv"
	"strings"

	"github.com/smgladkovskiy/structs"
)

type Float64 struct {
	Float64   float64
	Valid     bool
	Precision int
}

func NewFloat64(value interface{}, prc int) (*Float64, error) {
	var nf Float64
	nf.Precision = prc
	err := nf.Scan(value)
	return &nf, err
}

func (nf *Float64) Scan(value interface{}) error {
	//if value == nil {
	//	nf.Float64, nf.Valid = 0, false
	//	return
	//}
	//var (
	//	fv    float64
	//	valid bool
	//)
	switch v := value.(type) {
	case nil:
		return nil
	case string:
		var err error
		nf.Float64, err = strconv.ParseFloat(v, 64)
		nf.Valid = err == nil
		return err
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		i, ok := v.(int)
		if ok {
			nf.Float64, nf.Valid = float64(i), true
			return nil
		}

		i8, ok := v.(int8)
		if ok {
			nf.Float64, nf.Valid = float64(i8), true
			return nil
		}

		i16, ok := v.(int16)
		if ok {
			nf.Float64, nf.Valid = float64(i16), true
			return nil
		}

		i32, ok := v.(int32)
		if ok {
			nf.Float64, nf.Valid = float64(i32), true
			return nil
		}

		i64, ok := v.(int64)
		if ok {
			nf.Float64, nf.Valid = float64(i64), true
			return nil
		}

		ui, ok := v.(uint)
		if ok {
			nf.Float64, nf.Valid = float64(ui), true
			return nil
		}

		ui8, ok := v.(uint8)
		if ok {
			nf.Float64, nf.Valid = float64(ui8), true
			return nil
		}

		ui16, ok := v.(uint16)
		if ok {
			nf.Float64, nf.Valid = float64(ui16), true
			return nil
		}

		ui32, ok := v.(uint32)
		if ok {
			nf.Float64, nf.Valid = float64(ui32), true
			return nil
		}

		ui64, ok := v.(uint64)
		if ok {
			nf.Float64, nf.Valid = float64(ui64), true
			return nil
		}
	case float32, float64:
		f32, ok := v.(float32)
		if ok {
			nf.Float64, nf.Valid = float64(f32), true
			return nil
		}

		f64, ok := v.(float64)
		if ok {
			nf.Float64, nf.Valid = f64, true
			return nil
		}
	case []byte:
		if v == nil {
			return nil
		}

		err := nf.UnmarshalJSON(v)
		return err
	case Float64:
		nf.Float64, nf.Valid = v.Float64, v.Valid
		return nil
	case *Float64:
		nf.Float64, nf.Valid = v.Float64, v.Valid
		return nil
	}
	return structs.TypeIsNotAcceptable{CheckedValue: value, CheckedType: nf}
}

func (nf *Float64) Value() (driver.Value, error) {
	if !nf.Valid {
		return nil, nil
	}
	return nf.Float64, nil
}

func (nf *Float64) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	//str := strings.Replace(string(data), `"`, "", -1)
	if s == "null" || s == "" {
		nf.Float64, nf.Valid = 0.0, false
		return nil
	}
	var err error
	nf.Float64, err = strconv.ParseFloat(s, 64)
	nf.Valid = err == nil
	return err
}

func (nf Float64) MarshalJSON() ([]byte, error) {
	if !nf.Valid {
		return structs.NullString, nil
	}

	var b []byte
	b = strconv.AppendFloat(b, nf.Float64, 'f', nf.Precision, 64)

	return b, nil
}
