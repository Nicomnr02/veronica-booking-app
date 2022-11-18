package size

import (
	"pos/pkg/domain/service"
	"pos/pkg/usecase/size"

	"github.com/sirupsen/logrus"
)

type SizeService struct {
	SizeRepository size.SizeInputPort
}

func NewSizeService(ct size.SizeInputPort) *SizeService {
	return &SizeService{
		SizeRepository: ct,
	}
}
func (ct *SizeService) CreateSize(req service.CreateSizeRequest) error {
	err := ct.SizeRepository.CreateSize(req)
	if err != nil {
		return err
	}

	return nil
}

func (ct *SizeService) GetSize(page, limit int) (interface{}, error) {
	size, err := ct.SizeRepository.GetSize(page, limit)
	if err != nil {
		logrus.Error("[GetSize] error get Size in Controller")
		return nil, err
	}
	return size, err
}

func (ct *SizeService) GetSizeByID(ID int) (interface{}, error) {
	size, err := ct.SizeRepository.GetSizeByID(ID)
	if err != nil {
		logrus.Error("[GetSizeByID] error get size by id in Controller")
		return nil, err
	}
	return size, nil
}

func (ct *SizeService) DeleteSize(req service.DeleteSizeRequest) error {
	err := ct.SizeRepository.DeleteSize(req)
	if err != nil {
		return nil
	}
	return err
}

func (ct *SizeService) UpdateSize(req service.UpdateSizeRequest) error {
	err := ct.SizeRepository.UpdateSize(req)
	if err != nil {
		return err
	}

	return nil
}
