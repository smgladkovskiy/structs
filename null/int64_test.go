package null

import (
	"encoding/binary"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestNewInt64(t *testing.T) {
	t.Run("success NewInt64", func(t *testing.T) {
		i := int64(1)
		ni := NewInt64(i)
		assert.True(t, ni.Valid)
		assert.Equal(t, i, ni.Int64)
	})
	t.Run("error NewTime", func(t *testing.T) {
		ni := NewInt64(false)
		assert.False(t, ni.Valid)
		assert.Equal(t, int64(0), ni.Int64)
	})
}

func TestInt64_Value(t *testing.T) {
	t.Run("Return va", func(t *testing.T) {
		i := int64(1)
		ni := NewInt64(i)
		value, _ := ni.Value()
		assert.Equal(t, i, value)
	})
	t.Run("Return nil va", func(t *testing.T) {
		var ni Int64
		value, _ := ni.Value()
		assert.Nil(t, value)
	})
}

func TestInt64_Scan(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		var ni Int64
		i := 1
		_ = ni.Scan(i)
		assert.True(t, ni.Valid)
		assert.Equal(t, int64(i), ni.Int64)
	})
	t.Run("int32", func(t *testing.T) {
		var ni Int64
		i := int32(1)
		_ = ni.Scan(i)
		assert.True(t, ni.Valid)
		assert.Equal(t, int64(i), ni.Int64)
	})
	t.Run("int64", func(t *testing.T) {
		var ni Int64
		i := int64(1)
		_ = ni.Scan(i)
		assert.True(t, ni.Valid)
		assert.Equal(t, i, ni.Int64)
	})
	t.Run("zero int", func(t *testing.T) {
		var ni Int64
		i := 0
		_ = ni.Scan(i)
		assert.False(t, ni.Valid)
		assert.Equal(t, int64(i), ni.Int64)
	})
	t.Run("string", func(t *testing.T) {
		var ni Int64
		si := "1"
		i, _ := strconv.ParseInt(si, 10, 64)
		_ = ni.Scan(si)
		assert.True(t, ni.Valid)
		assert.Equal(t, i, ni.Int64)
	})
	t.Run("Int64", func(t *testing.T) {
		var ni Int64
		ni2 := *NewInt64(1)
		_ = ni.Scan(ni2)
		assert.True(t, ni.Valid)
		assert.Equal(t, ni2, ni)
	})
	t.Run("byte", func(t *testing.T) {
		var ni Int64
		i := int64(1)
		bi := make([]byte, 8)
		binary.BigEndian.PutUint64(bi, uint64(i))
		_ = ni.Scan(bi)
		assert.True(t, ni.Valid)
		assert.Equal(t, i, ni.Int64)
	})
	t.Run("nil", func(t *testing.T) {
		var ni Int64
		_ = ni.Scan(nil)
		assert.False(t, ni.Valid)
		assert.Equal(t, int64(0), ni.Int64)
	})
}

func TestInt64_MarshalJSON(t *testing.T) {
	t.Run("Success marshal", func(t *testing.T) {
		ni := NewInt64(1)
		b, _ := json.Marshal(1)
		jb, err := ni.MarshalJSON()
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		assert.Equal(t, b, jb)
	})

	t.Run("Null result", func(t *testing.T) {
		ni := NewInt64(nil)
		jb, err := ni.MarshalJSON()
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		assert.Equal(t, []byte("null"), jb)
	})
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
