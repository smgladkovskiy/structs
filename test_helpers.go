package nulls

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func AssertHasErrors(t *testing.T, errs interface{}) bool {
	switch e := errs.(type) {
	case error:
		return assert.Error(t, e)
	}

	return false
}

func AssertNoErrors(t *testing.T, errs interface{}) bool {
	switch e := errs.(type) {
	case error:
		return assert.NoError(t, e)
	case nil:
		return assert.NoError(t, nil)
	}

	return false
}
