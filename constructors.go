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
