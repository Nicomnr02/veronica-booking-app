package size

import (
	"pos/pkg/domain/entity"
	"pos/pkg/domain/repository"
	"pos/pkg/domain/service"

	"github.com/mitchellh/mapstructure"
)

type SizeUseCase struct {
	repo repository.SizeRepository
	out  SizeOutputPort
}

func NewSizeUseCase(r repository.SizeRepository, o SizeOutputPort) *SizeUseCase {
	return &SizeUseCase{
		repo: r,
		out:  o,
	}
}

func (ct *SizeUseCase) CreateSize(req service.CreateSizeRequest) error {
	var size entity.Size
	err := mapstructure.Decode(req, &size)
	if err != nil {
		return err
	}

	err = ct.repo.CreateSize(size)
	if err != nil {
		return err
	}
	return nil
}

func (ct *SizeUseCase) GetSize(page, limit int) (interface{}, error) {
	data, err := ct.repo.GetSize(page, limit)
	if err != nil {
		return nil, err
	}

	totalSizes, err := ct.repo.CountSize()
	if err != nil {
		return nil, err
	}

	return ct.out.GetResponseSizes(data, totalSizes)
}

func (ct *SizeUseCase) GetSizeByID(ID int) (interface{}, error) {
	data, err := ct.repo.GetSizeByID(ID)
	if err != nil {
		return nil, err
	}

	return ct.out.GetResponseSize(data)
}

func (ct *SizeUseCase) UpdateSize(req service.UpdateSizeRequest) error {
	var size entity.Size
	err := mapstructure.Decode(req, &size)
	if err != nil {
		return nil
	}

	err = ct.repo.UpdateSize(size)
	if err != nil {
		return err
	}

	return nil
}

func (ct *SizeUseCase) DeleteSize(req service.DeleteSizeRequest) error {
	var size entity.Size
	err := mapstructure.Decode(req, &size)
	if err != nil {
		return err
	}

	err = ct.repo.DeleteSize(size)
	if err != nil {
		return nil
	}

	return nil
}
