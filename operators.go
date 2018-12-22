package errors

// Is return true if err is an error value created by one of the constructor
// of this package and c has the same string representation (value returned by
// the String method) and the same message (value returned by the  Message
// method) to the error code, otherwise it returns false.
func Is(err error, c Code) bool {
	var derr, ok = err.(derror)
	if !ok {
		return false
	}

	return derr.c.String() == c.String() && derr.c.Message() == c.Message()
}
