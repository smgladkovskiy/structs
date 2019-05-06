package null

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/smgladkovskiy/structs/encoder"

	"github.com/stretchr/testify/assert"
)

func TestNewBool(t *testing.T) {
	t.Parallel()
	t.Run("Success", func(t *testing.T) {
		nullBool, err := NewBool(true)
		if !assert.NoError(t, err) {
			t.FailNow()
		}
		assert.True(t, nullBool.Valid)
		assert.Equal(t, true, nullBool.Bool)
	})

	t.Run("False on nil", func(t *testing.T) {
		nullBool, err := NewBool(nil)
		if !assert.NoError(t, err) {
			t.FailNow()
		}
		assert.False(t, nullBool.Valid)
		assert.Equal(t, false, nullBool.Bool)
	})
}

func BenchmarkNewBool(b *testing.B) {
	var nb Bool
	for i := 0; i < b.N; i++ {
		err := nb.Scan(i % 2)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func TestBool_Value(t *testing.T) {
	t.Parallel()
	t.Run("Return bool value", func(t *testing.T) {
		nullBool, err := NewBool(true)
		if !assert.NoError(t, err) {
			t.FailNow()
		}
		value, err := nullBool.Value()
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		assert.Equal(t, true, value)
	})
	t.Run("Return nil value", func(t *testing.T) {
		var nullBool Bool
		value, _ := nullBool.Value()

		assert.Nil(t, value)
	})
}

func BenchmarkBool_Value(b *testing.B) {
	nb, _ := NewBool(true)
	for i := 0; i < b.N; i++ {
		_, err := nb.Value()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func TestBool_Scan(t *testing.T) {
	nb1, _ := NewBool(true)
	nb2, _ := NewBool(false)
	nb3, _ := NewBool(nil)
	cases := TestCases{
		"ints": {
			{input: 1, expected: true, isValid: true, isError: false},
			{input: int8(1), expected: true, isValid: true, isError: false},
			{input: int16(1), expected: true, isValid: true, isError: false},
			{input: int32(1), expected: true, isValid: true, isError: false},
			{input: int64(1), expected: true, isValid: true, isError: false},

			{input: 0, expected: false, isValid: true, isError: false},
			{input: int8(0), expected: false, isValid: true, isError: false},
			{input: int16(0), expected: false, isValid: true, isError: false},
			{input: int32(0), expected: false, isValid: true, isError: false},
			{input: int64(0), expected: false, isValid: true, isError: false},

			{input: 5, expected: false, isValid: false, isError: false},
			{input: -5, expected: false, isValid: false, isError: false},
		},
		"strings": {
			{input: "1", expected: true, isValid: true, isError: false},
			{input: "t", expected: true, isValid: true, isError: false},
			{input: "T", expected: true, isValid: true, isError: false},
			{input: "true", expected: true, isValid: true, isError: false},
			{input: "TRUE", expected: true, isValid: true, isError: false},
			{input: "True", expected: true, isValid: true, isError: false},
			{input: "y", expected: true, isValid: true, isError: false},
			{input: "Y", expected: true, isValid: true, isError: false},
			{input: "YES", expected: true, isValid: true, isError: false},
			{input: "Yes", expected: true, isValid: true, isError: false},

			{input: "0", expected: false, isValid: true, isError: false},
			{input: "f", expected: false, isValid: true, isError: false},
			{input: "F", expected: false, isValid: true, isError: false},
			{input: "false", expected: false, isValid: true, isError: false},
			{input: "False", expected: false, isValid: true, isError: false},
			{input: "FALSE", expected: false, isValid: true, isError: false},
			{input: "n", expected: false, isValid: true, isError: false},
			{input: "N", expected: false, isValid: true, isError: false},
			{input: "NO", expected: false, isValid: true, isError: false},
			{input: "No", expected: false, isValid: true, isError: false},
			{input: "some string", expected: false, isValid: false, isError: false},
		},

		"booleans": {
			{input: true, expected: true, isValid: true, isError: false},
			{input: false, expected: false, isValid: true, isError: false},
			{input: nb1, expected: true, isValid: true, isError: false},
			{input: nb2, expected: false, isValid: true, isError: false},
			{input: nb3, expected: false, isValid: false, isError: false},
		},

		"byte slice": {
			{name: "bytes for true", input: []byte("true"), expected: true, isValid: true, isError: false},
			{name: "bytes for false", input: []byte("false"), expected: false, isValid: true, isError: false},
			{name: "bytes for true", input: encoder.StringToBytes("true"), expected: true, isValid: true, isError: false},
			{name: "bytes for false", input: encoder.StringToBytes("false"), expected: false, isValid: true, isError: false},
			{name: "bytes for nil", input: []byte("null"), expected: false, isValid: false, isError: false},
			{name: "bytes for yes", input: []byte("yes"), expected: true, isValid: true, isError: false},
			{name: "bytes for YES", input: []byte("YES"), expected: true, isValid: true, isError: false},
			{name: "bytes for y", input: []byte("y"), expected: true, isValid: true, isError: false},
			{name: "bytes for Y", input: []byte("Y"), expected: true, isValid: true, isError: false},
			{name: "bytes for no", input: []byte("no"), expected: false, isValid: true, isError: false},
			{name: "bytes for NO", input: []byte("NO"), expected: false, isValid: true, isError: false},
			{name: "bytes for n", input: []byte("n"), expected: false, isValid: true, isError: false},
			{name: "bytes for N", input: []byte("N"), expected: false, isValid: true, isError: false},
			{name: "bytes for 1", input: []byte("1"), expected: true, isValid: true, isError: false},
			{name: "bytes for 0", input: []byte("0"), expected: false, isValid: true, isError: false},
		},
		"nil": {
			{input: nil, expected: false, isValid: false, isError: false},
		},

		"errors": {},
	}
	checkCases(cases, t, Bool{})
}

func BenchmarkBool_Scan(b *testing.B) {
	var nb Bool
	for i := 0; i < b.N; i++ {
		err := nb.Scan(i)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func TestBool_MarshalJSON(t *testing.T) {
	t.Run("Success marshal", func(t *testing.T) {
		nullBool, err := NewBool(true)
		if !assert.NoError(t, err) {
			t.FailNow()
		}
		b, _ := json.Marshal("true")
		jb, err := nullBool.MarshalJSON()
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		assert.Equal(t, b, jb)
	})

	t.Run("Null result", func(t *testing.T) {
		nb, err := NewBool(nil)
		if !assert.NoError(t, err) {
			t.FailNow()
		}
		jb, _ := nb.MarshalJSON()
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		assert.Equal(t, []byte("null"), jb)
	})
}

func BenchmarkBool_MarshalJSON(b *testing.B) {
	nb, _ := NewBool("true")
	for i := 0; i < b.N; i++ {
		_, err := nb.MarshalJSON()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func TestBool_UnmarshalJSON(t *testing.T) {
	cases := TestCases{
		"bools": {
			{input: encoder.StringToBytes("false"), expected: false, isValid: true, isError: false},
			{input: encoder.StringToBytes("true"), expected: true, isValid: true, isError: false},
			{input: encoder.StringToBytes("null"), expected: false, isValid: false, isError: false},
		},
		"error": {
			{input: encoder.StringToBytes("t"), expected: false, isValid: false, isError: true},
		},
	}
	checkUnmarshalCases(t, cases, Bool{})
}

func BenchmarkBool_UnmarshalJSON(b *testing.B) {
	bytes := makeBytes(true)
	var nb Bool
	for i := 0; i < b.N; i++ {
		err := nb.UnmarshalJSON(bytes)
		if err != nil {
			log.Fatal(err)
		}
	}

}
