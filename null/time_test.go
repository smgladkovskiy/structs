package null

import (
	"github.com/smgladkovskiy/structs"
	"github.com/smgladkovskiy/structs/zero"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewTime(t *testing.T) {
	t.Run("success NewTime", func(t *testing.T) {
		ts := time.Now()
		nt := NewTime(ts)
		assert.True(t, nt.Valid)
		assert.Equal(t, ts, nt.Time)
	})
	t.Run("error NewTime", func(t *testing.T) {
		nt := NewTime(false)
		assert.False(t, nt.Valid)
		assert.Equal(t, time.Time{}, nt.Time)
	})
}

func TestTime_Scan(t *testing.T) {
	ts := time.Now()
	cases := TestCases{
		"time": {
			{na: "time", in: ts, va: ts, iv: true, ie: false},
			{na: "*time", in: &ts, va: ts, iv: true, ie: false},
			{na: "zero *time", in: &time.Time{}, va: time.Time{}, iv: false, ie: false},
		},
		"zero|null time": {
			{na: "zero time", in: zero.NewTime(time.Time{}), va: time.Time{}, iv: false, ie: false},
			{na: "Time", in: NewTime(ts), va: ts, iv: true, ie: false},
		},
		"strings": {
			{na: "string good format", in: ts.Format(structs.TimeFormat()), va: ts.Format(structs.TimeFormat()), iv: true, ie: false},
		},
		"nil": {
			{na: "nil", in: nil, va: time.Time{}, iv: false, ie: false},
		},
		"errors": {
			{na: "bool as input", in: false, va: false, iv: false, ie: true},
			{na: "bad format", in: ts.Format(time.ANSIC), va: ts, iv: false, ie: true},
		},
	}
	checkCases(cases, t, Time{}, ts)
}

func TestTime_Value(t *testing.T) {
	t.Run("Return va", func(t *testing.T) {
		ti := time.Now()
		nt := NewTime(ti)
		value, _ := nt.Value()
		assert.Equal(t, ti, value)
	})
	t.Run("Return nil va", func(t *testing.T) {
		var nt Time
		value, _ := nt.Value()
		assert.Nil(t, value)
	})
}

func TestTime_MarshalJSON(t *testing.T) {
	t.Run("Success marshal", func(t *testing.T) {
		ti := time.Now()
		timeJson := `"` + ti.Format(structs.TimeFormat()) + `"`
		nt := NewTime(ti)
		jb, err := nt.MarshalJSON()
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		assert.Equal(t, []byte(timeJson), jb)
	})

	t.Run("Null result", func(t *testing.T) {
		nt := NewTime(nil)
		jb, err := nt.MarshalJSON()
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		assert.Equal(t, []byte("null"), jb)
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
