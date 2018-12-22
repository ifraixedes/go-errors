package errors

import (
	"errors"
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDerror_Format(t *testing.T) {
	var f2 = func() error {
		return New(testCode(true), MD{K: "var1", V: "a string"}, MD{K: "var2", V: 10})
	}

	var f1 = func() error {
		return f2()
	}

	var (
		err  = f1()
		derr = err.(derror)
	)

	t.Run("code and message (%s)", func(t *testing.T) {
		var (
			s   = fmt.Sprintf("%s", err)
			exp = fmt.Sprintf("%s: %s", derr.c.String(), derr.c.Message())
		)

		assert.Equal(t, exp, s)
	})

	t.Run("code, message, id and metadata (%v)", func(t *testing.T) {
		var (
			s   = fmt.Sprintf("%v", err)
			exp = fmt.Sprintf("%s: %s\n\tid: %s\n\tmetadata: %v",
				derr.c.String(), derr.c.Message(), derr.id, derr.mds,
			)
		)

		assert.Equal(t, exp, s)
	})

	t.Run("code, message, id, metadata and call stack (%+v)", func(t *testing.T) {
		var (
			s   = fmt.Sprintf("%+v", err)
			exp = fmt.Sprintf("%s: %s\n\tid: %s\n\tmetadata: %v\n\tcall stack:\n%+v",
				derr.c.String(), derr.c.Message(), derr.id, derr.mds, derr.cs,
			)
		)

		assert.Equal(t, exp, s)
	})

	t.Run("code, message, id, metadata and call stack (compacted) (%+-v)", func(t *testing.T) {
		var (
			s   = fmt.Sprintf("%+-v", err)
			exp = fmt.Sprintf("%s: %s\n\tid: %s\n\tmetadata: %v\n\tcall stack (compacted):\n%-v",
				derr.c.String(), derr.c.Message(), derr.id, derr.mds, derr.cs,
			)
		)

		assert.Equal(t, exp, s)
	})

	t.Run("code, message, id, metadata, wrapped error and call stack (%+v)", func(t *testing.T) {
		var (
			err = Wrap(errors.New("some external error"), testCode(true), MD{K: "var1", V: "a string"})
			e   = err.(derror)
			s   = fmt.Sprintf("%+v", e)
			exp = fmt.Sprintf("%s: %s\n\tid: %s\n\tmetadata: %v\n\twrapped error: %+v\n\tcall stack:\n%+v",
				e.c.String(), e.c.Message(), e.id, e.mds, e.werr, e.cs,
			)
		)

		assert.Equal(t, exp, s)
	})

	t.Run("any other verb", func(t *testing.T) {
		var verbs = [...]string{
			"t", "b", "c", "d", "o", "q", "x", "X", "U", "e", "E", "f", "F", "g", "G", "q", "p",
		}

		var v = fmt.Sprintf("%%%s", verbs[rand.Intn(len(verbs))])
		var s = fmt.Sprintf(v, err)
		assert.Empty(t, s)
	})
}

func TestDerror_Error(t *testing.T) {
	var err = New(testCode(true), MD{K: "var1", V: "a string"}, MD{K: "var2", V: 10})
	assert.Equal(t, fmt.Sprintf("%s", err), err.Error())
}
