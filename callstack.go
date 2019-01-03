package errors

import (
	"fmt"
	"runtime"
)

type callStack []uintptr

// newCallStack creates a callStack of calls skipping the call to
// runtime.Callers, newCallStack and the caller of newCallStack.
// The newCallStack is skipped because it's meant to be used by the errors
// consturctors and they shouldn't appear in the error value call stack.
func newCallStack() callStack {
	var (
		depth = 20
		pcs   = make([]uintptr, depth)
		l     = runtime.Callers(3, pcs)
	)

	for l == depth {
		depth += 10
		pcs = make([]uintptr, depth)
		l = runtime.Callers(3, pcs)
	}

	pcs = pcs[:l]

	return callStack(pcs)
}

// Format satisfies the fmt.Formatter interface.
// It only prints the value if the verb 'v' and flag '+' are used, printing the
// complete function identifier (package path + name), the file and the line.
// When the additional '-' is used, then only the function identifier is
// printed.
func (cs callStack) Format(state fmt.State, verb rune) {
	if verb != 'v' {
		return
	}

	if len(cs) > 0 {
		var frs = runtime.CallersFrames(cs)
		for {
			var f, more = frs.Next()

			if state.Flag('-') {
				_, _ = fmt.Fprintf(state, "\t%s", f.Function)
			} else {
				_, _ = fmt.Fprintf(state, "\t%s\n\t\t%s:%d", f.Function, f.File, f.Line)
			}

			if !more {
				break
			}

			fmt.Fprint(state, "\n")
		}
	}
}
