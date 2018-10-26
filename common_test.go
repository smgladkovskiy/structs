package structs

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestDateFormat(t *testing.T) {
	expectedDateFormat := "02.01.2006"
	t.Run("success", func(t *testing.T) {
		var DateFormat = func() string {
			return expectedDateFormat
		}

		assert.Equal(t, expectedDateFormat, DateFormat())
	})
}

func TestTimeFormat(t *testing.T) {
	expectedTimeFormat := time.RFC1123
	t.Run("success", func(t *testing.T) {
		var TimeFormat = func() string {
			return expectedTimeFormat
		}

		assert.Equal(t, expectedTimeFormat, TimeFormat())
	})
}
