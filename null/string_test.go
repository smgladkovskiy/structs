package null

import (
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/smgladkovskiy/structs/encoder"
)

func TestNewString(t *testing.T) {
	cases := TestCases{
		"good": {
			{input: "string", expected: "string", isValid: true, isError: false},
		},
		"bad": {
			{input: true, expected: false, isValid: false, isError: true},
		},
		"nil": {
			{input: nil, expected: "", isValid: false, isError: false},
		},
	}
	checkCases(cases, t, String{})
}

func BenchmarkNewString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := NewString("some string")
		if err != nil {
			log.Fatal(err)
		}
	}
}

func TestNewStringf(t *testing.T) {
	f := "Some string by format: %s"
	str := "good"
	ns, err := NewStringf(f, str)
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	assert.True(t, ns.Valid)
	assert.Equal(t, fmt.Sprintf(f, str), ns.String)
}

func BenchmarkNewStringf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := NewStringf("string %d", i)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func TestString_Scan(t *testing.T) {
	nb, _ := NewBool(false)
	cases := TestCases{
		"integers": {
			{input: int8(1), expected: "1", isValid: true, isError: false},
			{input: uint8(1), expected: "1", isValid: true, isError: false},
			{input: int16(1), expected: "1", isValid: true, isError: false},
			{input: uint16(1), expected: "1", isValid: true, isError: false},
			{input: int32(1), expected: "1", isValid: true, isError: false},
			{input: uint32(1), expected: "1", isValid: true, isError: false},
			{input: int64(1), expected: "1", isValid: true, isError: false},
			{input: uint64(1), expected: "1", isValid: true, isError: false},
			{input: int8(0), expected: "0", isValid: true, isError: false},
			{input: uint8(0), expected: "0", isValid: true, isError: false},
			{input: int16(0), expected: "0", isValid: true, isError: false},
			{input: uint16(0), expected: "0", isValid: true, isError: false},
			{input: int32(0), expected: "0", isValid: true, isError: false},
			{input: uint32(0), expected: "0", isValid: true, isError: false},
			{input: int64(0), expected: "0", isValid: true, isError: false},
			{input: uint64(0), expected: "0", isValid: true, isError: false},
			{input: int8(-1), expected: "-1", isValid: true, isError: false},
			{input: int16(-1), expected: "-1", isValid: true, isError: false},
			{input: int32(-1), expected: "-1", isValid: true, isError: false},
			{input: int64(-1), expected: "-1", isValid: true, isError: false},
		},
		"strings": {
			{input: "string", expected: "string", isValid: true, isError: false},
			{input: "", expected: "", isValid: false, isError: false},
			{input: String{"string", true}, expected: "string", isValid: true, isError: false},
		},
		"bytes slice": {
			{input: makeBytes(false), expected: "", isValid: false, isError: false},
			{input: makeBytes(nil), expected: "", isValid: false, isError: false},
			{input: makeBytes(1), expected: "1", isValid: true, isError: false},
			{input: makeBytes(-1), expected: "-1", isValid: true, isError: false},
			{input: encoder.StringToBytes("string"), expected: "string", isValid: true, isError: false},
			{input: encoder.StringToBytes(""), expected: "", isValid: false, isError: false},
		},
		"nil": {
			{input: nil, expected: "", isValid: false, isError: false},
		},
		"errors": {
			{input: true, expected: false, isValid: false, isError: true},
			{input: false, expected: false, isValid: false, isError: true},
			{input: nb, expected: false, isValid: false, isError: true},
		},
	}

	checkCases(cases, t, String{})
}

func BenchmarkString_Scan(b *testing.B) {
	var ns String
	for i := 0; i < b.N; i++ {
		err := ns.Scan(i)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func TestString_Value(t *testing.T) {
	t.Run("Return value", func(t *testing.T) {
		s := "string"
		ns, err := NewString(s)
		if !assert.NoError(t, err) {
			t.FailNow()
		}
		value, _ := ns.Value()
		assert.Equal(t, s, value)
	})
	t.Run("Return nil value", func(t *testing.T) {
		var ns String
		value, _ := ns.Value()
		assert.Nil(t, value)
	})
}

func BenchmarkString_Value(b *testing.B) {
	ns, _ := NewString("string")
	for i := 0; i < b.N; i++ {
		_, err := ns.Value()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func TestString_MarshalJSON(t *testing.T) {
	t.Run("Success marshal", func(t *testing.T) {
		s := "string"
		ns, err := NewString(s)
		if !assert.NoError(t, err) {
			t.FailNow()
		}
		jb, err := ns.MarshalJSON()
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		assert.Equal(t, []byte(`"`+s+`"`), jb)
	})

	t.Run("Null result", func(t *testing.T) {
		ns, err := NewString(nil)
		if !assert.NoError(t, err) {
			t.FailNow()
		}
		jb, err := ns.MarshalJSON()
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		assert.Equal(t, []byte("null"), jb)
	})
}

func BenchmarkString_MarshalJSON(b *testing.B) {
	ns, _ := NewString("string")
	for i := 0; i < b.N; i++ {
		_, err := ns.MarshalJSON()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func TestString_UnmarshalJSON(t *testing.T) {
	t.Run("Success unmarshal", func(t *testing.T) {
		pt := "04231"
		escapedString := "\"" + pt + "\""
		b := []byte(escapedString)
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

func BenchmarkString_UnmarshalJSON(b *testing.B) {
	usc := "\"\u0410\u043b\u0435\u043a\u0441\u0435\u0439\""
	bytes := []byte(usc)
	var ns String
	for i := 0; i < b.N; i++ {
		err := ns.UnmarshalJSON(bytes)
		if err != nil {
			log.Fatal(err)
		}
	}

}
