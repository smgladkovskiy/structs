package nulls

import (
	"errors"
	"time"
)

var (
	TimeFormat = time.RFC3339
	nullString = []byte("null")
	errNilPtr  = errors.New("destination pointer is nil") // embedded in descriptive error
)

type RawBytes []byte
