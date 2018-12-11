package null

import (
	"github.com/smgladkovskiy/structs"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewDate(t *testing.T) {
	t.Run("success NewDate", func(t *testing.T) {
		ts := time.Now()
		nd, err := NewDate(ts)
		if !assert.NoError(t, err) {
			t.FailNow()
		}
		assert.True(t, nd.Valid)
		assert.Equal(t, ts, nd.Time)
	})
	t.Run("error NewDate", func(t *testing.T) {
		nd, err := NewDate(false)
		if !assert.Error(t, err) {
			t.FailNow()
		}
		assert.False(t, nd.Valid)
		assert.Equal(t, time.Time{}, nd.Time)
	})
}

func TestDate_Scan(t *testing.T) {
	tn := time.Now()
	nd, _ := NewDate(tn)
	cases := TestCases{
		"times": {
			{na: "time", in: tn, va: tn, iv: true, ie: false},
			{na: "*time", in: &tn, va: tn, iv: true, ie: false},
			{na: "zero time", in: time.Time{}, va: time.Time{}, iv: false, ie: false},
			{na: "zero *time", in: &time.Time{}, va: time.Time{}, iv: false, ie: false},
			{na: "Date", in: nd, va: tn, iv: true, ie: false},
		},
		"string": {
			{na: "string good format", in: tn.Format(structs.DateFormat()), va: tn.Format(structs.DateFormat()), iv: true, ie: false},
		},
		"nil": {
			{na: "nil", in: nil, va: time.Time{}, iv: false, ie: false},
		},
		"error": {
			{na: "string bad format", in: tn.Format(time.ANSIC), va: time.Time{}, iv: false, ie: true},
			{na: "error", in: false, va: time.Time{}, iv: false, ie: true},
		},
	}
	checkCases(cases, t, Date{}, tn)
}

func TestDate_Value(t *testing.T) {
	t.Run("Return va", func(t *testing.T) {
		ti := time.Now()
		nd, err := NewDate(ti)
		if !assert.NoError(t, err) {
			t.FailNow()
		}
		value, _ := nd.Value()
		assert.Equal(t, ti, value)
	})
	t.Run("Return nil va", func(t *testing.T) {
		var nd Date
		value, _ := nd.Value()
		assert.Nil(t, value)
	})
}

func TestDate_MarshalJSON(t *testing.T) {
	t.Run("Success marshal", func(t *testing.T) {
		ti := time.Now()
		timeJson := `"` + ti.Format(structs.DateFormat()) + `"`
		nd, err := NewDate(ti)
		if !assert.NoError(t, err) {
			t.FailNow()
		}
		jb, err := nd.MarshalJSON()
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		assert.Equal(t, []byte(timeJson), jb)
	})

	t.Run("Null result", func(t *testing.T) {
		nd, err := NewDate(nil)
		if !assert.NoError(t, err) {
			t.FailNow()
		}
		jb, err := nd.MarshalJSON()
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		assert.Equal(t, []byte("null"), jb)
	})
}

func TestDate_UnmarshalJSON(t *testing.T) {
	t.Run("Success unmarshal", func(t *testing.T) {
		ti := "2018-07-24"
		pt, _ := time.Parse(structs.DateFormat(), ti)
		var nd Date
		err := nd.UnmarshalJSON([]byte(ti))
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		assert.Equal(t, nd.Time, pt)
	})
	t.Run("Success unmarshal null", func(t *testing.T) {
		ti := "null"
		pt := time.Time{}
		var nd Date
		err := nd.UnmarshalJSON([]byte(ti))
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		assert.Equal(t, nd.Time, pt)
	})
	t.Run("Error wrong format", func(t *testing.T) {
		ti := "2018/07/24"
		pt := time.Time{}
		var nd Date
		err := nd.UnmarshalJSON([]byte(ti))
		if !assert.Error(t, err) {
			t.FailNow()
		}

		assert.Equal(t, nd.Time, pt)
	})
}
