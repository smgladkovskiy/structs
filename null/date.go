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

type Date struct {
	Time  time.Time
	Valid bool // iv is true if Time is not NULL
}

// NewDate Создание Date переменной
func NewDate(v interface{}) Date {
	var nt Date
	err := nt.Scan(v)
	if err != nil {
		log.Print(err)
	}
	return nt
}

// Scan implements the Scanner interface for Date
func (nt *Date) Scan(value interface{}) error {
	switch v := value.(type) {
	case Date:
		*nt = v
		return nil
	case nil:
		*nt = Date{Time: time.Time{}, Valid: false}
		return nil
	case string:
		t, err := time.Parse(structs.DateFormat(), v)
		if err != nil {
			*nt = Date{Time: time.Time{}, Valid: false}
			return err
		}
		*nt = Date{Time: t, Valid: true}
		return nil
	case time.Time:
		if v.IsZero() {
			*nt = Date{Time: time.Time{}, Valid: false}
			return nil
		}

		*nt = Date{Time: v, Valid: true}

		return nil
	case *time.Time:
		if v.IsZero() {
			*nt = Date{Time: time.Time{}, Valid: false}
			return nil
		}

		*nt = Date{Time: *v, Valid: true}

		return nil
	}

	return fmt.Errorf("unsupported Scan, storing driver.va type %T into type %T", value, nt)
}

// va implements the driver Valuer interface.
func (nt Date) Value() (driver.Value, error) {
	if !nt.Valid {
		return nil, nil
	}
	return nt.Time, nil
}

func (nt *Date) UnmarshalJSON(b []byte) (err error) {
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

func (nt Date) MarshalJSON() ([]byte, error) {
	if nt.Valid {
		return json.Marshal(nt.Time.Format(structs.DateFormat()))
	}

	return structs.NullString, nil
}
