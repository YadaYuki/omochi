package errors

import (
	"fmt"

	"github.com/YadaYuki/omochi/app/errors/code"
)

type Error struct {
	Code code.Code
	err  error
}

func NewError(code code.Code, err any) *Error {
	var e error
	switch err := err.(type) {
	case error:
		e = err
	default:
		e = fmt.Errorf("%v", err)
	}
	return &Error{
		Code: code,
		err:  e,
	}
}

func (e *Error) Error() string {
	return e.err.Error()
}

func (e *Error) Unwrap() error {
	return e.err
}
