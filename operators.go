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

// GetCode returns the Code of the err and true; if err isn't created by any of
// the constructors of this package, false is returned and Code value can be
// ignored.
func GetCode(err error) (Code, bool) {
	var derr, ok = err.(derror)
	if !ok {
		return nil, false
	}

	return derr.c, true
}
