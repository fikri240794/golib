package vms

import (
	"github.com/fikri240794/golib/errors"
)

type ResponseErrorVM struct {
	Code        int                     `json:"-"`
	Message     string                  `json:"message"`
	ErrorFields []*ResponseErrorFieldVM `json:"error_fields,omitempty"`
}

func NewResponseErrorVM() *ResponseErrorVM {
	return &ResponseErrorVM{}
}

func (vm *ResponseErrorVM) SetCode(code int) *ResponseErrorVM {
	vm.Code = code

	return vm
}

func (vm *ResponseErrorVM) SetMessage(message string) *ResponseErrorVM {
	vm.Message = message

	return vm
}

func (vm *ResponseErrorVM) SetErrorFields(errorField []*ResponseErrorFieldVM) *ResponseErrorVM {
	if vm.ErrorFields == nil {
		vm.ErrorFields = []*ResponseErrorFieldVM{}
	}

	vm.ErrorFields = append(vm.ErrorFields, errorField...)

	return vm
}

func (vm *ResponseErrorVM) AddErrorField(errorField *ResponseErrorFieldVM) *ResponseErrorVM {
	if vm.ErrorFields == nil {
		vm.ErrorFields = []*ResponseErrorFieldVM{}
	}

	vm.ErrorFields = append(vm.ErrorFields, errorField)

	return vm
}

func (vm *ResponseErrorVM) mapFromCustomError(err errors.Error) *ResponseErrorVM {
	vm = vm.SetCode(err.Code.ToHttpStatusCode()).
		SetMessage(err.Message)

	if len(err.ErrorFields) > 0 {
		for i := 0; i < len(err.ErrorFields); i++ {
			vm = vm.AddErrorField(NewResponseErrorFieldVM(err.ErrorFields[i].Field, err.ErrorFields[i].Message))
		}
	}

	return vm
}

func (vm *ResponseErrorVM) ParseError(err error) *ResponseErrorVM {
	var (
		customError   errors.Error
		isCustomError bool
	)

	if err != nil {
		customError, isCustomError = err.(errors.Error)

		if isCustomError {
			vm = vm.mapFromCustomError(customError)
		} else {
			vm = vm.SetMessage(err.Error())
		}
	}

	return vm
}
