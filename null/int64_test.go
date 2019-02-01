package null

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewInt64(t *testing.T) {
	t.Run("success NewInt64", func(t *testing.T) {
		i := int64(1)
		ni, err := NewInt64(i)
		if !assert.NoError(t, err) {
			t.FailNow()
		}
		assert.True(t, ni.Valid)
		assert.Equal(t, i, ni.Int64)
	})
	t.Run("error NewTime", func(t *testing.T) {
		ni, err := NewInt64(false)
		if !assert.Error(t, err) {
			t.FailNow()
		}
		assert.False(t, ni.Valid)
		assert.Equal(t, int64(0), ni.Int64)
	})
}

func BenchmarkNewInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := NewInt64(i)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func TestInt64_Scan(t *testing.T) {
	// enc := &encoder.Encoder{}
	ni, _ := NewInt64(1)

	cases := TestCases{
		"ints": {
			{in: 1, va: int64(1), iv: true, ie: false},
			{in: int8(1), va: int64(1), iv: true, ie: false},
			{in: int16(1), va: int64(1), iv: true, ie: false},
			{in: int32(1), va: int64(1), iv: true, ie: false},
			{in: int64(1), va: int64(1), iv: true, ie: false},
			{in: uint8(1), va: int64(1), iv: true, ie: false},
			{in: uint16(1), va: int64(1), iv: true, ie: false},
			{in: uint32(1), va: int64(1), iv: true, ie: false},
			{in: uint64(1), va: int64(1), iv: true, ie: false},
			{in: 0, va: int64(0), iv: true, ie: false},
			{in: int8(0), va: int64(0), iv: true, ie: false},
			{in: int16(0), va: int64(0), iv: true, ie: false},
			{in: int32(0), va: int64(0), iv: true, ie: false},
			{in: int64(0), va: int64(0), iv: true, ie: false},
			{in: uint8(0), va: int64(0), iv: true, ie: false},
			{in: uint16(0), va: int64(0), iv: true, ie: false},
			{in: uint32(0), va: int64(0), iv: true, ie: false},
			{in: uint64(0), va: int64(0), iv: true, ie: false},
			{in: -1, va: int64(-1), iv: true, ie: false},
			{in: int8(-1), va: int64(-1), iv: true, ie: false},
			{in: int16(-1), va: int64(-1), iv: true, ie: false},
			{in: int32(-1), va: int64(-1), iv: true, ie: false},
			{in: int64(-1), va: int64(-1), iv: true, ie: false},
			{in: ni, va: int64(1), iv: true, ie: false},
		},
		"bytes slice": {
			{in: makeBytes(int64(1)), va: int64(1), iv: true, ie: false},
			{in: makeBytes(int64(0)), va: int64(0), iv: true, ie: false},
			{in: makeBytes(int64(-1)), va: int64(-1), iv: true, ie: false},
		},
		"strings": {
			{in: "1", va: int64(1), iv: true, ie: false},
			{in: "0", va: int64(0), iv: true, ie: false},
			{in: "-1", va: int64(-1), iv: true, ie: false},
		},
		"null": {
			{in: nil, va: int64(0), iv: false, ie: false},
		},
		"errors": {
			{in: true, va: 0, iv: false, ie: true},
		},
	}
	checkCases(cases, t, Int64{}, ni)
}

func BenchmarkInt64_Scan(b *testing.B) {
	var ni Int64
	for i := 0; i < b.N; i++ {
		err := ni.Scan(i)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func TestInt64_Value(t *testing.T) {
	t.Run("Return va", func(t *testing.T) {
		i := int64(1)
		ni, err := NewInt64(i)
		if !assert.NoError(t, err) {
			t.FailNow()
		}
		value, _ := ni.Value()
		assert.Equal(t, i, value)
	})
	t.Run("Return nil va", func(t *testing.T) {
		var ni Int64
		value, _ := ni.Value()
		assert.Nil(t, value)
	})
}

func BenchmarkInt64_Value(b *testing.B) {
	ni, _ := NewInt64(1)
	for i := 0; i < b.N; i++ {
		_, err := ni.Value()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func TestInt64_MarshalJSON(t *testing.T) {
	t.Run("Success marshal", func(t *testing.T) {
		ni, err := NewInt64(1)
		if !assert.NoError(t, err) {
			t.FailNow()
		}
		b, _ := json.Marshal(int64(1))
		jb, err := ni.MarshalJSON()
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		assert.Equal(t, b, jb)
	})

	t.Run("Null result", func(t *testing.T) {
		ni, err := NewInt64(nil)
		if !assert.NoError(t, err) {
			t.FailNow()
		}
		jb, err := ni.MarshalJSON()
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		assert.Equal(t, []byte("null"), jb)
	})
}

func BenchmarkInt64_MarshalJSON(b *testing.B) {
	ni, _ := NewInt64(1)
	for i := 0; i < b.N; i++ {
		_, err := ni.MarshalJSON()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func TestInt64_UnmarshalJSON(t *testing.T) {
	t.Run("Success unmarshal", func(t *testing.T) {
		i := "1"
		var ni Int64
		err := ni.UnmarshalJSON([]byte(i))
		if !assert.NoError(t, err) {
			t.FailNow()
		}
		assert.True(t, ni.Valid)
		assert.Equal(t, int64(1), ni.Int64)
	})
	t.Run("Success unmarshal null", func(t *testing.T) {
		n := "null"
		var ni Int64
		err := ni.UnmarshalJSON([]byte(n))
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		assert.False(t, ni.Valid)
		assert.Equal(t, int64(0), ni.Int64)
	})
	t.Run("Unexpected value case", func(t *testing.T) {
		ti := "1.1"
		var ni Int64
		err := ni.UnmarshalJSON([]byte(ti))
		if !assert.Error(t, err) {
			t.FailNow()
		}
		assert.False(t, ni.Valid)
	})
}

func BenchmarkInt64_UnmarshalJSON(b *testing.B) {
	var ni Int64
	for i := 0; i < b.N; i++ {
		by := makeBytes(int64(i))
		err := ni.UnmarshalJSON(by)
		if err != nil {
			log.Fatal(err)
		}
	}
}
