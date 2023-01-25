package vms

type BaseResponseErrorFieldVM struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func NewBaseResponseErrorFieldVM(field string, message string) *BaseResponseErrorFieldVM {
	return &BaseResponseErrorFieldVM{
		Field:   field,
		Message: message,
	}
}
