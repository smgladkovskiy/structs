package structs

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"gitlab.teamc.io/teamc.io/calcifer/yubaba/support/logs"
)

type NullTime struct {
	Time  time.Time
	Valid bool // Valid is true if Time is not NULL
}

// NewNullTime Создание NullTime переменной
func NewNullTime(v interface{}) NullTime {
	var nt NullTime
	err := nt.Scan(v)
	if err != nil {
		logs.Logger.Error(err)
	}
	return nt
}

// Scan implements the Scanner interface for NullTime
func (nt *NullTime) Scan(value interface{}) error {
	switch v := value.(type) {
	case NullTime:
		*nt = v
		return nil
	case nil:
		*nt = NullTime{Time: time.Time{}, Valid: false}
		return nil
	case string:
		t, err := time.Parse(TimeFormat, v)
		if err != nil {
			*nt = NullTime{Time: time.Time{}, Valid: false}
			return err
		}
		*nt = NullTime{Time: t, Valid: true}
		return nil
	case time.Time:
		if v.IsZero() {
			*nt = NullTime{Time: time.Time{}, Valid: false}
			return nil
		}

		*nt = NullTime{Time: v, Valid: true}

		return nil
	case *time.Time:
		if v.IsZero() {
			*nt = NullTime{Time: time.Time{}, Valid: false}
			return nil
		}

		*nt = NullTime{Time: *v, Valid: true}

		return nil
	}

	return fmt.Errorf("unsupported Scan, storing driver.Value type %T into type %T", value, nt)
}

// Value implements the driver Valuer interface.
func (nt NullTime) Value() (driver.Value, error) {
	if !nt.Valid {
		return nil, nil
	}
	return nt.Time, nil
}

func (nt *NullTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		nt.Time = time.Time{}
		return
	}
	nt.Time, err = time.Parse(TimeFormat, s)
	if err == nil {
		nt.Valid = true
	}
	return
}

func (nt NullTime) MarshalJSON() ([]byte, error) {
	if nt.Valid {
		return json.Marshal(nt.Time.Format(TimeFormat))
	}

	return nullString, nil
}
