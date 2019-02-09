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
		if err != nil {
			log.Fatal(err)
		}

		assert.True(t, nf.Valid)
		assert.Equal(t, i, nf.Float64)
	})
	t.Run("error NewFloat64", func(t *testing.T) {
		nf, err := NewFloat64(wrongFloat, 2)
		if err != nil {
			log.Fatal(err)
		}

		assert.False(t, nf.Valid)
		assert.Equal(t, float64(0), nf.Float64)
	})
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

func TestFloat64_Scan(t *testing.T) {
	t.Run("valid string case", func(*testing.T) {
		var nf Float64
		s := "1.01"
		if err := nf.Scan(s); err != nil {
			t.Error(err)
		}
		assert.True(t, nf.Valid)
		assert.Equal(t, float64(1.01), nf.Float64)
	})
	t.Run("invalid string case", func(*testing.T) {
		var nf Float64
		s := "1.01z"
		if err := nf.Scan(s); err != nil {
			t.Error(err)
		}
		assert.False(t, nf.Valid)
	})
	t.Run("valid []byte case", func(*testing.T) {
		var nf Float64
		s := []byte{'1', '.', '0', '1'}
		if err := nf.Scan(s); err != nil {
			t.Error(err)
		}
		assert.True(t, nf.Valid)
		assert.Equal(t, float64(1.01), nf.Float64)
	})
	t.Run("invalid []byte case", func(*testing.T) {
		var nf Float64
		if err := nf.Scan(wrongFloat); err != nil {
			t.Error(err)
		}
		assert.False(t, nf.Valid)
	})
	t.Run("zero value case", func(*testing.T) {
		var nf Float64
		i := 0
		if err := nf.Scan(i); err != nil {
			t.Error(err)
		}
		assert.False(t, nf.Valid)
		assert.Equal(t, float64(0), nf.Float64)
	})
	t.Run("int case", func(*testing.T) {
		var nf Float64
		i := int(1)
		if err := nf.Scan(i); err != nil {
			t.Error(err)
		}
		assert.True(t, nf.Valid)
		assert.Equal(t, float64(i), nf.Float64)
	})
	t.Run("int8 case", func(*testing.T) {
		var nf Float64
		i := int8(1)
		if err := nf.Scan(i); err != nil {
			t.Error(err)
		}
		assert.True(t, nf.Valid)
		assert.Equal(t, float64(i), nf.Float64)
	})
	t.Run("int16 case", func(*testing.T) {
		var nf Float64
		i := int16(1)
		if err := nf.Scan(i); err != nil {
			t.Error(err)
		}
		assert.True(t, nf.Valid)
		assert.Equal(t, float64(i), nf.Float64)
	})
	t.Run("int32 case", func(*testing.T) {
		var nf Float64
		i := int32(1)
		if err := nf.Scan(i); err != nil {
			t.Error(err)
		}
		assert.True(t, nf.Valid)
		assert.Equal(t, float64(i), nf.Float64)
	})
	t.Run("int64 case", func(*testing.T) {
		var nf Float64
		i := int64(1)
		if err := nf.Scan(i); err != nil {
			t.Error(err)
		}
		assert.True(t, nf.Valid)
		assert.Equal(t, float64(i), nf.Float64)
	})
	t.Run("uint case", func(*testing.T) {
		var nf Float64
		i := uint(1)
		if err := nf.Scan(i); err != nil {
			t.Error(err)
		}
		assert.True(t, nf.Valid)
		assert.Equal(t, float64(i), nf.Float64)
	})
	t.Run("uint8 case", func(*testing.T) {
		var nf Float64
		i := uint8(1)
		if err := nf.Scan(i); err != nil {
			t.Error(err)
		}
		assert.True(t, nf.Valid)
		assert.Equal(t, float64(i), nf.Float64)
	})
	t.Run("uint16 case", func(*testing.T) {
		var nf Float64
		i := uint16(1)
		if err := nf.Scan(i); err != nil {
			t.Error(err)
		}
		assert.True(t, nf.Valid)
		assert.Equal(t, float64(i), nf.Float64)
	})
	t.Run("uint32 case", func(*testing.T) {
		var nf Float64
		i := uint32(1)
		if err := nf.Scan(i); err != nil {
			t.Error(err)
		}
		assert.True(t, nf.Valid)
		assert.Equal(t, float64(i), nf.Float64)
	})
	t.Run("uint64 case", func(*testing.T) {
		var nf Float64
		i := uint64(1)
		if err := nf.Scan(i); err != nil {
			t.Error(err)
		}
		assert.True(t, nf.Valid)
		assert.Equal(t, float64(i), nf.Float64)
	})
	t.Run("float32 case", func(*testing.T) {
		var nf Float64
		i := float32(1.01)
		if err := nf.Scan(i); err != nil {
			t.Error(err)
		}
		assert.True(t, nf.Valid)
		assert.Equal(t, float64(i), nf.Float64)
	})
	t.Run("float64 case", func(*testing.T) {
		var nf Float64
		i := float64(1.01)
		if err := nf.Scan(i); err != nil {
			t.Error(err)
		}
		assert.True(t, nf.Valid)
		assert.Equal(t, float64(i), nf.Float64)
	})
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
		val := int(0)
		nf, err := NewFloat64(val, 2)
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
