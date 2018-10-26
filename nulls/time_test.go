package nulls

import (
	"gitlab.teamc.io/teamc.io/golang/structs"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewNullTime(t *testing.T) {
	t.Run("success NewNullTime", func(t *testing.T) {
		ts := time.Now()
		nt := NewNullTime(ts)
		assert.True(t, nt.Valid)
		assert.Equal(t, ts, nt.Time)
	})
	t.Run("error NewNullTime", func(t *testing.T) {
		nt := NewNullTime(false)
		assert.False(t, nt.Valid)
		assert.Equal(t, time.Time{}, nt.Time)
	})
}

func TestNullTime_Scan(t *testing.T) {
	ts := time.Now()
	cases := []map[string]interface{}{
		{"name": "time", "input": ts, "value": ts, "valid": true, "err": false},
		{"name": "*time", "input": &ts, "value": ts, "valid": true, "err": false},
		{"name": "zero time", "input": time.Time{}, "value": time.Time{}, "valid": false, "err": false},
		{"name": "zero *time", "input": &time.Time{}, "value": time.Time{}, "valid": false, "err": false},
		{"name": "string good format", "input": ts.Format(structs.TimeFormat()), "value": ts.Format(structs.TimeFormat()), "valid": true, "err": false},
		{"name": "string bad format", "input": ts.Format(time.ANSIC), "value": time.Time{}, "valid": false, "err": true},
		{"name": "nil", "input": nil, "value": time.Time{}, "valid": false, "err": false},
		{"name": "NullTime", "input": NewNullTime(ts), "value": ts, "valid": true, "err": false},
		{"name": "error", "input": false, "value": time.Time{}, "valid": false, "err": true},
	}
	for _, testCase := range cases {
		var nullTime NullTime
		err := nullTime.Scan(testCase["input"])

		if testCase["err"].(bool) {
			assert.Error(t, err)
			break
		}

		switch testCase["input"].(type) {
		case string:
			assert.Equal(t, testCase["value"], nullTime.Time.Format(structs.TimeFormat()), "[%v] value param for intput %+v: %+v", testCase["name"], testCase["input"], testCase["value"])
		case *time.Time:
			if testCase["valid"].(bool) {
				assert.Equal(t, testCase["value"], ts, "[%v] value param for intput %+v: %+v", testCase["name"], testCase["input"], testCase["value"])
			} else {
				assert.Equal(t, testCase["value"], time.Time{}, "[%v] value param for intput %+v: %+v", testCase["name"], testCase["input"], testCase["value"])
			}

		default:
			assert.Equal(t, testCase["value"], nullTime.Time, "[%v] value param for intput %+v: %+v", testCase["name"], testCase["input"], testCase["value"])
		}

		assert.Equal(t, testCase["valid"], nullTime.Valid, "[%v] valid param for intput %+v: %+v", testCase["name"], testCase["input"], testCase["valid"])
	}
}

func TestNullTime_Value(t *testing.T) {
	t.Run("Return value", func(t *testing.T) {
		ti := time.Now()
		nt := NewNullTime(ti)
		value, _ := nt.Value()
		assert.Equal(t, ti, value)
	})
	t.Run("Return nil value", func(t *testing.T) {
		var nt NullTime
		value, _ := nt.Value()
		assert.Nil(t, value)
	})
}

func TestNullTime_MarshalJSON(t *testing.T) {
	t.Run("Success marshal", func(t *testing.T) {
		ti := time.Now()
		timeJson := `"` + ti.Format(structs.TimeFormat()) + `"`
		nt := NewNullTime(ti)
		jb, err := nt.MarshalJSON()
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		assert.Equal(t, []byte(timeJson), jb)
	})

	t.Run("Null result", func(t *testing.T) {
		nt := NewNullTime(nil)
		jb, err := nt.MarshalJSON()
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		assert.Equal(t, []byte("null"), jb)
	})
}

func TestNullTime_UnmarshalJSON(t *testing.T) {
	t.Run("Success unmarshal", func(t *testing.T) {
		ti := "2018-07-24T10:09:53+03:00"
		pt, _ := time.Parse(structs.TimeFormat(), ti)
		var nt NullTime
		err := nt.UnmarshalJSON([]byte(ti))
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		assert.Equal(t, nt.Time, pt)
	})
	t.Run("Success unmarshal null", func(t *testing.T) {
		ti := "null"
		pt := time.Time{}
		var nt NullTime
		err := nt.UnmarshalJSON([]byte(ti))
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		assert.Equal(t, nt.Time, pt)
	})
	t.Run("Error wrong format", func(t *testing.T) {
		ti := "2018-07-24"
		pt := time.Time{}
		var nt NullTime
		err := nt.UnmarshalJSON([]byte(ti))
		if !assert.Error(t, err) {
			t.FailNow()
		}

		assert.Equal(t, nt.Time, pt)
	})
}
