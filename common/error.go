package common

import (
    "bytes"
)

// Error represents more complex error than plain Go "error".
// It contains another Error, named cause, as well as error message.
type Error struct {
    message string
    cause   *Error
}

func NewError(message string, cause error) error {
    if err, ok := cause.(*Error); ok {
        return &Error{message, err}
    } else {
        return &Error{message, WrapError(cause)}
    }
}

func WrapError(err error) *Error {
    return &Error{err.Error(), nil}
}

func (e *Error) Message() string {
    return e.message
}

func (e *Error) Cause() *Error {
    return e.cause
}

func (e *Error) Error() string {
    return e.Message()
}

func (e *Error) String() string {
    return e.Trace()
}

func (e *Error) Trace() string {
    var buf bytes.Buffer

    buf.WriteString("Error: ")
    buf.WriteString(e.message)
    cause := e.cause
    for cause != nil {
        buf.WriteString("\nCaused by: ")
        buf.WriteString(cause.message)
        cause = cause.cause
    }

    return buf.String()
}
