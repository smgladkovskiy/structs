package nulls

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewNullBool(t *testing.T) {
	t.Parallel()
	t.Run("Success", func(t *testing.T) {
		nullBool := NewNullBool(true)
		assert.True(t, nullBool.Valid)
		assert.Equal(t, true, nullBool.Bool)
	})

	t.Run("False on nil", func(t *testing.T) {
		nullBool := NewNullBool(nil)
		assert.False(t, nullBool.Valid)
		assert.Equal(t, false, nullBool.Bool)
	})
}

func TestNullBool_Value(t *testing.T) {
	t.Parallel()
	t.Run("Return bool value", func(t *testing.T) {
		nullBool := NewNullBool(true)
		value, err := nullBool.Value()
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		assert.Equal(t, true, value)
	})
	t.Run("Return nil value", func(t *testing.T) {
		var nullBool NullBool
		value, _ := nullBool.Value()

		assert.Nil(t, value)
	})
}

func TestNullBool_Scan(t *testing.T) {
	cases := []map[string]interface{}{
		// ints
		{"input": 1, "value": true, "valid": true},
		{"input": int8(1), "value": true, "valid": true},
		{"input": int16(1), "value": true, "valid": true},
		{"input": int32(1), "value": true, "valid": true},
		{"input": int64(1), "value": true, "valid": true},

		{"input": 0, "value": false, "valid": true},
		{"input": int8(0), "value": false, "valid": true},
		{"input": int16(0), "value": false, "valid": true},
		{"input": int32(0), "value": false, "valid": true},
		{"input": int64(0), "value": false, "valid": true},

		{"input": 5, "value": false, "valid": false},
		{"input": -5, "value": false, "valid": false},

		// strings
		{"input": "1", "value": true, "valid": true},
		{"input": "t", "value": true, "valid": true},
		{"input": "T", "value": true, "valid": true},
		{"input": "true", "value": true, "valid": true},
		{"input": "TRUE", "value": true, "valid": true},
		{"input": "True", "value": true, "valid": true},

		{"input": "0", "value": false, "valid": true},
		{"input": "f", "value": false, "valid": true},
		{"input": "F", "value": false, "valid": true},
		{"input": "false", "value": false, "valid": true},
		{"input": "False", "value": false, "valid": true},
		{"input": "FALSE", "value": false, "valid": true},

		{"input": "some string", "value": false, "valid": false},

		// NullBool
		{"input": *NewNullBool(true), "value": true, "valid": true},
		{"input": *NewNullBool(false), "value": false, "valid": true},
		{"input": *NewNullBool(nil), "value": false, "valid": false},

		// []byte
		{"input": makeBytes(true), "value": true, "valid": true},
		{"input": makeBytes(false), "value": false, "valid": true},
		{"input": makeBytes(nil), "value": false, "valid": false},

		// nil
		{"input": nil, "value": false, "valid": false},
	}
	for _, testCase := range cases {
		var nullBool NullBool
		nullBool.Scan(testCase["input"])
		assert.Equal(t, testCase["value"], nullBool.Bool, "value param for intput %+v: %+v", testCase["input"], testCase["value"])
		assert.Equal(t, testCase["valid"], nullBool.Valid, "valid param for intput %+v: %+v", testCase["input"], testCase["valid"])
	}
}
func TestNullBool_MarshalJSON(t *testing.T) {
	t.Run("Success marshal", func(t *testing.T) {
		nullBool := NewNullBool(true)
		b, _ := json.Marshal(true)
		jb, err := nullBool.MarshalJSON()
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		assert.Equal(t, b, jb)
	})

	t.Run("Null result", func(t *testing.T) {
		ni := NewNullBool(nil)
		jb, err := ni.MarshalJSON()
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		assert.Equal(t, []byte("null"), jb)
	})
}

func TestNullBool_UnmarshalJSON(t *testing.T) {
	cases := []map[string]interface{}{
		{"input": []byte("true"), "value": true, "valid": true},
		{"input": []byte("1"), "value": true, "valid": true},
		{"input": []byte("t"), "value": true, "valid": true},
		{"input": []byte("T"), "value": true, "valid": true},
		{"input": []byte("TRUE"), "value": true, "valid": true},
		{"input": []byte("True"), "value": true, "valid": true},

		{"input": []byte("0"), "value": false, "valid": true},
		{"input": []byte("f"), "value": false, "valid": true},
		{"input": []byte("F"), "value": false, "valid": true},
		{"input": []byte("false"), "value": false, "valid": true},
		{"input": []byte("False"), "value": false, "valid": true},
		{"input": []byte("FALSE"), "value": false, "valid": true},

		{"input": []byte("not bool"), "value": false, "valid": false},
		{"input": []byte("null"), "value": false, "valid": false},
	}

	for _, testCase := range cases {
		var nullBool NullBool
		err := nullBool.UnmarshalJSON(testCase["input"].([]byte))
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		assert.Equal(t, testCase["value"], nullBool.Bool, "value param for intput %+v: %+v", testCase["input"], testCase["value"])
		assert.Equal(t, testCase["valid"], nullBool.Valid, "valid param for intput %+v: %+v", testCase["input"], testCase["valid"])
	}
}

func makeBytes(v interface{}) []byte {
	bytes, _ := json.Marshal(v)
	return bytes
}
