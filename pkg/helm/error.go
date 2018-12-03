package helm

import (
	"fmt"
)

// Error is a Basic struct to build helm errors from
type Error struct {
	message string
	cause   *error
}

// NewError creates a new Error
func NewError(message string, cause ...*error) *Error {
	var innerError *error
	var ok bool

	if len(cause) > 1 {
		panic("Should be just one cause")
	} else if len(cause) == 1 {
		innerError = cause[0]
		if !ok {
			panic(fmt.Sprintf("Incorrect type. Expected *error, got %v", cause[0]))
		}
	}

	return &Error{message, innerError}
}

func (err *Error) Error() string {
	if err.cause != nil {
		return fmt.Sprintf("%s\nCause:\n%v", err.message, err.cause)
	}
	return err.message
}
