package null

import (
	"database/sql/driver"
	"errors"
	"strings"
	"time"

	"github.com/smgladkovskiy/structs"
)

type Date struct {
	Time  time.Time
	Valid bool // is true if Time is not NULL
}

// NewDate Создание Date переменной
func NewDate(v interface{}) (*Date, error) {
	var nt Date
	err := nt.Scan(v)
	return &nt, err
}

// Scan implements the Scanner interface for Date
func (nd *Date) Scan(value interface{}) error {
	switch v := value.(type) {
	case nil:
		return nil
	case string:
		var err error
		nd.Time, err = time.Parse(structs.DateFormat(), v)
		nd.Valid = err == nil
		return err
	case time.Time:
		if v.IsZero() {
			return nil
		}
		nd.Time, nd.Valid = v, true
		return nil
	case *time.Time:
		if v.IsZero() {
			return nil
		}
		nd.Time, nd.Valid = *v, true
		return nil
	case Date:
		nd.Time, nd.Valid = v.Time, v.Valid
		return nil
	case *Date:
		nd.Time, nd.Valid = v.Time, v.Valid
		return nil
	}

	return structs.TypeIsNotAcceptable{CheckedValue: value, CheckedType: nd}
}

// Value implements the driver Valuer interface.
func (nd Date) Value() (driver.Value, error) {
	if !nd.Valid {
		return nil, nil
	}
	return nd.Time, nil
}

func (nd Date) MarshalJSON() ([]byte, error) {
	if !nd.Valid {
		return structs.NullString, nil
	}
	if y := nd.Time.Year(); y < 0 || y >= 10000 {
		// RFC 3339 is clear that years are 4 digits exactly.
		// See golang.org/issue/4556#c15 for more discussion.
		return nil, errors.New("Time.MarshalJSON: year outside of range [0,9999]")
	}

	b := make([]byte, 0, len(structs.DateFormat())+2)
	b = append(b, '"')
	b = nd.Time.AppendFormat(b, structs.DateFormat())
	b = append(b, '"')
	return b, nil
}

func (nd *Date) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		nd.Valid = false
		return
	}
	nd.Time, err = time.Parse(structs.DateFormat(), s)
	nd.Valid = err == nil
	return
}
