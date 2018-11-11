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

type Time struct {
	Time  time.Time
	Valid bool // iv is true if Time is not NULL
}

// NewTime Создание Time переменной
func NewTime(v interface{}) Time {
	var nt Time
	err := nt.Scan(v)
	if err != nil {
		log.Print(err)
	}
	return nt
}

// Scan implements the Scanner interface for Time
func (nt *Time) Scan(value interface{}) error {
	switch v := value.(type) {
	case Time:
		*nt = v
		return nil
	case nil:
		*nt = Time{Time: time.Time{}, Valid: false}
		return nil
	case string:
		t, err := time.Parse(structs.TimeFormat(), v)
		if err != nil {
			*nt = Time{Time: time.Time{}, Valid: false}
			return err
		}
		*nt = Time{Time: t, Valid: true}
		return nil
	case time.Time:
		if v.IsZero() {
			*nt = Time{Time: time.Time{}, Valid: false}
			return nil
		}

		*nt = Time{Time: v, Valid: true}

		return nil
	case *time.Time:
		if v.IsZero() {
			*nt = Time{Time: time.Time{}, Valid: false}
			return nil
		}

		*nt = Time{Time: *v, Valid: true}

		return nil
	}

	return fmt.Errorf("unsupported Scan, storing driver.va type %T into type %T", value, nt)
}

// va implements the driver Valuer interface.
func (nt Time) Value() (driver.Value, error) {
	if !nt.Valid {
		return nil, nil
	}
	return nt.Time, nil
}

func (nt *Time) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		nt.Time = time.Time{}
		return
	}
	nt.Time, err = time.Parse(structs.TimeFormat(), s)
	if err == nil {
		nt.Valid = true
	}
	return
}

func (nt Time) MarshalJSON() ([]byte, error) {
	if nt.Valid {
		return json.Marshal(nt.Time.Format(structs.TimeFormat()))
	}

	return structs.NullString, nil
}
