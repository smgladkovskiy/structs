package zero

import (
	"github.com/smgladkovskiy/structs"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
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
			case Time:
				var ti Time
				err := ti.Scan(testCase[in])

				if testCase[ie].(bool) {
					assert.Error(t, err, errorForErrorMsg, block, caseName, testCase[in], testCase[va])
					break
				}

				switch testCase[in].(type) {
				case string:
					assert.Equal(t, testCase[va], ti.Time.Format(structs.TimeFormat()), assertForValidMsg, block, caseName, testCase[in], testCase[va])
				case *time.Time:
					if testCase[iv].(bool) {
						assert.Equal(t, testCase[va], args[0], assertForValueMsg, block, caseName, testCase[in], testCase[va])
					} else {
						assert.Equal(t, testCase[va], time.Time{}, assertForValueMsg, block, caseName, testCase[in], testCase[va])
					}

				default:
					assert.Equal(t, testCase[va], ti.Time, assertForValueMsg, block, caseName, testCase[in], testCase[va])
				}
			}
		}
	}
}
