package zero

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"gitlab.teamc.io/teamc.io/golang/structs"
	"log"
	"strings"
	"time"
)

type Date struct {
	time.Time
}

// NewNullTime Создание Date переменной
func NewDate(v interface{}) Date {
	var t Date
	err := t.Scan(v)
	if err != nil {
		log.Print(err)
	}
	return t
}

// Scan implements the Scanner interface for NullTime
func (d *Date) Scan(value interface{}) error {
	switch v := value.(type) {
	case nil:
		*d = Date{Time: time.Time{}}
		return nil
	case string:
		pt, err := time.Parse(structs.DateFormat(), v)
		if err != nil {
			*d = Date{Time: time.Time{}}
			return err
		}
		*d = Date{Time: pt}
		return nil
	case time.Time:
		if v.IsZero() {
			*d = Date{Time: time.Time{}}
			return nil
		}

		*d = Date{Time: v}

		return nil
	}

	return fmt.Errorf("unsupported Scan, storing driver.Value type %T into type %T", value, d)
}

// Value implements the driver Valuer interface.
func (d Date) Value() (driver.Value, error) {
	return d.Time.Format(structs.DateFormat()), nil
}

func (d *Date) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	// Ignore null, like in the main JSON package.
	if s == "null" {
		d.Time = time.Time{}
		return
	}

	d.Time, err = time.Parse(structs.DateFormat(), s)
	return
}

func (d Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Time.Format(structs.DateFormat()))
}
