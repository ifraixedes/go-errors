package errors

import (
	"errors"
	"fmt"
	"math/rand"
	"testing"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDerror_Format(t *testing.T) {
	var f2 = func() error {
		var id, e = uuid.NewV4()
		require.NoError(t, e)

		var (
			err = derror{
				c:   testCode(true),
				mds: mDatas{{K: "var1", V: "a string"}, {K: "var2", V: 10}},
				id:  id,
				cs:  newCallStack(),
			}
		)

		return err
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
		var e = derr
		e.werr = errors.New("some external error")
		var (
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
	var id, e = uuid.NewV4()
	require.NoError(t, e)

	var err = derror{
		c:   testCode(true),
		mds: mDatas{{K: "var1", V: "a string"}, {K: "var2", V: 10}},
		id:  id,
		cs:  newCallStack(),
	}

	assert.Equal(t, fmt.Sprintf("%s", err), err.Error())
}

// testCode is a silly example of a Code implementation with the only purpose of
// testing derror type methods.
type testCode bool

func (testCode) String() string {
	return "TestCode"
}

func (testCode) Message() string {
	return "an test code error has happened"
}
