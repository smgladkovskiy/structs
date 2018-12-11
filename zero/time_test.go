package zero

import (
	"encoding/json"
	"github.com/smgladkovskiy/structs"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewTime(t *testing.T) {
	t.Run("success NewTime", func(t *testing.T) {
		ts := time.Now()
		tt, err := NewTime(ts)
		if !assert.NoError(t, err) {
			t.FailNow()
		}
		assert.Equal(t, ts, tt.Time)
	})
	t.Run("error NewTime", func(t *testing.T) {
		tt, err := NewTime(false)
		if !assert.Error(t, err) {
			t.FailNow()
		}
		assert.Equal(t, time.Time{}, tt.Time)
	})
}

func TestTime_Scan(t *testing.T) {
	ts := time.Now()
	titn, _ := NewTime(ts)
	cases := TestCases{
		"time": {
			{na: "time", in: ts, va: ts, iv: true, ie: false},
			{na: "*time", in: &ts, va: ts, iv: true, ie: false},
			{na: "zero time", in: time.Time{}, va: time.Time{}, iv: false, ie: false},
			{na: "zero *time", in: &time.Time{}, va: time.Time{}, iv: false, ie: false},
			{na: "string good format", in: ts.Format(structs.TimeFormat()), va: ts.Format(structs.TimeFormat()), iv: true, ie: false},
			{na: "string bad format", in: ts.Format(time.ANSIC), va: time.Time{}, iv: false, ie: true},
			{na: "Time", in: titn, va: ts, iv: true, ie: false},
		},
		"nil": {
			{na: "nil", in: nil, va: time.Time{}, iv: false, ie: false},
		},
		"error": {
			{na: "error", in: false, va: time.Time{}, iv: false, ie: true},
		},
	}
	checkCases(cases, t, Time{}, ts)
}

func TestTime_Value(t *testing.T) {
	t.Run("Return value", func(t *testing.T) {
		ti := time.Now().UTC()
		nt, err := NewTime(ti)
		if !assert.NoError(t, err) {
			t.FailNow()
		}
		value := nt.Time
		assert.Equal(t, ti, value)
	})
	t.Run("Return zero value", func(t *testing.T) {
		var nt Time
		value, _ := nt.Value()
		assert.Equal(t, "0001-01-01T00:00:00Z", value)
	})
}

func TestTime_MarshalJSON(t *testing.T) {
	t.Run("Success marshal", func(t *testing.T) {
		ti := time.Now()
		timeJson := `"` + ti.Format(structs.TimeFormat()) + `"`
		nt, err := NewTime(ti)
		if !assert.NoError(t, err) {
			t.FailNow()
		}
		jb, err := nt.MarshalJSON()
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		assert.Equal(t, []byte(timeJson), jb)
	})

	t.Run("Null result", func(t *testing.T) {
		nt, err := NewTime(nil)
		if !assert.NoError(t, err) {
			t.FailNow()
		}
		jb, err := nt.MarshalJSON()
		if !assert.NoError(t, err) {
			t.FailNow()
		}
		result, _ := json.Marshal("0001-01-01T00:00:00Z")
		assert.Equal(t, result, jb)
	})
}

func TestTime_UnmarshalJSON(t *testing.T) {
	t.Run("Success unmarshal", func(t *testing.T) {
		ti := "2018-07-24T10:09:53+03:00"
		pt, _ := time.Parse(structs.TimeFormat(), ti)
		var nt Time
		err := nt.UnmarshalJSON([]byte(ti))
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		assert.Equal(t, nt.Time, pt)
	})
	t.Run("Success unmarshal null", func(t *testing.T) {
		ti := "null"
		pt := time.Time{}
		var nt Time
		err := nt.UnmarshalJSON([]byte(ti))
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		assert.Equal(t, nt.Time, pt)
	})
	t.Run("Error wrong format", func(t *testing.T) {
		ti := "2018-07-24"
		pt := time.Time{}
		var nt Time
		err := nt.UnmarshalJSON([]byte(ti))
		if !assert.Error(t, err) {
			t.FailNow()
		}

		assert.Equal(t, nt.Time, pt)
	})
}
