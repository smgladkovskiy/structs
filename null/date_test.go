package null

import (
	"log"
	"testing"
	"time"

	"github.com/smgladkovskiy/structs"

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

func BenchmarkNewDate(b *testing.B) {
	td := time.Now()
	for i := 0; i < b.N; i++ {
		_, err := NewDate(td.Add(time.Duration(i)))
		if err != nil {
			log.Fatal(err)
		}
	}
}

func TestDate_Scan(t *testing.T) {
	tn := time.Now()
	nd, _ := NewDate(tn)
	cases := TestCases{
		"times": {
			{name: "time", input: tn, expected: tn, isValid: true, isError: false},
			{name: "*time", input: &tn, expected: tn, isValid: true, isError: false},
			{name: "zero time", input: time.Time{}, expected: time.Time{}, isValid: false, isError: false},
			{name: "zero *time", input: &time.Time{}, expected: time.Time{}, isValid: false, isError: false},
			{name: "Date", input: nd, expected: tn, isValid: true, isError: false},
		},
		"string": {
			{name: "string good format", input: tn.Format(structs.DateFormat()), expected: tn.Format(structs.DateFormat()), isValid: true, isError: false},
		},
		"nil": {
			{name: "nil", input: nil, expected: time.Time{}, isValid: false, isError: false},
		},
		"error": {
			{name: "string bad format", input: tn.Format(time.ANSIC), expected: time.Time{}, isValid: false, isError: true},
			{name: "error", input: false, expected: time.Time{}, isValid: false, isError: true},
		},
	}
	checkCases(cases, t, Date{}, tn)
}

func BenchmarkDate_Scan(b *testing.B) {
	var nd Date
	tn := time.Now()
	for i := 0; i < b.N; i++ {
		err := nd.Scan(tn.Add(time.Duration(i)))
		if err != nil {
			log.Fatal(err)
		}
	}
}

func TestDate_Value(t *testing.T) {
	t.Run("Return value", func(t *testing.T) {
		ti := time.Now()
		nd, err := NewDate(ti)
		if !assert.NoError(t, err) {
			t.FailNow()
		}
		value, _ := nd.Value()
		assert.Equal(t, ti, value)
	})
	t.Run("Return nil value", func(t *testing.T) {
		var nd Date
		value, _ := nd.Value()
		assert.Nil(t, value)
	})
}

func BenchmarkDate_Value(b *testing.B) {
	nd, _ := NewDate(time.Now())
	for i := 0; i < b.N; i++ {
		_, err := nd.Value()
		if err != nil {
			log.Fatal(err)
		}
	}
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

func BenchmarkDate_MarshalJSON(b *testing.B) {
	nd, _ := NewDate(time.Now())
	for i := 0; i < b.N; i++ {
		_, err := nd.MarshalJSON()
		if err != nil {
			log.Fatal(err)
		}
	}
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

func BenchmarkDate_UnmarshalJSON(b *testing.B) {
	ts := "2018-07-24"
	bytes := []byte(ts)
	var nd Date
	for i := 0; i < b.N; i++ {
		err := nd.UnmarshalJSON(bytes)
		if err != nil {
			log.Fatal(err)
		}
	}
}
