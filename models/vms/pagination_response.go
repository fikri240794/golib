package vms

type PaginationResponseVM struct {
	Count int64 `json:"count"`
}

func NewPaginationResponseVM() *PaginationResponseVM {
	return &PaginationResponseVM{}
}

func (vm *PaginationResponseVM) SetCount(count int64) *PaginationResponseVM {
	vm.Count = count

	return vm
}
