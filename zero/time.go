package zero

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"github.com/smgladkovskiy/structs"
)

type Time struct {
	time.Time
}

// NewNullTime Создание Time переменной
func NewTime(v interface{}) (*Time, error) {
	var t Time
	err := t.Scan(v)
	return &t, err
}

// Scan implements the Scanner interface for NullTime
func (t *Time) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	switch v := value.(type) {
	case nil:
		return nil
	case string:
		pt, err := time.Parse(structs.TimeFormat(), v)
		if err != nil {
			return err
		}
		t.Time = pt
		return nil
	case time.Time:
		if v.IsZero() {
			t.Time = time.Time{}
			return nil
		}
		t.Time = v
		return nil
	case *time.Time:
		if v.IsZero() {
			t.Time = time.Time{}
			return nil
		}
		t.Time = *v
		return nil
	case *Time:
		if v.IsZero() {
			t.Time = v.Time
			return nil
		}
		t.Time = v.Time

		*t = Time{Time: v}

		return nil
	case Time:
		*t = v

		return nil
	case *Time:
		*t = *v

		return nil
	}

	return structs.TypeIsNotAcceptable{CheckedValue: value, CheckedType: t}
}

// Value implements the driver Valuer interface.
func (t Time) Value() (driver.Value, error) {
	return t.Time.Format(structs.TimeFormat()), nil
}

func (t Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Time.Format(structs.TimeFormat()))
}

func (t *Time) UnmarshalJSON(b []byte) (err error) {
	// s := strings.Trim(string(b), "\"")
	// Ignore null, like in the main JSON package.
	if string(b) == "null" {
		t.Time = time.Time{}
		return nil
	}

	t.Time, err = time.Parse(structs.TimeFormat(), string(b))
	return
}
