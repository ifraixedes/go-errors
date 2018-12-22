package errors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIs(t *testing.T) {
	t.Run("is exactly the same error code", func(t *testing.T) {
		var derr = New(testCode(true))
		assert.True(t, Is(derr, testCode(true)))
	})

	t.Run("is a similar error code", func(t *testing.T) {
		var derr = New(testCode(true))
		assert.True(t, Is(derr, similarTestCode(false)))
	})

	t.Run("is a different error code", func(t *testing.T) {
		var derr = New(testCode(true))
		assert.False(t, Is(derr, differentTestCode(true)))
	})

	t.Run("external error", func(t *testing.T) {
		var err = errors.New("some error")
		assert.False(t, Is(err, testCode(true)))
	})
}

func TestGetCode(t *testing.T) {
	t.Run("created by this package", func(t *testing.T) {
		var (
			expc  = testCode(true)
			err   = New(expc)
			c, ok = GetCode(err)
		)

		assert.Equal(t, expc, c)
		assert.True(t, ok)
	})

	t.Run("created by other package", func(t *testing.T) {
		var (
			err   = errors.New("some error")
			_, ok = GetCode(err)
		)

		assert.False(t, ok)
	})
}

		assert.False(t, ok)
	})
}

// similarTestCode is a silly example of a Code implementation with the only
// purpose of testing the Is method
type similarTestCode bool

func (similarTestCode) String() string {
	var tc = testCode(true)
	return tc.String()
}

func (similarTestCode) Message() string {
	var tc = testCode(true)
	return tc.Message()
}

// differentTestCode is a silly example of a Code implementation with the only
// purpose of testing the Is method
type differentTestCode bool

func (differentTestCode) String() string {
	return "DifferentError"
}

func (differentTestCode) Message() string {
	return "This is a different error"
}
