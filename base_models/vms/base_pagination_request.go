package vms

type BasePaginationRequestVM struct {
	Skip int64 `json:"skip" query:"skip" validate:"required"`
	Take int64 `json:"take" query:"take" validate:"required"`
}

func NewBasePaginationRequestVM() *BasePaginationRequestVM {
	return &BasePaginationRequestVM{}
}
