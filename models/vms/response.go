package vms

type ResponseVM struct {
	Error *ResponseErrorVM `json:"error,omitempty"`
	Data  interface{}      `json:"data,omitempty"`
}

func NewResponseVM() *ResponseVM {
	return &ResponseVM{}
}

func (vm *ResponseVM) SetData(data interface{}) *ResponseVM {
	vm.Data = data

	return vm
}

func (vm *ResponseVM) SetError(err *ResponseErrorVM) *ResponseVM {
	vm.Error = err

	return vm
}

func (vm *ResponseVM) SetErrorFromError(err error) *ResponseVM {
	vm.Error = NewResponseErrorVM().
		ParseError(err)

	return vm
}
