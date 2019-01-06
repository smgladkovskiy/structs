package null

import (
	"bytes"
	"database/sql/driver"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Float64 struct {
	Float64   float64
	Valid     bool
	Precision int
}

func NewFloat64(value interface{}, prc int) *Float64 {
	var nf Float64
	nf.Precision = prc
	if err := nf.Scan(value); err != nil {
		log.Print(err)
	}
	return &nf
}

func (nf *Float64) Scan(value interface{}) (err error) {
	if value == nil {
		nf.Float64, nf.Valid = 0, false
		return
	}
	var (
		fv    float64
		valid bool
	)
	switch v := value.(type) {
	case string:
		if val, err := strconv.ParseFloat(v, 64); err == nil {
			fv, valid = val, true
		}
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		if f, ok := v.(int); ok {
			fv = float64(f)
			valid = true
		}
		if f, ok := v.(int8); ok {
			fv = float64(f)
			valid = true
		}
		if f, ok := v.(int16); ok {
			fv = float64(f)
			valid = true
		}
		if f, ok := v.(int32); ok {
			fv = float64(f)
			valid = true
		}
		if f, ok := v.(int64); ok {
			fv = float64(f)
			valid = true
		}
		if f, ok := v.(uint); ok {
			fv = float64(f)
			valid = true
		}
		if f, ok := v.(uint8); ok {
			fv = float64(f)
			valid = true
		}
		if f, ok := v.(uint16); ok {
			fv = float64(f)
			valid = true
		}
		if f, ok := v.(uint32); ok {
			fv = float64(f)
			valid = true
		}
		if f, ok := v.(uint64); ok {
			fv = float64(f)
			valid = true
		}
	case float32, float64:
		if f, ok := v.(float32); ok {
			fv = float64(f)
			valid = true
		}
		if f, ok := v.(float64); ok {
			fv = float64(f)
			valid = true
		}
	case []byte:
		if val, err := strconv.ParseFloat(string(v), 64); err == nil {
			fv = val
			valid = true
		}
	default:
		return fmt.Errorf("unsupported Scan, trying to store type %T as type %T", value, nf)
	}
	if fv == 0 {
		valid = false
	}
	nf.Float64, nf.Valid = fv, valid
	return
}

func (nf *Float64) Value() (driver.Value, error) {
	if !nf.Valid {
		return nil, nil
	}
	return nf.Float64, nil
}

func (nf *Float64) UnmarshalJSON(data []byte) (err error) {
	var (
		val   float64
		valid bool
	)
	str := strings.Replace(string(data), `"`, "", -1)
	if str == "null" || str == "" {
		valid = false
		nf.Float64, nf.Valid = val, valid
		return
	}
	if v, err := strconv.ParseFloat(str, 64); err != nil {
		return err
	} else {
		val, valid = v, true
	}

	nf.Float64, nf.Valid = val, valid
	return
}

func (nf Float64) MarshalJSON() ([]byte, error) {
	var buffer bytes.Buffer
	prc := strconv.Itoa(nf.Precision)
	if nf.Valid {
		buffer.WriteString(fmt.Sprintf(`%.`+prc+`f`, nf.Float64))
	} else {
		buffer.WriteString("null")
	}
	return buffer.Bytes(), nil
}
