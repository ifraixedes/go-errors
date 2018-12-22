package errors

import (
	"fmt"

	"github.com/gofrs/uuid"
)

// derror is the internal error type used by this package.
// Its name contains a d from decorated and for avoiding to clash with the
// standard error interface.
// derror holds a Code which identifies the error, an ID which uniquely
// identifies the error (it can be useful for correlating different log entries
// for identifying that it's the same error instance), metadata associated
// to the instance of the error and the call stack.
type derror struct {
	c    Code
	id   uuid.UUID
	mds  mDatas
	werr error
	cs   callStack
}

// Format satisfies the fmt.Formatter interface.
func (err derror) Format(state fmt.State, verb rune) {
	if verb == 's' {
		_, _ = fmt.Fprintf(state, "%s: %s", err.c.String(), err.c.Message())
		return
	}

	if verb != 'v' {
		return
	}

	_, _ = fmt.Fprintf(state, "%s: %s\n\tid: %s\n\tmetadata: %v", err.c.String(), err.c.Message(), err.id, err.mds)
	if !state.Flag('+') {
		return
	}

	if err.werr != nil {
		_, _ = fmt.Fprintf(state, "\n\twrapped error: %+v", err.werr)
	}

	if len(err.cs) == 0 {
		return
	}

	if state.Flag('-') {
		_, _ = fmt.Fprintf(state, "\n\tcall stack (compacted):\n%-v", err.cs)
	} else {
		_, _ = fmt.Fprintf(state, "\n\tcall stack:\n%v", err.cs)
	}
}

// Error satisfies the standard error interface.
// It returns a string which is the same output than fmt.Printf("%s", err).
func (err derror) Error() string {
	return fmt.Sprintf("%s", err)
}
