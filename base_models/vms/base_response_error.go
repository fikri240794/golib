package vms

import (
	"github.com/fikri240794/golib/errors"
)

type BaseResponseErrorVM struct {
	Code        int                         `json:"-"`
	Message     string                      `json:"message"`
	ErrorFields []*BaseResponseErrorFieldVM `json:"error_fields,omitempty"`
}

func NewBaseResponseErrorVM() *BaseResponseErrorVM {
	return &BaseResponseErrorVM{}
}

func (vm *BaseResponseErrorVM) SetCode(code int) *BaseResponseErrorVM {
	vm.Code = code

	return vm
}

func (vm *BaseResponseErrorVM) SetMessage(message string) *BaseResponseErrorVM {
	vm.Message = message

	return vm
}

func (vm *BaseResponseErrorVM) SetErrorFields(errorField []*BaseResponseErrorFieldVM) *BaseResponseErrorVM {
	if vm.ErrorFields == nil {
		vm.ErrorFields = []*BaseResponseErrorFieldVM{}
	}

	vm.ErrorFields = append(vm.ErrorFields, errorField...)

	return vm
}

func (vm *BaseResponseErrorVM) AddErrorField(errorField *BaseResponseErrorFieldVM) *BaseResponseErrorVM {
	if vm.ErrorFields == nil {
		vm.ErrorFields = []*BaseResponseErrorFieldVM{}
	}

	vm.ErrorFields = append(vm.ErrorFields, errorField)

	return vm
}

func (vm *BaseResponseErrorVM) mapFromCustomError(err errors.Error) *BaseResponseErrorVM {
	vm = vm.SetCode(err.Code.ToHttpStatusCode()).
		SetMessage(err.Message)

	if len(err.ErrorFields) > 0 {
		for i := 0; i < len(err.ErrorFields); i++ {
			vm = vm.AddErrorField(NewBaseResponseErrorFieldVM(err.ErrorFields[i].Field, err.ErrorFields[i].Message))
		}
	}

	return vm
}

func (vm *BaseResponseErrorVM) ParseError(err error) *BaseResponseErrorVM {
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
