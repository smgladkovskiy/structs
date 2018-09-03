package structs

import (
	"github.com/pkg/errors"
	"time"
)

var (
	TimeFormat = time.RFC3339
	nullString = []byte("null")
	errNilPtr  = errors.New("destination pointer is nil") // embedded in descriptive error
)

type RawBytes []byte

// Model realisation
type Model struct {
	ID        int64 `gorm:"PRIMARY_KEY;AUTO_INCREMENT" json:"id"`
	CreatedAt Time  `json:"created_at"`
	UpdatedAt Time  `json:"updated_at"`
}
