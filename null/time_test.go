package null

import (
	"log"
	"testing"
	"time"

	"github.com/smgladkovskiy/structs"
	"github.com/smgladkovskiy/structs/zero"

	"github.com/stretchr/testify/assert"
)

func TestNewTime(t *testing.T) {
	t.Run("success NewTime", func(t *testing.T) {
		ts := time.Now()
		nt, err := NewTime(ts)
		if !assert.NoError(t, err) {
			t.FailNow()
		}
		assert.True(t, nt.Valid)
		assert.Equal(t, ts, nt.Time)
	})
	t.Run("error NewTime", func(t *testing.T) {
		nt, err := NewTime(false)
		if !assert.Error(t, err) {
			t.FailNow()
		}
		assert.False(t, nt.Valid)
		assert.Equal(t, time.Time{}, nt.Time)
	})
}

func BenchmarkNewTime(b *testing.B) {
	tn := time.Now()
	for i := 0; i < b.N; i++ {
		_, err := NewTime(tn.Add(time.Duration(i)))
		if err != nil {
			log.Fatal(err)
		}
	}
}

func TestTime_Scan(t *testing.T) {
	ts := time.Now()
	nt, _ := NewTime(ts)
	ztn, _ := zero.NewTime(time.Time{})
	cases := TestCases{
		"time": {
			{name: "time", input: ts, expected: ts, isValid: true, isError: false},
			{name: "*time", input: &ts, expected: ts, isValid: true, isError: false},
			{name: "empty *time", input: &time.Time{}, expected: time.Time{}, isValid: false, isError: false},
			{name: "zero.Time", input: ztn, expected: time.Time{}, isValid: false, isError: false},
			{name: "null.Time now", input: nt, expected: ts, isValid: true, isError: false},
		},
		"strings": {
			{name: "string good format", input: ts.Format(structs.TimeFormat()), expected: ts.Format(structs.TimeFormat()), isValid: true, isError: false},
		},
		"nil": {
			{name: "nil", input: nil, expected: time.Time{}, isValid: false, isError: false},
		},
		"errors": {
			{name: "bool as input", input: false, expected: false, isValid: false, isError: true},
			{name: "bad format", input: ts.Format(time.ANSIC), expected: ts, isValid: false, isError: true},
		},
	}
	checkCases(cases, t, Time{}, ts)
}

func BenchmarkTime_Scan(b *testing.B) {
	var nt Time
	tn := time.Now()
	for i := 0; i < b.N; i++ {
		err := nt.Scan(tn.Add(time.Duration(i)))
		if err != nil {
			log.Fatal(err)
		}
	}
}

func TestTime_Value(t *testing.T) {
	t.Run("Return value", func(t *testing.T) {
		ti := time.Now()
		nt, err := NewTime(ti)
		if !assert.NoError(t, err) {
			t.FailNow()
		}
		value, _ := nt.Value()
		assert.Equal(t, ti, value)
	})
	t.Run("Return nil value", func(t *testing.T) {
		var nt Time
		value, _ := nt.Value()
		assert.Nil(t, value)
	})
}

func BenchmarkTime_Value(b *testing.B) {
	nt, _ := NewTime(time.Now())
	for i := 0; i < b.N; i++ {
		_, err := nt.Value()
		if err != nil {
			log.Fatal(err)
		}
	}
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

		assert.Equal(t, []byte("null"), jb)
	})
}

func BenchmarkTime_MarshalJSON(b *testing.B) {
	nt, _ := NewTime(time.Now())
	for i := 0; i < b.N; i++ {
		_, err := nt.MarshalJSON()
		if err != nil {
			log.Fatal(err)
		}
	}
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

func BenchmarkTime_UnmarshalJSON(b *testing.B) {
	ts := "2018-07-24T10:09:53+03:00"
	bytes := []byte(ts)
	var nt Time
	for i := 0; i < b.N; i++ {
		err := nt.UnmarshalJSON(bytes)
		if err != nil {
			log.Fatal(err)
		}
	}
}
