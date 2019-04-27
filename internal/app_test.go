package internal

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

var TA = AppHelper{mode: Testing}

func TestAppHelper_Errors_nil(t *testing.T) {
	err := errors.New("The world is flat")
	defer func() {
		if r := recover(); r != nil {
			assert.NotNil(t, r)
		} else {
			assert.Fail(t, "")
		}
	}()
	TA.Errors(err, nil)
}

func TestAppHelper_Errors_fill(t *testing.T) {
	err := errors.New("42 is the answer")
	ok := false
	TA.Errors(err, func(err error) {
		ok = true
	})

	assert.True(t, ok)
}
