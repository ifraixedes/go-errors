package errors

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMD_Format(t *testing.T) {
	var tcases = []struct {
		desc   string
		format string
		md     MD
		expout string
	}{
		{
			desc:   "'v' verb",
			format: "%v",
			md:     MD{K: "a-key", V: "a-value"},
			expout: "{\"a-key\": a-value}",
		},
		{
			desc:   "'v' verb when value is a struct",
			format: "%v",
			md:     MD{K: "a-key", V: struct{ Name string }{Name: "Ivan"}},
			expout: "{\"a-key\": {Name:Ivan}}",
		},
		{
			desc:   "'v' verb with '+' flag",
			format: "%+v",
			md:     MD{K: "a key", V: 10.5},
			expout: "{\"a key\": 10.5}",
		},
		{
			desc: "any other verb",
			format: func() string {
				var verbs = [...]string{
					"t", "b", "c", "d", "o", "q", "x", "X", "U", "e", "E", "f", "F", "g", "G", "q", "p", "s",
				}

				return fmt.Sprintf("%%%s", verbs[rand.Intn(len(verbs))])
			}(),
			md:     MD{K: "some-key", V: "some-value"},
			expout: "",
		},
	}

	for i := range tcases {
		var tc = tcases[i]
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()

			var out = fmt.Sprintf(tc.format, tc.md)
			assert.Equal(t, tc.expout, out)
		})
	}
}
