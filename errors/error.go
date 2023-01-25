package errors

import "github.com/fikri240794/golib/status_code"

type Error struct {
	Code        status_code.StatusCode
	Message     string
	ErrorFields []ErrorField
}

func NewError(code status_code.StatusCode, message string, errorFields ...ErrorField) Error {
	var err Error = Error{
		Code:        code,
		Message:     message,
		ErrorFields: errorFields,
	}

	return err
}

func (e Error) Error() string {
	return e.Message
}
