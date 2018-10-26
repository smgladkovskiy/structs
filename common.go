package structs

import (
	"errors"
	"time"
)

type RawBytes []byte

var (
	TimeFormat = func() string {
		return time.RFC3339
	}
	DateFormat = func() string {
		return "2006-01-02"
	}
	NullString = []byte("null")
	ErrNilPtr  = errors.New("destination pointer is nil") // embedded in descriptive error
)
