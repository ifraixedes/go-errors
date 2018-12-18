package errors

import "fmt"

// MD is a key/value pair to add metadata to an error.
type MD struct {
	K string
	V interface{}
}

// Format satisfies the fmt.Formatter interface.
// It only prints when 'v' verb is used.
func (md MD) Format(state fmt.State, verb rune) {
	if verb != 'v' {
		return
	}

	_, _ = fmt.Fprintf(state, "{%q: %+v}", md.K, md.V)
}
