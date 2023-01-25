package dtos

type BasePaginationResultDTO struct {
	Count int64
}

func NewBasePaginationResultDTO() *BasePaginationResultDTO {
	return &BasePaginationResultDTO{}
}

func (dto *BasePaginationResultDTO) SetCount(count int64) *BasePaginationResultDTO {
	dto.Count = count

	return dto
}
