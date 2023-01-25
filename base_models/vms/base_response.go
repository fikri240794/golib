package vms

type BaseResponseVM struct {
	Error *BaseResponseErrorVM `json:"error,omitempty"`
	Data  interface{}          `json:"data,omitempty"`
}

func NewBaseResponseVM() *BaseResponseVM {
	return &BaseResponseVM{}
}

func (vm *BaseResponseVM) SetData(data interface{}) *BaseResponseVM {
	vm.Data = data

	return vm
}

func (vm *BaseResponseVM) SetError(err *BaseResponseErrorVM) *BaseResponseVM {
	vm.Error = err

	return vm
}

func (vm *BaseResponseVM) SetErrorFromError(err error) *BaseResponseVM {
	vm.Error = NewBaseResponseErrorVM().
		ParseError(err)

	return vm
}
