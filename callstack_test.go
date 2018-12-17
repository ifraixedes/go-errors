package errors

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCallStack_Format(t *testing.T) {
	var (
		ptName = t.Name()
		cstk   callStack
		f2     = func() { cstk = newCallStack() }
		f1     = func() { f2() }
	)
	f1()

	t.Run("'v' verb", func(t *testing.T) {
		var s = fmt.Sprintf("%v", cstk)
		var sls = strings.Split(s, "\n")

		assert.Contains(t, sls[0], fmt.Sprintf("errors.%s.func1", ptName))
		assert.Contains(t, sls[1], "errors/callstack_test.go:16")
		assert.Contains(t, sls[2], fmt.Sprintf("errors.%s.func2", ptName))
		assert.Contains(t, sls[3], "errors/callstack_test.go:17")
	})

	t.Run("'v' verb and '-' flags", func(t *testing.T) {
		var s = fmt.Sprintf("%-v", cstk)
		var sls = strings.Split(s, "\n")

		assert.Contains(t, sls[0], fmt.Sprintf("errors.%s.func1", ptName))
		assert.Contains(t, sls[1], fmt.Sprintf("errors.%s.func2", ptName))
	})

	t.Run("any other verb", func(t *testing.T) {
		var verbs = [...]string{
			"t", "b", "c", "d", "o", "q", "x", "X", "U", "e", "E", "f", "F", "g", "G", "q", "p", "s",
		}

		var f = fmt.Sprintf("%%%s", verbs[rand.Intn(len(verbs))])
		var s = fmt.Sprintf(f, cstk)
		assert.Empty(t, s)
	})

	t.Run("empty stack", func(t *testing.T) {
		var (
			ecstk = callStack(nil)
			s     = fmt.Sprintf("%v", ecstk)
		)

		assert.Empty(t, s)
	})
}
