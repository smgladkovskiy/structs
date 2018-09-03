package nulls_git

import (
	"errors"
	"time"
)

var (
	TimeFormat = func() string {
		return time.RFC3339
	}
	nullString = []byte("null")
	errNilPtr  = errors.New("destination pointer is nil") // embedded in descriptive error
)

type RawBytes []byte
