package zero

import (
	"github.com/smgladkovskiy/structs"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewDate(t *testing.T) {
	t.Run("success NewDate", func(t *testing.T) {
		t.Parallel()
		ts := time.Now()
		nt, err := NewDate(ts)
		if !assert.NoError(t, err) {
			t.FailNow()
		}
		assert.Equal(t, ts, nt.Time)
	})
	t.Run("error NewDate", func(t *testing.T) {
		t.Parallel()
		nt, err := NewDate(false)
		if !assert.Error(t, err) {
			t.FailNow()
		}
		assert.Equal(t, time.Time{}, nt.Time)
	})
}

func TestDate_Scan(t *testing.T) {
	t.Run("Scan Timestamp", func(t *testing.T) {
		t.Parallel()
		var nt Date
		tn := time.Now()
		_ = nt.Scan(tn)
		assert.Equal(t, tn, nt.Time)
	})
	t.Run("Scan zero Timestamp", func(t *testing.T) {
		t.Parallel()
		var nt Date
		tn := time.Time{}
		_ = nt.Scan(tn)
		assert.Equal(t, tn, nt.Time)
	})
	t.Run("Scan String", func(t *testing.T) {
		t.Parallel()
		var nt Date
		tn := time.Now().Format(structs.DateFormat())
		_ = nt.Scan(tn)
		assert.Equal(t, tn, nt.Time.Format(structs.DateFormat()))
	})
	t.Run("Scan String without expected format", func(t *testing.T) {
		t.Parallel()
		var nt Date
		tn := time.Now().Format(time.ANSIC)
		_ = nt.Scan(tn)
		assert.Equal(t, time.Time{}, nt.Time)
	})
	t.Run("Scan nil", func(t *testing.T) {
		t.Parallel()
		var nt Date
		_ = nt.Scan(nil)
		assert.Equal(t, time.Time{}, nt.Time)
	})
	t.Run("Scan error", func(t *testing.T) {
		t.Parallel()
		var nt Date
		err := nt.Scan(false)
		assert.Error(t, err)
	})
}

func TestDate_Value(t *testing.T) {
	t.Run("Return value", func(t *testing.T) {
		t.Parallel()
		ti := time.Now()
		nt, err := NewDate(ti)
		if !assert.NoError(t, err) {
			t.FailNow()
		}
		value, _ := nt.Value()
		assert.Equal(t, ti.Format(structs.DateFormat()), value)
	})
}

func TestDate_MarshalJSON(t *testing.T) {
	t.Run("Success marshal", func(t *testing.T) {
		t.Parallel()
		ti := time.Now()
		timeJson := `"` + ti.Format(structs.DateFormat()) + `"`
		nt := Date{ti}
		jb, err := nt.MarshalJSON()
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		assert.Equal(t, []byte(timeJson), jb)
	})

	// t.Run("Null result", func(t *testing.T) {
	// 	nt := NewDate(nil)
	// 	jb, err := nt.MarshalJSON()
	// 	if !assert.NoError(t, err) {
	// 		t.FailNow()
	// 	}
	//
	// 	assert.Equal(t, []byte("null"), jb)
	// })
}

func TestDate_UnmarshalJSON(t *testing.T) {
	t.Run("Success unmarshal", func(t *testing.T) {
		t.Parallel()
		ti := "2018-07-24"
		pt, _ := time.Parse(structs.DateFormat(), ti)
		nt := Date{pt}
		err := nt.UnmarshalJSON([]byte(ti))
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		assert.Equal(t, nt.Time, pt)
	})

	t.Run("Success unmarshal null", func(t *testing.T) {
		t.Parallel()
		ti := "null"
		pt := time.Time{}
		var timeObject Date
		err := timeObject.UnmarshalJSON([]byte(ti))
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		assert.Equal(t, timeObject.Time, pt)
	})
}
