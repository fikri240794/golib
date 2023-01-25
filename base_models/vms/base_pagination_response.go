package vms

type BasePaginationResponseVM struct {
	Count int64 `json:"count"`
}

func NewBasePaginationResponseVM() *BasePaginationResponseVM {
	return &BasePaginationResponseVM{}
}

func (vm *BasePaginationResponseVM) SetCount(count int64) *BasePaginationResponseVM {
	vm.Count = count

	return vm
}
