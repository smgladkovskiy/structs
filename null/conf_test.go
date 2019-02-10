package null

import (
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/smgladkovskiy/structs"
)

const (
	in = "input"
	va = "value"
	iv = "isValid"
	na = "name"
	ie = "isError"
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
			caseName := testCase[na]
			if caseName == nil {
				caseName = testCase[in]
			}
			switch valType.(type) {
			case String:
				var ns String
				err := ns.Scan(testCase[in])

				if testCase[ie].(bool) {
					assert.Error(t, err, errorForErrorMsg, block, caseName, testCase[in], testCase[va])
					break
				}

				assert.Equal(t, testCase[va], ns.String, assertForValueMsg, block, caseName, testCase[in], testCase[va])
				assert.Equal(t, testCase[iv], ns.Valid, assertForValidMsg, block, caseName, testCase[in], testCase[iv])
			case Int64:
				var ni Int64
				err := ni.Scan(testCase[in])

				if testCase[ie].(bool) {
					assert.Error(t, err, errorForErrorMsg, block, caseName, testCase[in], testCase[va])
					break
				}

				assert.Equal(t, testCase[va], ni.Int64, assertForValueMsg, block, caseName, testCase[in], testCase[va])
				assert.Equal(t, testCase[iv], ni.Valid, assertForValidMsg, block, caseName, testCase[in], testCase[iv])
			case Float64:
				var nf Float64
				err := nf.Scan(testCase[in])

				if testCase[ie].(bool) {
					assert.Error(t, err, errorForErrorMsg, block, caseName, testCase[in], testCase[va])
					break
				}

				assert.Equal(t, testCase[va], nf.Float64, assertForValueMsg, block, caseName, testCase[in], testCase[va])
				assert.Equal(t, testCase[iv], nf.Valid, assertForValidMsg, block, caseName, testCase[in], testCase[iv])
			case Bool:
				var nb Bool
				err := nb.Scan(testCase[in])
				if testCase[ie].(bool) {
					assert.Error(t, err, errorForErrorMsg, block, caseName, testCase[in], testCase[va])
					break
				}

				assert.Equal(t, testCase[va], nb.Bool, assertForValueMsg, block, caseName, testCase[in], testCase[va])
				assert.Equal(t, testCase[iv], nb.Valid, assertForValidMsg, block, caseName, testCase[in], testCase[iv])
			case Time:
				var nt Time
				err := nt.Scan(testCase[in])

				if testCase[ie].(bool) {
					assert.Error(t, err, errorForErrorMsg, block, caseName, testCase[in], testCase[va])
					break
				}

				switch testCase[in].(type) {
				case string:
					assert.Equal(t, testCase[va], nt.Time.Format(structs.TimeFormat()), assertForValidMsg, block, caseName, testCase[in], testCase[va])
				case *time.Time:
					if testCase[iv].(bool) {
						assert.Equal(t, testCase[va], args[0], assertForValueMsg, block, caseName, testCase[in], testCase[va])
					} else {
						assert.Equal(t, testCase[va], time.Time{}, assertForValueMsg, block, caseName, testCase[in], testCase[va])
					}

				default:
					assert.Equal(t, testCase[va], nt.Time, assertForValueMsg, block, caseName, testCase[in], testCase[va])
				}

				assert.Equal(t, testCase[iv], nt.Valid, assertForValidMsg, block, caseName, testCase[in], testCase[iv])
			case Date:
				var nd Date
				err := nd.Scan(testCase[in])

				if testCase[ie].(bool) {
					assert.Error(t, err, errorForErrorMsg, block, caseName, testCase[in], testCase[va])
					break
				}

				switch testCase[in].(type) {
				case string:
					assert.Equal(t, testCase[va], nd.Time.Format(structs.DateFormat()), assertForValidMsg, block, caseName, testCase[in], testCase[va])
				case *time.Time:
					if testCase[iv].(bool) {
						assert.Equal(t, testCase[va], args[0], assertForValueMsg, block, caseName, testCase[in], testCase[va])
					} else {
						assert.Equal(t, testCase[va], time.Time{}, assertForValueMsg, block, caseName, testCase[in], testCase[va])
					}

				default:
					assert.Equal(t, testCase[va], nd.Time, assertForValueMsg, block, caseName, testCase[in], testCase[va])
				}

				assert.Equal(t, testCase[iv], nd.Valid, assertForValidMsg, block, caseName, testCase[in], testCase[iv])
			}
		}
	}
}

func checkUnmarshalCases(t *testing.T, cases TestCases, valType interface{}, args ...interface{}) {
	for block, tcs := range cases {
		for _, testCase := range tcs {
			caseName := testCase[na]
			if caseName == nil {
				caseName = string(testCase[in].([]byte))
			}
			switch valType.(type) {
			case Bool:
				var nb Bool
				err := nb.UnmarshalJSON(testCase[in].([]byte))
				if testCase[ie].(bool) {
					assert.Error(t, err, errorForErrorMsg, block, caseName, testCase[in], testCase[va])
					break
				}

				assert.Equal(t, testCase[va], nb.Bool, assertForValueMsg, block, caseName, testCase[in], testCase[va])
				assert.Equal(t, testCase[iv], nb.Valid, assertForValidMsg, block, caseName, testCase[in], testCase[iv])
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
