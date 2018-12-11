package null

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestNewString(t *testing.T) {
	cases := TestCases{
		"good": {
			{in: "string", va: "string", iv: true, ie: false},
		},
		"bad": {
			{in: true, va: false, iv: false, ie: true},
		},
		"nil": {
			{in: nil, va: "", iv: false, ie: false},
		},
	}
	checkCases(cases, t, String{})
}

func BenchmarkNewString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := NewString(i)
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
			{in: int8(1), va: "1", iv: true, ie: false},
			{in: uint8(1), va: "1", iv: true, ie: false},
			{in: int16(1), va: "1", iv: true, ie: false},
			{in: uint16(1), va: "1", iv: true, ie: false},
			{in: int32(1), va: "1", iv: true, ie: false},
			{in: uint32(1), va: "1", iv: true, ie: false},
			{in: int64(1), va: "1", iv: true, ie: false},
			{in: uint64(1), va: "1", iv: true, ie: false},
			{in: int8(0), va: "0", iv: true, ie: false},
			{in: uint8(0), va: "0", iv: true, ie: false},
			{in: int16(0), va: "0", iv: true, ie: false},
			{in: uint16(0), va: "0", iv: true, ie: false},
			{in: int32(0), va: "0", iv: true, ie: false},
			{in: uint32(0), va: "0", iv: true, ie: false},
			{in: int64(0), va: "0", iv: true, ie: false},
			{in: uint64(0), va: "0", iv: true, ie: false},
			{in: int8(-1), va: "-1", iv: true, ie: false},
			{in: int16(-1), va: "-1", iv: true, ie: false},
			{in: int32(-1), va: "-1", iv: true, ie: false},
			{in: int64(-1), va: "-1", iv: true, ie: false},
		},
		"strings": {
			{in: "string", va: "string", iv: true, ie: false},
			{in: "", va: "", iv: false, ie: false},
			{in: String{"string", true}, va: "string", iv: true, ie: false},
		},
		"bytes slice": {
			{in: makeBytes("string"), va: "string", iv: true, ie: false},
			{in: makeBytes(""), va: "", iv: false, ie: false},
			{in: makeBytes(1), va: "1", iv: true, ie: false},
			{in: makeBytes(0), va: "0", iv: true, ie: false},
			{in: makeBytes(-1), va: "-1", iv: true, ie: false},
			{in: makeBytes(nil), va: "", iv: false, ie: false},
		},
		"nil": {
			{in: nil, va: "", iv: false, ie: false},
		},
		"errors": {
			{in: true, va: false, iv: false, ie: true},
			{in: false, va: false, iv: false, ie: true},
			{in: nb, va: false, iv: false, ie: true},
			{in: makeBytes(false), va: false, iv: false, ie: true},
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
