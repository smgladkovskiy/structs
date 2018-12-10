package null

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

type TestCases map[string][]map[string]interface{}

func checkCases(cases TestCases, t *testing.T, valType interface{}, args ...interface{}) {
	for block, tcs := range cases {
		for _, testCase := range tcs {
			switch valType.(type) {
			case String:
				var ns String
				err := ns.Scan(testCase[in])

				if testCase[ie].(bool) {
					assert.Error(t, err, "[%s | %s] Has error for input %+v: %+v", block, testCase[na], testCase[in], testCase[va])
					break
				}

				assert.Equal(t, testCase[va], ns.String, "[%s | %s] String param for input %+v: %+v", block, testCase[na], testCase[in], testCase[va])
				assert.Equal(t, testCase[iv], ns.Valid, "[%s | %s] isValue param for input %+v: %+v", block, testCase[na], testCase[in], testCase[iv])
			case Time:
				var nt Time
				err := nt.Scan(testCase[in])

				if testCase[ie].(bool) {
					assert.Error(t, err, "[%s | %s] Has error for input %+v: %+v", block, testCase[na], testCase[in], testCase[va])
					break
				}

				switch testCase[in].(type) {
				case string:
					assert.Equal(t, testCase[va], nt.Time.Format(structs.TimeFormat()), "[%s | %s] Valid param for input %+v: %+v", block, testCase[na], testCase[in], testCase[va])
				case *time.Time:
					if testCase[iv].(bool) {
						assert.Equal(t, testCase[va], args[0], "[%s | %s] Time param for input %+v: %+v", block, testCase[na], testCase[in], testCase[va])
					} else {
						assert.Equal(t, testCase[va], time.Time{}, "[%s | %s] Time param for input %+v: %+v", block, testCase[na], testCase[in], testCase[va])
					}

				default:
					assert.Equal(t, testCase[va], nt.Time, "[%s | %s] Time param for input %+v: %+v", block, testCase[na], testCase[in], testCase[va])
				}

				assert.Equal(t, testCase[iv], nt.Valid, "[%s | %s] isValue param for input %+v: %+v", block, testCase[na], testCase[in], testCase[iv])
			}
		}
	}
}
