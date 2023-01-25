package dtos

import (
	"fmt"
	"net/http"

	"github.com/fikri240794/golib/errors"
)

type BasePaginationParamsDTO struct {
	Skip int64
	Take int64
}

func NewBasePaginationParamsDTO() *BasePaginationParamsDTO {
	return &BasePaginationParamsDTO{
		Skip: 0,
		Take: 100,
	}
}

func (dto *BasePaginationParamsDTO) SetSkip(skip int64) *BasePaginationParamsDTO {
	dto.Skip = skip

	return dto
}

func (dto *BasePaginationParamsDTO) SetTake(take int64) *BasePaginationParamsDTO {
	dto.Take = take

	return dto
}

func (dto *BasePaginationParamsDTO) Validate() error {
	var errFields []errors.ErrorField = []errors.ErrorField{}

	if dto.Skip < 0 {
		errFields = append(errFields, errors.NewErrorField("skip", fmt.Sprintf("minimum value is %d", 0)))
	}

	if (dto.Skip + dto.Take) > 1000 {
		errFields = append(errFields, errors.NewErrorField("skip", fmt.Sprintf("maximum value is %d", (dto.Skip-((dto.Skip+dto.Take)-1000)))))
	}

	if dto.Take < 1 {
		errFields = append(errFields, errors.NewErrorField("take", fmt.Sprintf("minimum value is %d", 1)))
	}

	if dto.Take > 100 {
		errFields = append(errFields, errors.NewErrorField("take", fmt.Sprintf("maximum value is %d", 100)))
	}

	if len(errFields) > 0 {
		return errors.NewError(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), errFields...)
	}

	return nil
}
