package null

import (
	"database/sql/driver"
	"encoding/json"
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
func (nd *Date) Scan(value interface{}) error {
	switch v := value.(type) {
	case Date:
		*nd = v
		return nil
	case nil:
		*nd = Date{Time: time.Time{}, Valid: false}
		return nil
	case string:
		t, err := time.Parse(structs.DateFormat(), v)
		if err != nil {
			*nd = Date{Time: time.Time{}, Valid: false}
			return err
		}
		*nd = Date{Time: t, Valid: true}
		return nil
	case time.Time:
		if v.IsZero() {
			*nd = Date{Time: time.Time{}, Valid: false}
			return nil
		}

		*nd = Date{Time: v, Valid: true}

		return nil
	case *time.Time:
		if v.IsZero() {
			*nd = Date{Time: time.Time{}, Valid: false}
			return nil
		}

		*nd = Date{Time: *v, Valid: true}

		return nil
	}

	return structs.TypeIsNotAcceptable{CheckedValue: value, CheckedType: nd}
}

// va implements the driver Valuer interface.
func (nd Date) Value() (driver.Value, error) {
	if !nd.Valid {
		return nil, nil
	}
	return nd.Time, nil
}

func (nd *Date) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		nd.Time = time.Time{}
		return
	}
	nd.Time, err = time.Parse(structs.DateFormat(), s)
	if err == nil {
		nd.Valid = true
	}
	return
}

func (nd Date) MarshalJSON() ([]byte, error) {
	if nd.Valid {
		return json.Marshal(nd.Time.Format(structs.DateFormat()))
	}

	return structs.NullString, nil
}
