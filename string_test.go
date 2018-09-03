package structs

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewNullString(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		str := "Some string"
		ns := NewNullString(str)
		assert.True(t, ns.Valid)
		assert.Equal(t, str, ns.String)
	})
	t.Run("nil", func(t *testing.T) {
		ns := NewNullString(nil)
		assert.False(t, ns.Valid)
		assert.Equal(t, "", ns.String)
	})
	t.Run("false", func(t *testing.T) {
		ns := NewNullString(false)
		assert.False(t, ns.Valid)
		assert.Equal(t, "", ns.String)
	})
}

func TestNewNullStringf(t *testing.T) {
	f := "Some string by format: %s"
	str := "good"
	ns := NewNullStringf(f, str)
	assert.True(t, ns.Valid)
	assert.Equal(t, fmt.Sprintf(f, str), ns.String)
}

func TestNullString_Scan(t *testing.T) {
	t.Run("String", func(t *testing.T) {
		var ns NullString
		s := "string"
		ns.Scan(s)
		assert.True(t, ns.Valid)
		assert.Equal(t, s, ns.String)
	})
	t.Run("nil", func(t *testing.T) {
		var ns NullString
		ns.Scan(nil)
		assert.False(t, ns.Valid)
		assert.Equal(t, "", ns.String)
	})
	t.Run("bytes", func(t *testing.T) {
		var ns NullString
		s := "string"
		b := []byte(s)
		ns.Scan(b)
		assert.True(t, ns.Valid)
		assert.Equal(t, s, ns.String)
	})
	t.Run("nil bytes error", func(t *testing.T) {
		var ns NullString
		b := []byte(nil)
		ns.Scan(b)
		assert.False(t, ns.Valid)
		assert.Equal(t, "", ns.String)
	})
	t.Run("raw bytes", func(t *testing.T) {
		var ns NullString
		s := "string"
		b := RawBytes(s)
		ns.Scan(b)
		assert.True(t, ns.Valid)
		assert.Equal(t, s, ns.String)
	})

	t.Run("nil raw bytes error", func(t *testing.T) {
		var ns NullString
		b := RawBytes(nil)
		ns.Scan(b)
		assert.False(t, ns.Valid)
		assert.Equal(t, "", ns.String)
	})
	t.Run("Time", func(t *testing.T) {
		var ns NullString
		ti := time.Now()
		ns.Scan(ti)
		assert.True(t, ns.Valid)
		assert.Equal(t, ti.Format(TimeFormat), ns.String)
	})
	t.Run("NullTime", func(t *testing.T) {
		var ns NullString
		nt := NewNullTime(time.Now())
		ns.Scan(nt)
		assert.True(t, ns.Valid)
		assert.Equal(t, nt.Time.Format(TimeFormat), ns.String)
	})
	t.Run("NullString", func(t *testing.T) {
		var ns2 NullString
		ns1 := NullString{"string", true}
		ns2.Scan(ns1)
		assert.True(t, ns2.Valid)
		assert.Equal(t, ns1, ns2)
	})
	t.Run("error", func(t *testing.T) {
		var ns NullString
		err := ns.Scan(false)
		assert.Error(t, err)
	})
}

func TestNullString_Value(t *testing.T) {
	t.Run("Return value", func(t *testing.T) {
		s := "string"
		ns := NewNullString(s)
		value, _ := ns.Value()
		assert.Equal(t, s, value)
	})
	t.Run("Return nil value", func(t *testing.T) {
		var ns NullString
		value, _ := ns.Value()
		assert.Nil(t, value)
	})
}

func TestNullString_MarshalJSON(t *testing.T) {
	t.Run("Success marshal", func(t *testing.T) {
		s := "string"
		ns := NewNullString(s)
		jb, err := ns.MarshalJSON()
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		assert.Equal(t, []byte(`"`+s+`"`), jb)
	})

	t.Run("Null result", func(t *testing.T) {
		ns := NewNullString(nil)
		jb, err := ns.MarshalJSON()
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		assert.Equal(t, []byte("null"), jb)
	})
}

func TestNullString_UnmarshalJSON(t *testing.T) {
	t.Run("Success unmarshal", func(t *testing.T) {
		pt := "04231"
		b := []byte(pt)
		var ns NullString
		err := ns.UnmarshalJSON(b)
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		assert.Equal(t, pt, ns.String)
	})
	t.Run("Success unmarshal null", func(t *testing.T) {
		s := "null"
		var ns NullString
		err := ns.UnmarshalJSON([]byte(s))
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		assert.Equal(t, ns.String, "")
	})
}
