package errors

import "github.com/fikri240794/golib/models/enums"

type Error struct {
	Code        enums.StatusCode
	Message     string
	ErrorFields []ErrorField
}

func NewError(code enums.StatusCode, message string, errorFields ...ErrorField) Error {
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
