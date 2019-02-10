package null

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	nullBytes   = []byte{'n', 'u', 'l', 'l'}
	wrongFloat  = []byte{'z', 'e', 'r', 'o'}
	wrongString = "zero"
)

func TestNewFloat64(t *testing.T) {
	t.Run("success NewFloat64", func(t *testing.T) {
		i := float64(1.01)
		nf, err := NewFloat64(i, 2)
		if !assert.NoError(t, err) {
			t.FailNow()
		}
		assert.True(t, nf.Valid)
		assert.Equal(t, i, nf.Float64)
	})
	t.Run("error NewFloat64", func(t *testing.T) {
		nf, err := NewFloat64(wrongFloat, 2)
		if !assert.Error(t, err) {
			t.FailNow()
		}
		assert.False(t, nf.Valid)
		assert.Equal(t, float64(0), nf.Float64)
	})
}

func BenchmarkNewFloat64(b *testing.B) {
	f := 0.3
	for i := 0; i < b.N; i++ {
		f *= float64(i)
		_, err := NewFloat64(f, 1)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func TestFloat64_Scan(t *testing.T) {
	nf, _ := NewFloat64(1.01, 2)
	f32 := float32(1.01)
	f64 := float64(f32)
	cases := TestCases{
		"numbers": {
			{in: 0, va: float64(0), iv: true, ie: false},
			{in: int(1), va: float64(1), iv: true, ie: false},
			{in: int8(1), va: float64(1), iv: true, ie: false},
			{in: int16(1), va: float64(1), iv: true, ie: false},
			{in: int32(1), va: float64(1), iv: true, ie: false},
			{in: int64(1), va: float64(1), iv: true, ie: false},
			{in: uint(1), va: float64(1), iv: true, ie: false},
			{in: uint8(1), va: float64(1), iv: true, ie: false},
			{in: uint16(1), va: float64(1), iv: true, ie: false},
			{in: uint32(1), va: float64(1), iv: true, ie: false},
			{in: uint64(1), va: float64(1), iv: true, ie: false},
			{in: f32, va: f64, iv: true, ie: false},
			{in: float64(1.02), va: float64(1.02), iv: true, ie: false},
			{in: nf, va: float64(1.01), iv: true, ie: false},
		},
		"strings": {
			{in: "1.01", va: float64(1.01), iv: true, ie: false},
			{in: "1.01z", va: nil, iv: false, ie: true},
		},
		"bytes": {
			{in: []byte{'1', '.', '0', '1'}, va: float64(1.01), iv: true, ie: false},
			{in: wrongFloat, va: nil, iv: false, ie: true},
		},
	}

	checkCases(cases, t, Float64{}, nf)
}

func BenchmarkFloat64_Scan(b *testing.B) {
	var nf Float64
	f := 0.3
	for i := 0; i < b.N; i++ {
		f *= float64(i)
		err := nf.Scan(f)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func TestFloat64_Value(t *testing.T) {
	t.Run("Actual value case", func(t *testing.T) {
		i := float64(1)
		nf, err := NewFloat64(i, 2)
		if !assert.NoError(t, err) {
			t.FailNow()
		}
		value, _ := nf.Value()
		assert.Equal(t, i, value)
	})
	t.Run("nil value case", func(t *testing.T) {
		nf, err := NewFloat64(nullBytes, 2)
		if !assert.NoError(t, err) {
			t.FailNow()
		}
		value, _ := nf.Value()
		assert.Nil(t, value)
	})
}

func BenchmarkFloat64_Value(b *testing.B) {
	nf, _ := NewFloat64(1.03, 2)
	for i := 0; i < b.N; i++ {
		_, err := nf.Value()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func TestFloat64_MarshalJSON(t *testing.T) {
	t.Run("success float case", func(*testing.T) {
		val := float64(1.01)
		nf, err := NewFloat64(val, 2)
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		bt, err := nf.MarshalJSON()
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, []byte{'1', '.', '0', '1'}, bt)
	})
	t.Run("success int case", func(*testing.T) {
		val := int(1)
		nf, err := NewFloat64(val, 2)
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		bt, err := nf.MarshalJSON()
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, []byte{'1', '.', '0', '0'}, bt)
	})
	t.Run("success string case", func(*testing.T) {
		val := "1.01"
		nf, err := NewFloat64(val, 2)
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		bt, err := nf.MarshalJSON()
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, []byte{'1', '.', '0', '1'}, bt)
	})
	t.Run("null case", func(*testing.T) {
		nf, err := NewFloat64(nil, 2)
		if !assert.NoError(t, err) {
			t.FailNow()
		}
		bt, err := nf.MarshalJSON()
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, nullBytes, bt)
	})
	t.Run("test precision", func(t *testing.T) {
		nf, err := NewFloat64(1, 3)
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		b, err := nf.MarshalJSON()
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, []byte{'1', '.', '0', '0', '0'}, b)
	})
}

func BenchmarkFloat64_MarshalJSON(b *testing.B) {
	nf, _ := NewFloat64(1.03, 2)
	for i := 0; i < b.N; i++ {
		_, err := nf.MarshalJSON()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func TestFloat64_UnmarshalJSON(t *testing.T) {
	t.Run("unmarshal succeeded", func(t *testing.T) {
		i := "1"
		var nf Float64
		err := nf.UnmarshalJSON([]byte(i))
		if !assert.NoError(t, err) {
			t.FailNow()
		}
		assert.True(t, nf.Valid)
		assert.Equal(t, float64(1), nf.Float64)
	})
	t.Run("null successfully unmarhsaled", func(t *testing.T) {
		var nf Float64
		err := nf.UnmarshalJSON(nullBytes)
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		assert.False(t, nf.Valid)
		assert.Equal(t, float64(0), nf.Float64)
	})
	t.Run("Unexpected value case", func(t *testing.T) {
		var ni Float64
		err := ni.UnmarshalJSON([]byte(wrongString))
		if !assert.Error(t, err) {
			t.FailNow()
		}
		assert.False(t, ni.Valid)
	})
}

func BenchmarkFloat64_UnmarshalJSON(b *testing.B) {
	var nf Float64
	f := 0.3
	for i := 0; i < b.N; i++ {
		by := makeBytes(float64(i) * f)
		err := nf.UnmarshalJSON(by)
		if err != nil {
			log.Fatal(err)
		}
	}
}
