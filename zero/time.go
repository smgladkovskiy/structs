package zero

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/smgladkovskiy/structs"
	"log"
	"strings"
	"time"
)

type Time struct {
	time.Time
}

// NewNullTime Создание Time переменной
func NewTime(v interface{}) Time {
	var t Time
	err := t.Scan(v)
	if err != nil {
		log.Print(err)
	}
	return t
}

// Scan implements the Scanner interface for NullTime
func (t *Time) Scan(value interface{}) error {
	switch v := value.(type) {
	case nil:
		*t = Time{Time: time.Time{}}
		return nil
	case string:
		pt, err := time.Parse(structs.TimeFormat(), v)
		if err != nil {
			*t = Time{Time: time.Time{}}
			return err
		}
		*t = Time{Time: pt}
		return nil
	case time.Time:
		if v.IsZero() {
			*t = Time{Time: time.Time{}}
			return nil
		}

		*t = Time{Time: v}

		return nil
	}

	return fmt.Errorf("unsupported Scan, storing driver.Value type %T into type %T", value, t)
}

// Value implements the driver Valuer interface.
func (t Time) Value() (driver.Value, error) {
	return t.Time.Format(structs.TimeFormat()), nil
}

func (t *Time) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	// Ignore null, like in the main JSON package.
	if s == "null" {
		t.Time = time.Time{}
		return
	}

	t.Time, err = time.Parse(structs.TimeFormat(), s)
	return
}

func (t Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Time.Format(structs.TimeFormat()))
}
