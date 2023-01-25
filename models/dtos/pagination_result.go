package dtos

type PaginationResultDTO struct {
	Count int64
}

func NewPaginationResultDTO() *PaginationResultDTO {
	return &PaginationResultDTO{}
}

func (dto *PaginationResultDTO) SetCount(count int64) *PaginationResultDTO {
	dto.Count = count

	return dto
}
