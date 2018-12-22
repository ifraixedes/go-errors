package errors

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	t.Run("without metadata", func(t *testing.T) {
		var err = New(testCode(true))
		require.IsType(t, derror{}, err)

		var derr = err.(derror)
		assert.Equal(t, derr.c, testCode(true))
		assert.NotEqual(t, uuid.UUID{}, derr.id)
		assert.Empty(t, derr.mds)
		assert.Nil(t, derr.werr)
		assert.NotEmpty(t, derr.cs)
		assert.NotContains(t, fmt.Sprintf("%+v", err), "errors.New")
	})

	t.Run("with metadata", func(t *testing.T) {
		var err = New(testCode(true), MD{K: "a", V: "va"}, MD{K: "b", V: "vb"})
		require.IsType(t, derror{}, err)

		var derr = err.(derror)
		assert.Equal(t, derr.c, testCode(true))
		assert.NotEqual(t, uuid.UUID{}, derr.id)
		assert.Equal(t, derr.mds, mDatas{{K: "a", V: "va"}, {K: "b", V: "vb"}})
		assert.Nil(t, derr.werr)
		assert.NotEmpty(t, derr.cs)
		assert.NotContains(t, fmt.Sprintf("%+v", err), "errors.New")
	})
}

func TestWrap(t *testing.T) {
	t.Run("ext error & without metadata", func(t *testing.T) {
		var (
			extErr = errors.New("ext error: " + t.Name())
			err    = Wrap(extErr, testCode(true))
		)
		require.IsType(t, derror{}, err)

		var derr = err.(derror)
		assert.Equal(t, derr.c, testCode(true))
		assert.NotEqual(t, uuid.UUID{}, derr.id)
		assert.Empty(t, derr.mds)
		assert.Equal(t, extErr, derr.werr)
		assert.NotEmpty(t, derr.cs)
		assert.NotContains(t, fmt.Sprintf("%+v", err), "errors.Wrap")
	})

	t.Run("ext error & with metadata", func(t *testing.T) {
		var (
			extErr = errors.New("ext error: " + t.Name())
			err    = Wrap(extErr, testCode(true), MD{K: "a", V: "va"}, MD{K: "b", V: "vb"})
		)
		require.IsType(t, derror{}, err)

		var derr = err.(derror)
		assert.Equal(t, derr.c, testCode(true))
		assert.NotEqual(t, uuid.UUID{}, derr.id)
		assert.Equal(t, derr.mds, mDatas{{K: "a", V: "va"}, {K: "b", V: "vb"}})
		assert.Equal(t, extErr, derr.werr)
		assert.NotEmpty(t, derr.cs)
		assert.NotContains(t, fmt.Sprintf("%+v", err), "errors.Wrap")
	})

	t.Run("derror & without metadata", func(t *testing.T) {
		var (
			dErr = New(testCode(true), MD{K: "da", V: "vda"}, MD{K: "db", V: "vdb"})
			err  = Wrap(dErr, testCode(true))
		)
		require.IsType(t, derror{}, err)

		var derr = err.(derror)
		assert.Equal(t, derr.c, testCode(true))
		assert.NotEqual(t, uuid.UUID{}, derr.id)
		assert.Empty(t, derr.mds)
		assert.NotEmpty(t, derr.cs)

		var (
			cdErr   = dErr.(derror)
			expDErr = derror{
				c:    cdErr.c,
				id:   cdErr.id,
				mds:  cdErr.mds,
				werr: cdErr.werr,
			}
		)
		assert.Equal(t, expDErr, derr.werr)

		var prcs = fmt.Sprintf("%+v", err)
		assert.NotContains(t, prcs, "errors.Wrap")

		// Make sure that call stack of derr is not printed
		var parts = strings.Split(prcs, "call stack")
		assert.Len(t, parts, 2)
	})

	t.Run("derror & with metadata", func(t *testing.T) {
		var (
			dErr = New(testCode(true), MD{K: "da", V: "vda"}, MD{K: "db", V: "vdb"})
			err  = Wrap(dErr, testCode(true), MD{K: "a", V: "va"}, MD{K: "b", V: "vb"})
		)
		require.IsType(t, derror{}, err)

		var derr = err.(derror)
		assert.Equal(t, derr.c, testCode(true))
		assert.NotEqual(t, uuid.UUID{}, derr.id)
		assert.Equal(t, derr.mds, mDatas{{K: "a", V: "va"}, {K: "b", V: "vb"}})
		assert.NotEmpty(t, derr.cs)

		var (
			cdErr   = dErr.(derror)
			expDErr = derror{
				c:    cdErr.c,
				id:   cdErr.id,
				mds:  cdErr.mds,
				werr: cdErr.werr,
			}
		)
		assert.Equal(t, expDErr, derr.werr)

		var prcs = fmt.Sprintf("%+v", err)
		assert.NotContains(t, prcs, "errors.Wrap")
		// Make sure that call stack of derr is not printed
		var parts = strings.Split(prcs, "call stack")
		assert.Len(t, parts, 2)
	})
}
