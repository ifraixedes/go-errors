package errors

import (
	"fmt"
	"strings"
)

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

// mDatas is a MD slice which allows to internally break the logic between
// dealing with a MD and a slice of MD
type mDatas []MD

// Format satisfies the fmt.Formatter interface.
// It only prints when 'v' verb is used.
func (mds mDatas) Format(state fmt.State, verb rune) {
	if verb != 'v' {
		return
	}

	_, _ = fmt.Fprint(state, "[")

	var mss = make([]string, len(mds))
	for i, m := range mds {
		mss[i] = fmt.Sprintf("%v", m)
	}

	_, _ = fmt.Fprintf(state, "%s]", strings.Join(mss, ","))
}
