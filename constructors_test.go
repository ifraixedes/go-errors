package errors

import (
	"fmt"
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
