package null

import (
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/smgladkovskiy/structs"
)

const (
	input    = "input"
	expected = "value"
	isValid  = "isValid"
	name     = "name"
	isError  = "isError"
)

const (
	errorForErrorMsg  = "[%s | %+v] Has error for input %+v: %+v"
	assertForValueMsg = "[%s | %+v] Value param for input %+v: %+v"
	assertForValidMsg = "[%s | %+v] isValid param for input %+v: %+v"
)

type TestCases map[string][]map[string]interface{}

func checkCases(cases TestCases, t *testing.T, valType interface{}, args ...interface{}) {
	for block, tcs := range cases {
		for _, testCase := range tcs {
			caseName := testCase[name]
			if caseName == nil {
				caseName = testCase[input]
			}
			switch valType.(type) {
			case String:
				var ns String
				err := ns.Scan(testCase[input])

				if testCase[isError].(bool) {
					assert.Error(t, err, errorForErrorMsg, block, caseName, testCase[input], testCase[expected])
					break
				}

				assert.Equal(t, testCase[expected], ns.String, assertForValueMsg, block, caseName, testCase[input], testCase[expected])
				assert.Equal(t, testCase[isValid], ns.Valid, assertForValidMsg, block, caseName, testCase[input], testCase[isValid])
			case Int64:
				var ni Int64
				err := ni.Scan(testCase[input])

				if testCase[isError].(bool) {
					assert.Error(t, err, errorForErrorMsg, block, caseName, testCase[input], testCase[expected])
					break
				}

				assert.Equal(t, testCase[expected], ni.Int64, assertForValueMsg, block, caseName, testCase[input], testCase[expected])
				assert.Equal(t, testCase[isValid], ni.Valid, assertForValidMsg, block, caseName, testCase[input], testCase[isValid])
			case Float64:
				var nf Float64
				err := nf.Scan(testCase[input])

				if testCase[isError].(bool) {
					assert.Error(t, err, errorForErrorMsg, block, caseName, testCase[input], testCase[expected])
					break
				}

				assert.Equal(t, testCase[expected], nf.Float64, assertForValueMsg, block, caseName, testCase[input], testCase[expected])
				assert.Equal(t, testCase[isValid], nf.Valid, assertForValidMsg, block, caseName, testCase[input], testCase[isValid])
			case Bool:
				var nb Bool
				err := nb.Scan(testCase[input])
				if testCase[isError].(bool) {
					assert.Error(t, err, errorForErrorMsg, block, caseName, testCase[input], testCase[expected])
					break
				}

				assert.Equal(t, testCase[expected], nb.Bool, assertForValueMsg, block, caseName, testCase[input], testCase[expected])
				assert.Equal(t, testCase[isValid], nb.Valid, assertForValidMsg, block, caseName, testCase[input], testCase[isValid])
			case Time:
				var nt Time
				err := nt.Scan(testCase[input])

				if testCase[isError].(bool) {
					assert.Error(t, err, errorForErrorMsg, block, caseName, testCase[input], testCase[expected])
					break
				}

				switch testCase[input].(type) {
				case string:
					assert.Equal(t, testCase[expected], nt.Time.Format(structs.TimeFormat()), assertForValidMsg, block, caseName, testCase[input], testCase[expected])
				case *time.Time:
					if testCase[isValid].(bool) {
						assert.Equal(t, testCase[expected], args[0], assertForValueMsg, block, caseName, testCase[input], testCase[expected])
					} else {
						assert.Equal(t, testCase[expected], time.Time{}, assertForValueMsg, block, caseName, testCase[input], testCase[expected])
					}

				default:
					assert.Equal(t, testCase[expected], nt.Time, assertForValueMsg, block, caseName, testCase[input], testCase[expected])
				}

				assert.Equal(t, testCase[isValid], nt.Valid, assertForValidMsg, block, caseName, testCase[input], testCase[isValid])
			case Date:
				var nd Date
				err := nd.Scan(testCase[input])

				if testCase[isError].(bool) {
					assert.Error(t, err, errorForErrorMsg, block, caseName, testCase[input], testCase[expected])
					break
				}

				switch testCase[input].(type) {
				case string:
					assert.Equal(t, testCase[expected], nd.Time.Format(structs.DateFormat()), assertForValidMsg, block, caseName, testCase[input], testCase[expected])
				case *time.Time:
					if testCase[isValid].(bool) {
						assert.Equal(t, testCase[expected], args[0], assertForValueMsg, block, caseName, testCase[input], testCase[expected])
					} else {
						assert.Equal(t, testCase[expected], time.Time{}, assertForValueMsg, block, caseName, testCase[input], testCase[expected])
					}

				default:
					assert.Equal(t, testCase[expected], nd.Time, assertForValueMsg, block, caseName, testCase[input], testCase[expected])
				}

				assert.Equal(t, testCase[isValid], nd.Valid, assertForValidMsg, block, caseName, testCase[input], testCase[isValid])
			}
		}
	}
}

func checkUnmarshalCases(t *testing.T, cases TestCases, valType interface{}) {
	for block, tcs := range cases {
		for _, testCase := range tcs {
			caseName := testCase[name]
			if caseName == nil {
				caseName = string(testCase[input].([]byte))
			}
			switch valType.(type) {
			case Bool:
				var nb Bool
				err := nb.UnmarshalJSON(testCase[input].([]byte))
				if testCase[isError].(bool) {
					assert.Error(t, err, errorForErrorMsg, block, caseName, testCase[input], testCase[expected])
					break
				}

				assert.Equal(t, testCase[expected], nb.Bool, assertForValueMsg, block, caseName, testCase[input], testCase[expected])
				assert.Equal(t, testCase[isValid], nb.Valid, assertForValidMsg, block, caseName, testCase[input], testCase[isValid])
			}
		}
	}
}

func makeBytes(v interface{}) []byte {
	var b []byte
	switch val := v.(type) {
	case string:
		b = append(b, val...)
	case int:
		b = strconv.AppendInt(b, int64(val), 10)
	case int8:
		b = strconv.AppendInt(b, int64(val), 10)
	case int16:
		b = strconv.AppendInt(b, int64(val), 10)
	case int32:
		b = strconv.AppendInt(b, int64(val), 10)
	case int64:
		b = strconv.AppendInt(b, int64(val), 10)
	}
	return b
}
