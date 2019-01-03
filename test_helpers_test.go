package errors

// testCode is a silly example of a Code implementation with the only purpose of
// testing derror type methods.
type testCode bool

func (testCode) String() string {
	return "TestCode"
}

func (testCode) Message() string {
	return "an test code error has happened"
}
