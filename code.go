package errors

// Code is the interface that any error code must satisfies.
type Code interface {
	String() string
	Message() string
}
