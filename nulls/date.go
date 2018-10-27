package nulls

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/smgladkovskiy/go-structs"
	"log"
	"strings"
	"time"
)

type NullDate struct {
	Time  time.Time
	Valid bool // Valid is true if Time is not NULL
}

// NewNullDate Создание NullDate переменной
func NewNullDate(v interface{}) NullDate {
	var nt NullDate
	err := nt.Scan(v)
	if err != nil {
		log.Print(err)
	}
	return nt
}

// Scan implements the Scanner interface for NullDate
func (nt *NullDate) Scan(value interface{}) error {
	switch v := value.(type) {
	case NullDate:
		*nt = v
		return nil
	case nil:
		*nt = NullDate{Time: time.Time{}, Valid: false}
		return nil
	case string:
		t, err := time.Parse(structs.DateFormat(), v)
		if err != nil {
			*nt = NullDate{Time: time.Time{}, Valid: false}
			return err
		}
		*nt = NullDate{Time: t, Valid: true}
		return nil
	case time.Time:
		if v.IsZero() {
			*nt = NullDate{Time: time.Time{}, Valid: false}
			return nil
		}

		*nt = NullDate{Time: v, Valid: true}

		return nil
	case *time.Time:
		if v.IsZero() {
			*nt = NullDate{Time: time.Time{}, Valid: false}
			return nil
		}

		*nt = NullDate{Time: *v, Valid: true}

		return nil
	}

	return fmt.Errorf("unsupported Scan, storing driver.Value type %T into type %T", value, nt)
}

// Value implements the driver Valuer interface.
func (nt NullDate) Value() (driver.Value, error) {
	if !nt.Valid {
		return nil, nil
	}
	return nt.Time, nil
}

func (nt *NullDate) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		nt.Time = time.Time{}
		return
	}
	nt.Time, err = time.Parse(structs.DateFormat(), s)
	if err == nil {
		nt.Valid = true
	}
	return
}

func (nt NullDate) MarshalJSON() ([]byte, error) {
	if nt.Valid {
		return json.Marshal(nt.Time.Format(structs.DateFormat()))
	}

	return structs.NullString, nil
}
