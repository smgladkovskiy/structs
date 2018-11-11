package null

import (
	"fmt"
	"github.com/smgladkovskiy/structs"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewString(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		str := "Some string"
		ns := NewString(str)
		assert.True(t, ns.Valid)
		assert.Equal(t, str, ns.String)
	})
	t.Run("nil", func(t *testing.T) {
		ns := NewString(nil)
		assert.False(t, ns.Valid)
		assert.Equal(t, "", ns.String)
	})
	t.Run("false", func(t *testing.T) {
		ns := NewString(false)
		assert.False(t, ns.Valid)
		assert.Equal(t, "", ns.String)
	})
}

func TestNewStringf(t *testing.T) {
	f := "Some string by format: %s"
	str := "good"
	ns := NewStringf(f, str)
	assert.True(t, ns.Valid)
	assert.Equal(t, fmt.Sprintf(f, str), ns.String)
}

func TestString_Scan(t *testing.T) {
	t.Run("String", func(t *testing.T) {
		var ns String
		s := "string"
		_ = ns.Scan(s)
		assert.True(t, ns.Valid)
		assert.Equal(t, s, ns.String)
	})
	t.Run("nil", func(t *testing.T) {
		var ns String
		_ = ns.Scan(nil)
		assert.False(t, ns.Valid)
		assert.Equal(t, "", ns.String)
	})
	t.Run("bytes", func(t *testing.T) {
		var ns String
		s := "string"
		b := []byte(s)
		_ = ns.Scan(b)
		assert.True(t, ns.Valid)
		assert.Equal(t, s, ns.String)
	})
	t.Run("nil bytes error", func(t *testing.T) {
		var ns String
		b := []byte(nil)
		_ = ns.Scan(b)
		assert.False(t, ns.Valid)
		assert.Equal(t, "", ns.String)
	})
	t.Run("raw bytes", func(t *testing.T) {
		var ns String
		s := "string"
		b := structs.RawBytes(s)
		_ = ns.Scan(b)
		assert.True(t, ns.Valid)
		assert.Equal(t, s, ns.String)
	})

	t.Run("nil raw bytes error", func(t *testing.T) {
		var ns String
		b := structs.RawBytes(nil)
		_ = ns.Scan(b)
		assert.False(t, ns.Valid)
		assert.Equal(t, "", ns.String)
	})
	t.Run("Time", func(t *testing.T) {
		var ns String
		ti := time.Now()
		_ = ns.Scan(ti)
		assert.True(t, ns.Valid)
		assert.Equal(t, ti.Format(structs.TimeFormat()), ns.String)
	})
	t.Run("Time", func(t *testing.T) {
		var ns String
		nt := NewTime(time.Now())
		_ = ns.Scan(nt)
		assert.True(t, ns.Valid)
		assert.Equal(t, nt.Time.Format(structs.TimeFormat()), ns.String)
	})
	t.Run("String", func(t *testing.T) {
		var ns2 String
		ns1 := String{"string", true}
		_ = ns2.Scan(ns1)
		assert.True(t, ns2.Valid)
		assert.Equal(t, ns1, ns2)
	})
	t.Run("error", func(t *testing.T) {
		var ns String
		err := ns.Scan(false)
		assert.Error(t, err)
	})
}

func TestString_Value(t *testing.T) {
	t.Run("Return va", func(t *testing.T) {
		s := "string"
		ns := NewString(s)
		value, _ := ns.Value()
		assert.Equal(t, s, value)
	})
	t.Run("Return nil va", func(t *testing.T) {
		var ns String
		value, _ := ns.Value()
		assert.Nil(t, value)
	})
}

func TestString_MarshalJSON(t *testing.T) {
	t.Run("Success marshal", func(t *testing.T) {
		s := "string"
		ns := NewString(s)
		jb, err := ns.MarshalJSON()
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		assert.Equal(t, []byte(`"`+s+`"`), jb)
	})

	t.Run("Null result", func(t *testing.T) {
		ns := NewString(nil)
		jb, err := ns.MarshalJSON()
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		assert.Equal(t, []byte("null"), jb)
	})
}

func TestString_UnmarshalJSON(t *testing.T) {
	t.Run("Success unmarshal", func(t *testing.T) {
		pt := "04231"
		b := []byte(pt)
		var ns String
		err := ns.UnmarshalJSON(b)
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		assert.Equal(t, pt, ns.String)
	})
	t.Run("Success unmarshal unicode", func(t *testing.T) {
		us := "Алексей"
		usc := "\"\u0410\u043b\u0435\u043a\u0441\u0435\u0439\""
		b := []byte(usc)
		var ns String
		err := ns.UnmarshalJSON(b)
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		assert.Equal(t, us, ns.String)
	})
	t.Run("Success unmarshal null", func(t *testing.T) {
		s := "null"
		var ns String
		err := ns.UnmarshalJSON([]byte(s))
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		assert.Equal(t, ns.String, "")
	})
}
