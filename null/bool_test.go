package null

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBool(t *testing.T) {
	t.Parallel()
	t.Run("Success", func(t *testing.T) {
		nullBool := NewBool(true)
		assert.True(t, nullBool.Valid)
		assert.Equal(t, true, nullBool.Bool)
	})

	t.Run("False on nil", func(t *testing.T) {
		nullBool := NewBool(nil)
		assert.False(t, nullBool.Valid)
		assert.Equal(t, false, nullBool.Bool)
	})
}

func TestBool_Value(t *testing.T) {
	t.Parallel()
	t.Run("Return bool va", func(t *testing.T) {
		nullBool := NewBool(true)
		value, err := nullBool.Value()
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		assert.Equal(t, true, value)
	})
	t.Run("Return nil va", func(t *testing.T) {
		var nullBool Bool
		value, _ := nullBool.Value()

		assert.Nil(t, value)
	})
}

func TestBool_Scan(t *testing.T) {
	cases := []map[string]interface{}{
		// ints
		{in: 1, va: true, iv: true},
		{in: int8(1), va: true, iv: true},
		{in: int16(1), va: true, iv: true},
		{in: int32(1), va: true, iv: true},
		{in: int64(1), va: true, iv: true},

		{in: 0, va: false, iv: true},
		{in: int8(0), va: false, iv: true},
		{in: int16(0), va: false, iv: true},
		{in: int32(0), va: false, iv: true},
		{in: int64(0), va: false, iv: true},

		{in: 5, va: false, iv: false},
		{in: -5, va: false, iv: false},

		// strings
		{in: "1", va: true, iv: true},
		{in: "t", va: true, iv: true},
		{in: "T", va: true, iv: true},
		{in: "true", va: true, iv: true},
		{in: "TRUE", va: true, iv: true},
		{in: "True", va: true, iv: true},
		{in: "y", va: true, iv: true},
		{in: "Y", va: true, iv: true},
		{in: "YES", va: true, iv: true},
		{in: "Yes", va: true, iv: true},

		{in: "0", va: false, iv: true},
		{in: "f", va: false, iv: true},
		{in: "F", va: false, iv: true},
		{in: "false", va: false, iv: true},
		{in: "False", va: false, iv: true},
		{in: "FALSE", va: false, iv: true},
		{in: "na", va: false, iv: true},
		{in: "N", va: false, iv: true},
		{in: "NO", va: false, iv: true},
		{in: "No", va: false, iv: true},

		{in: "some string", va: false, iv: false},

		// Bool
		{in: *NewBool(true), va: true, iv: true},
		{in: *NewBool(false), va: false, iv: true},
		{in: *NewBool(nil), va: false, iv: false},

		// []byte
		{in: makeBytes(true), va: true, iv: true},
		{in: makeBytes(false), va: false, iv: true},
		{in: makeBytes(nil), va: false, iv: false},

		// nil
		{in: nil, va: false, iv: false},
	}
	for _, testCase := range cases {
		var nullBool Bool
		_ = nullBool.Scan(testCase[in])
		assert.Equal(t, testCase[va], nullBool.Bool, "va param for intput %+v: %+v", testCase[in], testCase[va])
		assert.Equal(t, testCase[iv], nullBool.Valid, "iv param for intput %+v: %+v", testCase[in], testCase[iv])
	}
}
func TestBool_MarshalJSON(t *testing.T) {
	t.Run("Success marshal", func(t *testing.T) {
		nullBool := NewBool(true)
		b, _ := json.Marshal(true)
		jb, err := nullBool.MarshalJSON()
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		assert.Equal(t, b, jb)
	})

	t.Run("Null result", func(t *testing.T) {
		ni := NewBool(nil)
		jb, err := ni.MarshalJSON()
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		assert.Equal(t, []byte("null"), jb)
	})
}

func TestBool_UnmarshalJSON(t *testing.T) {
	cases := []map[string]interface{}{
		{in: []byte("true"), va: true, iv: true},
		{in: []byte("1"), va: true, iv: true},
		{in: []byte("t"), va: true, iv: true},
		{in: []byte("T"), va: true, iv: true},
		{in: []byte("TRUE"), va: true, iv: true},
		{in: []byte("True"), va: true, iv: true},

		{in: []byte("0"), va: false, iv: true},
		{in: []byte("f"), va: false, iv: true},
		{in: []byte("F"), va: false, iv: true},
		{in: []byte("false"), va: false, iv: true},
		{in: []byte("False"), va: false, iv: true},
		{in: []byte("FALSE"), va: false, iv: true},

		{in: []byte("not bool"), va: false, iv: false},
		{in: []byte("null"), va: false, iv: false},
	}

	for _, testCase := range cases {
		var nullBool Bool
		err := nullBool.UnmarshalJSON(testCase[in].([]byte))
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		assert.Equal(t, testCase[va], nullBool.Bool, "va param for intput %+v: %+v", testCase[in], testCase[va])
		assert.Equal(t, testCase[iv], nullBool.Valid, "iv param for intput %+v: %+v", testCase[in], testCase[iv])
	}
}

func makeBytes(v interface{}) []byte {
	bytes, _ := json.Marshal(v)
	return bytes
}
