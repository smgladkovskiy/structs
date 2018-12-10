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
		tt := NewTime(ts)
		assert.Equal(t, ts, tt.Time)
	})
	t.Run("error NewTime", func(t *testing.T) {
		tt := NewTime(false)
		assert.Equal(t, time.Time{}, tt.Time)
	})
}

func TestTime_Scan(t *testing.T) {
	ts := time.Now()
	cases := []map[string]interface{}{
		{na: "time", in: ts, va: ts, iv: true, ie: false},
		{na: "*time", in: &ts, va: ts, iv: true, ie: false},
		{na: "zero time", in: time.Time{}, va: time.Time{}, iv: false, ie: false},
		{na: "zero *time", in: &time.Time{}, va: time.Time{}, iv: false, ie: false},
		{na: "string good format", in: ts.Format(structs.TimeFormat()), va: ts.Format(structs.TimeFormat()), iv: true, ie: false},
		{na: "string bad format", in: ts.Format(time.ANSIC), va: time.Time{}, iv: false, ie: true},
		{na: "nil", in: nil, va: time.Time{}, iv: false, ie: false},
		{na: "Time", in: NewTime(ts), va: ts, iv: true, ie: false},
		{na: "error", in: false, va: time.Time{}, iv: false, ie: true},
	}
	for _, testCase := range cases {
		var testTime Time
		err := testTime.Scan(testCase[in])

		if testCase[ie].(bool) {
			assert.Error(t, err)
			break
		}

		switch testCase[in].(type) {
		case string:
			assert.Equal(t, testCase[va], testTime.Time.Format(structs.TimeFormat()), "[%v] value param for intput %+v: %+v", testCase[na], testCase[in], testCase[va])
		case *time.Time:
			if testCase[iv].(bool) {
				assert.Equal(t, testCase[va], ts, "[%v] value param for intput %+v: %+v", testCase[na], testCase[in], testCase[va])
			} else {
				assert.Equal(t, testCase[va], time.Time{}, "[%v] value param for intput %+v: %+v", testCase[na], testCase[in], testCase[va])
			}

		default:
			assert.Equal(t, testCase[va], testTime.Time, "[%v] value param for intput %+v: %+v", testCase[na], testCase[in], testCase[va])
		}
	}
}

func TestTime_Value(t *testing.T) {
	t.Run("Return value", func(t *testing.T) {
		ti := time.Now().UTC()
		nt := NewTime(ti)
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