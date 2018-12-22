package errors

import "github.com/gofrs/uuid"

// New creates a new error with c and mds.
func New(c Code, mds ...MD) error {
	var id, _ = uuid.NewV4()

	return derror{
		c:   c,
		id:  id,
		mds: mds,
		cs:  newCallStack(),
	}
}

// Wrap creates a new error with c and mds, wrapping err.
func Wrap(err error, c Code, mds ...MD) error {
	var (
		id, _ = uuid.NewV4()
		derr  = derror{
			c:   c,
			id:  id,
			mds: mds,
			cs:  newCallStack(),
		}
	)

	if de, ok := err.(derror); ok {
		de.cs = nil

		derr.werr = de
	} else {
		derr.werr = err
	}

	return derr
}
